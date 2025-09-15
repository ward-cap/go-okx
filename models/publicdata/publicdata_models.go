package publicdata

import (
	"github.com/ward-cap/go-okx"
	"github.com/ward-cap/go-okx/decimal"
)

type (
	Instrument struct {
		InstID    string                `json:"instId"`
		Uly       string                `json:"uly,omitempty"`
		BaseCcy   string                `json:"baseCcy,omitempty"`
		QuoteCcy  string                `json:"quoteCcy,omitempty"`
		SettleCcy string                `json:"settleCcy,omitempty"`
		CtValCcy  string                `json:"ctValCcy,omitempty"`
		CtVal     decimal.NullDecimalV2 `json:"ctVal"`
		CtMult    decimal.NullDecimalV2 `json:"ctMult"`
		Stk       decimal.NullDecimalV2 `json:"stk"`
		TickSz    decimal.NullDecimalV2 `json:"tickSz"`
		LotSz     decimal.NullDecimalV2 `json:"lotSz"`
		MinSz     decimal.NullDecimalV2 `json:"minSz"`
		Lever     decimal.NullDecimalV2 `json:"lever"`
		MaxMktSz  decimal.NullDecimalV2 `json:"maxMktSz"`
		InstType  okex.InstrumentType   `json:"instType"`
		Category  okex.FeeCategory      `json:"category,string"`
		OptType   okex.OptionType       `json:"optType,omitempty"`
		ListTime  okex.JSONTime         `json:"listTime"`
		ExpTime   *okex.JSONTime        `json:"expTime,omitempty"`
		CtType    okex.ContractType     `json:"ctType,omitempty"`
		Alias     okex.AliasType        `json:"alias,omitempty"`
		State     okex.InstrumentState  `json:"state"`
	}
	DeliveryExerciseHistory struct {
		Details []*DeliveryExerciseHistoryDetails `json:"details"`
		TS      okex.JSONTime                     `json:"ts"`
	}
	DeliveryExerciseHistoryDetails struct {
		InstID string                    `json:"instId"`
		Px     decimal.NullDecimalV2     `json:"px"`
		Type   okex.DeliveryExerciseType `json:"type"`
	}
	OpenInterest struct {
		InstID   string                `json:"instId"`
		Oi       decimal.NullDecimalV2 `json:"oi"`
		OiCcy    decimal.NullDecimalV2 `json:"oiCcy"`
		InstType okex.InstrumentType   `json:"instType"`
		TS       okex.JSONTime         `json:"ts"`
	}
	FundingRate struct {
		InstID          string                `json:"instId"`
		InstType        okex.InstrumentType   `json:"instType"`
		FundingRate     decimal.NullDecimalV2 `json:"fundingRate"`
		NextFundingRate decimal.NullDecimalV2 `json:"NextFundingRate"`
		FundingTime     okex.JSONTime         `json:"fundingTime"`
		NextFundingTime okex.JSONTime         `json:"nextFundingTime"`
	}
	LimitPrice struct {
		InstID   string                `json:"instId"`
		InstType okex.InstrumentType   `json:"instType"`
		BuyLmt   decimal.NullDecimalV2 `json:"buyLmt"`
		SellLmt  decimal.NullDecimalV2 `json:"sellLmt"`
		TS       okex.JSONTime         `json:"ts"`
	}
	EstimatedDeliveryExercisePrice struct {
		InstID   string                `json:"instId"`
		InstType okex.InstrumentType   `json:"instType"`
		SettlePx decimal.NullDecimalV2 `json:"settlePx"`
		TS       okex.JSONTime         `json:"ts"`
	}
	OptionMarketData struct {
		InstID   string                `json:"instId"`
		Uly      string                `json:"uly"`
		InstType okex.InstrumentType   `json:"instType"`
		Delta    decimal.NullDecimalV2 `json:"delta"`
		Gamma    decimal.NullDecimalV2 `json:"gamma"`
		Vega     decimal.NullDecimalV2 `json:"vega"`
		Theta    decimal.NullDecimalV2 `json:"theta"`
		DeltaBS  decimal.NullDecimalV2 `json:"deltaBS"`
		GammaBS  decimal.NullDecimalV2 `json:"gammaBS"`
		VegaBS   decimal.NullDecimalV2 `json:"vegaBS"`
		ThetaBS  decimal.NullDecimalV2 `json:"thetaBS"`
		Lever    decimal.NullDecimalV2 `json:"lever"`
		MarkVol  decimal.NullDecimalV2 `json:"markVol"`
		BidVol   decimal.NullDecimalV2 `json:"bidVol"`
		AskVol   decimal.NullDecimalV2 `json:"askVol"`
		RealVol  decimal.NullDecimalV2 `json:"realVol"`
		TS       okex.JSONTime         `json:"ts"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Ccy          string                `json:"ccy"`
		Amt          decimal.NullDecimalV2 `json:"amt"`
		DiscountLv   decimal.NullDecimalV2 `json:"discountLv"`
		DiscountInfo []*DiscountInfo       `json:"discountInfo"`
	}
	DiscountInfo struct {
		DiscountRate decimal.NullDecimalV2 `json:"discountRate"`
		MaxAmt       decimal.NullDecimalV2 `json:"maxAmt"`
		MinAmt       decimal.NullDecimalV2 `json:"minAmt"`
	}
	SystemTime struct {
		TS okex.JSONTime `json:"ts"`
	}
	LiquidationOrder struct {
		InstID    string                    `json:"instId"`
		Uly       string                    `json:"uly,omitempty"`
		InstType  okex.InstrumentType       `json:"instType"`
		TotalLoss decimal.NullDecimalV2     `json:"totalLoss"`
		Details   []*LiquidationOrderDetail `json:"details"`
	}
	LiquidationOrderDetail struct {
		Ccy     string                `json:"ccy,omitempty"`
		Side    okex.OrderSide        `json:"side"`
		OosSide okex.PositionSide     `json:"posSide"`
		BkPx    decimal.NullDecimalV2 `json:"bkPx"`
		Sz      decimal.NullDecimalV2 `json:"sz"`
		BkLoss  decimal.NullDecimalV2 `json:"bkLoss"`
		TS      okex.JSONTime         `json:"ts"`
	}
	MarkPrice struct {
		InstID   string                `json:"instId"`
		InstType okex.InstrumentType   `json:"instType"`
		MarkPx   decimal.NullDecimalV2 `json:"markPx"`
		TS       okex.JSONTime         `json:"ts"`
	}
	PositionTier struct {
		InstID       string                `json:"instId"`
		Uly          string                `json:"uly,omitempty"`
		InstType     okex.InstrumentType   `json:"instType"`
		Tier         decimal.NullDecimalV2 `json:"tier"`
		MinSz        decimal.NullDecimalV2 `json:"minSz"`
		MaxSz        decimal.NullDecimalV2 `json:"maxSz"`
		Mmr          decimal.NullDecimalV2 `json:"mmr"`
		Imr          decimal.NullDecimalV2 `json:"imr"`
		OptMgnFactor decimal.NullDecimalV2 `json:"optMgnFactor,omitempty"`
		QuoteMaxLoan decimal.NullDecimalV2 `json:"quoteMaxLoan,omitempty"`
		BaseMaxLoan  decimal.NullDecimalV2 `json:"baseMaxLoan,omitempty"`
		MaxLever     decimal.NullDecimalV2 `json:"maxLever"`
		TS           okex.JSONTime         `json:"ts"`
	}
	InterestRateAndLoanQuota struct {
		Basic   []*InterestRateAndLoanBasic `json:"basic"`
		Vip     []*InterestRateAndLoanUser  `json:"vip"`
		Regular []*InterestRateAndLoanUser  `json:"regular"`
	}
	InterestRateAndLoanBasic struct {
		Ccy   string                `json:"ccy"`
		Rate  decimal.NullDecimalV2 `json:"rate"`
		Quota decimal.NullDecimalV2 `json:"quota"`
	}
	InterestRateAndLoanUser struct {
		Level         string                `json:"level"`
		IrDiscount    decimal.NullDecimalV2 `json:"irDiscount"`
		LoanQuotaCoef int                   `json:"loanQuotaCoef,string"`
	}
	State struct {
		Title       string        `json:"title"`
		State       string        `json:"state"`
		Href        string        `json:"href"`
		ServiceType string        `json:"serviceType"`
		System      string        `json:"system"`
		ScheDesc    string        `json:"scheDesc"`
		Begin       okex.JSONTime `json:"begin"`
		End         okex.JSONTime `json:"end"`
	}
)
