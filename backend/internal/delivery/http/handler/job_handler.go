package handler

import (
	dtoRequest "github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/dto/request"
	dtoResponse "github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/dto/response"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/usecase"
	pkgResponse "github.com/FyaEdu/JOB-SHARE/backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type JobHandler struct {
	jobUsecase usecase.JobUsecase
}

func NewJobHandler(jobUsecase usecase.JobUsecase) *JobHandler {
	return &JobHandler{jobUsecase: jobUsecase}
}

// Create godoc
// @Summary Create a new job vacancy (UMKM)
// @Description Post a new job vacancy into the system by an UMKM.
// @Tags Jobs
// @Accept json
// @Produce json
// @Param request body request.CreateJobRequest true "Create Job Payload"
// @Success 201 {object} response.JobResponse
// @Failure 400 {object} map[string]interface{}
// @Router /jobs [post]
func (h *JobHandler) Create(c *fiber.Ctx) error {
	var req dtoRequest.CreateJobRequest
	if err := c.BodyParser(&req); err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	job, err := h.jobUsecase.Create(c.Context(), req.UmkmID, req.Title, req.Description, req.Budget, req.Status)
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, err.Error())
	}

	res := dtoResponse.JobResponse{
		ID:          job.ID,
		UmkmID:      job.UmkmID,
		Title:       job.Title,
		Description: job.Description,
		Budget:      job.Budget,
		Status:      job.Status,
		CreatedAt:   job.CreatedAt,
		UpdatedAt:   job.UpdatedAt,
	}

	return pkgResponse.Success(c, fiber.StatusCreated, "Job created successfully", res)
}

// GetAll godoc
// @Summary Get all job vacancies
// @Description Retrieve a list of all available job vacancies.
// @Tags Jobs
// @Accept json
// @Produce json
// @Success 200 {object} []response.JobResponse
// @Failure 500 {object} map[string]interface{}
// @Router /jobs [get]
func (h *JobHandler) GetAll(c *fiber.Ctx) error {
	jobs, err := h.jobUsecase.GetAll(c.Context())
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	var resList []dtoResponse.JobResponse
	for _, job := range jobs {
		resList = append(resList, dtoResponse.JobResponse{
			ID:          job.ID,
			UmkmID:      job.UmkmID,
			Title:       job.Title,
			Description: job.Description,
			Budget:      job.Budget,
			Status:      job.Status,
			CreatedAt:   job.CreatedAt,
			UpdatedAt:   job.UpdatedAt,
		})
	}

	return pkgResponse.Success(c, fiber.StatusOK, "Jobs retrieved successfully", resList)
}

// GetByID godoc
// @Summary Get job vacancy by ID
// @Description Retrieve details of a specific job vacancy using its unique ID.
// @Tags Jobs
// @Accept json
// @Produce json
// @Param id path string true "Job ID"
// @Success 200 {object} response.JobResponse
// @Failure 404 {object} map[string]interface{}
// @Router /jobs/{id} [get]
func (h *JobHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	job, err := h.jobUsecase.GetByID(c.Context(), id)
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	if job == nil {
		return pkgResponse.Error(c, fiber.StatusNotFound, "Job not found")
	}

	res := dtoResponse.JobResponse{
		ID:          job.ID,
		UmkmID:      job.UmkmID,
		Title:       job.Title,
		Description: job.Description,
		Budget:      job.Budget,
		Status:      job.Status,
		CreatedAt:   job.CreatedAt,
		UpdatedAt:   job.UpdatedAt,
	}

	return pkgResponse.Success(c, fiber.StatusOK, "Job retrieved successfully", res)
}
