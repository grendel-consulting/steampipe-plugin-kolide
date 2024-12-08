package kolide_client

import (
	"encoding/json"
	"time"
)

type LiveQueryCampaignListResponse struct {
	LiveQueryCampaigns []LiveQueryCampaign `json:"data"`
	Pagination         Pagination          `json:"pagination"`
}

type LiveQueryCampaign struct {
	Id        string    `json:"id"`
	Query     string    `json:"query"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Client) GetLiveQueryCampaigns(cursor string, limit int32, searches ...Search) (interface{}, error) {
	return c.fetchCollection("/live_query_campaigns/", cursor, limit, searches, new(LiveQueryCampaignListResponse))
}

func (c *Client) GetLiveQueryCampaignById(id string) (interface{}, error) {
	return c.fetchResource("/live_query_campaigns/", id, new(LiveQueryCampaign))
}
