package utils

import (
	"fmt"

	"github.com/Arjun259194/wecup-go/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SendErrResponse(err error, message string, status int, c *fiber.Ctx) error {
	fmt.Printf("Error - %v", err)
	return c.Status(status).JSON(types.ErrorResponse{
		Status:  status,
		Message: message,
	})
}


func UserErrorHandler(err error, c *fiber.Ctx) error {
	var responseStatus, responseMessage = fiber.StatusInternalServerError, "Error while fetching data"

	if err == mongo.ErrNoDocuments {
		responseStatus, responseMessage = fiber.StatusNotFound, "Document not found"
	}

	return SendErrResponse(err, responseMessage, responseStatus, c)
}
