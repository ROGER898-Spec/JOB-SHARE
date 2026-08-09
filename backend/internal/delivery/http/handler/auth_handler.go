package handler

import (
	pkgResponse "github.com/FyaEdu/JOB-SHARE/backend/pkg/response"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/dto/request"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/dto/response"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

// Register godoc
// @Summary Register a new account (UMKM / Freelancer)
// @Description Register a new user into the system with a specific role (admin, umkm, freelancer).
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.RegisterRequest true "Register Payload"
// @Success 201 {object} response.UserResponse
// @Failure 400 {object} map[string]interface{}
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req request.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	user, err := h.authUsecase.Register(c.Context(), req.Email, req.Password, req.Role)
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, err.Error())
	}

	res := response.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Role:      user.Role,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return pkgResponse.Success(c, fiber.StatusCreated, "User registered successfully", res)
}

// Login godoc
// @Summary User Login
// @Description Authenticate user based on email and password.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "Login Payload"
// @Success 200 {object} response.UserResponse
// @Failure 401 {object} map[string]interface{}
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req request.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	user, err := h.authUsecase.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusUnauthorized, err.Error())
	}

	res := response.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Role:      user.Role,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return pkgResponse.Success(c, fiber.StatusOK, "Login successful", res)
}
