package kolide_client

import (
	"time"
)

type PackageListResponse struct {
	Packages   []Package  `json:"data"`
	Pagination Pagination `json:"pagination"`
}
type Package struct {
	Id      string    `json:"id"`
	BuiltAt time.Time `json:"built_at"`
	Url     string    `json:"url"`
	Version string    `json:"version"`
}

func (c *Client) GetPackages(cursor string, limit int32, searches ...Search) (interface{}, error) {
	return c.fetchCollection("/packages/", cursor, limit, searches, new(PackageListResponse))
}

func (c *Client) GetPackageById(id string) (interface{}, error) {
	return c.fetchResource("/packages/", id, new(Package))
}
