package kolide_client

import (
	"time"
)

type PeopleListResponse struct {
	People     []Person   `json:"data"`
	Pagination Pagination `json:"pagination"`
}
type Person struct {
	Id                  string    `json:"id"`
	Name                string    `json:"name"`
	Email               string    `json:"email"`
	CreatedAt           time.Time `json:"created_at"`
	LastAuthenticatedAt time.Time `json:"last_authenticated_at,omitempty"`
	HasRegisteredDevice bool      `json:"has_registered_device"`
	Usernames           []int     `json:"usernames,omitempty"`
}

func (c *Client) GetPeople(cursor string, limit int32, searches ...Search) (interface{}, error) {
	return c.fetchCollection("/people/", cursor, limit, searches, new(PeopleListResponse))
}

func (c *Client) GetPersonById(id string) (interface{}, error) {
	return c.fetchResource("/people/", id, new(Person))
}

// For /people/{personId}/open_issues endpoint, see issues.go
// For /people/{personId}/registered_devices endpoint, see devices.go
