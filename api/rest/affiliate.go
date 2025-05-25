package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ward-cap/go-okx/responses/affiliate"
	"net/http"
)

type Affiliate struct {
	client *ClientRest
}

func NewAffiliate(c *ClientRest) *Affiliate {
	return &Affiliate{c}
}

func (c *Affiliate) GetInvitees(ctx context.Context, uid string) (t affiliate.InviteeResponse, err error) {
	p := "/api/v5/affiliate/invitee/detail"

	if uid != "" {
		p += fmt.Sprintf("?uid=%s", uid)
	}

	res, err := c.client.DoWithContext(ctx, http.MethodGet, p, true)
	if err != nil {
		return
	}

	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&t)

	return
}

func (c *Affiliate) IsMyRefer(ctx context.Context, apiKey string) (t affiliate.IsReferResponse, err error) {
	if len(apiKey) == 0 {
		return affiliate.IsReferResponse{}, errors.New("api key is empty")
	}

	p := "/api/v5/users/partner/if-rebate"

	res, err := c.client.DoWithContext(ctx, http.MethodGet, p, true, map[string]string{"apiKey": apiKey})
	if err != nil {
		return
	}

	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&t)

	return
}
