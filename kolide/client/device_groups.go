package kolide_client

import (
	"time"
)

type DeviceGroupListResponse struct {
	DeviceGroups []DeviceGroup `json:"data"`
	Pagination   Pagination    `json:"pagination"`
}
type DeviceGroup struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	Description  string    `json:"description,omitempty"`
	MembersCount int       `json:"members_count"`
}

func (c *Client) GetDeviceGroups(cursor string, limit int32, searches ...Search) (interface{}, error) {
	return c.fetchCollection("/device_groups/", cursor, limit, searches, new(DeviceGroupListResponse))
}

func (c *Client) GetDeviceGroupById(id string) (interface{}, error) {
	return c.fetchResource("/device_groups/", id, new(DeviceGroup))
}

// For /device_groups/{deviceGroupId}/devices endpoint, see devices.go
