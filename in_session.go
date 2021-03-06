package quickfix

import (
	"time"

	"github.com/quickfixgo/quickfix/enum"
)

type inSession struct {
}

func (state inSession) String() string   { return "In Session" }
func (state inSession) IsLoggedOn() bool { return true }

func (state inSession) VerifyMsgIn(session *session, msg Message) (err MessageRejectError) {
	var msgType FIXString
	if err := msg.Header.GetField(tagMsgType, &msgType); err == nil {
		switch string(msgType) {
		case enum.MsgType_LOGON:
			return session.verifyLogon(msg)
		case enum.MsgType_LOGOUT:
			return nil
		case enum.MsgType_RESEND_REQUEST:
			return session.verifyIgnoreSeqNumTooHighOrLow(msg)
		case enum.MsgType_SEQUENCE_RESET:
			var gapFillFlag FIXBoolean
			msg.Body.GetField(tagGapFillFlag, &gapFillFlag)
			return session.verifySelect(msg, bool(gapFillFlag), bool(gapFillFlag))
		default:
			return session.verify(msg)
		}
	}
	return nil
}

func (state inSession) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	var msgType FIXString
	if err := msg.Header.GetField(tagMsgType, &msgType); err == nil {
		switch string(msgType) {
		case enum.MsgType_LOGON:
			session.handleLogon(msg)
			return state
		case enum.MsgType_LOGOUT:
			session.log.OnEvent("Received logout request")
			session.log.OnEvent("Sending logout response")
			session.sendLogout("")
			return latentState{}
		case enum.MsgType_TEST_REQUEST:
			return state.handleTestRequest(session, msg)
		case enum.MsgType_RESEND_REQUEST:
			return state.handleResendRequest(session, msg)
		case enum.MsgType_SEQUENCE_RESET:
			return state.handleSequenceReset(session, msg)
		}
	}

	session.store.IncrNextTargetMsgSeqNum()
	return state
}

func (state inSession) FixMsgInRej(session *session, msg Message, rej MessageRejectError) (nextState sessionState) {
	var msgType FIXString
	if err := msg.Header.GetField(tagMsgType, &msgType); err == nil {
		switch string(msgType) {
		case enum.MsgType_LOGON:
			session.initiateLogout("")
			return logoutState{}
		case enum.MsgType_LOGOUT:
			return latentState{}
		}
	}
	return state.processReject(session, msg, rej)
}

func (state inSession) Timeout(session *session, event event) (nextState sessionState) {
	switch event {
	case needHeartbeat:
		heartBt := NewMessage()
		heartBt.Header.SetField(tagMsgType, FIXString("0"))
		session.send(heartBt)
	case peerTimeout:
		testReq := NewMessage()
		testReq.Header.SetField(tagMsgType, FIXString("1"))
		testReq.Body.SetField(tagTestReqID, FIXString("TEST"))
		session.send(testReq)
		session.peerTimer.Reset(time.Duration(int64(1.2 * float64(session.heartBeatTimeout))))
		return pendingTimeout{}
	}
	return state
}

func (state inSession) handleTestRequest(session *session, msg Message) (nextState sessionState) {
	var testReq FIXString
	if err := msg.Body.GetField(tagTestReqID, &testReq); err != nil {
		session.log.OnEvent("Test Request with no testRequestID")
	} else {
		heartBt := NewMessage()
		heartBt.Header.SetField(tagMsgType, FIXString("0"))
		heartBt.Body.SetField(tagTestReqID, testReq)
		session.send(heartBt)
	}

	session.store.IncrNextTargetMsgSeqNum()
	return state
}

func (state inSession) handleSequenceReset(session *session, msg Message) (nextState sessionState) {
	var newSeqNo FIXInt
	if err := msg.Body.GetField(tagNewSeqNo, &newSeqNo); err == nil {
		expectedSeqNum := FIXInt(session.store.NextTargetMsgSeqNum())
		session.log.OnEventf("Received SequenceReset FROM: %v TO: %v", expectedSeqNum, newSeqNo)

		switch {
		case newSeqNo > expectedSeqNum:
			session.store.SetNextTargetMsgSeqNum(int(newSeqNo))
		case newSeqNo < expectedSeqNum:
			//FIXME: to be compliant with legacy tests, do not include tag in reftagid? (11c_NewSeqNoLess)
			session.doReject(msg, valueIsIncorrectNoTag())
		}
	}
	return state
}

func (state inSession) handleResendRequest(session *session, msg Message) (nextState sessionState) {
	var err error
	var beginSeqNoField FIXInt
	if err = msg.Body.GetField(tagBeginSeqNo, &beginSeqNoField); err != nil {
		return state.processReject(session, msg, RequiredTagMissing(tagBeginSeqNo))
	}

	beginSeqNo := beginSeqNoField

	var endSeqNoField FIXInt
	if err = msg.Body.GetField(tagEndSeqNo, &endSeqNoField); err != nil {
		return state.processReject(session, msg, RequiredTagMissing(tagEndSeqNo))
	}

	endSeqNo := int(endSeqNoField)

	session.log.OnEventf("Received ResendRequest FROM: %d TO: %d", beginSeqNo, endSeqNo)
	expectedSeqNum := session.store.NextTargetMsgSeqNum()

	if (session.sessionID.BeginString >= enum.BeginStringFIX42 && endSeqNo == 0) ||
		(session.sessionID.BeginString <= enum.BeginStringFIX42 && endSeqNo == 999999) ||
		(endSeqNo >= expectedSeqNum) {
		endSeqNo = expectedSeqNum - 1
	}

	state.resendMessages(session, int(beginSeqNo), endSeqNo)
	session.store.IncrNextTargetMsgSeqNum()
	return state
}

func (state inSession) resendMessages(session *session, beginSeqNo, endSeqNo int) {
	msgs, err := session.store.GetMessages(beginSeqNo, endSeqNo)
	if err != nil {
		session.log.OnEventf("error retrieving messages from store: %s", err.Error())
		panic(err)
	}

	seqNum := beginSeqNo
	nextSeqNum := seqNum
	for _, msgBytes := range msgs {
		msg, _ := ParseMessage(msgBytes)
		msgType, _ := msg.Header.GetString(tagMsgType)
		sentMessageSeqNum, _ := msg.Header.GetInt(tagMsgSeqNum)

		if isAdminMessageType(msgType) {
			nextSeqNum = sentMessageSeqNum + 1
			continue
		}

		if seqNum != sentMessageSeqNum {
			state.generateSequenceReset(session, seqNum, sentMessageSeqNum)
		}
		session.resend(msg)
		seqNum = sentMessageSeqNum + 1
		nextSeqNum = seqNum
	}

	if seqNum != nextSeqNum { // gapfill for catch-up
		state.generateSequenceReset(session, seqNum, nextSeqNum)
	}
}

func (state inSession) processReject(session *session, msg Message, rej MessageRejectError) (nextState sessionState) {
	switch TypedError := rej.(type) {
	case targetTooHigh:

		switch session.sessionState.(type) {
		default:
			session.doTargetTooHigh(TypedError)
		case resendState:
			//assumes target too high reject already sent
		}

		session.messageStash[TypedError.ReceivedTarget] = msg
		return resendState{}

	case targetTooLow:
		return state.doTargetTooLow(session, msg, TypedError)
	case incorrectBeginString:
		session.initiateLogout(rej.Error())
		return logoutState{}
	}

	switch rej.RejectReason() {
	case rejectReasonCompIDProblem, rejectReasonSendingTimeAccuracyProblem:
		session.doReject(msg, rej)
		session.initiateLogout("")
		return logoutState{}
	default:
		session.doReject(msg, rej)
		session.store.IncrNextTargetMsgSeqNum()
		return state
	}
}

func (state inSession) doTargetTooLow(session *session, msg Message, rej targetTooLow) (nextState sessionState) {
	var posDupFlag FIXBoolean
	if err := msg.Header.GetField(tagPossDupFlag, &posDupFlag); err == nil && posDupFlag {

		origSendingTime := new(FIXUTCTimestamp)
		if err = msg.Header.GetField(tagOrigSendingTime, origSendingTime); err != nil {
			session.doReject(msg, RequiredTagMissing(tagOrigSendingTime))
			return state
		}

		sendingTime := new(FIXUTCTimestamp)
		msg.Header.GetField(tagSendingTime, sendingTime)

		if sendingTime.Before(origSendingTime.Time) {
			session.doReject(msg, sendingTimeAccuracyProblem())
			session.initiateLogout("")
			return logoutState{}
		}

		if appReject := session.fromCallback(msg); appReject != nil {
			session.doReject(msg, appReject)
			session.initiateLogout("")
			return logoutState{}
		}
	} else {
		session.initiateLogout(rej.Error())
		return logoutState{}
	}

	return state
}

func (state *inSession) generateSequenceReset(session *session, beginSeqNo int, endSeqNo int) {
	sequenceReset := NewMessage()
	session.fillDefaultHeader(sequenceReset)

	sequenceReset.Header.SetField(tagMsgType, FIXString("4"))
	sequenceReset.Header.SetField(tagMsgSeqNum, FIXInt(beginSeqNo))
	sequenceReset.Header.SetField(tagPossDupFlag, FIXBoolean(true))
	sequenceReset.Body.SetField(tagNewSeqNo, FIXInt(endSeqNo))
	sequenceReset.Body.SetField(tagGapFillFlag, FIXBoolean(true))

	var origSendingTime FIXString
	if err := sequenceReset.Header.GetField(tagSendingTime, &origSendingTime); err == nil {
		sequenceReset.Header.SetField(tagOrigSendingTime, origSendingTime)
	}

	//FIXME error check?
	msgBytes, _ := sequenceReset.Build()
	session.sendBytes(msgBytes)
}
