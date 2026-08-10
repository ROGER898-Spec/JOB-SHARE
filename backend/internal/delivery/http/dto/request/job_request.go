package request

type CreateJobRequest struct {
	UmkmID      int     `json:"umkm_id" validate:"required"`
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Budget      float64 `json:"budget" validate:"required,gt=0"`
	Status      string  `json:"status" validate:"omitempty,oneof=open in_progress completed closed"`
}
