package kolide_client

type CheckListResponse struct {
	Checks     []Check    `json:"data"`
	Pagination Pagination `json:"pagination"`
}
type Check struct {
	Id                      string                  `json:"id"`
	Name                    string                  `json:"name"`
	CompatiblePlatforms     []string                `json:"compatible_platforms"`
	Description             string                  `json:"description,omitempty"`
	Topics                  []string                `json:"topics,omitempty"`
	CheckTags               []CheckTag              `json:"check_tags,omitempty"`
	BlocksAuthConfiguration BlocksAuthConfiguration `json:"blocks_auth_configuration"`
	TargetingConfiguration  TargetingConfiguration  `json:"targeting_configuration"`
}

type CheckTag struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

type BlocksAuthConfiguration struct {
	BlockingEnabled          bool     `json:"blocking_enabled"`
	GracePeriodDays          int      `json:"grace_period_days"`
	BlockingGroupNames       []string `json:"blocking_group_names,omitempty"`
	BlockingExemptGroupNames []string `json:"blocking_exempt_group_names,omitempty"`
}

type TargetingConfiguration struct {
	ExcludedGroups []string `json:"excluded_groups,omitempty"`
	TargetedGroups []string `json:"targeted_groups,omitempty"`
}

func (c *Client) GetChecks(cursor string, limit int32, searches ...Search) (interface{}, error) {
	var friendlies = map[string]string{
		"description": "check_description",
		"check_tag":   "check_tag_name",
		// Kolide K2 supports filtering by check_tag_id and check_tag_description as well
	}

	return c.fetchCollection("/checks/", cursor, limit, searches, new(CheckListResponse), friendlies)
}

func (c *Client) GetCheckById(id string) (interface{}, error) {
	return c.fetchResource("/checks/", id, new(Check))
}
