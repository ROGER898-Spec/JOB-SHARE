package handler

import (
	"strconv"

	dtoRequest "github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/dto/request"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/usecase"
	pkgResponse "github.com/FyaEdu/JOB-SHARE/backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type JobHandler struct {
	jobUsecase usecase.JobUsecase
}

func NewJobHandler(uc usecase.JobUsecase) *JobHandler {
	return &JobHandler{jobUsecase: uc}
}

// Create godoc
// @Summary Create Job Posting
// @Description UMKM creates a new job post and specifies required skills.
// @Tags Jobs
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body request.CreateJobRequest true "Job Payload"
// @Success 201 {object} response.JobResponse
// @Router /jobs [post]
func (h *JobHandler) Create(c *fiber.Ctx) error {
	var req dtoRequest.CreateJobRequest
	if err := c.BodyParser(&req); err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid request")
	}

	jobDomain := &domain.Job{
		UmkmID:       req.UmkmID,
		CategoryID:   &req.CategoryID,
		Title:        req.Title,
		Description:  req.Description,
		BudgetAmount: req.BudgetAmount,
	}

	job, err := h.jobUsecase.Create(c.Context(), jobDomain, req.SkillIDs)
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkgResponse.Success(c, fiber.StatusCreated, "Job posted successfully", job)
}

// GetAll godoc
// @Summary Get All Jobs
// @Tags Jobs
// @Produce json
// @Success 200 {object} []response.JobResponse
// @Router /jobs [get]
func (h *JobHandler) GetAll(c *fiber.Ctx) error {
	jobs, err := h.jobUsecase.GetAll(c.Context())
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return pkgResponse.Success(c, fiber.StatusOK, "Jobs retrieved", jobs)
}

// GetByID godoc
// @Summary Get Job Details
// @Tags Jobs
// @Param id path int true "Job ID"
// @Produce json
// @Success 200 {object} response.JobResponse
// @Router /jobs/{id} [get]
func (h *JobHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	job, err := h.jobUsecase.GetByID(c.Context(), id)
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusNotFound, "Job not found")
	}
	return pkgResponse.Success(c, fiber.StatusOK, "Job details retrieved", job)
}
