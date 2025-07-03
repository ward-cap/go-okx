package funding

import (
	"github.com/ward-cap/go-okx"
	"github.com/ward-cap/go-okx/decimal"
)

type (
	Currency struct {
		Ccy         string `json:"ccy"`
		Name        string `json:"name"`
		Chain       string `json:"chain"`
		MinWd       string `json:"minWd"`
		MinFee      string `json:"minFee"`
		MaxFee      string `json:"maxFee"`
		CanDep      bool   `json:"canDep"`
		CanWd       bool   `json:"canWd"`
		CanInternal bool   `json:"canInternal"`
	}
	Balance struct {
		Ccy       string `json:"ccy"`
		Bal       string `json:"bal"`
		FrozenBal string `json:"frozenBal"`
		AvailBal  string `json:"availBal"`
	}
	Transfer struct {
		TransID string                `json:"transId"`
		Ccy     string                `json:"ccy"`
		Amt     decimal.NullDecimalV2 `json:"amt"`
		From    okex.AccountType      `json:"from,string"`
		To      okex.AccountType      `json:"to,string"`
	}
	Bill struct {
		BillID string                `json:"billId"`
		Ccy    string                `json:"ccy"`
		Bal    decimal.NullDecimalV2 `json:"bal"`
		BalChg decimal.NullDecimalV2 `json:"balChg"`
		Type   okex.BillType         `json:"type,string"`
		TS     okex.JSONTime         `json:"ts"`
	}
	DepositAddress struct {
		Addr     string           `json:"addr"`
		Tag      string           `json:"tag,omitempty"`
		Memo     string           `json:"memo,omitempty"`
		PmtID    string           `json:"pmtId,omitempty"`
		Ccy      string           `json:"ccy"`
		Chain    string           `json:"chain"`
		CtAddr   string           `json:"ctAddr"`
		Selected bool             `json:"selected"`
		To       okex.AccountType `json:"to,string"`
		TS       okex.JSONTime    `json:"ts"`
	}
	DepositHistory struct {
		Ccy   string                `json:"ccy"`
		Chain string                `json:"chain"`
		TxID  string                `json:"txId"`
		From  string                `json:"from"`
		To    string                `json:"to"`
		DepId string                `json:"depId"`
		Amt   decimal.NullDecimalV2 `json:"amt"`
		State okex.DepositState     `json:"state,string"`
		TS    okex.JSONTime         `json:"ts"`
	}
	Withdrawal struct {
		Ccy   string                `json:"ccy"`
		Chain string                `json:"chain"`
		WdID  decimal.NullDecimalV2 `json:"wdId"`
		Amt   decimal.NullDecimalV2 `json:"amt"`
	}
	WithdrawalHistory struct {
		Ccy   string                `json:"ccy"`
		Chain string                `json:"chain"`
		TxID  string                `json:"txId"`
		From  string                `json:"from"`
		To    string                `json:"to"`
		Tag   string                `json:"tag,omitempty"`
		PmtID string                `json:"pmtId,omitempty"`
		Memo  string                `json:"memo,omitempty"`
		Amt   decimal.NullDecimalV2 `json:"amt"`
		Fee   decimal.NullDecimalV2 `json:"fee"`
		WdID  decimal.NullDecimalV2 `json:"wdId"`
		State okex.WithdrawalState  `json:"state,string"`
		TS    okex.JSONTime         `json:"ts"`
	}
	PiggyBank struct {
		Ccy  string                `json:"ccy"`
		Amt  decimal.NullDecimalV2 `json:"amt"`
		Side okex.ActionType       `json:"side,string"`
	}
	PiggyBankBalance struct {
		Ccy      string                `json:"ccy"`
		Amt      decimal.NullDecimalV2 `json:"amt"`
		Earnings decimal.NullDecimalV2 `json:"earnings"`
	}
)
