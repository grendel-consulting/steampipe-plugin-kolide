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
	LastAuthenticatedAt time.Time `json:"last_authenticated_at"`
	HasRegisteredDevice bool      `json:"has_registered_device"`
	Usernames           []int     `json:"usernames"`
}

func (c *Client) GetPeople(cursor string, limit int32, searches ...Search) (interface{}, error) {
	return c.fetchCollection("/people/", cursor, limit, searches, new(PeopleListResponse))
}

func (c *Client) GetPersonById(id string) (interface{}, error) {
	return c.fetchResource("/people/", id, new(Person))
}

func (c *Client) GetDevicesByPerson(id string, cursor string, limit int32, searches ...Search) (interface{}, error) {
	return c.fetchCollectionWithResourceId("/people/{resourceId}/registered_devices", id, cursor, limit, searches, new(IssueListResponse))
}
