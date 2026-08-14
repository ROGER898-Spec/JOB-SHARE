package request

type CreateJobRequest struct {
	UmkmID        int     `json:"umkm_id" validate:"required"`
	CategoryID    int     `json:"category_id" validate:"required"`
	Title         string  `json:"title" validate:"required"`
	Description   string  `json:"description" validate:"required"`
	BudgetAmount  float64 `json:"budget_amount" validate:"required"`
	Location      string  `json:"location"`
	RadiusKm      int     `json:"radius_km"`
	DurationLabel string  `json:"duration_label"`
	SkillIDs      []int   `json:"skill_ids" validate:"required"`
}

type UpdateJobStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=open in_progress completed cancelled"`
}