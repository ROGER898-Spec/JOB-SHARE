package request

type CreateUmkmProfileRequest struct {
	UserID       string `json:"user_id" validate:"required"`
	BusinessName string `json:"business_name" validate:"required"`
	OwnerName    string `json:"owner_name" validate:"required"`
	PhoneNumber  string `json:"phone_number" validate:"required"`
	City         string `json:"city" validate:"required"`
	FullAddress  string `json:"full_address" validate:"required"`
}
