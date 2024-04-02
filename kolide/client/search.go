package kolide_client

import (
	"fmt"
	"strings"
)

type Search struct {
	Field    string
	Operator OperatorType
	Value    string
}

type OperatorType string

const (
	Equals         OperatorType = ":"
	SubstringMatch OperatorType = "~"
	GreaterThan    OperatorType = ">"
	LessThan       OperatorType = "<"
)

// See: https://www.kolide.com/docs/developers/api#search
func serializeSearches(searches []Search, friendlies ...map[string]string) string {
	var builder strings.Builder

	if len(searches) == 0 {
		return ""
	}

	for _, s := range searches {
		field := s.Field

		if len(friendlies) > 0 {
			for _, m := range friendlies {
				if value, ok := m[field]; ok {
					field = value
				}
			}
		}
		serialized := fmt.Sprintf("%s%s%s", field, string(s.Operator), s.Value)
		if builder.Len() > 0 {
			builder.WriteString(" AND ")
		}
		builder.WriteString(serialized)
	}

	return builder.String()
}
