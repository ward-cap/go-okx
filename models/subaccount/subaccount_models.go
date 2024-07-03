package subaccount

import (
	"github.com/shopspring/decimal"
	"github.com/ward-cap/go-okx"
)

type (
	SubAccount struct {
		SubAcct string        `json:"subAcct,omitempty"`
		Label   string        `json:"label,omitempty"`
		Mobile  string        `json:"mobile,omitempty"`
		GAuth   bool          `json:"gAuth"`
		Enable  bool          `json:"enable"`
		TS      okex.JSONTime `json:"ts"`
	}
	APIKey struct {
		SubAcct    string        `json:"subAcct,omitempty"`
		Label      string        `json:"label,omitempty"`
		APIKey     string        `json:"apiKey,omitempty"`
		SecretKey  string        `json:"secretKey,omitempty"`
		Passphrase string        `json:"Passphrase,omitempty"`
		Perm       string        `json:"perm,omitempty"`
		IP         string        `json:"ip,omitempty"`
		TS         okex.JSONTime `json:"ts,omitempty"`
	}
	HistoryTransfer struct {
		SubAcct string              `json:"subAcct,omitempty"`
		Ccy     string              `json:"ccy,omitempty"`
		BillID  decimal.NullDecimal `json:"billId,omitempty"`
		Type    okex.BillType       `json:"type,omitempty,string"`
		TS      okex.JSONTime       `json:"ts,omitempty"`
	}
	Transfer struct {
		TransID decimal.NullDecimal `json:"transId"`
	}
)
