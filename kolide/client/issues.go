package kolide_client

import (
	"encoding/json"
	"time"
)

type IssueListResponse struct {
	Issues     []Issue    `json:"data"`
	Pagination Pagination `json:"pagination"`
}
type Issue struct {
	Id                string            `json:"id"`
	IssueKey          string            `json:"issue_key,omitempty"`
	IssueValue        string            `json:"issue_value,omitempty"`
	Title             string            `json:"title"`
	Value             json.RawMessage   `json:"value,omitempty"`
	Exempted          bool              `json:"exempted"`
	ResolvedAt        time.Time         `json:"resolved_at,omitempty"`
	DetectedAt        time.Time         `json:"detected_at"`
	BlocksDeviceAt    time.Time         `json:"blocks_device_at,omitempty"`
	DeviceInformation DeviceInformation `json:"device_information,omitempty"`
	CheckInformation  CheckInformation  `json:"check_information,omitempty"`
	LastRecheckedAt   time.Time         `json:"last_rechecked_at,omitempty"`
}

type DeviceInformation struct {
	Identifier string `json:"identifier,omitempty"`
}

type CheckInformation struct {
	Identifier string `json:"identifier,omitempty"`
}

func (c *Client) GetIssues(cursor string, limit int32, searches ...Search) (interface{}, error) {
	return c.fetchCollection("/issues/", cursor, limit, searches, new(IssueListResponse))
}

func (c *Client) GetIssueById(id string) (interface{}, error) {
	return c.fetchResource("/issues/", id, new(Issue))
}

func (c *Client) GetIssuesByDevice(id string, cursor string, limit int32, searches ...Search) (interface{}, error) {
	return c.fetchCollectionWithResourceId("/devices/{resourceId}/open_issues", id, cursor, limit, searches, new(IssueListResponse))
}

func (c *Client) GetIssuesByPerson(id string, cursor string, limit int32, searches ...Search) (interface{}, error) {
	return c.fetchCollectionWithResourceId("/people/{resourceId}/open_issues", id, cursor, limit, searches, new(IssueListResponse))
}
