package handler

import (
	"strconv"

	dtoRequest "github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/dto/request"
	dtoResponse "github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/dto/response"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
	pkgResponse "github.com/FyaEdu/JOB-SHARE/backend/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type ReviewHandler struct {
	reviewUsecase domain.ReviewUsecase
}

func NewReviewHandler(uc domain.ReviewUsecase) *ReviewHandler {
	return &ReviewHandler{reviewUsecase: uc}
}

// Create godoc
// @Summary Create a Review
// @Description UMKM gives a rating and feedback to the freelancer after job completion.
// @Tags Reviews
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body request.CreateReviewRequest true "Review Payload"
// @Success 201 {object} response.ReviewResponse
// @Router /reviews [post]
func (h *ReviewHandler) Create(c *fiber.Ctx) error {
	var req dtoRequest.CreateReviewRequest
	if err := c.BodyParser(&req); err != nil {
		return pkgResponse.Error(c, fiber.StatusBadRequest, "Invalid request payload")
	}

	review := &domain.Review{
		JobID:        req.JobID,
		UmkmID:       req.UmkmID,
		FreelancerID: req.FreelancerID,
		Rating:       req.Rating,
		Feedback:     req.Feedback,
	}

	if err := h.reviewUsecase.CreateReview(c.Context(), review); err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	res := dtoResponse.ReviewResponse{
		ID:           review.ID,
		JobID:        review.JobID,
		UmkmID:       review.UmkmID,
		FreelancerID: review.FreelancerID,
		Rating:       review.Rating,
		Feedback:     review.Feedback,
		CreatedAt:    review.CreatedAt,
	}

	return pkgResponse.Success(c, fiber.StatusCreated, "Review created successfully", res)
}

// GetByJobID godoc
// @Summary Get review by Job ID
// @Tags Reviews
// @Security BearerAuth
// @Param job_id path int true "Job ID"
// @Produce json
// @Success 200 {object} response.ReviewResponse
// @Router /reviews/job/{job_id} [get]
func (h *ReviewHandler) GetByJobID(c *fiber.Ctx) error {
	jobID, _ := strconv.Atoi(c.Params("job_id"))

	review, err := h.reviewUsecase.GetReviewByJobID(c.Context(), jobID)
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	if review == nil {
		return pkgResponse.Error(c, fiber.StatusNotFound, "Review not found")
	}

	res := dtoResponse.ReviewResponse{
		ID: review.ID, JobID: review.JobID, UmkmID: review.UmkmID, FreelancerID: review.FreelancerID,
		Rating: review.Rating, Feedback: review.Feedback, CreatedAt: review.CreatedAt,
	}

	return pkgResponse.Success(c, fiber.StatusOK, "Review retrieved", res)
}

// GetByFreelancerID godoc
// @Summary Get reviews by Freelancer ID
// @Tags Reviews
// @Security BearerAuth
// @Param freelancer_id path int true "Freelancer ID"
// @Produce json
// @Success 200 {object} []response.ReviewResponse
// @Router /reviews/freelancer/{freelancer_id} [get]
func (h *ReviewHandler) GetByFreelancerID(c *fiber.Ctx) error {
	freelancerID, _ := strconv.Atoi(c.Params("freelancer_id"))

	reviews, err := h.reviewUsecase.GetReviewsByFreelancerID(c.Context(), freelancerID)
	if err != nil {
		return pkgResponse.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	var resList []dtoResponse.ReviewResponse
	for _, rev := range reviews {
		resList = append(resList, dtoResponse.ReviewResponse{
			ID: rev.ID, JobID: rev.JobID, UmkmID: rev.UmkmID, FreelancerID: rev.FreelancerID,
			Rating: rev.Rating, Feedback: rev.Feedback, CreatedAt: rev.CreatedAt,
		})
	}

	return pkgResponse.Success(c, fiber.StatusOK, "Reviews retrieved", resList)
}
