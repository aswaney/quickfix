package quickfix

import "github.com/quickfixgo/quickfix/enum"

type logoutState struct{}

func (state logoutState) String() string { return "Logout State" }
func (s logoutState) IsLoggedOn() bool   { return false }

func (state logoutState) VerifyMsgIn(session *session, msg Message) MessageRejectError { return nil }

func (state logoutState) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	var msgType FIXString
	if err := msg.Header.GetField(tagMsgType, &msgType); err != nil {
		return latentState{}
	}

	switch string(msgType) {
	case enum.MsgType_LOGOUT:
		session.log.OnEvent("Received logout response")
		return latentState{}
	default:
		return state
	}
}

func (state logoutState) FixMsgInRej(session *session, msg Message, err MessageRejectError) sessionState {
	return state.FixMsgIn(session, msg)
}

func (state logoutState) Timeout(session *session, event event) (nextState sessionState) {
	switch event {
	case logoutTimeout:
		session.log.OnEvent("Timed out waiting for logout response")
		return latentState{}
	}

	return state
}
