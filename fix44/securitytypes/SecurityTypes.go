//Package securitytypes msg type = w.
package securitytypes

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a SecurityTypes wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//SecurityReqID is a required field for SecurityTypes.
func (m Message) SecurityReqID() (*field.SecurityReqIDField, errors.MessageRejectError) {
	f := &field.SecurityReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from SecurityTypes.
func (m Message) GetSecurityReqID(f *field.SecurityReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseID is a required field for SecurityTypes.
func (m Message) SecurityResponseID() (*field.SecurityResponseIDField, errors.MessageRejectError) {
	f := &field.SecurityResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseID reads a SecurityResponseID from SecurityTypes.
func (m Message) GetSecurityResponseID(f *field.SecurityResponseIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseType is a required field for SecurityTypes.
func (m Message) SecurityResponseType() (*field.SecurityResponseTypeField, errors.MessageRejectError) {
	f := &field.SecurityResponseTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseType reads a SecurityResponseType from SecurityTypes.
func (m Message) GetSecurityResponseType(f *field.SecurityResponseTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoSecurityTypes is a non-required field for SecurityTypes.
func (m Message) TotNoSecurityTypes() (*field.TotNoSecurityTypesField, errors.MessageRejectError) {
	f := &field.TotNoSecurityTypesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoSecurityTypes reads a TotNoSecurityTypes from SecurityTypes.
func (m Message) GetTotNoSecurityTypes(f *field.TotNoSecurityTypesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for SecurityTypes.
func (m Message) LastFragment() (*field.LastFragmentField, errors.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from SecurityTypes.
func (m Message) GetLastFragment(f *field.LastFragmentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityTypes is a non-required field for SecurityTypes.
func (m Message) NoSecurityTypes() (*field.NoSecurityTypesField, errors.MessageRejectError) {
	f := &field.NoSecurityTypesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityTypes reads a NoSecurityTypes from SecurityTypes.
func (m Message) GetNoSecurityTypes(f *field.NoSecurityTypesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for SecurityTypes.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from SecurityTypes.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for SecurityTypes.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from SecurityTypes.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for SecurityTypes.
func (m Message) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from SecurityTypes.
func (m Message) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for SecurityTypes.
func (m Message) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from SecurityTypes.
func (m Message) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for SecurityTypes.
func (m Message) TradingSessionSubID() (*field.TradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from SecurityTypes.
func (m Message) GetTradingSessionSubID(f *field.TradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for SecurityTypes.
func (m Message) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, errors.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from SecurityTypes.
func (m Message) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds SecurityTypes messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for SecurityTypes.
func Builder(
	securityreqid *field.SecurityReqIDField,
	securityresponseid *field.SecurityResponseIDField,
	securityresponsetype *field.SecurityResponseTypeField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header().Set(field.NewMsgType("w"))
	builder.Body().Set(securityreqid)
	builder.Body().Set(securityresponseid)
	builder.Body().Set(securityresponsetype)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX44, "w", r
}
