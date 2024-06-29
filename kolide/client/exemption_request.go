package kolide_client

import "time"

type ExemptionRequestListResponse struct {
	ExemptionRequests []ExemptionRequest `json:"data"`
	Pagination        Pagination         `json:"pagination"`
}
type ExemptionRequest struct {
	Id                   string               `json:"id"`
	Status               string               `json:"status"`
	RequesterMessage     string               `json:"requester_message,omitempty"`
	InternalExplanation  string               `json:"internal_explanation,omitempty"`
	RequestedAt          time.Time            `json:"requested_at,omitempty"`
	DenialExplanation    string               `json:"denial_explanation,omitempty"`
	RequesterInformation RequesterInformation `json:"requester_information,omitempty"`
	DeviceInformation    DeviceInformation    `json:"device_information,omitempty"`
	Issues               []IssueInformation   `json:"issues,omitempty"`
}

type IssueInformation struct {
	Identifier string `json:"identifier,omitempty"`
}

func (c *Client) GetExemptionRequests(cursor string, limit int32, searches ...Search) (interface{}, error) {
	return c.fetchCollection("/exemption_requests/", cursor, limit, searches, new(ExemptionRequestListResponse))
}

func (c *Client) GetExemptionRequestById(id string) (interface{}, error) {
	return c.fetchResource("/exemption_requests/", id, new(ExemptionRequest))
}
