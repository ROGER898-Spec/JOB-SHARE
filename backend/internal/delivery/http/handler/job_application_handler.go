package handler

import (
	"strconv"

	dtoRequest "github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/dto/request"
	dtoResponse "github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/dto/response"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
	pkgResponse "github.com/FyaEdu/JOB-SHARE/backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type JobApplicationHandler struct {
	appUsecase domain.JobApplicationUsecase
}

func NewJobApplicationHandler(uc domain.JobApplicationUsecase) *JobApplicationHandler {
	return &JobApplicationHandler{appUsecase: uc}
}

// Apply godoc
// @Summary Submit Job Application
// @Description Freelancer applies for a job
// @Tags Job Applications
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body request.CreateJobAppRequest true "Application Payload"
// @Success 201 {object} response.JobApplicationResponse
// @Router /applications [post]
func (h *JobApplicationHandler) Apply(c *fiber.Ctx) error {
	var req dtoRequest.CreateJobAppRequest
	if err := c.BodyParser(&req); err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid request")
	}

	app := &domain.JobApplication{
		JobID:        req.JobID,
		FreelancerID: req.FreelancerID,
		PitchMessage: req.PitchMessage,
	}

	if err := h.appUsecase.ApplyJob(c.Context(), app); err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	res := dtoResponse.JobApplicationResponse{
		ID:           app.ID,
		JobID:        app.JobID,
		FreelancerID: app.FreelancerID,
		PitchMessage: app.PitchMessage,
		Status:       app.Status,
		AppliedAt:    app.AppliedAt,
	}

	return pkgResponse.Success(c, fiber.StatusCreated, "Applied successfully", res)
}

// GetByJobID godoc
// @Summary Get applications by Job ID
// @Tags Job Applications
// @Security BearerAuth
// @Param job_id path int true "Job ID"
// @Produce json
// @Success 200 {object} []response.JobApplicationResponse
// @Router /applications/job/{job_id} [get]
func (h *JobApplicationHandler) GetByJobID(c *fiber.Ctx) error {
	jobID, err := strconv.Atoi(c.Params("job_id"))
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid job ID")
	}

	apps, err := h.appUsecase.GetByJobID(c.Context(), jobID)
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	var resList []dtoResponse.JobApplicationResponse
	for _, a := range apps {
		resList = append(resList, dtoResponse.JobApplicationResponse{
			ID:           a.ID,
			JobID:        a.JobID,
			FreelancerID: a.FreelancerID,
			PitchMessage: a.PitchMessage,
			Status:       a.Status,
			AppliedAt:    a.AppliedAt,
		})
	}

	return pkgResponse.Success(c, fiber.StatusOK, "Retrieved successfully", resList)
}

// UpdateStatus godoc
// @Summary Update application status
// @Tags Job Applications
// @Security BearerAuth
// @Param id path int true "Application ID"
// @Param request body request.UpdateStatusRequest true "Status Payload"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /applications/{id}/status [patch]
func (h *JobApplicationHandler) UpdateStatus(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid application ID")
	}

	var req dtoRequest.UpdateStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid status payload")
	}

	if err := h.appUsecase.UpdateStatus(c.Context(), id, req.Status); err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkgResponse.Success(c, fiber.StatusOK, "Status updated successfully", nil)
}

// GetByFreelancerID godoc
// @Summary Get applications by Freelancer ID
// @Description Retrieve a list of applications submitted by a specific freelancer
// @Tags Job Applications
// @Security BearerAuth
// @Param freelancer_id path int true "Freelancer ID"
// @Produce json
// @Success 200 {object} []response.JobApplicationResponse
// @Router /applications/freelancer/{freelancer_id} [get]
func (h *JobApplicationHandler) GetByFreelancerID(c *fiber.Ctx) error {
	freelancerID, err := strconv.Atoi(c.Params("freelancer_id"))
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid freelancer ID")
	}

	apps, err := h.appUsecase.GetByFreelancerID(c.Context(), freelancerID)
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	var resList []dtoResponse.JobApplicationResponse
	for _, a := range apps {
		resList = append(resList, dtoResponse.JobApplicationResponse{
			ID:           a.ID,
			JobID:        a.JobID,
			FreelancerID: a.FreelancerID,
			PitchMessage: a.PitchMessage,
			Status:       a.Status,
			AppliedAt:    a.AppliedAt,
		})
	}

	return pkgResponse.Success(c, fiber.StatusOK, "Retrieved successfully", resList)
}
