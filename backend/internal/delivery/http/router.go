package http

import (
	_ "github.com/FyaEdu/JOB-SHARE/backend/docs"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/handler"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/middleware"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func RegisterRoutes(
	app *fiber.App,
	authHandler *handler.AuthHandler,
	umkmProfileHandler *handler.UmkmProfileHandler,
	freelancerProfileHandler *handler.FreelancerProfileHandler,
	jobHandler *handler.JobHandler,
	jobAppHandler *handler.JobApplicationHandler,
	masterDataHandler *handler.MasterDataHandler,
	kanbanHandler *handler.KanbanTaskHandler,
	trxHandler *handler.TransactionHandler,
	reviewHandler *handler.ReviewHandler,
	auditLogRepo domain.AuditLogRepository,
) {
	// Swagger Documentation
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Health Check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "success",
		})
	})

	api := app.Group("/api/v1")

	api.Use(middleware.AuditLogMiddleware(auditLogRepo))

	// ==========================================
	// PUBLIC ROUTES (Tidak butuh token)
	// ==========================================

	// Auth Public
	authRoutes := api.Group("/auth")
	authRoutes.Post("/register", authHandler.Register)
	authRoutes.Post("/login", authHandler.Login)

	// Master Data Read-Only (Public)
	api.Get("/categories", masterDataHandler.GetAllCategories)
	api.Get("/skills/category/:category_id", masterDataHandler.GetSkillsByCategory)

	// Job Read (Public / Optional - semua yang login atau umum bisa lihat list job)
	api.Get("/jobs", jobHandler.GetAll)
	api.Get("/jobs/:id", jobHandler.GetByID)

	// ==========================================
	// PROTECTED ROUTES (Wajib bawa token JWT)
	// ==========================================
	protected := api.Group("/", middleware.Protected())

	// UMKM Profile Routes
	protected.Post("/umkm/profile", middleware.RoleGuard("umkm"), umkmProfileHandler.CreateProfile)
	protected.Get("/umkm/profile/:user_id", umkmProfileHandler.GetProfile)

	// Freelancer Profile Routes
	protected.Post("/freelancers/profile", middleware.RoleGuard("freelancer"), freelancerProfileHandler.CreateProfile)
	protected.Get("/freelancers/profile/:user_id", freelancerProfileHandler.GetProfile)

	// Job Posting (Hanya UMKM)
	protected.Post("/jobs", middleware.RoleGuard("umkm"), jobHandler.Create)

	// Job Application Routes
	protected.Post("/applications", middleware.RoleGuard("freelancer"), jobAppHandler.Apply)
	protected.Get("/applications/job/:job_id", middleware.RoleGuard("umkm"), jobAppHandler.GetByJobID)
	protected.Get("/applications/freelancer/:freelancer_id", middleware.RoleGuard("freelancer"), jobAppHandler.GetByFreelancerID)
	protected.Patch("/applications/:id/status", middleware.RoleGuard("umkm"), jobAppHandler.UpdateStatus)

	// Master Data Create (Hanya Admin)
	protected.Post("/categories", middleware.RoleGuard("admin"), masterDataHandler.CreateCategory)
	protected.Post("/skills", middleware.RoleGuard("admin"), masterDataHandler.CreateSkill)

	// Kanban Routes
	protected.Post("/kanban/tasks", kanbanHandler.Create)
	protected.Get("/kanban/jobs/:job_id/tasks", kanbanHandler.GetByJobID)
	protected.Patch("/kanban/tasks/:id/status", kanbanHandler.UpdateStatus)

	// Transaction Routes
	protected.Post("/transactions", middleware.RoleGuard("umkm"), trxHandler.Create)
	protected.Get("/transactions/job/:job_id", trxHandler.GetByJobID)
	protected.Patch("/transactions/:id/release", middleware.RoleGuard("umkm"), trxHandler.ReleaseEscrow)

	// Review Routes
	protected.Post("/reviews", middleware.RoleGuard("umkm"), reviewHandler.Create)
	protected.Get("/reviews/job/:job_id", reviewHandler.GetByJobID)
	protected.Get("/reviews/freelancer/:freelancer_id", reviewHandler.GetByFreelancerID)
}
