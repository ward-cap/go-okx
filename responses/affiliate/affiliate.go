package affiliate

import (
	okex "github.com/ward-cap/go-okx"
	"github.com/ward-cap/go-okx/decimal"
	"github.com/ward-cap/go-okx/responses"
)

type InviteeData struct {
	AccFee            decimal.NullDecimalV2 `json:"accFee"`
	AffiliateCode     string                `json:"affiliateCode"`
	DepAmt            decimal.NullDecimalV2 `json:"depAmt"`
	FirstTradeTime    okex.JSONTime         `json:"firstTradeTime"`
	InviteeLevel      string                `json:"inviteeLevel"`
	InviteeRebateRate decimal.NullDecimalV2 `json:"inviteeRebateRate"`
	JoinTime          okex.JSONTime         `json:"joinTime"`
	KycTime           okex.JSONTime         `json:"kycTime"`
	Level             string                `json:"level"`
	Region            string                `json:"region"`
	TotalCommission   decimal.NullDecimalV2 `json:"totalCommission"`
	VolMonth          decimal.NullDecimalV2 `json:"volMonth"`
}

type InviteeResponse struct {
	responses.Basic
	Data []InviteeData `json:"data"`
}

type IsReferResponse struct {
	responses.Basic
	Data struct {
		Result bool   `json:"result"`
		Type   string `json:"type"`
	} `json:"data"`
}
