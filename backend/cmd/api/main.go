// @title Jobshare Backend API
// @version 1.0
// @description This is the official API documentation for the Jobshare platform (Auth, UMKM Profiles, and Job Vacancies).
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @BasePath /api/v1
package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	httpDelivery "github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/handler"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/repository/postgres"
	"github.com/FyaEdu/JOB-SHARE/backend/internal/usecase"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	dsn := "postgres://root:password@localhost:5432/jobshare?sslmode=disable"
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

	httpDelivery.RegisterRoutes(app, authHandler, jobHandler)

	log.Fatal(app.Listen(":8080"))
}
