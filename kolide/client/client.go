package kolide_client

import (
	"fmt"
	"net/url"
	"os"

	"github.com/imroc/req/v3"
)

type Client struct {
	c *req.Client
}

type ClientOption func(c *Client)

type ApiError struct {
	Message string `json:"error"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("Kolide API Error: %s", e.Message)
}

func New(options ...ClientOption) *Client {
	c := &Client{
		c: req.C(),
	}

	var baseUrl string = os.Getenv("KOLIDE_API_URL")

	_, err := url.Parse(baseUrl)
	if err != nil || baseUrl == "" {
		// Log the error but and use the default URL
		fmt.Printf("Invalid or empty KOLIDE_API_URL provided: %v. Using default URL.\n", err)
		baseUrl = "https://api.kolide.com"
	}

	c.c.SetBaseURL(baseUrl).
		SetCommonHeader("accept", "application/json").
		SetCommonErrorResult(&ApiError{}).
		EnableDumpEachRequest().
		OnAfterResponse(func(client *req.Client, res *req.Response) error {
			if res.Err != nil {
				return nil // Skip the following logic if there is an underlying error.
			}
			// Return a human-readable error if server api returned an 401 or 403 error message.
			if err, ok := res.ErrorResult().(*ApiError); ok {
				res.Err = err
				return nil
			}
			if !res.IsSuccessState() {
				res.Err = fmt.Errorf("bad response, raw content:\n%s", res.Dump())
				return nil
			}
			return nil
		})
	for _, o := range options {
		o(c)
	}

	return c
}

func (c *Client) setAuth(apiToken string) *Client {
	c.c.SetCommonBearerAuthToken(apiToken).
		SetCommonHeader("x-kolide-api-version", "2023-05-26")
	return c
}

func WithAuth(apiToken string) ClientOption {
	return func(c *Client) {
		c.setAuth(apiToken)
	}
}
