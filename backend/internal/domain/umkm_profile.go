package domain

import (
	"context"
	"time"
)

type UmkmProfile struct {
	ID                  int        `json:"id"`
	UserID              string     `json:"user_id"`
	BusinessName        string     `json:"business_name"`
	OwnerName           string     `json:"owner_name"`
	PhoneNumber         string     `json:"phone_number"`
	City                string     `json:"city"`
	FullAddress         string     `json:"full_address"`
	IdentityDocumentURL *string    `json:"identity_document_url,omitempty"`
	VerificationStatus  string     `json:"verification_status"`
	VerifiedByAdminID   *string    `json:"verified_by_admin_id,omitempty"`
	VerifiedAt          *time.Time `json:"verified_at,omitempty"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           *time.Time `json:"updated_at,omitempty"`
}

type UmkmProfileRepository interface {
	Create(ctx context.Context, profile *UmkmProfile) error
	FindByUserID(ctx context.Context, userID string) (*UmkmProfile, error)
}
