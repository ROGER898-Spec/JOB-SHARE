package handler

import (
	"strconv"

	dtoRequest "github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/dto/request"
	dtoResponse "github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/dto/response"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
	pkgResponse "github.com/FyaEdu/JOB-SHARE/backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type KanbanTaskHandler struct {
	taskUsecase domain.KanbanTaskUsecase
}

func NewKanbanTaskHandler(uc domain.KanbanTaskUsecase) *KanbanTaskHandler {
	return &KanbanTaskHandler{taskUsecase: uc}
}

// Create godoc
// @Summary Create Kanban Task
// @Description Add a new task to a job's workspace
// @Tags Kanban
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body request.CreateKanbanTaskRequest true "Task Payload"
// @Success 201 {object} response.KanbanTaskResponse
// @Router /kanban/tasks [post]
func (h *KanbanTaskHandler) Create(c *fiber.Ctx) error {
	var req dtoRequest.CreateKanbanTaskRequest
	if err := c.BodyParser(&req); err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	task := &domain.KanbanTask{
		JobID:       req.JobID,
		Title:       req.Title,
		Description: req.Description,
	}

	if err := h.taskUsecase.CreateTask(c.Context(), task); err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	res := dtoResponse.KanbanTaskResponse{
		ID:          task.ID,
		JobID:       task.JobID,
		Title:       task.Title,
		Description: task.Description,
		Status:      "todo",
		CreatedAt:   task.CreatedAt,
	}

	return pkgResponse.Success(c, fiber.StatusCreated, "Task created successfully", res)
}

// GetByJobID godoc
// @Summary Get Tasks by Job ID
// @Description Retrieve all kanban tasks for a specific job
// @Tags Kanban
// @Security BearerAuth
// @Param job_id path int true "Job ID"
// @Produce json
// @Success 200 {object} []response.KanbanTaskResponse
// @Router /kanban/jobs/{job_id}/tasks [get]
func (h *KanbanTaskHandler) GetByJobID(c *fiber.Ctx) error {
	jobID, err := strconv.Atoi(c.Params("job_id"))
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid job ID")
	}

	tasks, err := h.taskUsecase.GetTasksByJobID(c.Context(), jobID)
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	var resList []dtoResponse.KanbanTaskResponse
	for _, t := range tasks {
		resList = append(resList, dtoResponse.KanbanTaskResponse{
			ID:          t.ID,
			JobID:       t.JobID,
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
			CreatedAt:   t.CreatedAt,
		})
	}

	return pkgResponse.Success(c, fiber.StatusOK, "Tasks retrieved successfully", resList)
}

// UpdateStatus godoc
// @Summary Update Task Status
// @Description Move a task to todo, in_progress, or done
// @Tags Kanban
// @Security BearerAuth
// @Param id path int true "Task ID"
// @Param request body request.UpdateKanbanTaskStatusRequest true "Status Payload"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /kanban/tasks/{id}/status [patch]
func (h *KanbanTaskHandler) UpdateStatus(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid task ID")
	}

	var req dtoRequest.UpdateKanbanTaskStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid status payload")
	}

	if err := h.taskUsecase.UpdateTaskStatus(c.Context(), id, req.Status); err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkgResponse.Success(c, fiber.StatusOK, "Task status updated successfully", nil)
}
