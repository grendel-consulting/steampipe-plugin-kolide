package kolide_client

type PersonGroupListResponse struct {
	PersonGroups []PersonGroup `json:"data"`
	Pagination   Pagination    `json:"pagination"`
}
type PersonGroup struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (c *Client) GetPersonGroups(cursor string, limit int32, searches ...Search) (interface{}, error) {
	return c.fetchCollection("/person_groups/", cursor, limit, searches, new(PersonGroupListResponse))
}

func (c *Client) GetPersonGroupById(id string) (interface{}, error) {
	return c.fetchResource("/person_groups/", id, new(PersonGroup))
}
