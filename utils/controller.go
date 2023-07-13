package utils

import (
	"fmt"

	"github.com/Arjun259194/wecup-go/types"
	"github.com/gofiber/fiber/v2"
)

func SendErrResponse(err error, message string, status int, c *fiber.Ctx) error {
	fmt.Printf("Error - %v", err)
	return c.Status(status).JSON(types.ErrorResponse{
		Status:  status,
		Message: message,
	})
}
