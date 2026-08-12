package request

type CreateReviewRequest struct {
	JobID        int    `json:"job_id" validate:"required"`
	UmkmID       int    `json:"umkm_id" validate:"required"`
	FreelancerID int    `json:"freelancer_id" validate:"required"`
	Rating       int    `json:"rating" validate:"required,min=1,max=5"`
	Feedback     string `json:"feedback"`
}
