package response

import "time"

type ReviewResponse struct {
	ID           int       `json:"id"`
	JobID        int       `json:"job_id"`
	UmkmID       int       `json:"umkm_id"`
	FreelancerID int       `json:"freelancer_id"`
	Rating       int       `json:"rating"`
	Feedback     string    `json:"feedback"`
	CreatedAt    time.Time `json:"created_at"`
}
