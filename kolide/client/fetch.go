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

	res, err := c.r().SetQueryParams(params).Get(path)
	if err != nil {
		return nil, fmt.Errorf("error retrieving collection at %s with params %v: %q", path, params, err)
	}

	if !res.IsSuccessState() {
		return nil, fmt.Errorf("error retrieving collection at %s: %q", path, res.Status)
	}

	err = res.UnmarshalJson(&target)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON for collection at %s with response %s: %q", path, res.String(), err)
	}

	return target, nil
}

func (c *Client) fetchResource(path string, resourceId string, target interface{}) (interface{}, error) {
	res, err := c.r().SetPathParam("id", resourceId).Get(path + "{id}")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve resource at %s{id} with ID %s: %v", path, resourceId, err)
	}

	if !res.IsSuccessState() {
		return nil, fmt.Errorf("error retrieving resource: %s", res.Status)
	}

	err = res.UnmarshalJson(&target)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON for resource at %s{id} with ID %s: %v", path, resourceId, err)
	}

	return &target, nil
}
