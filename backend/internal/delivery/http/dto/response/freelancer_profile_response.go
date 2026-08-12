package response

import "time"

type FreelancerProfileResponse struct {
	ID            int        `json:"id"`
	UserID        string     `json:"user_id"`
	FullName      string     `json:"full_name"`
	BioSummary    *string    `json:"bio_summary,omitempty"`
	PhoneNumber   *string    `json:"phone_number,omitempty"`
	City          *string    `json:"city,omitempty"`
	PortfolioLink *string    `json:"portfolio_link,omitempty"`
	CvURL         *string    `json:"cv_url,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}
