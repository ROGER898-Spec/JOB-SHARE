package domain

import (
	"context"
	"time"
)

type AuditLog struct {
	ID         int       `json:"id"`
	UserID     *string   `json:"user_id"` // Pointer karena bisa jadi null (misal saat register)
	Action     string    `json:"action"`
	EntityType string    `json:"entity_type,omitempty"`
	EntityID   string    `json:"entity_id,omitempty"`
	IPAddress  string    `json:"ip_address,omitempty"`
	Details    string    `json:"details,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

type AuditLogRepository interface {
	Create(ctx context.Context, log *AuditLog) error
}
