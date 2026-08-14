package middleware

import (
	"context"
	"encoding/json"

	"github.com/FyaEdu/JOB-SHARE/backend/internal/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuditLogMiddleware(repo domain.AuditLogRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		go func() {
			var userID *string

			userToken := c.Locals("user")
			if userToken != nil {
				token := userToken.(*jwt.Token)
				claims := token.Claims.(jwt.MapClaims)
				if id, ok := claims["user_id"].(string); ok {
					userID = &id
				}
			}

			method := c.Method()
			path := c.Path()
			action := method + " " + path

			var details map[string]interface{}
			if method == "POST" || method == "PUT" || method == "PATCH" {
				_ = json.Unmarshal(c.Body(), &details)
				if _, exists := details["password"]; exists {
					details["password"] = "***MASKED***"
				}
			}
			detailsJSON, _ := json.Marshal(details)
			if string(detailsJSON) == "null" {
				detailsJSON = []byte("{}")
			}

			logData := &domain.AuditLog{
				UserID:    userID,
				Action:    action,
				IPAddress: c.IP(),
				Details:   string(detailsJSON),
			}

			repo.Create(context.Background(), logData)
		}()

		return err
	}
}
