package account

import (
	"github.com/ward-cap/go-okx"
	"github.com/ward-cap/go-okx/decimal"
)

type (
	Balance struct {
		TotalEq     decimal.NullDecimalV2 `json:"totalEq"`
		IsoEq       decimal.NullDecimalV2 `json:"isoEq"`
		AdjEq       decimal.NullDecimalV2 `json:"adjEq,omitempty"`
		OrdFroz     decimal.NullDecimalV2 `json:"ordFroz,omitempty"`
		Imr         decimal.NullDecimalV2 `json:"imr,omitempty"`
		Mmr         decimal.NullDecimalV2 `json:"mmr,omitempty"`
		MgnRatio    decimal.NullDecimalV2 `json:"mgnRatio,omitempty"`
		NotionalUsd decimal.NullDecimalV2 `json:"notionalUsd,omitempty"`
		Details     []*BalanceDetails     `json:"details,omitempty"`
		UTime       okex.JSONTime         `json:"uTime"`
	}
	BalanceDetails struct {
		Ccy           string                `json:"ccy"`
		Eq            decimal.NullDecimalV2 `json:"eq"`
		CashBal       decimal.NullDecimalV2 `json:"cashBal"`
		IsoEq         decimal.NullDecimalV2 `json:"isoEq,omitempty"`
		AvailEq       decimal.NullDecimalV2 `json:"availEq,omitempty"`
		DisEq         decimal.NullDecimalV2 `json:"disEq"`
		AvailBal      decimal.NullDecimalV2 `json:"availBal"`
		FrozenBal     decimal.NullDecimalV2 `json:"frozenBal"`
		OrdFrozen     decimal.NullDecimalV2 `json:"ordFrozen"`
		Liab          decimal.NullDecimalV2 `json:"liab,omitempty"`
		Upl           decimal.NullDecimalV2 `json:"upl,omitempty"`
		UplLib        decimal.NullDecimalV2 `json:"uplLib,omitempty"`
		CrossLiab     decimal.NullDecimalV2 `json:"crossLiab,omitempty"`
		IsoLiab       decimal.NullDecimalV2 `json:"isoLiab,omitempty"`
		MgnRatio      decimal.NullDecimalV2 `json:"mgnRatio,omitempty"`
		Interest      decimal.NullDecimalV2 `json:"interest,omitempty"`
		Twap          decimal.NullDecimalV2 `json:"twap,omitempty"`
		MaxLoan       decimal.NullDecimalV2 `json:"maxLoan,omitempty"`
		EqUsd         decimal.NullDecimalV2 `json:"eqUsd"`
		NotionalLever decimal.NullDecimalV2 `json:"notionalLever,omitempty"`
		StgyEq        decimal.NullDecimalV2 `json:"stgyEq"`
		IsoUpl        decimal.NullDecimalV2 `json:"isoUpl,omitempty"`
		UTime         okex.JSONTime         `json:"uTime"`
	}
	Position struct {
		InstID      string                `json:"instId"`
		PosCcy      string                `json:"posCcy,omitempty"`
		LiabCcy     string                `json:"liabCcy,omitempty"`
		OptVal      string                `json:"optVal,omitempty"`
		Ccy         string                `json:"ccy"`
		PosID       string                `json:"posId"`
		TradeID     string                `json:"tradeId"`
		Pos         decimal.NullDecimalV2 `json:"pos"`
		AvailPos    decimal.NullDecimalV2 `json:"availPos,omitempty"`
		AvgPx       decimal.NullDecimalV2 `json:"avgPx"`
		Upl         decimal.NullDecimalV2 `json:"upl"`
		UplRatio    decimal.NullDecimalV2 `json:"uplRatio"`
		Lever       decimal.NullDecimalV2 `json:"lever"`
		LiqPx       decimal.NullDecimalV2 `json:"liqPx,omitempty"`
		Imr         decimal.NullDecimalV2 `json:"imr,omitempty"`
		Margin      decimal.NullDecimalV2 `json:"margin,omitempty"`
		MgnRatio    decimal.NullDecimalV2 `json:"mgnRatio"`
		Mmr         decimal.NullDecimalV2 `json:"mmr"`
		Liab        decimal.NullDecimalV2 `json:"liab,omitempty"`
		Interest    decimal.NullDecimalV2 `json:"interest"`
		NotionalUsd decimal.NullDecimalV2 `json:"notionalUsd"`
		ADL         decimal.NullDecimalV2 `json:"adl"`
		Last        decimal.NullDecimalV2 `json:"last"`
		DeltaBS     decimal.NullDecimalV2 `json:"deltaBS"`
		DeltaPA     decimal.NullDecimalV2 `json:"deltaPA"`
		GammaBS     decimal.NullDecimalV2 `json:"gammaBS"`
		GammaPA     decimal.NullDecimalV2 `json:"gammaPA"`
		ThetaBS     decimal.NullDecimalV2 `json:"thetaBS"`
		ThetaPA     decimal.NullDecimalV2 `json:"thetaPA"`
		VegaBS      decimal.NullDecimalV2 `json:"vegaBS"`
		VegaPA      decimal.NullDecimalV2 `json:"vegaPA"`
		PosSide     okex.PositionSide     `json:"posSide"`
		MgnMode     okex.MarginMode       `json:"mgnMode"`
		InstType    okex.InstrumentType   `json:"instType"`
		CTime       okex.JSONTime         `json:"cTime"`
		UTime       okex.JSONTime         `json:"uTime"`
	}
	BalanceAndPosition struct {
		EventType okex.EventType    `json:"eventType"`
		PTime     okex.JSONTime     `json:"pTime"`
		UTime     okex.JSONTime     `json:"uTime"`
		PosData   []*Position       `json:"posData"`
		BalData   []*BalanceDetails `json:"balData"`
	}
	PositionAndAccountRisk struct {
		AdjEq   decimal.NullDecimalV2                `json:"adjEq,omitempty"`
		BalData []*PositionAndAccountRiskBalanceData `json:"balData"`
		PosData []*PositionAndAccountRiskBalanceData `json:"posData"`
		TS      okex.JSONTime                        `json:"ts"`
	}
	PositionAndAccountRiskBalanceData struct {
		Ccy   string                `json:"ccy"`
		Eq    decimal.NullDecimalV2 `json:"eq"`
		DisEq decimal.NullDecimalV2 `json:"disEq"`
	}
	PositionAndAccountRiskPositionData struct {
		InstID      string                `json:"instId"`
		PosCcy      string                `json:"posCcy,omitempty"`
		Ccy         string                `json:"ccy"`
		NotionalCcy decimal.NullDecimalV2 `json:"notionalCcy"`
		Pos         decimal.NullDecimalV2 `json:"pos"`
		NotionalUsd decimal.NullDecimalV2 `json:"notionalUsd"`
		PosSide     okex.PositionSide     `json:"posSide"`
		InstType    okex.InstrumentType   `json:"instType"`
		MgnMode     okex.MarginMode       `json:"mgnMode"`
	}
	Bill struct {
		Ccy       string                `json:"ccy"`
		InstID    string                `json:"instId"`
		Notes     string                `json:"notes"`
		BillID    string                `json:"billId"`
		OrdID     string                `json:"ordId"`
		BalChg    decimal.NullDecimalV2 `json:"balChg"`
		PosBalChg decimal.NullDecimalV2 `json:"posBalChg"`
		Bal       decimal.NullDecimalV2 `json:"bal"`
		PosBal    decimal.NullDecimalV2 `json:"posBal"`
		Sz        decimal.NullDecimalV2 `json:"sz"`
		Pnl       decimal.NullDecimalV2 `json:"pnl"`
		Fee       decimal.NullDecimalV2 `json:"fee"`
		From      okex.AccountType      `json:"from,string"`
		To        okex.AccountType      `json:"to,string"`
		InstType  okex.InstrumentType   `json:"instType"`
		MgnMode   okex.MarginMode       `json:"MgnMode"`
		Type      okex.BillType         `json:"type,string"`
		SubType   okex.BillSubType      `json:"subType,string"`
		TS        okex.JSONTime         `json:"ts"`
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
		InstID  string                `json:"instId"`
		Lever   decimal.NullDecimalV2 `json:"lever"`
		MgnMode okex.MarginMode       `json:"mgnMode"`
		PosSide okex.PositionSide     `json:"posSide"`
	}
	MaxBuySellAmount struct {
		InstID  string                `json:"instId"`
		Ccy     string                `json:"ccy"`
		MaxBuy  decimal.NullDecimalV2 `json:"maxBuy"`
		MaxSell decimal.NullDecimalV2 `json:"maxSell"`
	}
	MaxAvailableTradeAmount struct {
		InstID    string                `json:"instId"`
		AvailBuy  decimal.NullDecimalV2 `json:"availBuy"`
		AvailSell decimal.NullDecimalV2 `json:"availSell"`
	}
	MarginBalanceAmount struct {
		InstID  string                `json:"instId"`
		Amt     decimal.NullDecimalV2 `json:"amt"`
		PosSide okex.PositionSide     `json:"posSide,string"`
		Type    okex.CountAction      `json:"type,string"`
	}
	Loan struct {
		InstID  string                `json:"instId"`
		MgnCcy  string                `json:"mgnCcy"`
		Ccy     string                `json:"ccy"`
		MaxLoan decimal.NullDecimalV2 `json:"maxLoan"`
		MgnMode okex.MarginMode       `json:"mgnMode"`
		Side    okex.OrderSide        `json:"side,string"`
	}
	Fee struct {
		Level    string                `json:"level"`
		Taker    decimal.NullDecimalV2 `json:"taker"`
		Maker    decimal.NullDecimalV2 `json:"maker"`
		Delivery decimal.NullDecimalV2 `json:"delivery,omitempty"`
		Exercise decimal.NullDecimalV2 `json:"exercise,omitempty"`
		Category okex.FeeCategory      `json:"category,string"`
		InstType okex.InstrumentType   `json:"instType"`
		TS       okex.JSONTime         `json:"ts"`
	}
	InterestAccrued struct {
		InstID       string                `json:"instId"`
		Ccy          string                `json:"ccy"`
		Interest     decimal.NullDecimalV2 `json:"interest"`
		InterestRate decimal.NullDecimalV2 `json:"interestRate"`
		Liab         decimal.NullDecimalV2 `json:"liab"`
		MgnMode      okex.MarginMode       `json:"mgnMode"`
		TS           okex.JSONTime         `json:"ts"`
	}
	InterestRate struct {
		Ccy          string                `json:"ccy"`
		InterestRate decimal.NullDecimalV2 `json:"interestRate"`
	}
	Greek struct {
		GreeksType string `json:"greeksType"`
	}
	MaxWithdrawal struct {
		Ccy   string                `json:"ccy"`
		MaxWd decimal.NullDecimalV2 `json:"maxWd"`
	}
)
