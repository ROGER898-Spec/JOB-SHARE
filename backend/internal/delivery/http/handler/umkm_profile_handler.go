package handler

import (
	dtoRequest "github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/dto/request"
	dtoResponse "github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/dto/response"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/usecase"
	pkgResponse "github.com/FyaEdu/JOB-SHARE/backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type UmkmProfileHandler struct {
	profileUsecase usecase.UmkmProfileUsecase
}

func NewUmkmProfileHandler(uc usecase.UmkmProfileUsecase) *UmkmProfileHandler {
	return &UmkmProfileHandler{profileUsecase: uc}
}

// CreateProfile godoc
// @Summary Create UMKM Profile
// @Description Complete the UMKM profile for a newly registered user.
// @Tags UMKM Profile
// @Accept json
// @Produce json
// @Param request body request.CreateUmkmProfileRequest true "Profile Payload"
// @Success 201 {object} response.UmkmProfileResponse
// @Failure 400 {object} map[string]interface{}
// @Router /umkm/profile [post]
func (h *UmkmProfileHandler) CreateProfile(c *fiber.Ctx) error {
	var req dtoRequest.CreateUmkmProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	domainReq := &domain.UmkmProfile{
		UserID:       req.UserID,
		BusinessName: req.BusinessName,
		OwnerName:    req.OwnerName,
		PhoneNumber:  req.PhoneNumber,
		City:         req.City,
		FullAddress:  req.FullAddress,
	}

	profile, err := h.profileUsecase.CreateProfile(c.Context(), domainReq)
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, err.Error())
	}

	res := dtoResponse.UmkmProfileResponse{
		ID:                 profile.ID,
		UserID:             profile.UserID,
		BusinessName:       profile.BusinessName,
		OwnerName:          profile.OwnerName,
		PhoneNumber:        profile.PhoneNumber,
		City:               profile.City,
		FullAddress:        profile.FullAddress,
		VerificationStatus: profile.VerificationStatus,
		CreatedAt:          profile.CreatedAt,
		UpdatedAt:          profile.UpdatedAt,
	}

	return pkgResponse.Success(c, fiber.StatusCreated, "UMKM Profile created successfully", res)
}

// GetProfile godoc
// @Summary Get UMKM Profile by User ID
// @Description Retrieve UMKM profile details using User ID.
// @Tags UMKM Profile
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} response.UmkmProfileResponse
// @Failure 404 {object} map[string]interface{}
// @Router /umkm/profile/{user_id} [get]
func (h *UmkmProfileHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	profile, err := h.profileUsecase.GetProfileByUserID(c.Context(), userID)
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	if profile == nil {
		return pkgResponse.Error(c, fiber.StatusNotFound, "Profile not found")
	}

	res := dtoResponse.UmkmProfileResponse{
		ID:                  profile.ID,
		UserID:              profile.UserID,
		BusinessName:        profile.BusinessName,
		OwnerName:           profile.OwnerName,
		PhoneNumber:         profile.PhoneNumber,
		City:                profile.City,
		FullAddress:         profile.FullAddress,
		IdentityDocumentURL: profile.IdentityDocumentURL,
		VerificationStatus:  profile.VerificationStatus,
		VerifiedByAdminID:   profile.VerifiedByAdminID,
		VerifiedAt:          profile.VerifiedAt,
		CreatedAt:           profile.CreatedAt,
		UpdatedAt:           profile.UpdatedAt,
	}

	return pkgResponse.Success(c, fiber.StatusOK, "Profile retrieved successfully", res)
}
