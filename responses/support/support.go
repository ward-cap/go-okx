package support

import "github.com/ward-cap/go-okx/responses"

type AnnouncementsResponse struct {
	responses.Basic
	Data []AnnouncementData `json:"data"`
}

type AnnouncementData struct {
	Details   []AnnouncementDetail `json:"details"`
	TotalPage string               `json:"totalPage"`
}

type AnnouncementDetail struct {
	AnnType string `json:"annType"`
	PTime   string `json:"pTime"`
	Title   string `json:"title"`
	Url     string `json:"url"`
}
