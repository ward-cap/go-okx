package rest

import (
	"context"
	"encoding/json"
	"github.com/ward-cap/go-okx/responses/support"
	"net/http"
)

type Support struct {
	client *ClientRest
}

func NewSupport(c *ClientRest) *Support {
	return &Support{c}
}

func (c *Support) GetAnnouncement(
	ctx context.Context,
	annType, page *string,
) (t support.AnnouncementsResponse, err error) {
	const p = "/api/v5/support/announcements"

	var params = make(map[string]string)
	if annType != nil && *annType != "" {
		params["annType"] = *annType
	}
	if page != nil && *page != "" {
		params["page"] = *page
	}

	res, err := c.client.DoWithContext(ctx, http.MethodGet, p, true, params)
	if err != nil {
		return
	}

	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&t)

	return
}
