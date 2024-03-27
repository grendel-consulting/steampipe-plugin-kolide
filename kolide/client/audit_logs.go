package kolide_client

import (
	"time"
)

type AuditLogListResponse struct {
	AuditLogs  []AuditLog `json:"data"`
	Pagination Pagination `json:"pagination"`
}
type AuditLog struct {
	Id          string    `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	ActorName   string    `json:"actor_name"`
	Description string    `json:"description"`
}

func (c *Client) GetAuditLogs(cursor string, limit int32, searches ...Search) (interface{}, error) {
	return c.fetchCollection("/audit_logs/", cursor, limit, searches, new(AuditLogListResponse))
}

func (c *Client) GetAuditLogById(id string) (interface{}, error) {
	return c.fetchResource("/audit_logs/", id, new(AuditLog))
}
