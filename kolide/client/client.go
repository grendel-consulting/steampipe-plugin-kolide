package kolide_client

import (
	"fmt"
	"strings"

	"github.com/imroc/req/v3"
)

type Client struct {
	c *req.Client
}

type ClientOption func(c *Client)

type OperatorType string

const (
	Equals         OperatorType = ":"
	SubstringMatch OperatorType = "~"
	GreaterThan    OperatorType = ">"
	LessThan       OperatorType = "<"
)

type Search struct {
	Field    string
	Operator OperatorType
	Value    string
}

func New(options ...ClientOption) *Client {
	c := &Client{
		c: req.C(),
	}
	c.c.SetBaseURL("https://api.kolide.com/")

	for _, o := range options {
		o(c)
	}

	return c
}

func (c *Client) r() *req.Request {
	return c.c.R().
		SetHeader("accept", "application/json").
		SetHeader("x-kolide-api-version", "2023-05-26")
}

func (c *Client) setAuth(apiToken string) *Client {
	c.c.SetCommonBearerAuthToken(apiToken)
	return c
}

func WithAuth(apiToken string) ClientOption {
	return func(c *Client) {
		c.setAuth(apiToken)
	}
}

func serializeSearches(searches []Search) string {
	var serializedStrings []string

	if len(searches) == 0 {
		return ""
	}

	for _, s := range searches {
		serialized := fmt.Sprintf("%s%s%s", s.Field, string(s.Operator), s.Value)
		serializedStrings = append(serializedStrings, serialized)
	}

	result := strings.Join(serializedStrings, " AND ")

	return result
}
