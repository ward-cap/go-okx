package affiliate

import (
	"github.com/shopspring/decimal"
	okex "github.com/ward-cap/go-okx"
)

type InviteeData struct {
	AccFee            decimal.NullDecimal `json:"accFee"`
	AffiliateCode     string              `json:"affiliateCode"`
	DepAmt            decimal.NullDecimal `json:"depAmt"`
	FirstTradeTime    okex.JSONTime       `json:"firstTradeTime"`
	InviteeLevel      string              `json:"inviteeLevel"`
	InviteeRebateRate decimal.NullDecimal `json:"inviteeRebateRate"`
	JoinTime          okex.JSONTime       `json:"joinTime"`
	KycTime           okex.JSONTime       `json:"kycTime"`
	Level             string              `json:"level"`
	Region            string              `json:"region"`
	TotalCommission   decimal.NullDecimal `json:"totalCommission"`
	VolMonth          decimal.NullDecimal `json:"volMonth"`
}

type InviteeResponse struct {
	Msg  string        `json:"msg"`
	Code string        `json:"code"`
	Data []InviteeData `json:"data"`
}
