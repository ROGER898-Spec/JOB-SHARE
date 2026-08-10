package response

import "time"

type JobResponse struct {
	ID          string     `json:"id"`
	UmkmID      int        `json:"umkm_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Budget      float64    `json:"budget"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
