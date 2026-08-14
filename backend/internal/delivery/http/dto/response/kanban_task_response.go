package response

import "time"

type KanbanTaskResponse struct {
	ID          int       `json:"id"`
	JobID       int       `json:"job_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
