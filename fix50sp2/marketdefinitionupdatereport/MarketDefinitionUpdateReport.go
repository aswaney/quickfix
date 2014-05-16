//Package marketdefinitionupdatereport msg type = BV.
package marketdefinitionupdatereport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a MarketDefinitionUpdateReport wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//MarketReportID is a required field for MarketDefinitionUpdateReport.
func (m Message) MarketReportID() (*field.MarketReportIDField, errors.MessageRejectError) {
	f := &field.MarketReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketReportID reads a MarketReportID from MarketDefinitionUpdateReport.
func (m Message) GetMarketReportID(f *field.MarketReportIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketReqID is a non-required field for MarketDefinitionUpdateReport.
func (m Message) MarketReqID() (*field.MarketReqIDField, errors.MessageRejectError) {
	f := &field.MarketReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketReqID reads a MarketReqID from MarketDefinitionUpdateReport.
func (m Message) GetMarketReqID(f *field.MarketReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketUpdateAction is a non-required field for MarketDefinitionUpdateReport.
func (m Message) MarketUpdateAction() (*field.MarketUpdateActionField, errors.MessageRejectError) {
	f := &field.MarketUpdateActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketUpdateAction reads a MarketUpdateAction from MarketDefinitionUpdateReport.
func (m Message) GetMarketUpdateAction(f *field.MarketUpdateActionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a required field for MarketDefinitionUpdateReport.
func (m Message) MarketID() (*field.MarketIDField, errors.MessageRejectError) {
	f := &field.MarketIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from MarketDefinitionUpdateReport.
func (m Message) GetMarketID(f *field.MarketIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for MarketDefinitionUpdateReport.
func (m Message) MarketSegmentID() (*field.MarketSegmentIDField, errors.MessageRejectError) {
	f := &field.MarketSegmentIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from MarketDefinitionUpdateReport.
func (m Message) GetMarketSegmentID(f *field.MarketSegmentIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentDesc is a non-required field for MarketDefinitionUpdateReport.
func (m Message) MarketSegmentDesc() (*field.MarketSegmentDescField, errors.MessageRejectError) {
	f := &field.MarketSegmentDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentDesc reads a MarketSegmentDesc from MarketDefinitionUpdateReport.
func (m Message) GetMarketSegmentDesc(f *field.MarketSegmentDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedMktSegmDescLen is a non-required field for MarketDefinitionUpdateReport.
func (m Message) EncodedMktSegmDescLen() (*field.EncodedMktSegmDescLenField, errors.MessageRejectError) {
	f := &field.EncodedMktSegmDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedMktSegmDescLen reads a EncodedMktSegmDescLen from MarketDefinitionUpdateReport.
func (m Message) GetEncodedMktSegmDescLen(f *field.EncodedMktSegmDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedMktSegmDesc is a non-required field for MarketDefinitionUpdateReport.
func (m Message) EncodedMktSegmDesc() (*field.EncodedMktSegmDescField, errors.MessageRejectError) {
	f := &field.EncodedMktSegmDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedMktSegmDesc reads a EncodedMktSegmDesc from MarketDefinitionUpdateReport.
func (m Message) GetEncodedMktSegmDesc(f *field.EncodedMktSegmDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ParentMktSegmID is a non-required field for MarketDefinitionUpdateReport.
func (m Message) ParentMktSegmID() (*field.ParentMktSegmIDField, errors.MessageRejectError) {
	f := &field.ParentMktSegmIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetParentMktSegmID reads a ParentMktSegmID from MarketDefinitionUpdateReport.
func (m Message) GetParentMktSegmID(f *field.ParentMktSegmIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for MarketDefinitionUpdateReport.
func (m Message) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from MarketDefinitionUpdateReport.
func (m Message) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTickRules is a non-required field for MarketDefinitionUpdateReport.
func (m Message) NoTickRules() (*field.NoTickRulesField, errors.MessageRejectError) {
	f := &field.NoTickRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTickRules reads a NoTickRules from MarketDefinitionUpdateReport.
func (m Message) GetNoTickRules(f *field.NoTickRulesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLotTypeRules is a non-required field for MarketDefinitionUpdateReport.
func (m Message) NoLotTypeRules() (*field.NoLotTypeRulesField, errors.MessageRejectError) {
	f := &field.NoLotTypeRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLotTypeRules reads a NoLotTypeRules from MarketDefinitionUpdateReport.
func (m Message) GetNoLotTypeRules(f *field.NoLotTypeRulesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceLimitType is a non-required field for MarketDefinitionUpdateReport.
func (m Message) PriceLimitType() (*field.PriceLimitTypeField, errors.MessageRejectError) {
	f := &field.PriceLimitTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceLimitType reads a PriceLimitType from MarketDefinitionUpdateReport.
func (m Message) GetPriceLimitType(f *field.PriceLimitTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LowLimitPrice is a non-required field for MarketDefinitionUpdateReport.
func (m Message) LowLimitPrice() (*field.LowLimitPriceField, errors.MessageRejectError) {
	f := &field.LowLimitPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLowLimitPrice reads a LowLimitPrice from MarketDefinitionUpdateReport.
func (m Message) GetLowLimitPrice(f *field.LowLimitPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HighLimitPrice is a non-required field for MarketDefinitionUpdateReport.
func (m Message) HighLimitPrice() (*field.HighLimitPriceField, errors.MessageRejectError) {
	f := &field.HighLimitPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHighLimitPrice reads a HighLimitPrice from MarketDefinitionUpdateReport.
func (m Message) GetHighLimitPrice(f *field.HighLimitPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingReferencePrice is a non-required field for MarketDefinitionUpdateReport.
func (m Message) TradingReferencePrice() (*field.TradingReferencePriceField, errors.MessageRejectError) {
	f := &field.TradingReferencePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingReferencePrice reads a TradingReferencePrice from MarketDefinitionUpdateReport.
func (m Message) GetTradingReferencePrice(f *field.TradingReferencePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpirationCycle is a non-required field for MarketDefinitionUpdateReport.
func (m Message) ExpirationCycle() (*field.ExpirationCycleField, errors.MessageRejectError) {
	f := &field.ExpirationCycleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpirationCycle reads a ExpirationCycle from MarketDefinitionUpdateReport.
func (m Message) GetExpirationCycle(f *field.ExpirationCycleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinTradeVol is a non-required field for MarketDefinitionUpdateReport.
func (m Message) MinTradeVol() (*field.MinTradeVolField, errors.MessageRejectError) {
	f := &field.MinTradeVolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinTradeVol reads a MinTradeVol from MarketDefinitionUpdateReport.
func (m Message) GetMinTradeVol(f *field.MinTradeVolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxTradeVol is a non-required field for MarketDefinitionUpdateReport.
func (m Message) MaxTradeVol() (*field.MaxTradeVolField, errors.MessageRejectError) {
	f := &field.MaxTradeVolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxTradeVol reads a MaxTradeVol from MarketDefinitionUpdateReport.
func (m Message) GetMaxTradeVol(f *field.MaxTradeVolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxPriceVariation is a non-required field for MarketDefinitionUpdateReport.
func (m Message) MaxPriceVariation() (*field.MaxPriceVariationField, errors.MessageRejectError) {
	f := &field.MaxPriceVariationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxPriceVariation reads a MaxPriceVariation from MarketDefinitionUpdateReport.
func (m Message) GetMaxPriceVariation(f *field.MaxPriceVariationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ImpliedMarketIndicator is a non-required field for MarketDefinitionUpdateReport.
func (m Message) ImpliedMarketIndicator() (*field.ImpliedMarketIndicatorField, errors.MessageRejectError) {
	f := &field.ImpliedMarketIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetImpliedMarketIndicator reads a ImpliedMarketIndicator from MarketDefinitionUpdateReport.
func (m Message) GetImpliedMarketIndicator(f *field.ImpliedMarketIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingCurrency is a non-required field for MarketDefinitionUpdateReport.
func (m Message) TradingCurrency() (*field.TradingCurrencyField, errors.MessageRejectError) {
	f := &field.TradingCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingCurrency reads a TradingCurrency from MarketDefinitionUpdateReport.
func (m Message) GetTradingCurrency(f *field.TradingCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundLot is a non-required field for MarketDefinitionUpdateReport.
func (m Message) RoundLot() (*field.RoundLotField, errors.MessageRejectError) {
	f := &field.RoundLotField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundLot reads a RoundLot from MarketDefinitionUpdateReport.
func (m Message) GetRoundLot(f *field.RoundLotField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultilegModel is a non-required field for MarketDefinitionUpdateReport.
func (m Message) MultilegModel() (*field.MultilegModelField, errors.MessageRejectError) {
	f := &field.MultilegModelField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultilegModel reads a MultilegModel from MarketDefinitionUpdateReport.
func (m Message) GetMultilegModel(f *field.MultilegModelField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultilegPriceMethod is a non-required field for MarketDefinitionUpdateReport.
func (m Message) MultilegPriceMethod() (*field.MultilegPriceMethodField, errors.MessageRejectError) {
	f := &field.MultilegPriceMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultilegPriceMethod reads a MultilegPriceMethod from MarketDefinitionUpdateReport.
func (m Message) GetMultilegPriceMethod(f *field.MultilegPriceMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for MarketDefinitionUpdateReport.
func (m Message) PriceType() (*field.PriceTypeField, errors.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from MarketDefinitionUpdateReport.
func (m Message) GetPriceType(f *field.PriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrdTypeRules is a non-required field for MarketDefinitionUpdateReport.
func (m Message) NoOrdTypeRules() (*field.NoOrdTypeRulesField, errors.MessageRejectError) {
	f := &field.NoOrdTypeRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrdTypeRules reads a NoOrdTypeRules from MarketDefinitionUpdateReport.
func (m Message) GetNoOrdTypeRules(f *field.NoOrdTypeRulesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTimeInForceRules is a non-required field for MarketDefinitionUpdateReport.
func (m Message) NoTimeInForceRules() (*field.NoTimeInForceRulesField, errors.MessageRejectError) {
	f := &field.NoTimeInForceRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTimeInForceRules reads a NoTimeInForceRules from MarketDefinitionUpdateReport.
func (m Message) GetNoTimeInForceRules(f *field.NoTimeInForceRulesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoExecInstRules is a non-required field for MarketDefinitionUpdateReport.
func (m Message) NoExecInstRules() (*field.NoExecInstRulesField, errors.MessageRejectError) {
	f := &field.NoExecInstRulesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoExecInstRules reads a NoExecInstRules from MarketDefinitionUpdateReport.
func (m Message) GetNoExecInstRules(f *field.NoExecInstRulesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for MarketDefinitionUpdateReport.
func (m Message) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from MarketDefinitionUpdateReport.
func (m Message) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for MarketDefinitionUpdateReport.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from MarketDefinitionUpdateReport.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for MarketDefinitionUpdateReport.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from MarketDefinitionUpdateReport.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for MarketDefinitionUpdateReport.
func (m Message) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from MarketDefinitionUpdateReport.
func (m Message) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for MarketDefinitionUpdateReport.
func (m Message) ApplID() (*field.ApplIDField, errors.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from MarketDefinitionUpdateReport.
func (m Message) GetApplID(f *field.ApplIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for MarketDefinitionUpdateReport.
func (m Message) ApplSeqNum() (*field.ApplSeqNumField, errors.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from MarketDefinitionUpdateReport.
func (m Message) GetApplSeqNum(f *field.ApplSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for MarketDefinitionUpdateReport.
func (m Message) ApplLastSeqNum() (*field.ApplLastSeqNumField, errors.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from MarketDefinitionUpdateReport.
func (m Message) GetApplLastSeqNum(f *field.ApplLastSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for MarketDefinitionUpdateReport.
func (m Message) ApplResendFlag() (*field.ApplResendFlagField, errors.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from MarketDefinitionUpdateReport.
func (m Message) GetApplResendFlag(f *field.ApplResendFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds MarketDefinitionUpdateReport messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for MarketDefinitionUpdateReport.
func Builder(
	marketreportid *field.MarketReportIDField,
	marketid *field.MarketIDField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("BV"))
	builder.Body().Set(marketreportid)
	builder.Body().Set(marketid)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "BV", r
}
