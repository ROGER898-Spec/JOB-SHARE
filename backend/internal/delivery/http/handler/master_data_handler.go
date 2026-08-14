package handler

import (
	"strconv"

	dtoRequest "github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/dto/request"
	dtoResponse "github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/dto/response"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/usecase"
	pkgResponse "github.com/FyaEdu/JOB-SHARE/backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type MasterDataHandler struct {
	categoryUsecase usecase.CategoryUsecase
	skillUsecase    usecase.SkillUsecase
}

func NewMasterDataHandler(cu usecase.CategoryUsecase, su usecase.SkillUsecase) *MasterDataHandler {
	return &MasterDataHandler{
		categoryUsecase: cu,
		skillUsecase:    su,
	}
}

// CreateCategory godoc
// @Summary Create Project Category
// @Description Add a new project category (e.g., IT, Design, Writing).
// @Tags Master Data
// @Accept json
// @Produce json
// @Param request body request.CreateCategoryRequest true "Category Payload"
// @Success 201 {object} response.CategoryResponse
// @Failure 400 {object} map[string]interface{}
// @Router /categories [post]
func (h *MasterDataHandler) CreateCategory(c *fiber.Ctx) error {
	var req dtoRequest.CreateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	category, err := h.categoryUsecase.Create(c.Context(), req.Name, req.Description)
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	res := dtoResponse.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return pkgResponse.Success(c, fiber.StatusCreated, "Category created successfully", res)
}

// GetAllCategories godoc
// @Summary Get All Categories
// @Description Retrieve a list of all project categories.
// @Tags Master Data
// @Produce json
// @Success 200 {object} []response.CategoryResponse
// @Failure 500 {object} map[string]interface{}
// @Router /categories [get]
func (h *MasterDataHandler) GetAllCategories(c *fiber.Ctx) error {
	categories, err := h.categoryUsecase.GetAll(c.Context())
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	var resList []dtoResponse.CategoryResponse
	for _, cat := range categories {
		resList = append(resList, dtoResponse.CategoryResponse{
			ID:          cat.ID,
			Name:        cat.Name,
			Description: cat.Description,
		})
	}

	return pkgResponse.Success(c, fiber.StatusOK, "Categories retrieved successfully", resList)
}

// CreateSkill godoc
// @Summary Create Skill
// @Description Add a new skill under a specific category.
// @Tags Master Data
// @Accept json
// @Produce json
// @Param request body request.CreateSkillRequest true "Skill Payload"
// @Success 201 {object} response.SkillResponse
// @Failure 400 {object} map[string]interface{}
// @Router /skills [post]
func (h *MasterDataHandler) CreateSkill(c *fiber.Ctx) error {
	var req dtoRequest.CreateSkillRequest
	if err := c.BodyParser(&req); err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	skill, err := h.skillUsecase.Create(c.Context(), req.CategoryID, req.Name)
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	res := dtoResponse.SkillResponse{
		ID:         skill.ID,
		CategoryID: skill.CategoryID,
		Name:       skill.Name,
	}

	return pkgResponse.Success(c, fiber.StatusCreated, "Skill created successfully", res)
}

// GetSkillsByCategory godoc
// @Summary Get Skills by Category
// @Description Retrieve all skills associated with a specific category ID.
// @Tags Master Data
// @Produce json
// @Param category_id path int true "Category ID"
// @Success 200 {object} []response.SkillResponse
// @Failure 500 {object} map[string]interface{}
// @Router /skills/category/{category_id} [get]
func (h *MasterDataHandler) GetSkillsByCategory(c *fiber.Ctx) error {
	categoryID, err := strconv.Atoi(c.Params("category_id"))
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid category ID")
	}

	skills, err := h.skillUsecase.GetByCategory(c.Context(), categoryID)
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	var resList []dtoResponse.SkillResponse
	for _, skill := range skills {
		resList = append(resList, dtoResponse.SkillResponse{
			ID:         skill.ID,
			CategoryID: skill.CategoryID,
			Name:       skill.Name,
		})
	}

	return pkgResponse.Success(c, fiber.StatusOK, "Skills retrieved successfully", resList)
}
