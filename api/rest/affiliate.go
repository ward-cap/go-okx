package rest

import (
	"context"
	"encoding/json"
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
