package kolide_client

import (
	"fmt"
	"strconv"
	"time"
)

type AdminUserListResponse struct {
	AdminUsers []AdminUser `json:"data"`
	Pagination Pagination  `json:"pagination"`
}
type AdminUser struct {
	Id        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Access    string    `json:"access"`
}

func (c *Client) GetAdminUsers(cursor string, limit int32, searches ...Search) (interface{}, error) {
	params := make(map[string]string)
	params["query"] = serializeSearches(searches)

	if cursor != "" {
		params["per_page"] = strconv.Itoa(int(limit))
		params["cursor"] = cursor
	}

	res, err := c.r().SetQueryParams(params).Get("/admin_users/")

	if err != nil {
		return nil, fmt.Errorf("error retrieving deprovisioned people: %q", err)
	}

	if !res.IsSuccessState() {
		return nil, fmt.Errorf("error retrieving deprovisioned people: %q", res.Status)
	}
	var response AdminUserListResponse

	err = res.UnmarshalJson(&response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %q", err)
	}

	return &response, nil
}

func (c *Client) GetAdminUserById(id string) (*AdminUser, error) {
	res, err := c.r().SetPathParam("adminUserId", id).Get("/admin_users/{adminUserId}")

	if err != nil {
		return nil, fmt.Errorf("error retrieving admin users: %q", err)
	}

	if !res.IsSuccessState() {
		return nil, fmt.Errorf("error retrieving admin user: %q", res.Status)
	}
	var response AdminUser

	err = res.UnmarshalJson(&response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %q", err)
	}

	return &response, nil
}
