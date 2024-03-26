package kolide_client

import (
	"fmt"
	"time"
)

type DeviceListResponse struct {
	Devices []Device `json:"data"`
}
type Device struct {
	Id                  string    `json:"id"`
	Name                string    `json:"name"`
	RegisteredAt        time.Time `json:"registered_at,omitempty"`
	LastAuthenticatedAt time.Time `json:"last_authenticated_at,omitempty"`
	OperatingSystem     string    `json:"operating_system"`
	HardwareModel       string    `json:"hardware_model"`
	Serial              string    `json:"serial,omitempty"`
	HardwareUuid        string    `json:"hardware_uuid,omitempty"`
	Note                string    `json:"note,omitempty"`
	AuthState           string    `json:"auth_state"`
	WillBlockAt         time.Time `json:"will_block_at,omitempty"`
	ProductImageUrl     string    `json:"product_image_url"`
	DeviceType          string    `json:"device_type"`
	FormFactor          string    `json:"form_factor"`
}

func (c *Client) GetDevices(searches ...Search) (*DeviceListResponse, error) {
	query := serializeSearches(searches)

	res, err := c.r().SetQueryParam("query", query).Get("/devices/")

	if err != nil {
		return nil, fmt.Errorf("error retrieving devices: %q", err)
	}

	if !res.IsSuccessState() {
		return nil, fmt.Errorf("error retrieving devices: %q", res.Status)
	}
	var response DeviceListResponse

	err = res.UnmarshalJson(&response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %q", err)
	}

	return &response, nil
}

func (c *Client) GetDeviceById(id string) (*Device, error) {
	res, err := c.r().SetPathParam("deviceId", id).Get("/devices/{deviceId}")

	if err != nil {
		return nil, fmt.Errorf("error retrieving devices: %q", err)
	}

	if !res.IsSuccessState() {
		return nil, fmt.Errorf("error retrieving devices: %q", res.Status)
	}
	var response Device

	err = res.UnmarshalJson(&response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %q", err)
	}

	return &response, nil
}
