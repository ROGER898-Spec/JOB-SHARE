package request

type CreateSkillRequest struct {
	CategoryID int    `json:"category_id" validate:"required"`
	Name       string `json:"name" validate:"required"`
}
