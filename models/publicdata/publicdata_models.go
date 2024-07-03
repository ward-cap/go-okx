package publicdata

import (
	"github.com/shopspring/decimal"
	"github.com/ward-cap/go-okx"
)

type (
	Instrument struct {
		InstID    string               `json:"instId"`
		Uly       string               `json:"uly,omitempty"`
		BaseCcy   string               `json:"baseCcy,omitempty"`
		QuoteCcy  string               `json:"quoteCcy,omitempty"`
		SettleCcy string               `json:"settleCcy,omitempty"`
		CtValCcy  string               `json:"ctValCcy,omitempty"`
		CtVal     decimal.NullDecimal  `json:"ctVal,omitempty"`
		CtMult    decimal.NullDecimal  `json:"ctMult,omitempty"`
		Stk       decimal.NullDecimal  `json:"stk,omitempty"`
		TickSz    decimal.NullDecimal  `json:"tickSz,omitempty"`
		LotSz     decimal.NullDecimal  `json:"lotSz,omitempty"`
		MinSz     decimal.NullDecimal  `json:"minSz,omitempty"`
		Lever     decimal.NullDecimal  `json:"lever"`
		InstType  okex.InstrumentType  `json:"instType"`
		Category  okex.FeeCategory     `json:"category,string"`
		OptType   okex.OptionType      `json:"optType,omitempty"`
		ListTime  okex.JSONTime        `json:"listTime"`
		ExpTime   okex.JSONTime        `json:"expTime,omitempty"`
		CtType    okex.ContractType    `json:"ctType,omitempty"`
		Alias     okex.AliasType       `json:"alias,omitempty"`
		State     okex.InstrumentState `json:"state"`
	}
	DeliveryExerciseHistory struct {
		Details []*DeliveryExerciseHistoryDetails `json:"details"`
		TS      okex.JSONTime                     `json:"ts"`
	}
	DeliveryExerciseHistoryDetails struct {
		InstID string                    `json:"instId"`
		Px     decimal.NullDecimal       `json:"px"`
		Type   okex.DeliveryExerciseType `json:"type"`
	}
	OpenInterest struct {
		InstID   string              `json:"instId"`
		Oi       decimal.NullDecimal `json:"oi"`
		OiCcy    decimal.NullDecimal `json:"oiCcy"`
		InstType okex.InstrumentType `json:"instType"`
		TS       okex.JSONTime       `json:"ts"`
	}
	FundingRate struct {
		InstID          string              `json:"instId"`
		InstType        okex.InstrumentType `json:"instType"`
		FundingRate     decimal.NullDecimal `json:"fundingRate"`
		NextFundingRate decimal.NullDecimal `json:"NextFundingRate"`
		FundingTime     okex.JSONTime       `json:"fundingTime"`
		NextFundingTime okex.JSONTime       `json:"nextFundingTime"`
	}
	LimitPrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		BuyLmt   decimal.NullDecimal `json:"buyLmt"`
		SellLmt  decimal.NullDecimal `json:"sellLmt"`
		TS       okex.JSONTime       `json:"ts"`
	}
	EstimatedDeliveryExercisePrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		SettlePx decimal.NullDecimal `json:"settlePx"`
		TS       okex.JSONTime       `json:"ts"`
	}
	OptionMarketData struct {
		InstID   string              `json:"instId"`
		Uly      string              `json:"uly"`
		InstType okex.InstrumentType `json:"instType"`
		Delta    decimal.NullDecimal `json:"delta"`
		Gamma    decimal.NullDecimal `json:"gamma"`
		Vega     decimal.NullDecimal `json:"vega"`
		Theta    decimal.NullDecimal `json:"theta"`
		DeltaBS  decimal.NullDecimal `json:"deltaBS"`
		GammaBS  decimal.NullDecimal `json:"gammaBS"`
		VegaBS   decimal.NullDecimal `json:"vegaBS"`
		ThetaBS  decimal.NullDecimal `json:"thetaBS"`
		Lever    decimal.NullDecimal `json:"lever"`
		MarkVol  decimal.NullDecimal `json:"markVol"`
		BidVol   decimal.NullDecimal `json:"bidVol"`
		AskVol   decimal.NullDecimal `json:"askVol"`
		RealVol  decimal.NullDecimal `json:"realVol"`
		TS       okex.JSONTime       `json:"ts"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Ccy          string              `json:"ccy"`
		Amt          decimal.NullDecimal `json:"amt"`
		DiscountLv   decimal.NullDecimal `json:"discountLv"`
		DiscountInfo []*DiscountInfo     `json:"discountInfo"`
	}
	DiscountInfo struct {
		DiscountRate decimal.NullDecimal `json:"discountRate"`
		MaxAmt       decimal.NullDecimal `json:"maxAmt"`
		MinAmt       decimal.NullDecimal `json:"minAmt"`
	}
	SystemTime struct {
		TS okex.JSONTime `json:"ts"`
	}
	LiquidationOrder struct {
		InstID    string                    `json:"instId"`
		Uly       string                    `json:"uly,omitempty"`
		InstType  okex.InstrumentType       `json:"instType"`
		TotalLoss decimal.NullDecimal       `json:"totalLoss"`
		Details   []*LiquidationOrderDetail `json:"details"`
	}
	LiquidationOrderDetail struct {
		Ccy     string              `json:"ccy,omitempty"`
		Side    okex.OrderSide      `json:"side"`
		OosSide okex.PositionSide   `json:"posSide"`
		BkPx    decimal.NullDecimal `json:"bkPx"`
		Sz      decimal.NullDecimal `json:"sz"`
		BkLoss  decimal.NullDecimal `json:"bkLoss"`
		TS      okex.JSONTime       `json:"ts"`
	}
	MarkPrice struct {
		InstID   string              `json:"instId"`
		InstType okex.InstrumentType `json:"instType"`
		MarkPx   decimal.NullDecimal `json:"markPx"`
		TS       okex.JSONTime       `json:"ts"`
	}
	PositionTier struct {
		InstID       string              `json:"instId"`
		Uly          string              `json:"uly,omitempty"`
		InstType     okex.InstrumentType `json:"instType"`
		Tier         decimal.NullDecimal `json:"tier"`
		MinSz        decimal.NullDecimal `json:"minSz"`
		MaxSz        decimal.NullDecimal `json:"maxSz"`
		Mmr          decimal.NullDecimal `json:"mmr"`
		Imr          decimal.NullDecimal `json:"imr"`
		OptMgnFactor decimal.NullDecimal `json:"optMgnFactor,omitempty"`
		QuoteMaxLoan decimal.NullDecimal `json:"quoteMaxLoan,omitempty"`
		BaseMaxLoan  decimal.NullDecimal `json:"baseMaxLoan,omitempty"`
		MaxLever     decimal.NullDecimal `json:"maxLever"`
		TS           okex.JSONTime       `json:"ts"`
	}
	InterestRateAndLoanQuota struct {
		Basic   []*InterestRateAndLoanBasic `json:"basic"`
		Vip     []*InterestRateAndLoanUser  `json:"vip"`
		Regular []*InterestRateAndLoanUser  `json:"regular"`
	}
	InterestRateAndLoanBasic struct {
		Ccy   string              `json:"ccy"`
		Rate  decimal.NullDecimal `json:"rate"`
		Quota decimal.NullDecimal `json:"quota"`
	}
	InterestRateAndLoanUser struct {
		Level         string              `json:"level"`
		IrDiscount    decimal.NullDecimal `json:"irDiscount"`
		LoanQuotaCoef int                 `json:"loanQuotaCoef,string"`
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
