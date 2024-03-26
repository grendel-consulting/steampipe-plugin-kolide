package kolide_client

// See: https://www.kolide.com/docs/developers/api#pagination
type Pagination struct {
	Next          string `json:"next,omitempty"`
	CurrentCursor string `json:"current_cursor,omitempty"`
	NextCursor    string `json:"next_cursor,omitempty"`
	Count         int    `json:"count"`
}

const (
	DefaultPaging int32 = 25
	MaxPaging     int32 = 100
)
