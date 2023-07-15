package utils

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func DatabaseInsertionHandler(c *fiber.Ctx, err error) error {
	return SendErrResponse(err, "Error while inserting into database", fiber.StatusInternalServerError, c)
}

func InvalidRequestBodyHandler(c *fiber.Ctx, err error) error {
	return SendErrResponse(err, "Request Body Is Not Valid", fiber.StatusBadRequest, c)
}

func ReqBodyFailedToDecodeHandler(c *fiber.Ctx, err error) error {
	return SendErrResponse(err, "Please Check Request Body", fiber.StatusBadRequest, c)
}

func NotValidIDHandler(c *fiber.Ctx, err error) error {
	return SendErrResponse(err, "ID not valid", fiber.StatusBadRequest, c)
}

func SingleUserErrorHandler(err error, c *fiber.Ctx) error {
	var responseStatus, responseMessage = fiber.StatusInternalServerError, "Error while fetching data"

	if err == mongo.ErrNoDocuments {
		responseStatus, responseMessage = fiber.StatusNotFound, "Document not found"
	}

	return SendErrResponse(err, responseMessage, responseStatus, c)
}
