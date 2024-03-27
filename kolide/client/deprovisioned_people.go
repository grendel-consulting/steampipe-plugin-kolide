package kolide_client

import (
	"fmt"
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
	params := make(map[string]string)
	params["query"] = serializeSearches(searches)

	if cursor != "" {
		params["per_page"] = string(limit)
		params["cursor"] = cursor
	}

	res, err := c.r().SetQueryParams(params).Get("/deprovisioned_people/")

	if err != nil {
		return nil, fmt.Errorf("error retrieving deprovisioned people: %q", err)
	}

	if !res.IsSuccessState() {
		return nil, fmt.Errorf("error retrieving deprovisioned people: %q", res.Status)
	}
	var response DeprovisionedPeopleListResponse

	err = res.UnmarshalJson(&response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %q", err)
	}

	return &response, nil
}
