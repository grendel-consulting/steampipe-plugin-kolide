package kolide_client

import "time"

type RegistrationRequestListResponse struct {
	RegistrationRequests []RegistrationRequest `json:"data"`
	Pagination           Pagination            `json:"pagination"`
}
type RegistrationRequest struct {
	Id                   string               `json:"id"`
	Status               string               `json:"status"`
	RequesterMessage     string               `json:"requester_message,omitempty"`
	InternalDenialNote   string               `json:"internal_denial_note,omitempty"`
	RequestedAt          time.Time            `json:"requested_at,omitempty"`
	EndUserDenialNote    string               `json:"end_user_denial_note,omitempty"`
	RequesterInformation RequesterInformation `json:"requester_information,omitempty"`
	DeviceInformation    DeviceInformation    `json:"device_information,omitempty"`
}

type RequesterInformation struct {
	Identifier string `json:"identifier,omitempty"`
}

func (c *Client) GetRegistrationRequests(cursor string, limit int32, searches ...Search) (interface{}, error) {
	return c.fetchCollection("/registration_request/", cursor, limit, searches, new(RegistrationRequestListResponse))
}

func (c *Client) GetRegistrationRequestById(id string) (interface{}, error) {
	return c.fetchResource("/registration_request/", id, new(RegistrationRequest))
}
