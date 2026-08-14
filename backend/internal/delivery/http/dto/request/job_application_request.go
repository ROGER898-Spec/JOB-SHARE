package request

type CreateJobAppRequest struct {
	JobID        int    `json:"job_id" validate:"required"`
	FreelancerID int    `json:"freelancer_id" validate:"required"`
	PitchMessage string `json:"pitch_message" validate:"required"`
}

type UpdateStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=accepted rejected withdrawn"`
}
