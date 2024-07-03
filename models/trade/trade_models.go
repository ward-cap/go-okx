package trade

import (
	"github.com/shopspring/decimal"
	"github.com/ward-cap/go-okx"
)

type (
	PlaceOrder struct {
		ClOrdID string              `json:"clOrdId"`
		Tag     string              `json:"tag"`
		SMsg    string              `json:"sMsg"`
		SCode   decimal.NullDecimal `json:"sCode"`
		OrdID   decimal.NullDecimal `json:"ordId"`
	}
	CancelOrder struct {
		OrdID   string              `json:"ordId"`
		ClOrdID string              `json:"clOrdId"`
		SMsg    string              `json:"sMsg"`
		SCode   decimal.NullDecimal `json:"sCode"`
	}
	AmendOrder struct {
		OrdID   string              `json:"ordId"`
		ClOrdID string              `json:"clOrdId"`
		ReqID   string              `json:"reqId"`
		SMsg    string              `json:"sMsg"`
		SCode   decimal.NullDecimal `json:"sCode"`
	}
	ClosePosition struct {
		InstID  string            `json:"instId"`
		PosSide okex.PositionSide `json:"posSide"`
	}
	Order struct {
		InstID      string              `json:"instId"`
		Ccy         string              `json:"ccy"`
		OrdID       string              `json:"ordId"`
		ClOrdID     string              `json:"clOrdId"`
		TradeID     string              `json:"tradeId"`
		Tag         string              `json:"tag"`
		Category    string              `json:"category"`
		FeeCcy      string              `json:"feeCcy"`
		RebateCcy   string              `json:"rebateCcy"`
		Px          decimal.NullDecimal `json:"px"`
		Sz          decimal.NullDecimal `json:"sz"`
		Pnl         decimal.NullDecimal `json:"pnl"`
		AccFillSz   decimal.NullDecimal `json:"accFillSz"`
		FillPx      decimal.NullDecimal `json:"fillPx"`
		FillSz      decimal.NullDecimal `json:"fillSz"`
		FillTime    decimal.NullDecimal `json:"fillTime"`
		AvgPx       decimal.NullDecimal `json:"avgPx"`
		Lever       decimal.NullDecimal `json:"lever"`
		TpTriggerPx decimal.NullDecimal `json:"tpTriggerPx"`
		TpOrdPx     decimal.NullDecimal `json:"tpOrdPx"`
		SlTriggerPx decimal.NullDecimal `json:"slTriggerPx"`
		SlOrdPx     decimal.NullDecimal `json:"slOrdPx"`
		Fee         decimal.NullDecimal `json:"fee"`
		Rebate      decimal.NullDecimal `json:"rebate"`
		State       okex.OrderState     `json:"state"`
		TdMode      okex.TradeMode      `json:"tdMode"`
		PosSide     okex.PositionSide   `json:"posSide"`
		Side        okex.OrderSide      `json:"side"`
		OrdType     okex.OrderType      `json:"ordType"`
		InstType    okex.InstrumentType `json:"instType"`
		TgtCcy      okex.QuantityType   `json:"tgtCcy"`
		UTime       okex.JSONTime       `json:"uTime"`
		CTime       okex.JSONTime       `json:"cTime"`
	}
	TransactionDetail struct {
		InstID   string              `json:"instId"`
		OrdID    string              `json:"ordId"`
		TradeID  string              `json:"tradeId"`
		ClOrdID  string              `json:"clOrdId"`
		BillID   string              `json:"billId"`
		Tag      decimal.NullDecimal `json:"tag"`
		FillPx   decimal.NullDecimal `json:"fillPx"`
		FillSz   decimal.NullDecimal `json:"fillSz"`
		FeeCcy   decimal.NullDecimal `json:"feeCcy"`
		Fee      decimal.NullDecimal `json:"fee"`
		InstType okex.InstrumentType `json:"instType"`
		Side     okex.OrderSide      `json:"side"`
		PosSide  okex.PositionSide   `json:"posSide"`
		ExecType okex.OrderFlowType  `json:"execType"`
		TS       okex.JSONTime       `json:"ts"`
	}
	PlaceAlgoOrder struct {
		AlgoID string              `json:"algoId"`
		SMsg   string              `json:"sMsg"`
		SCode  decimal.NullDecimal `json:"sCode"`
	}
	CancelAlgoOrder struct {
		AlgoID string              `json:"algoId"`
		SMsg   string              `json:"sMsg"`
		SCode  decimal.NullDecimal `json:"sCode"`
	}
	AlgoOrder struct {
		InstID       string              `json:"instId"`
		Ccy          string              `json:"ccy"`
		OrdID        string              `json:"ordId"`
		AlgoID       string              `json:"algoId"`
		ClOrdID      string              `json:"clOrdId"`
		TradeID      string              `json:"tradeId"`
		Tag          string              `json:"tag"`
		Category     string              `json:"category"`
		FeeCcy       string              `json:"feeCcy"`
		RebateCcy    string              `json:"rebateCcy"`
		TimeInterval string              `json:"timeInterval"`
		Px           decimal.NullDecimal `json:"px"`
		PxVar        decimal.NullDecimal `json:"pxVar"`
		PxSpread     decimal.NullDecimal `json:"pxSpread"`
		PxLimit      decimal.NullDecimal `json:"pxLimit"`
		Sz           decimal.NullDecimal `json:"sz"`
		SzLimit      decimal.NullDecimal `json:"szLimit"`
		ActualSz     decimal.NullDecimal `json:"actualSz"`
		ActualPx     decimal.NullDecimal `json:"actualPx"`
		Pnl          decimal.NullDecimal `json:"pnl"`
		AccFillSz    decimal.NullDecimal `json:"accFillSz"`
		FillPx       decimal.NullDecimal `json:"fillPx"`
		FillSz       decimal.NullDecimal `json:"fillSz"`
		FillTime     decimal.NullDecimal `json:"fillTime"`
		AvgPx        decimal.NullDecimal `json:"avgPx"`
		Lever        decimal.NullDecimal `json:"lever"`
		TpTriggerPx  decimal.NullDecimal `json:"tpTriggerPx"`
		TpOrdPx      decimal.NullDecimal `json:"tpOrdPx"`
		SlTriggerPx  decimal.NullDecimal `json:"slTriggerPx"`
		SlOrdPx      decimal.NullDecimal `json:"slOrdPx"`
		OrdPx        decimal.NullDecimal `json:"ordPx"`
		Fee          decimal.NullDecimal `json:"fee"`
		Rebate       decimal.NullDecimal `json:"rebate"`
		State        okex.OrderState     `json:"state"`
		TdMode       okex.TradeMode      `json:"tdMode"`
		ActualSide   okex.PositionSide   `json:"actualSide"`
		PosSide      okex.PositionSide   `json:"posSide"`
		Side         okex.OrderSide      `json:"side"`
		OrdType      okex.AlgoOrderType  `json:"ordType"`
		InstType     okex.InstrumentType `json:"instType"`
		TgtCcy       okex.QuantityType   `json:"tgtCcy"`
		CTime        okex.JSONTime       `json:"cTime"`
		TriggerTime  okex.JSONTime       `json:"triggerTime"`
	}
)
