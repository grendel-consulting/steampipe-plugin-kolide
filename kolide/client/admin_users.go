package kolide_client

import (
	"fmt"
	"time"
)

type AdminUserListResponse struct {
	AdminUsers []AdminUser `json:"data"`
}
type AdminUser struct {
	Id        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Access    string    `json:"access"`
}

func (c *Client) GetAdminUsers() (*AdminUserListResponse, error) {
	res, err := c.r().Get("/admin_users/")

	if err != nil {
		return nil, fmt.Errorf("error retrieving AdminUsers: %q", err)
	}

	if !res.IsSuccessState() {
		return nil, fmt.Errorf("error retrieving AdminUsers: %q", res.Status)
	}
	var response AdminUserListResponse

	err = res.UnmarshalJson(&response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %q", err)
	}

	return &response, nil
}
