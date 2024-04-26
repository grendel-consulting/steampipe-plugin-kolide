package kolide_client

import (
	"time"
)

type DeviceListResponse struct {
	Devices    []Device   `json:"data"`
	Pagination Pagination `json:"pagination"`
}
type Device struct {
	Id                  string          `json:"id"`
	Name                string          `json:"name"`
	RegisteredAt        time.Time       `json:"registered_at,omitempty"`
	LastAuthenticatedAt time.Time       `json:"last_authenticated_at,omitempty"`
	RegisteredOwnerInfo RegisteredOwner `json:"registered_owner_info,omitempty"`
	OperatingSystem     string          `json:"operating_system"`
	HardwareModel       string          `json:"hardware_model"`
	Serial              string          `json:"serial,omitempty"`
	HardwareUuid        string          `json:"hardware_uuid,omitempty"`
	Note                string          `json:"note,omitempty"`
	AuthState           string          `json:"auth_state"`
	WillBlockAt         time.Time       `json:"will_block_at,omitempty"`
	ProductImageUrl     string          `json:"product_image_url"`
	DeviceType          string          `json:"device_type"`
	FormFactor          string          `json:"form_factor"`
}

type RegisteredOwner struct {
	Identifier string `json:"identifier,omitempty"`
}

type AuthConfiguration struct {
	DeviceId           string        `json:"id"`
	AuthenticationMode string        `json:"authentication_mode"`
	PersonGroups       []PersonGroup `json:"person_groups"`
}

func (c *Client) GetDevices(cursor string, limit int32, searches ...Search) (interface{}, error) {
	return c.fetchCollection("/devices/", cursor, limit, searches, new(DeviceListResponse))
}

func (c *Client) GetDeviceById(id string) (interface{}, error) {
	return c.fetchResource("/devices/", id, new(Device))
}

func (c *Client) GetIssuesByDevice(id string, cursor string, limit int32, searches ...Search) (interface{}, error) {
	return c.fetchCollectionWithResourceId("/devices/{resourceId}/open_issues", id, cursor, limit, searches, new(IssueListResponse))
}
