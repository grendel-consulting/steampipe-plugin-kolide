package kolide_client

import (
	"time"
)

type DeprovisionedPeopleListResponse struct {
	DeprovisionedPeople []DeprovisionedPerson `json:"data"`
	Pagination          Pagination            `json:"pagination"`
}
type DeprovisionedPerson struct {
	Id                  string    `json:"id"`
	Name                string    `json:"name"`
	Email               string    `json:"email,omitempty"`
	CreatedAt           time.Time `json:"registered_at,omitempty"`
	LastAuthenticatedAt time.Time `json:"last_authenticated_at,omitempty"`
	HasRegisteredDevice bool      `json:"has_registered_device"`
}

func (c *Client) GetDeprovisionedPeople(cursor string, limit int32, searches ...Search) (interface{}, error) {
	return c.fetchCollection("/deprovisioned_people/", cursor, limit, searches, new(DeprovisionedPeopleListResponse))
}
