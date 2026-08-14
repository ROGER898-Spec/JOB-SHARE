package response

import "time"

type JobApplicationResponse struct {
	ID           int       `json:"id"`
	JobID        int       `json:"job_id"`
	FreelancerID int       `json:"freelancer_id"`
	PitchMessage string    `json:"pitch_message"`
	Status       string    `json:"status"`
	AppliedAt    time.Time `json:"applied_at"`
}
