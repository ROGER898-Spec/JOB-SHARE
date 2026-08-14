package request

type CreateFreelancerProfileRequest struct {
	UserID        string  `json:"user_id" validate:"required"`
	FullName      string  `json:"full_name" validate:"required"`
	BioSummary    *string `json:"bio_summary,omitempty"`
	PhoneNumber   *string `json:"phone_number,omitempty"`
	City          *string `json:"city,omitempty"`
	PortfolioLink *string `json:"portfolio_link,omitempty"`
	CvURL         *string `json:"cv_url,omitempty"`
}
