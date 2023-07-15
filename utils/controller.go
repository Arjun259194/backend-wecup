package utils

import (
	"fmt"

	"github.com/Arjun259194/wecup-go/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SendErrResponse(err error, message string, status int, c *fiber.Ctx) error {
	fmt.Printf("Error - %v\n", err)
	return c.Status(status).JSON(types.ErrorResponse{
		Status:  status,
		Message: message,
	})
}


func GetIDFromParams(c *fiber.Ctx) (primitive.ObjectID, error) {
	strID := c.Params("id")
	return primitive.ObjectIDFromHex(strID)
}
