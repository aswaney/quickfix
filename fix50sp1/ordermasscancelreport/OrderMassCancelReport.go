//Package ordermasscancelreport msg type = r.
package ordermasscancelreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a OrderMassCancelReport wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//ClOrdID is a non-required field for OrderMassCancelReport.
func (m Message) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from OrderMassCancelReport.
func (m Message) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for OrderMassCancelReport.
func (m Message) SecondaryClOrdID() (*field.SecondaryClOrdIDField, errors.MessageRejectError) {
	f := &field.SecondaryClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from OrderMassCancelReport.
func (m Message) GetSecondaryClOrdID(f *field.SecondaryClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderID is a required field for OrderMassCancelReport.
func (m Message) OrderID() (*field.OrderIDField, errors.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from OrderMassCancelReport.
func (m Message) GetOrderID(f *field.OrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryOrderID is a non-required field for OrderMassCancelReport.
func (m Message) SecondaryOrderID() (*field.SecondaryOrderIDField, errors.MessageRejectError) {
	f := &field.SecondaryOrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryOrderID reads a SecondaryOrderID from OrderMassCancelReport.
func (m Message) GetSecondaryOrderID(f *field.SecondaryOrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MassCancelRequestType is a required field for OrderMassCancelReport.
func (m Message) MassCancelRequestType() (*field.MassCancelRequestTypeField, errors.MessageRejectError) {
	f := &field.MassCancelRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMassCancelRequestType reads a MassCancelRequestType from OrderMassCancelReport.
func (m Message) GetMassCancelRequestType(f *field.MassCancelRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MassCancelResponse is a required field for OrderMassCancelReport.
func (m Message) MassCancelResponse() (*field.MassCancelResponseField, errors.MessageRejectError) {
	f := &field.MassCancelResponseField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMassCancelResponse reads a MassCancelResponse from OrderMassCancelReport.
func (m Message) GetMassCancelResponse(f *field.MassCancelResponseField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MassCancelRejectReason is a non-required field for OrderMassCancelReport.
func (m Message) MassCancelRejectReason() (*field.MassCancelRejectReasonField, errors.MessageRejectError) {
	f := &field.MassCancelRejectReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMassCancelRejectReason reads a MassCancelRejectReason from OrderMassCancelReport.
func (m Message) GetMassCancelRejectReason(f *field.MassCancelRejectReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalAffectedOrders is a non-required field for OrderMassCancelReport.
func (m Message) TotalAffectedOrders() (*field.TotalAffectedOrdersField, errors.MessageRejectError) {
	f := &field.TotalAffectedOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotalAffectedOrders reads a TotalAffectedOrders from OrderMassCancelReport.
func (m Message) GetTotalAffectedOrders(f *field.TotalAffectedOrdersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAffectedOrders is a non-required field for OrderMassCancelReport.
func (m Message) NoAffectedOrders() (*field.NoAffectedOrdersField, errors.MessageRejectError) {
	f := &field.NoAffectedOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAffectedOrders reads a NoAffectedOrders from OrderMassCancelReport.
func (m Message) GetNoAffectedOrders(f *field.NoAffectedOrdersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for OrderMassCancelReport.
func (m Message) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from OrderMassCancelReport.
func (m Message) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for OrderMassCancelReport.
func (m Message) TradingSessionSubID() (*field.TradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from OrderMassCancelReport.
func (m Message) GetTradingSessionSubID(f *field.TradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for OrderMassCancelReport.
func (m Message) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from OrderMassCancelReport.
func (m Message) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for OrderMassCancelReport.
func (m Message) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from OrderMassCancelReport.
func (m Message) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for OrderMassCancelReport.
func (m Message) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from OrderMassCancelReport.
func (m Message) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for OrderMassCancelReport.
func (m Message) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from OrderMassCancelReport.
func (m Message) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for OrderMassCancelReport.
func (m Message) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from OrderMassCancelReport.
func (m Message) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for OrderMassCancelReport.
func (m Message) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from OrderMassCancelReport.
func (m Message) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for OrderMassCancelReport.
func (m Message) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from OrderMassCancelReport.
func (m Message) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for OrderMassCancelReport.
func (m Message) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from OrderMassCancelReport.
func (m Message) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for OrderMassCancelReport.
func (m Message) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from OrderMassCancelReport.
func (m Message) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for OrderMassCancelReport.
func (m Message) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from OrderMassCancelReport.
func (m Message) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for OrderMassCancelReport.
func (m Message) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from OrderMassCancelReport.
func (m Message) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for OrderMassCancelReport.
func (m Message) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from OrderMassCancelReport.
func (m Message) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for OrderMassCancelReport.
func (m Message) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from OrderMassCancelReport.
func (m Message) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for OrderMassCancelReport.
func (m Message) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from OrderMassCancelReport.
func (m Message) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for OrderMassCancelReport.
func (m Message) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from OrderMassCancelReport.
func (m Message) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for OrderMassCancelReport.
func (m Message) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from OrderMassCancelReport.
func (m Message) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for OrderMassCancelReport.
func (m Message) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from OrderMassCancelReport.
func (m Message) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for OrderMassCancelReport.
func (m Message) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from OrderMassCancelReport.
func (m Message) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for OrderMassCancelReport.
func (m Message) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from OrderMassCancelReport.
func (m Message) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for OrderMassCancelReport.
func (m Message) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from OrderMassCancelReport.
func (m Message) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for OrderMassCancelReport.
func (m Message) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from OrderMassCancelReport.
func (m Message) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for OrderMassCancelReport.
func (m Message) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from OrderMassCancelReport.
func (m Message) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for OrderMassCancelReport.
func (m Message) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from OrderMassCancelReport.
func (m Message) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for OrderMassCancelReport.
func (m Message) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from OrderMassCancelReport.
func (m Message) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for OrderMassCancelReport.
func (m Message) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from OrderMassCancelReport.
func (m Message) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for OrderMassCancelReport.
func (m Message) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from OrderMassCancelReport.
func (m Message) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for OrderMassCancelReport.
func (m Message) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from OrderMassCancelReport.
func (m Message) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for OrderMassCancelReport.
func (m Message) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from OrderMassCancelReport.
func (m Message) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for OrderMassCancelReport.
func (m Message) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from OrderMassCancelReport.
func (m Message) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for OrderMassCancelReport.
func (m Message) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from OrderMassCancelReport.
func (m Message) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for OrderMassCancelReport.
func (m Message) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from OrderMassCancelReport.
func (m Message) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for OrderMassCancelReport.
func (m Message) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from OrderMassCancelReport.
func (m Message) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for OrderMassCancelReport.
func (m Message) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from OrderMassCancelReport.
func (m Message) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for OrderMassCancelReport.
func (m Message) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from OrderMassCancelReport.
func (m Message) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for OrderMassCancelReport.
func (m Message) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from OrderMassCancelReport.
func (m Message) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for OrderMassCancelReport.
func (m Message) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from OrderMassCancelReport.
func (m Message) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for OrderMassCancelReport.
func (m Message) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from OrderMassCancelReport.
func (m Message) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for OrderMassCancelReport.
func (m Message) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from OrderMassCancelReport.
func (m Message) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for OrderMassCancelReport.
func (m Message) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from OrderMassCancelReport.
func (m Message) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for OrderMassCancelReport.
func (m Message) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from OrderMassCancelReport.
func (m Message) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for OrderMassCancelReport.
func (m Message) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from OrderMassCancelReport.
func (m Message) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for OrderMassCancelReport.
func (m Message) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from OrderMassCancelReport.
func (m Message) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for OrderMassCancelReport.
func (m Message) SecurityStatus() (*field.SecurityStatusField, errors.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from OrderMassCancelReport.
func (m Message) GetSecurityStatus(f *field.SecurityStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for OrderMassCancelReport.
func (m Message) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, errors.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from OrderMassCancelReport.
func (m Message) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for OrderMassCancelReport.
func (m Message) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, errors.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from OrderMassCancelReport.
func (m Message) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for OrderMassCancelReport.
func (m Message) StrikeMultiplier() (*field.StrikeMultiplierField, errors.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from OrderMassCancelReport.
func (m Message) GetStrikeMultiplier(f *field.StrikeMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for OrderMassCancelReport.
func (m Message) StrikeValue() (*field.StrikeValueField, errors.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from OrderMassCancelReport.
func (m Message) GetStrikeValue(f *field.StrikeValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for OrderMassCancelReport.
func (m Message) MinPriceIncrement() (*field.MinPriceIncrementField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from OrderMassCancelReport.
func (m Message) GetMinPriceIncrement(f *field.MinPriceIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for OrderMassCancelReport.
func (m Message) PositionLimit() (*field.PositionLimitField, errors.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from OrderMassCancelReport.
func (m Message) GetPositionLimit(f *field.PositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for OrderMassCancelReport.
func (m Message) NTPositionLimit() (*field.NTPositionLimitField, errors.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from OrderMassCancelReport.
func (m Message) GetNTPositionLimit(f *field.NTPositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for OrderMassCancelReport.
func (m Message) NoInstrumentParties() (*field.NoInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from OrderMassCancelReport.
func (m Message) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for OrderMassCancelReport.
func (m Message) UnitOfMeasure() (*field.UnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from OrderMassCancelReport.
func (m Message) GetUnitOfMeasure(f *field.UnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for OrderMassCancelReport.
func (m Message) TimeUnit() (*field.TimeUnitField, errors.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from OrderMassCancelReport.
func (m Message) GetTimeUnit(f *field.TimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for OrderMassCancelReport.
func (m Message) MaturityTime() (*field.MaturityTimeField, errors.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from OrderMassCancelReport.
func (m Message) GetMaturityTime(f *field.MaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for OrderMassCancelReport.
func (m Message) SecurityGroup() (*field.SecurityGroupField, errors.MessageRejectError) {
	f := &field.SecurityGroupField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from OrderMassCancelReport.
func (m Message) GetSecurityGroup(f *field.SecurityGroupField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for OrderMassCancelReport.
func (m Message) MinPriceIncrementAmount() (*field.MinPriceIncrementAmountField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from OrderMassCancelReport.
func (m Message) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for OrderMassCancelReport.
func (m Message) UnitOfMeasureQty() (*field.UnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from OrderMassCancelReport.
func (m Message) GetUnitOfMeasureQty(f *field.UnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for OrderMassCancelReport.
func (m Message) SecurityXMLLen() (*field.SecurityXMLLenField, errors.MessageRejectError) {
	f := &field.SecurityXMLLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from OrderMassCancelReport.
func (m Message) GetSecurityXMLLen(f *field.SecurityXMLLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for OrderMassCancelReport.
func (m Message) SecurityXML() (*field.SecurityXMLField, errors.MessageRejectError) {
	f := &field.SecurityXMLField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from OrderMassCancelReport.
func (m Message) GetSecurityXML(f *field.SecurityXMLField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for OrderMassCancelReport.
func (m Message) SecurityXMLSchema() (*field.SecurityXMLSchemaField, errors.MessageRejectError) {
	f := &field.SecurityXMLSchemaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from OrderMassCancelReport.
func (m Message) GetSecurityXMLSchema(f *field.SecurityXMLSchemaField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for OrderMassCancelReport.
func (m Message) ProductComplex() (*field.ProductComplexField, errors.MessageRejectError) {
	f := &field.ProductComplexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from OrderMassCancelReport.
func (m Message) GetProductComplex(f *field.ProductComplexField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for OrderMassCancelReport.
func (m Message) PriceUnitOfMeasure() (*field.PriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from OrderMassCancelReport.
func (m Message) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for OrderMassCancelReport.
func (m Message) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from OrderMassCancelReport.
func (m Message) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for OrderMassCancelReport.
func (m Message) SettlMethod() (*field.SettlMethodField, errors.MessageRejectError) {
	f := &field.SettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from OrderMassCancelReport.
func (m Message) GetSettlMethod(f *field.SettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for OrderMassCancelReport.
func (m Message) ExerciseStyle() (*field.ExerciseStyleField, errors.MessageRejectError) {
	f := &field.ExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from OrderMassCancelReport.
func (m Message) GetExerciseStyle(f *field.ExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayAmount is a non-required field for OrderMassCancelReport.
func (m Message) OptPayAmount() (*field.OptPayAmountField, errors.MessageRejectError) {
	f := &field.OptPayAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayAmount reads a OptPayAmount from OrderMassCancelReport.
func (m Message) GetOptPayAmount(f *field.OptPayAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for OrderMassCancelReport.
func (m Message) PriceQuoteMethod() (*field.PriceQuoteMethodField, errors.MessageRejectError) {
	f := &field.PriceQuoteMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from OrderMassCancelReport.
func (m Message) GetPriceQuoteMethod(f *field.PriceQuoteMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for OrderMassCancelReport.
func (m Message) ListMethod() (*field.ListMethodField, errors.MessageRejectError) {
	f := &field.ListMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from OrderMassCancelReport.
func (m Message) GetListMethod(f *field.ListMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for OrderMassCancelReport.
func (m Message) CapPrice() (*field.CapPriceField, errors.MessageRejectError) {
	f := &field.CapPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from OrderMassCancelReport.
func (m Message) GetCapPrice(f *field.CapPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for OrderMassCancelReport.
func (m Message) FloorPrice() (*field.FloorPriceField, errors.MessageRejectError) {
	f := &field.FloorPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from OrderMassCancelReport.
func (m Message) GetFloorPrice(f *field.FloorPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for OrderMassCancelReport.
func (m Message) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from OrderMassCancelReport.
func (m Message) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for OrderMassCancelReport.
func (m Message) FlexibleIndicator() (*field.FlexibleIndicatorField, errors.MessageRejectError) {
	f := &field.FlexibleIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from OrderMassCancelReport.
func (m Message) GetFlexibleIndicator(f *field.FlexibleIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for OrderMassCancelReport.
func (m Message) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicatorField, errors.MessageRejectError) {
	f := &field.FlexProductEligibilityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from OrderMassCancelReport.
func (m Message) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FuturesValuationMethod is a non-required field for OrderMassCancelReport.
func (m Message) FuturesValuationMethod() (*field.FuturesValuationMethodField, errors.MessageRejectError) {
	f := &field.FuturesValuationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFuturesValuationMethod reads a FuturesValuationMethod from OrderMassCancelReport.
func (m Message) GetFuturesValuationMethod(f *field.FuturesValuationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbol is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingSymbol() (*field.UnderlyingSymbolField, errors.MessageRejectError) {
	f := &field.UnderlyingSymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbol reads a UnderlyingSymbol from OrderMassCancelReport.
func (m Message) GetUnderlyingSymbol(f *field.UnderlyingSymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbolSfx is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingSymbolSfx() (*field.UnderlyingSymbolSfxField, errors.MessageRejectError) {
	f := &field.UnderlyingSymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbolSfx reads a UnderlyingSymbolSfx from OrderMassCancelReport.
func (m Message) GetUnderlyingSymbolSfx(f *field.UnderlyingSymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityID is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingSecurityID() (*field.UnderlyingSecurityIDField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityID reads a UnderlyingSecurityID from OrderMassCancelReport.
func (m Message) GetUnderlyingSecurityID(f *field.UnderlyingSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityIDSource is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingSecurityIDSource() (*field.UnderlyingSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityIDSource reads a UnderlyingSecurityIDSource from OrderMassCancelReport.
func (m Message) GetUnderlyingSecurityIDSource(f *field.UnderlyingSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingSecurityAltID is a non-required field for OrderMassCancelReport.
func (m Message) NoUnderlyingSecurityAltID() (*field.NoUnderlyingSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoUnderlyingSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingSecurityAltID reads a NoUnderlyingSecurityAltID from OrderMassCancelReport.
func (m Message) GetNoUnderlyingSecurityAltID(f *field.NoUnderlyingSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingProduct is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingProduct() (*field.UnderlyingProductField, errors.MessageRejectError) {
	f := &field.UnderlyingProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingProduct reads a UnderlyingProduct from OrderMassCancelReport.
func (m Message) GetUnderlyingProduct(f *field.UnderlyingProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCFICode is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingCFICode() (*field.UnderlyingCFICodeField, errors.MessageRejectError) {
	f := &field.UnderlyingCFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCFICode reads a UnderlyingCFICode from OrderMassCancelReport.
func (m Message) GetUnderlyingCFICode(f *field.UnderlyingCFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityType is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingSecurityType() (*field.UnderlyingSecurityTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityType reads a UnderlyingSecurityType from OrderMassCancelReport.
func (m Message) GetUnderlyingSecurityType(f *field.UnderlyingSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecuritySubType is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingSecuritySubType() (*field.UnderlyingSecuritySubTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingSecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecuritySubType reads a UnderlyingSecuritySubType from OrderMassCancelReport.
func (m Message) GetUnderlyingSecuritySubType(f *field.UnderlyingSecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityMonthYear is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingMaturityMonthYear() (*field.UnderlyingMaturityMonthYearField, errors.MessageRejectError) {
	f := &field.UnderlyingMaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityMonthYear reads a UnderlyingMaturityMonthYear from OrderMassCancelReport.
func (m Message) GetUnderlyingMaturityMonthYear(f *field.UnderlyingMaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityDate is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingMaturityDate() (*field.UnderlyingMaturityDateField, errors.MessageRejectError) {
	f := &field.UnderlyingMaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityDate reads a UnderlyingMaturityDate from OrderMassCancelReport.
func (m Message) GetUnderlyingMaturityDate(f *field.UnderlyingMaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponPaymentDate is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingCouponPaymentDate() (*field.UnderlyingCouponPaymentDateField, errors.MessageRejectError) {
	f := &field.UnderlyingCouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponPaymentDate reads a UnderlyingCouponPaymentDate from OrderMassCancelReport.
func (m Message) GetUnderlyingCouponPaymentDate(f *field.UnderlyingCouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssueDate is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingIssueDate() (*field.UnderlyingIssueDateField, errors.MessageRejectError) {
	f := &field.UnderlyingIssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssueDate reads a UnderlyingIssueDate from OrderMassCancelReport.
func (m Message) GetUnderlyingIssueDate(f *field.UnderlyingIssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepoCollateralSecurityType is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingRepoCollateralSecurityType() (*field.UnderlyingRepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingRepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepoCollateralSecurityType reads a UnderlyingRepoCollateralSecurityType from OrderMassCancelReport.
func (m Message) GetUnderlyingRepoCollateralSecurityType(f *field.UnderlyingRepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseTerm is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingRepurchaseTerm() (*field.UnderlyingRepurchaseTermField, errors.MessageRejectError) {
	f := &field.UnderlyingRepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseTerm reads a UnderlyingRepurchaseTerm from OrderMassCancelReport.
func (m Message) GetUnderlyingRepurchaseTerm(f *field.UnderlyingRepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseRate is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingRepurchaseRate() (*field.UnderlyingRepurchaseRateField, errors.MessageRejectError) {
	f := &field.UnderlyingRepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseRate reads a UnderlyingRepurchaseRate from OrderMassCancelReport.
func (m Message) GetUnderlyingRepurchaseRate(f *field.UnderlyingRepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFactor is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingFactor() (*field.UnderlyingFactorField, errors.MessageRejectError) {
	f := &field.UnderlyingFactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFactor reads a UnderlyingFactor from OrderMassCancelReport.
func (m Message) GetUnderlyingFactor(f *field.UnderlyingFactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCreditRating is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingCreditRating() (*field.UnderlyingCreditRatingField, errors.MessageRejectError) {
	f := &field.UnderlyingCreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCreditRating reads a UnderlyingCreditRating from OrderMassCancelReport.
func (m Message) GetUnderlyingCreditRating(f *field.UnderlyingCreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingInstrRegistry is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingInstrRegistry() (*field.UnderlyingInstrRegistryField, errors.MessageRejectError) {
	f := &field.UnderlyingInstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingInstrRegistry reads a UnderlyingInstrRegistry from OrderMassCancelReport.
func (m Message) GetUnderlyingInstrRegistry(f *field.UnderlyingInstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCountryOfIssue is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingCountryOfIssue() (*field.UnderlyingCountryOfIssueField, errors.MessageRejectError) {
	f := &field.UnderlyingCountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCountryOfIssue reads a UnderlyingCountryOfIssue from OrderMassCancelReport.
func (m Message) GetUnderlyingCountryOfIssue(f *field.UnderlyingCountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingStateOrProvinceOfIssue() (*field.UnderlyingStateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.UnderlyingStateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStateOrProvinceOfIssue reads a UnderlyingStateOrProvinceOfIssue from OrderMassCancelReport.
func (m Message) GetUnderlyingStateOrProvinceOfIssue(f *field.UnderlyingStateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingLocaleOfIssue is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingLocaleOfIssue() (*field.UnderlyingLocaleOfIssueField, errors.MessageRejectError) {
	f := &field.UnderlyingLocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingLocaleOfIssue reads a UnderlyingLocaleOfIssue from OrderMassCancelReport.
func (m Message) GetUnderlyingLocaleOfIssue(f *field.UnderlyingLocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRedemptionDate is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingRedemptionDate() (*field.UnderlyingRedemptionDateField, errors.MessageRejectError) {
	f := &field.UnderlyingRedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRedemptionDate reads a UnderlyingRedemptionDate from OrderMassCancelReport.
func (m Message) GetUnderlyingRedemptionDate(f *field.UnderlyingRedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikePrice is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingStrikePrice() (*field.UnderlyingStrikePriceField, errors.MessageRejectError) {
	f := &field.UnderlyingStrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikePrice reads a UnderlyingStrikePrice from OrderMassCancelReport.
func (m Message) GetUnderlyingStrikePrice(f *field.UnderlyingStrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikeCurrency is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingStrikeCurrency() (*field.UnderlyingStrikeCurrencyField, errors.MessageRejectError) {
	f := &field.UnderlyingStrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikeCurrency reads a UnderlyingStrikeCurrency from OrderMassCancelReport.
func (m Message) GetUnderlyingStrikeCurrency(f *field.UnderlyingStrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingOptAttribute is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingOptAttribute() (*field.UnderlyingOptAttributeField, errors.MessageRejectError) {
	f := &field.UnderlyingOptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingOptAttribute reads a UnderlyingOptAttribute from OrderMassCancelReport.
func (m Message) GetUnderlyingOptAttribute(f *field.UnderlyingOptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingContractMultiplier is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingContractMultiplier() (*field.UnderlyingContractMultiplierField, errors.MessageRejectError) {
	f := &field.UnderlyingContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingContractMultiplier reads a UnderlyingContractMultiplier from OrderMassCancelReport.
func (m Message) GetUnderlyingContractMultiplier(f *field.UnderlyingContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponRate is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingCouponRate() (*field.UnderlyingCouponRateField, errors.MessageRejectError) {
	f := &field.UnderlyingCouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponRate reads a UnderlyingCouponRate from OrderMassCancelReport.
func (m Message) GetUnderlyingCouponRate(f *field.UnderlyingCouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityExchange is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingSecurityExchange() (*field.UnderlyingSecurityExchangeField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityExchange reads a UnderlyingSecurityExchange from OrderMassCancelReport.
func (m Message) GetUnderlyingSecurityExchange(f *field.UnderlyingSecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssuer is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingIssuer() (*field.UnderlyingIssuerField, errors.MessageRejectError) {
	f := &field.UnderlyingIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssuer reads a UnderlyingIssuer from OrderMassCancelReport.
func (m Message) GetUnderlyingIssuer(f *field.UnderlyingIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuerLen is a non-required field for OrderMassCancelReport.
func (m Message) EncodedUnderlyingIssuerLen() (*field.EncodedUnderlyingIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuerLen reads a EncodedUnderlyingIssuerLen from OrderMassCancelReport.
func (m Message) GetEncodedUnderlyingIssuerLen(f *field.EncodedUnderlyingIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuer is a non-required field for OrderMassCancelReport.
func (m Message) EncodedUnderlyingIssuer() (*field.EncodedUnderlyingIssuerField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuer reads a EncodedUnderlyingIssuer from OrderMassCancelReport.
func (m Message) GetEncodedUnderlyingIssuer(f *field.EncodedUnderlyingIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityDesc is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingSecurityDesc() (*field.UnderlyingSecurityDescField, errors.MessageRejectError) {
	f := &field.UnderlyingSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityDesc reads a UnderlyingSecurityDesc from OrderMassCancelReport.
func (m Message) GetUnderlyingSecurityDesc(f *field.UnderlyingSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDescLen is a non-required field for OrderMassCancelReport.
func (m Message) EncodedUnderlyingSecurityDescLen() (*field.EncodedUnderlyingSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDescLen reads a EncodedUnderlyingSecurityDescLen from OrderMassCancelReport.
func (m Message) GetEncodedUnderlyingSecurityDescLen(f *field.EncodedUnderlyingSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDesc is a non-required field for OrderMassCancelReport.
func (m Message) EncodedUnderlyingSecurityDesc() (*field.EncodedUnderlyingSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedUnderlyingSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDesc reads a EncodedUnderlyingSecurityDesc from OrderMassCancelReport.
func (m Message) GetEncodedUnderlyingSecurityDesc(f *field.EncodedUnderlyingSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPProgram is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingCPProgram() (*field.UnderlyingCPProgramField, errors.MessageRejectError) {
	f := &field.UnderlyingCPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPProgram reads a UnderlyingCPProgram from OrderMassCancelReport.
func (m Message) GetUnderlyingCPProgram(f *field.UnderlyingCPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPRegType is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingCPRegType() (*field.UnderlyingCPRegTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingCPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPRegType reads a UnderlyingCPRegType from OrderMassCancelReport.
func (m Message) GetUnderlyingCPRegType(f *field.UnderlyingCPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrency is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingCurrency() (*field.UnderlyingCurrencyField, errors.MessageRejectError) {
	f := &field.UnderlyingCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrency reads a UnderlyingCurrency from OrderMassCancelReport.
func (m Message) GetUnderlyingCurrency(f *field.UnderlyingCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingQty is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingQty() (*field.UnderlyingQtyField, errors.MessageRejectError) {
	f := &field.UnderlyingQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingQty reads a UnderlyingQty from OrderMassCancelReport.
func (m Message) GetUnderlyingQty(f *field.UnderlyingQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPx is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingPx() (*field.UnderlyingPxField, errors.MessageRejectError) {
	f := &field.UnderlyingPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPx reads a UnderlyingPx from OrderMassCancelReport.
func (m Message) GetUnderlyingPx(f *field.UnderlyingPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingDirtyPrice is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingDirtyPrice() (*field.UnderlyingDirtyPriceField, errors.MessageRejectError) {
	f := &field.UnderlyingDirtyPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingDirtyPrice reads a UnderlyingDirtyPrice from OrderMassCancelReport.
func (m Message) GetUnderlyingDirtyPrice(f *field.UnderlyingDirtyPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndPrice is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingEndPrice() (*field.UnderlyingEndPriceField, errors.MessageRejectError) {
	f := &field.UnderlyingEndPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndPrice reads a UnderlyingEndPrice from OrderMassCancelReport.
func (m Message) GetUnderlyingEndPrice(f *field.UnderlyingEndPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStartValue is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingStartValue() (*field.UnderlyingStartValueField, errors.MessageRejectError) {
	f := &field.UnderlyingStartValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStartValue reads a UnderlyingStartValue from OrderMassCancelReport.
func (m Message) GetUnderlyingStartValue(f *field.UnderlyingStartValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrentValue is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingCurrentValue() (*field.UnderlyingCurrentValueField, errors.MessageRejectError) {
	f := &field.UnderlyingCurrentValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrentValue reads a UnderlyingCurrentValue from OrderMassCancelReport.
func (m Message) GetUnderlyingCurrentValue(f *field.UnderlyingCurrentValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndValue is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingEndValue() (*field.UnderlyingEndValueField, errors.MessageRejectError) {
	f := &field.UnderlyingEndValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndValue reads a UnderlyingEndValue from OrderMassCancelReport.
func (m Message) GetUnderlyingEndValue(f *field.UnderlyingEndValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingStips is a non-required field for OrderMassCancelReport.
func (m Message) NoUnderlyingStips() (*field.NoUnderlyingStipsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingStipsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingStips reads a NoUnderlyingStips from OrderMassCancelReport.
func (m Message) GetNoUnderlyingStips(f *field.NoUnderlyingStipsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAllocationPercent is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingAllocationPercent() (*field.UnderlyingAllocationPercentField, errors.MessageRejectError) {
	f := &field.UnderlyingAllocationPercentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAllocationPercent reads a UnderlyingAllocationPercent from OrderMassCancelReport.
func (m Message) GetUnderlyingAllocationPercent(f *field.UnderlyingAllocationPercentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlementType is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingSettlementType() (*field.UnderlyingSettlementTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingSettlementTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlementType reads a UnderlyingSettlementType from OrderMassCancelReport.
func (m Message) GetUnderlyingSettlementType(f *field.UnderlyingSettlementTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashAmount is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingCashAmount() (*field.UnderlyingCashAmountField, errors.MessageRejectError) {
	f := &field.UnderlyingCashAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashAmount reads a UnderlyingCashAmount from OrderMassCancelReport.
func (m Message) GetUnderlyingCashAmount(f *field.UnderlyingCashAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashType is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingCashType() (*field.UnderlyingCashTypeField, errors.MessageRejectError) {
	f := &field.UnderlyingCashTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashType reads a UnderlyingCashType from OrderMassCancelReport.
func (m Message) GetUnderlyingCashType(f *field.UnderlyingCashTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasure is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingUnitOfMeasure() (*field.UnderlyingUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnderlyingUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasure reads a UnderlyingUnitOfMeasure from OrderMassCancelReport.
func (m Message) GetUnderlyingUnitOfMeasure(f *field.UnderlyingUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingTimeUnit is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingTimeUnit() (*field.UnderlyingTimeUnitField, errors.MessageRejectError) {
	f := &field.UnderlyingTimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingTimeUnit reads a UnderlyingTimeUnit from OrderMassCancelReport.
func (m Message) GetUnderlyingTimeUnit(f *field.UnderlyingTimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCapValue is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingCapValue() (*field.UnderlyingCapValueField, errors.MessageRejectError) {
	f := &field.UnderlyingCapValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCapValue reads a UnderlyingCapValue from OrderMassCancelReport.
func (m Message) GetUnderlyingCapValue(f *field.UnderlyingCapValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUndlyInstrumentParties is a non-required field for OrderMassCancelReport.
func (m Message) NoUndlyInstrumentParties() (*field.NoUndlyInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoUndlyInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUndlyInstrumentParties reads a NoUndlyInstrumentParties from OrderMassCancelReport.
func (m Message) GetNoUndlyInstrumentParties(f *field.NoUndlyInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlMethod is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingSettlMethod() (*field.UnderlyingSettlMethodField, errors.MessageRejectError) {
	f := &field.UnderlyingSettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlMethod reads a UnderlyingSettlMethod from OrderMassCancelReport.
func (m Message) GetUnderlyingSettlMethod(f *field.UnderlyingSettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAdjustedQuantity is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingAdjustedQuantity() (*field.UnderlyingAdjustedQuantityField, errors.MessageRejectError) {
	f := &field.UnderlyingAdjustedQuantityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAdjustedQuantity reads a UnderlyingAdjustedQuantity from OrderMassCancelReport.
func (m Message) GetUnderlyingAdjustedQuantity(f *field.UnderlyingAdjustedQuantityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRate is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingFXRate() (*field.UnderlyingFXRateField, errors.MessageRejectError) {
	f := &field.UnderlyingFXRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRate reads a UnderlyingFXRate from OrderMassCancelReport.
func (m Message) GetUnderlyingFXRate(f *field.UnderlyingFXRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRateCalc is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingFXRateCalc() (*field.UnderlyingFXRateCalcField, errors.MessageRejectError) {
	f := &field.UnderlyingFXRateCalcField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRateCalc reads a UnderlyingFXRateCalc from OrderMassCancelReport.
func (m Message) GetUnderlyingFXRateCalc(f *field.UnderlyingFXRateCalcField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityTime is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingMaturityTime() (*field.UnderlyingMaturityTimeField, errors.MessageRejectError) {
	f := &field.UnderlyingMaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityTime reads a UnderlyingMaturityTime from OrderMassCancelReport.
func (m Message) GetUnderlyingMaturityTime(f *field.UnderlyingMaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPutOrCall is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingPutOrCall() (*field.UnderlyingPutOrCallField, errors.MessageRejectError) {
	f := &field.UnderlyingPutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPutOrCall reads a UnderlyingPutOrCall from OrderMassCancelReport.
func (m Message) GetUnderlyingPutOrCall(f *field.UnderlyingPutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingExerciseStyle is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingExerciseStyle() (*field.UnderlyingExerciseStyleField, errors.MessageRejectError) {
	f := &field.UnderlyingExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingExerciseStyle reads a UnderlyingExerciseStyle from OrderMassCancelReport.
func (m Message) GetUnderlyingExerciseStyle(f *field.UnderlyingExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasureQty is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingUnitOfMeasureQty() (*field.UnderlyingUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnderlyingUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasureQty reads a UnderlyingUnitOfMeasureQty from OrderMassCancelReport.
func (m Message) GetUnderlyingUnitOfMeasureQty(f *field.UnderlyingUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceUnitOfMeasure is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingPriceUnitOfMeasure() (*field.UnderlyingPriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnderlyingPriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceUnitOfMeasure reads a UnderlyingPriceUnitOfMeasure from OrderMassCancelReport.
func (m Message) GetUnderlyingPriceUnitOfMeasure(f *field.UnderlyingPriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceUnitOfMeasureQty is a non-required field for OrderMassCancelReport.
func (m Message) UnderlyingPriceUnitOfMeasureQty() (*field.UnderlyingPriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnderlyingPriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceUnitOfMeasureQty reads a UnderlyingPriceUnitOfMeasureQty from OrderMassCancelReport.
func (m Message) GetUnderlyingPriceUnitOfMeasureQty(f *field.UnderlyingPriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for OrderMassCancelReport.
func (m Message) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from OrderMassCancelReport.
func (m Message) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for OrderMassCancelReport.
func (m Message) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from OrderMassCancelReport.
func (m Message) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for OrderMassCancelReport.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from OrderMassCancelReport.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for OrderMassCancelReport.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from OrderMassCancelReport.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for OrderMassCancelReport.
func (m Message) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from OrderMassCancelReport.
func (m Message) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for OrderMassCancelReport.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from OrderMassCancelReport.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MassActionReportID is a required field for OrderMassCancelReport.
func (m Message) MassActionReportID() (*field.MassActionReportIDField, errors.MessageRejectError) {
	f := &field.MassActionReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMassActionReportID reads a MassActionReportID from OrderMassCancelReport.
func (m Message) GetMassActionReportID(f *field.MassActionReportIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoNotAffectedOrders is a non-required field for OrderMassCancelReport.
func (m Message) NoNotAffectedOrders() (*field.NoNotAffectedOrdersField, errors.MessageRejectError) {
	f := &field.NoNotAffectedOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoNotAffectedOrders reads a NoNotAffectedOrders from OrderMassCancelReport.
func (m Message) GetNoNotAffectedOrders(f *field.NoNotAffectedOrdersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for OrderMassCancelReport.
func (m Message) MarketID() (*field.MarketIDField, errors.MessageRejectError) {
	f := &field.MarketIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from OrderMassCancelReport.
func (m Message) GetMarketID(f *field.MarketIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for OrderMassCancelReport.
func (m Message) MarketSegmentID() (*field.MarketSegmentIDField, errors.MessageRejectError) {
	f := &field.MarketSegmentIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from OrderMassCancelReport.
func (m Message) GetMarketSegmentID(f *field.MarketSegmentIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds OrderMassCancelReport messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for OrderMassCancelReport.
func Builder(
	orderid *field.OrderIDField,
	masscancelrequesttype *field.MassCancelRequestTypeField,
	masscancelresponse *field.MassCancelResponseField,
	massactionreportid *field.MassActionReportIDField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header().Set(field.NewMsgType("r"))
	builder.Body().Set(orderid)
	builder.Body().Set(masscancelrequesttype)
	builder.Body().Set(masscancelresponse)
	builder.Body().Set(massactionreportid)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "r", r
}
