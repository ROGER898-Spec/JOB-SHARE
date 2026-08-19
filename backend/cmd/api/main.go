// @title Jobshare Backend API
// @version 1.0
// @description This is the official API documentation for the Jobshare platform
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	httpDelivery "github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/handler"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/repository/postgres"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/usecase"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	if os.Getenv("DATABASE_URL") == "" {
		if err := godotenv.Load("../.env"); err != nil {
			log.Printf("Warning: .env file not found: %v", err)
		}
	}

	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}
	fmt.Println("Database connected successfully!")

	userRepo := postgres.NewUserRepository(db)
	authUsecase := usecase.NewAuthUsecase(userRepo, 5*time.Second)
	authHandler := handler.NewAuthHandler(authUsecase)

	app := fiber.New(fiber.Config{
		AppName: "Jobshare API v1.0",
	})

	jobRepo := postgres.NewJobRepository(db)
	jobUsecase := usecase.NewJobUsecase(jobRepo, 5*time.Second)
	jobHandler := handler.NewJobHandler(jobUsecase)

	umkmProfileRepo := postgres.NewUmkmProfileRepository(db)
	umkmProfileUsecase := usecase.NewUmkmProfileUsecase(umkmProfileRepo, 5*time.Second)
	umkmProfileHandler := handler.NewUmkmProfileHandler(umkmProfileUsecase)

	freelancerProfileRepo := postgres.NewFreelancerProfileRepository(db)
	freelancerProfileUsecase := usecase.NewFreelancerProfileUsecase(freelancerProfileRepo, 5*time.Second)
	freelancerProfileHandler := handler.NewFreelancerProfileHandler(freelancerProfileUsecase)

	jobApplicationRepo := postgres.NewJobApplicationRepository(db)
	jobApplicationUsecase := usecase.NewJobApplicationUsecase(jobApplicationRepo, 5*time.Second)
	jobApplicationHandler := handler.NewJobApplicationHandler(jobApplicationUsecase)

	categoryRepo := postgres.NewCategoryRepository(db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo, 5*time.Second)

	skillRepo := postgres.NewSkillRepository(db)
	skillUsecase := usecase.NewSkillUsecase(skillRepo, 5*time.Second)

	masterDataHandler := handler.NewMasterDataHandler(categoryUsecase, skillUsecase)

	kanbanTaskRepo := postgres.NewKanbanTaskRepository(db)
	kanbanTaskUsecase := usecase.NewKanbanTaskUsecase(kanbanTaskRepo, 5*time.Second)
	kanbanHandler := handler.NewKanbanTaskHandler(kanbanTaskUsecase)

	trxRepo := postgres.NewTransactionRepository(db)
	trxUsecase := usecase.NewTransactionUsecase(trxRepo, 5*time.Second)
	trxHandler := handler.NewTransactionHandler(trxUsecase)

	reviewRepo := postgres.NewReviewRepository(db)
	reviewUsecase := usecase.NewReviewUsecase(reviewRepo, 5*time.Second)
	reviewHandler := handler.NewReviewHandler(reviewUsecase)

	auditLogRepo := postgres.NewAuditLogRepository(db)

	httpDelivery.RegisterRoutes(app, authHandler, umkmProfileHandler, freelancerProfileHandler, jobHandler, jobApplicationHandler, masterDataHandler, kanbanHandler, trxHandler, reviewHandler, auditLogRepo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
