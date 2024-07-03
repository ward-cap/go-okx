package account

import (
	"github.com/shopspring/decimal"
	"github.com/ward-cap/go-okx"
)

type (
	Balance struct {
		TotalEq     decimal.NullDecimal `json:"totalEq"`
		IsoEq       decimal.NullDecimal `json:"isoEq"`
		AdjEq       decimal.NullDecimal `json:"adjEq,omitempty"`
		OrdFroz     decimal.NullDecimal `json:"ordFroz,omitempty"`
		Imr         decimal.NullDecimal `json:"imr,omitempty"`
		Mmr         decimal.NullDecimal `json:"mmr,omitempty"`
		MgnRatio    decimal.NullDecimal `json:"mgnRatio,omitempty"`
		NotionalUsd decimal.NullDecimal `json:"notionalUsd,omitempty"`
		Details     []*BalanceDetails   `json:"details,omitempty"`
		UTime       okex.JSONTime       `json:"uTime"`
	}
	BalanceDetails struct {
		Ccy           string              `json:"ccy"`
		Eq            decimal.NullDecimal `json:"eq"`
		CashBal       decimal.NullDecimal `json:"cashBal"`
		IsoEq         decimal.NullDecimal `json:"isoEq,omitempty"`
		AvailEq       decimal.NullDecimal `json:"availEq,omitempty"`
		DisEq         decimal.NullDecimal `json:"disEq"`
		AvailBal      decimal.NullDecimal `json:"availBal"`
		FrozenBal     decimal.NullDecimal `json:"frozenBal"`
		OrdFrozen     decimal.NullDecimal `json:"ordFrozen"`
		Liab          decimal.NullDecimal `json:"liab,omitempty"`
		Upl           decimal.NullDecimal `json:"upl,omitempty"`
		UplLib        decimal.NullDecimal `json:"uplLib,omitempty"`
		CrossLiab     decimal.NullDecimal `json:"crossLiab,omitempty"`
		IsoLiab       decimal.NullDecimal `json:"isoLiab,omitempty"`
		MgnRatio      decimal.NullDecimal `json:"mgnRatio,omitempty"`
		Interest      decimal.NullDecimal `json:"interest,omitempty"`
		Twap          decimal.NullDecimal `json:"twap,omitempty"`
		MaxLoan       decimal.NullDecimal `json:"maxLoan,omitempty"`
		EqUsd         decimal.NullDecimal `json:"eqUsd"`
		NotionalLever decimal.NullDecimal `json:"notionalLever,omitempty"`
		StgyEq        decimal.NullDecimal `json:"stgyEq"`
		IsoUpl        decimal.NullDecimal `json:"isoUpl,omitempty"`
		UTime         okex.JSONTime       `json:"uTime"`
	}
	Position struct {
		InstID      string              `json:"instId"`
		PosCcy      string              `json:"posCcy,omitempty"`
		LiabCcy     string              `json:"liabCcy,omitempty"`
		OptVal      string              `json:"optVal,omitempty"`
		Ccy         string              `json:"ccy"`
		PosID       string              `json:"posId"`
		TradeID     string              `json:"tradeId"`
		Pos         decimal.NullDecimal `json:"pos"`
		AvailPos    decimal.NullDecimal `json:"availPos,omitempty"`
		AvgPx       decimal.NullDecimal `json:"avgPx"`
		Upl         decimal.NullDecimal `json:"upl"`
		UplRatio    decimal.NullDecimal `json:"uplRatio"`
		Lever       decimal.NullDecimal `json:"lever"`
		LiqPx       decimal.NullDecimal `json:"liqPx,omitempty"`
		Imr         decimal.NullDecimal `json:"imr,omitempty"`
		Margin      decimal.NullDecimal `json:"margin,omitempty"`
		MgnRatio    decimal.NullDecimal `json:"mgnRatio"`
		Mmr         decimal.NullDecimal `json:"mmr"`
		Liab        decimal.NullDecimal `json:"liab,omitempty"`
		Interest    decimal.NullDecimal `json:"interest"`
		NotionalUsd decimal.NullDecimal `json:"notionalUsd"`
		ADL         decimal.NullDecimal `json:"adl"`
		Last        decimal.NullDecimal `json:"last"`
		DeltaBS     decimal.NullDecimal `json:"deltaBS"`
		DeltaPA     decimal.NullDecimal `json:"deltaPA"`
		GammaBS     decimal.NullDecimal `json:"gammaBS"`
		GammaPA     decimal.NullDecimal `json:"gammaPA"`
		ThetaBS     decimal.NullDecimal `json:"thetaBS"`
		ThetaPA     decimal.NullDecimal `json:"thetaPA"`
		VegaBS      decimal.NullDecimal `json:"vegaBS"`
		VegaPA      decimal.NullDecimal `json:"vegaPA"`
		PosSide     okex.PositionSide   `json:"posSide"`
		MgnMode     okex.MarginMode     `json:"mgnMode"`
		InstType    okex.InstrumentType `json:"instType"`
		CTime       okex.JSONTime       `json:"cTime"`
		UTime       okex.JSONTime       `json:"uTime"`
	}
	BalanceAndPosition struct {
		EventType okex.EventType    `json:"eventType"`
		PTime     okex.JSONTime     `json:"pTime"`
		UTime     okex.JSONTime     `json:"uTime"`
		PosData   []*Position       `json:"posData"`
		BalData   []*BalanceDetails `json:"balData"`
	}
	PositionAndAccountRisk struct {
		AdjEq   decimal.NullDecimal                  `json:"adjEq,omitempty"`
		BalData []*PositionAndAccountRiskBalanceData `json:"balData"`
		PosData []*PositionAndAccountRiskBalanceData `json:"posData"`
		TS      okex.JSONTime                        `json:"ts"`
	}
	PositionAndAccountRiskBalanceData struct {
		Ccy   string              `json:"ccy"`
		Eq    decimal.NullDecimal `json:"eq"`
		DisEq decimal.NullDecimal `json:"disEq"`
	}
	PositionAndAccountRiskPositionData struct {
		InstID      string              `json:"instId"`
		PosCcy      string              `json:"posCcy,omitempty"`
		Ccy         string              `json:"ccy"`
		NotionalCcy decimal.NullDecimal `json:"notionalCcy"`
		Pos         decimal.NullDecimal `json:"pos"`
		NotionalUsd decimal.NullDecimal `json:"notionalUsd"`
		PosSide     okex.PositionSide   `json:"posSide"`
		InstType    okex.InstrumentType `json:"instType"`
		MgnMode     okex.MarginMode     `json:"mgnMode"`
	}
	Bill struct {
		Ccy       string              `json:"ccy"`
		InstID    string              `json:"instId"`
		Notes     string              `json:"notes"`
		BillID    string              `json:"billId"`
		OrdID     string              `json:"ordId"`
		BalChg    decimal.NullDecimal `json:"balChg"`
		PosBalChg decimal.NullDecimal `json:"posBalChg"`
		Bal       decimal.NullDecimal `json:"bal"`
		PosBal    decimal.NullDecimal `json:"posBal"`
		Sz        decimal.NullDecimal `json:"sz"`
		Pnl       decimal.NullDecimal `json:"pnl"`
		Fee       decimal.NullDecimal `json:"fee"`
		From      okex.AccountType    `json:"from,string"`
		To        okex.AccountType    `json:"to,string"`
		InstType  okex.InstrumentType `json:"instType"`
		MgnMode   okex.MarginMode     `json:"MgnMode"`
		Type      okex.BillType       `json:"type,string"`
		SubType   okex.BillSubType    `json:"subType,string"`
		TS        okex.JSONTime       `json:"ts"`
	}
	Config struct {
		Level      string            `json:"level"`
		LevelTmp   string            `json:"levelTmp"`
		AcctLv     string            `json:"acctLv"`
		AutoLoan   bool              `json:"autoLoan"`
		UID        string            `json:"uid"`
		GreeksType okex.GreekType    `json:"greeksType"`
		PosMode    okex.PositionType `json:"posMode"`
	}
	PositionMode struct {
		PosMode okex.PositionType `json:"posMode"`
	}
	Leverage struct {
		InstID  string              `json:"instId"`
		Lever   decimal.NullDecimal `json:"lever"`
		MgnMode okex.MarginMode     `json:"mgnMode"`
		PosSide okex.PositionSide   `json:"posSide"`
	}
	MaxBuySellAmount struct {
		InstID  string              `json:"instId"`
		Ccy     string              `json:"ccy"`
		MaxBuy  decimal.NullDecimal `json:"maxBuy"`
		MaxSell decimal.NullDecimal `json:"maxSell"`
	}
	MaxAvailableTradeAmount struct {
		InstID    string              `json:"instId"`
		AvailBuy  decimal.NullDecimal `json:"availBuy"`
		AvailSell decimal.NullDecimal `json:"availSell"`
	}
	MarginBalanceAmount struct {
		InstID  string              `json:"instId"`
		Amt     decimal.NullDecimal `json:"amt"`
		PosSide okex.PositionSide   `json:"posSide,string"`
		Type    okex.CountAction    `json:"type,string"`
	}
	Loan struct {
		InstID  string              `json:"instId"`
		MgnCcy  string              `json:"mgnCcy"`
		Ccy     string              `json:"ccy"`
		MaxLoan decimal.NullDecimal `json:"maxLoan"`
		MgnMode okex.MarginMode     `json:"mgnMode"`
		Side    okex.OrderSide      `json:"side,string"`
	}
	Fee struct {
		Level    string              `json:"level"`
		Taker    decimal.NullDecimal `json:"taker"`
		Maker    decimal.NullDecimal `json:"maker"`
		Delivery decimal.NullDecimal `json:"delivery,omitempty"`
		Exercise decimal.NullDecimal `json:"exercise,omitempty"`
		Category okex.FeeCategory    `json:"category,string"`
		InstType okex.InstrumentType `json:"instType"`
		TS       okex.JSONTime       `json:"ts"`
	}
	InterestAccrued struct {
		InstID       string              `json:"instId"`
		Ccy          string              `json:"ccy"`
		Interest     decimal.NullDecimal `json:"interest"`
		InterestRate decimal.NullDecimal `json:"interestRate"`
		Liab         decimal.NullDecimal `json:"liab"`
		MgnMode      okex.MarginMode     `json:"mgnMode"`
		TS           okex.JSONTime       `json:"ts"`
	}
	InterestRate struct {
		Ccy          string              `json:"ccy"`
		InterestRate decimal.NullDecimal `json:"interestRate"`
	}
	Greek struct {
		GreeksType string `json:"greeksType"`
	}
	MaxWithdrawal struct {
		Ccy   string              `json:"ccy"`
		MaxWd decimal.NullDecimal `json:"maxWd"`
	}
)
