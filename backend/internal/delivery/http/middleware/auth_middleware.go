package middleware

import (
	"strings"

	jwtPkg "github.com/FyaEdu/JOB-SHARE/backend/pkg/jwt"
	pkgResponse "github.com/FyaEdu/JOB-SHARE/backend/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return pkgResponse.Error(c, fiber.StatusUnauthorized, "Missing or malformed JWT")
		}

		// Format token harus "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return pkgResponse.Error(c, fiber.StatusUnauthorized, "Invalid token format")
		}

		tokenString := parts[1]
		claims, err := jwtPkg.ValidateToken(tokenString)
		if err != nil {
			return pkgResponse.Error(c, fiber.StatusUnauthorized, "Invalid or expired token")
		}

		c.Locals("user_id", claims.UserID)
		c.Locals("role", claims.Role)

		return c.Next()
	}
}

func RoleGuard(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole, ok := c.Locals("role").(string)
		if !ok {
			return pkgResponse.Error(c, fiber.StatusInternalServerError, "Role not found in context")
		}

		for _, role := range allowedRoles {
			if role == userRole {
				return c.Next()
			}
		}

		return pkgResponse.Error(c, fiber.StatusForbidden, "You do not have permission to access this resource")
	}
}
