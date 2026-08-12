package request

type CreateKanbanTaskRequest struct {
	JobID       int    `json:"job_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

type UpdateKanbanTaskStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=todo in_progress done"`
}
