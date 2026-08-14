package response

import "time"

type JobResponse struct {
	ID            int             `json:"id"`
	UmkmID        int             `json:"umkm_id"`
	CategoryID    int             `json:"category_id"`
	Title         string          `json:"title"`
	Description   string          `json:"description"`
	BudgetAmount  float64         `json:"budget_amount"`
	Location      string          `json:"location,omitempty"`
	RadiusKm      int             `json:"radius_km,omitempty"`
	DurationLabel string          `json:"duration_label,omitempty"`
	Status        string          `json:"status"`
	CreatedAt     time.Time       `json:"created_at"`
	Skills        []SkillResponse `json:"skills,omitempty"`
}