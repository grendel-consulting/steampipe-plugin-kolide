package kolide_client

import (
	"fmt"
	"strconv"
)

func (c *Client) fetchCollection(path string, cursor string, limit int32, searches []Search, target interface{}) (interface{}, error) {
	params := make(map[string]string)
	params["query"] = serializeSearches(searches)

	if cursor != "" {
		params["per_page"] = strconv.Itoa(int(limit))
		params["cursor"] = cursor
	}

	err := c.c.Get(path).SetQueryParams(params).Do().Into(&target)

	if err != nil {
		return nil, fmt.Errorf("faled to retrieve collection at %s with response: %q", path, err)
	}

	return target, nil
}

func (c *Client) fetchResource(path string, resourceId string, target interface{}) (interface{}, error) {
	err := c.c.Get(path+"{resourceId}").SetPathParam("resourceId", resourceId).Do().Into(&target)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve resource at %s{resouceId} with ID %s: %v", path, resourceId, err)
	}

	return target, nil
}
