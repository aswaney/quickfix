//Package quotecancel msg type = Z.
package quotecancel

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a QuoteCancel wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//QuoteReqID is a non-required field for QuoteCancel.
func (m Message) QuoteReqID() (*field.QuoteReqIDField, errors.MessageRejectError) {
	f := &field.QuoteReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from QuoteCancel.
func (m Message) GetQuoteReqID(f *field.QuoteReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a required field for QuoteCancel.
func (m Message) QuoteID() (*field.QuoteIDField, errors.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from QuoteCancel.
func (m Message) GetQuoteID(f *field.QuoteIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteCancelType is a required field for QuoteCancel.
func (m Message) QuoteCancelType() (*field.QuoteCancelTypeField, errors.MessageRejectError) {
	f := &field.QuoteCancelTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteCancelType reads a QuoteCancelType from QuoteCancel.
func (m Message) GetQuoteCancelType(f *field.QuoteCancelTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteResponseLevel is a non-required field for QuoteCancel.
func (m Message) QuoteResponseLevel() (*field.QuoteResponseLevelField, errors.MessageRejectError) {
	f := &field.QuoteResponseLevelField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteResponseLevel reads a QuoteResponseLevel from QuoteCancel.
func (m Message) GetQuoteResponseLevel(f *field.QuoteResponseLevelField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for QuoteCancel.
func (m Message) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from QuoteCancel.
func (m Message) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoQuoteEntries is a required field for QuoteCancel.
func (m Message) NoQuoteEntries() (*field.NoQuoteEntriesField, errors.MessageRejectError) {
	f := &field.NoQuoteEntriesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoQuoteEntries reads a NoQuoteEntries from QuoteCancel.
func (m Message) GetNoQuoteEntries(f *field.NoQuoteEntriesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds QuoteCancel messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for QuoteCancel.
func Builder(
	quoteid *field.QuoteIDField,
	quotecanceltype *field.QuoteCancelTypeField,
	noquoteentries *field.NoQuoteEntriesField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX42))
	builder.Header().Set(field.NewMsgType("Z"))
	builder.Body().Set(quoteid)
	builder.Body().Set(quotecanceltype)
	builder.Body().Set(noquoteentries)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX42, "Z", r
}
