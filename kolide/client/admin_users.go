package kolide_client

import (
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
	return c.fetchCollection("/admin_users/", cursor, limit, searches, new(AdminUserListResponse))
}

func (c *Client) GetAdminUserById(id string) (interface{}, error) {
	return c.fetchResource("/admin_users/", id, new(AdminUser))
}
