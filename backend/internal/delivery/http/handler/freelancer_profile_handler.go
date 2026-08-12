package handler

import (
	dtoRequest "github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/dto/request"
	dtoResponse "github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/dto/response"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/usecase"
	pkgResponse "github.com/FyaEdu/JOB-SHARE/backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type FreelancerProfileHandler struct {
	profileUsecase usecase.FreelancerProfileUsecase
}

func NewFreelancerProfileHandler(uc usecase.FreelancerProfileUsecase) *FreelancerProfileHandler {
	return &FreelancerProfileHandler{profileUsecase: uc}
}

// CreateProfile godoc
// @Summary Create Freelancer Profile
// @Description Complete the freelancer profile for a newly registered user.
// @Tags Freelancer Profile
// @Accept json
// @Produce json
// @Param request body request.CreateFreelancerProfileRequest true "Profile Payload"
// @Success 201 {object} response.FreelancerProfileResponse
// @Failure 400 {object} map[string]interface{}
// @Router /freelancers/profile [post]
func (h *FreelancerProfileHandler) CreateProfile(c *fiber.Ctx) error {
	var req dtoRequest.CreateFreelancerProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	domainReq := &domain.FreelancerProfile{
		UserID:        req.UserID,
		FullName:      req.FullName,
		BioSummary:    req.BioSummary,
		PhoneNumber:   req.PhoneNumber,
		City:          req.City,
		PortfolioLink: req.PortfolioLink,
		CvURL:         req.CvURL,
	}

	profile, err := h.profileUsecase.CreateProfile(c.Context(), domainReq)
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, err.Error())
	}

	res := dtoResponse.FreelancerProfileResponse{
		ID:            profile.ID,
		UserID:        profile.UserID,
		FullName:      profile.FullName,
		BioSummary:    profile.BioSummary,
		PhoneNumber:   profile.PhoneNumber,
		City:          profile.City,
		PortfolioLink: profile.PortfolioLink,
		CvURL:         profile.CvURL,
		CreatedAt:     profile.CreatedAt,
		UpdatedAt:     profile.UpdatedAt,
	}

	return pkgResponse.Success(c, fiber.StatusCreated, "Freelancer Profile created successfully", res)
}

// GetProfile godoc
// @Summary Get Freelancer Profile by User ID
// @Description Retrieve freelancer profile details using User ID.
// @Tags Freelancer Profile
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} response.FreelancerProfileResponse
// @Failure 404 {object} map[string]interface{}
// @Router /freelancers/profile/{user_id} [get]
func (h *FreelancerProfileHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	profile, err := h.profileUsecase.GetProfileByUserID(c.Context(), userID)
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	if profile == nil {
		return pkgResponse.Error(c, fiber.StatusNotFound, "Profile not found")
	}

	res := dtoResponse.FreelancerProfileResponse{
		ID:            profile.ID,
		UserID:        profile.UserID,
		FullName:      profile.FullName,
		BioSummary:    profile.BioSummary,
		PhoneNumber:   profile.PhoneNumber,
		City:          profile.City,
		PortfolioLink: profile.PortfolioLink,
		CvURL:         profile.CvURL,
		CreatedAt:     profile.CreatedAt,
		UpdatedAt:     profile.UpdatedAt,
	}

	return pkgResponse.Success(c, fiber.StatusOK, "Profile retrieved successfully", res)
}
