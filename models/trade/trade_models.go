package trade

import (
	"github.com/ward-cap/go-okx"
	"github.com/ward-cap/go-okx/decimal"
)

type (
	PlaceOrder struct {
		ClOrdID string                `json:"clOrdId"`
		Tag     string                `json:"tag"`
		SMsg    string                `json:"sMsg"`
		SCode   decimal.NullDecimalV2 `json:"sCode"`
		OrdID   decimal.NullDecimalV2 `json:"ordId"`
	}
	CancelOrder struct {
		OrdID   string                `json:"ordId"`
		ClOrdID string                `json:"clOrdId"`
		SMsg    string                `json:"sMsg"`
		SCode   decimal.NullDecimalV2 `json:"sCode"`
	}
	AmendOrder struct {
		OrdID   string                `json:"ordId"`
		ClOrdID string                `json:"clOrdId"`
		ReqID   string                `json:"reqId"`
		SMsg    string                `json:"sMsg"`
		SCode   decimal.NullDecimalV2 `json:"sCode"`
	}
	ClosePosition struct {
		InstID  string            `json:"instId"`
		PosSide okex.PositionSide `json:"posSide"`
	}
	Order struct {
		InstID      string                `json:"instId"`
		Ccy         string                `json:"ccy"`
		OrdID       string                `json:"ordId"`
		ClOrdID     string                `json:"clOrdId"`
		TradeID     string                `json:"tradeId"`
		Tag         string                `json:"tag"`
		Category    string                `json:"category"`
		FeeCcy      string                `json:"feeCcy"`
		RebateCcy   string                `json:"rebateCcy"`
		Px          decimal.NullDecimalV2 `json:"px"`
		Sz          decimal.NullDecimalV2 `json:"sz"`
		Pnl         decimal.NullDecimalV2 `json:"pnl"`
		AccFillSz   decimal.NullDecimalV2 `json:"accFillSz"`
		FillPx      decimal.NullDecimalV2 `json:"fillPx"`
		FillSz      decimal.NullDecimalV2 `json:"fillSz"`
		FillTime    decimal.NullDecimalV2 `json:"fillTime"`
		AvgPx       decimal.NullDecimalV2 `json:"avgPx"`
		Lever       decimal.NullDecimalV2 `json:"lever"`
		TpTriggerPx decimal.NullDecimalV2 `json:"tpTriggerPx"`
		TpOrdPx     decimal.NullDecimalV2 `json:"tpOrdPx"`
		SlTriggerPx decimal.NullDecimalV2 `json:"slTriggerPx"`
		SlOrdPx     decimal.NullDecimalV2 `json:"slOrdPx"`
		Fee         decimal.NullDecimalV2 `json:"fee"`
		Rebate      decimal.NullDecimalV2 `json:"rebate"`
		State       okex.OrderState       `json:"state"`
		TdMode      okex.TradeMode        `json:"tdMode"`
		PosSide     okex.PositionSide     `json:"posSide"`
		Side        okex.OrderSide        `json:"side"`
		OrdType     okex.OrderType        `json:"ordType"`
		InstType    okex.InstrumentType   `json:"instType"`
		TgtCcy      okex.QuantityType     `json:"tgtCcy"`
		UTime       okex.JSONTime         `json:"uTime"`
		CTime       okex.JSONTime         `json:"cTime"`
	}
	TransactionDetail struct {
		InstID   string                `json:"instId"`
		OrdID    string                `json:"ordId"`
		TradeID  string                `json:"tradeId"`
		ClOrdID  string                `json:"clOrdId"`
		BillID   string                `json:"billId"`
		Tag      decimal.NullDecimalV2 `json:"tag"`
		FillPx   decimal.NullDecimalV2 `json:"fillPx"`
		FillSz   decimal.NullDecimalV2 `json:"fillSz"`
		FeeCcy   string                `json:"feeCcy"`
		FillTime int64                 `json:"fillTime,string"`
		Fee      decimal.NullDecimalV2 `json:"fee"`
		InstType okex.InstrumentType   `json:"instType"`
		Side     okex.OrderSide        `json:"side"`
		PosSide  okex.PositionSide     `json:"posSide"`
		ExecType okex.OrderFlowType    `json:"execType"`
		TS       okex.JSONTime         `json:"ts"`
	}
	PlaceAlgoOrder struct {
		AlgoID string                `json:"algoId"`
		SMsg   string                `json:"sMsg"`
		SCode  decimal.NullDecimalV2 `json:"sCode"`
	}
	CancelAlgoOrder struct {
		AlgoID string                `json:"algoId"`
		SMsg   string                `json:"sMsg"`
		SCode  decimal.NullDecimalV2 `json:"sCode"`
	}
	AlgoOrder struct {
		InstID       string                `json:"instId"`
		Ccy          string                `json:"ccy"`
		OrdID        string                `json:"ordId"`
		AlgoID       string                `json:"algoId"`
		ClOrdID      string                `json:"clOrdId"`
		TradeID      string                `json:"tradeId"`
		Tag          string                `json:"tag"`
		Category     string                `json:"category"`
		FeeCcy       string                `json:"feeCcy"`
		RebateCcy    string                `json:"rebateCcy"`
		TimeInterval string                `json:"timeInterval"`
		Px           decimal.NullDecimalV2 `json:"px"`
		PxVar        decimal.NullDecimalV2 `json:"pxVar"`
		PxSpread     decimal.NullDecimalV2 `json:"pxSpread"`
		PxLimit      decimal.NullDecimalV2 `json:"pxLimit"`
		Sz           decimal.NullDecimalV2 `json:"sz"`
		SzLimit      decimal.NullDecimalV2 `json:"szLimit"`
		ActualSz     decimal.NullDecimalV2 `json:"actualSz"`
		ActualPx     decimal.NullDecimalV2 `json:"actualPx"`
		Pnl          decimal.NullDecimalV2 `json:"pnl"`
		AccFillSz    decimal.NullDecimalV2 `json:"accFillSz"`
		FillPx       decimal.NullDecimalV2 `json:"fillPx"`
		FillSz       decimal.NullDecimalV2 `json:"fillSz"`
		FillTime     decimal.NullDecimalV2 `json:"fillTime"`
		AvgPx        decimal.NullDecimalV2 `json:"avgPx"`
		Lever        decimal.NullDecimalV2 `json:"lever"`
		TpTriggerPx  decimal.NullDecimalV2 `json:"tpTriggerPx"`
		TpOrdPx      decimal.NullDecimalV2 `json:"tpOrdPx"`
		SlTriggerPx  decimal.NullDecimalV2 `json:"slTriggerPx"`
		SlOrdPx      decimal.NullDecimalV2 `json:"slOrdPx"`
		OrdPx        decimal.NullDecimalV2 `json:"ordPx"`
		Fee          decimal.NullDecimalV2 `json:"fee"`
		Rebate       decimal.NullDecimalV2 `json:"rebate"`
		State        okex.OrderState       `json:"state"`
		TdMode       okex.TradeMode        `json:"tdMode"`
		ActualSide   okex.PositionSide     `json:"actualSide"`
		PosSide      okex.PositionSide     `json:"posSide"`
		Side         okex.OrderSide        `json:"side"`
		OrdType      okex.AlgoOrderType    `json:"ordType"`
		InstType     okex.InstrumentType   `json:"instType"`
		TgtCcy       okex.QuantityType     `json:"tgtCcy"`
		CTime        okex.JSONTime         `json:"cTime"`
		TriggerTime  okex.JSONTime         `json:"triggerTime"`
	}
)
