package response

import "github.com/gofiber/fiber/v2"

type WebResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

func Success(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(WebResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func SuccessWithPagination(c *fiber.Ctx, statusCode int, message string, data interface{}, meta interface{}) error {
	return c.Status(statusCode).JSON(WebResponse{
		Status:  "success",
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func Error(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(WebResponse{
		Status:  "error",
		Message: message,
	})
}
