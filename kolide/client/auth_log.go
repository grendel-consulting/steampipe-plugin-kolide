package kolide_client

import (
	"time"
)

type AuthLogListResponse struct {
	AuthLogs   []AuthLog  `json:"data"`
	Pagination Pagination `json:"pagination"`
}
type AuthLog struct {
	Id               string                     `json:"id"`
	Timestamp        time.Time                  `json:"timestamp"`
	PersonName       string                     `json:"person_name"`
	PersonEmail      string                     `json:"person_email"`
	PersonInfo       PersonInformation          `json:"person_info"`
	DeviceInfo       DeviceInformation          `json:"device_information,omitempty"`
	Result           string                     `json:"result"`
	InitialStatus    string                     `json:"initial_status"`
	IpAddress        string                     `json:"ip_address"`
	AgentVersion     string                     `json:"agent_version,omitempty"`
	Country          string                     `json:"country,omitempty"`
	City             string                     `json:"city,omitempty"`
	BrowserName      string                     `json:"browser_name"`
	BrowserUserAgent string                     `json:"browser_user_agent"`
	IssuesDisplayed  []DetailedIssueInformation `json:"issues_displayed"`
	Events           []EventInformation         `json:"events"`
}

type PersonInformation struct {
	Identifier string `json:"identifier,omitempty"`
}

type DetailedIssueInformation struct {
	Title          string `json:"title,omitempty"`
	BlockingStatus string `json:"blocking_status,omitempty"`
	Identifier     string `json:"identifier,omitempty"`
}

type EventInformation struct {
	Timestamp        time.Time `json:"timestamp"`
	EventType        string    `json:"event_type,omitempty"`
	EventDescription string    `json:"event_description,omitempty"`
}

func (c *Client) GetAuthLogs(cursor string, limit int32, searches ...Search) (interface{}, error) {
	return c.fetchCollection("/auth_logs/", cursor, limit, searches, new(AuthLogListResponse))
}

func (c *Client) GetAuthLogById(id string) (interface{}, error) {
	return c.fetchResource("/auth_logs/", id, new(AuthLog))
}
