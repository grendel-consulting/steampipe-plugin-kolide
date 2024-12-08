package kolide_client

import (
	"time"
)

type LiveQueryCampaignListResponse struct {
	LiveQueryCampaigns []LiveQueryCampaign `json:"data"`
	Pagination         Pagination          `json:"pagination"`
}

type LiveQueryCampaign struct {
	Id                     string    `json:"id"`
	Name                   string    `json:"name"`
	OsquerySql             string    `json:"osquery_sql"`
	CreatedAt              time.Time `json:"created_at"`
	Published              bool      `json:"published"`
	Revision               int       `json:"revision"`
	TablesUsed             []string  `json:"tables_used,omitempty"`
	SuccessfulDevicesCount int       `json:"successful_devices_count"`
	ErroredDevicesCount    int       `json:"errored_devices_count"`
	WaitingDevicesCount    int       `json:"waiting_devices_count"`
}

func (c *Client) GetLiveQueryCampaigns(cursor string, limit int32, searches ...Search) (interface{}, error) {
	return c.fetchCollection("/live_query_campaigns/", cursor, limit, searches, new(LiveQueryCampaignListResponse))
}

func (c *Client) GetLiveQueryCampaignById(id string) (interface{}, error) {
	return c.fetchResource("/live_query_campaigns/", id, new(LiveQueryCampaign))
}
