package request

type CreateJobRequest struct {
	UmkmID       int     `json:"umkm_id" validate:"required"`
	CategoryID   int     `json:"category_id" validate:"required"`
	Title        string  `json:"title" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	BudgetAmount float64 `json:"budget_amount" validate:"required"`
	SkillIDs     []int   `json:"skill_ids" validate:"required"`
}
