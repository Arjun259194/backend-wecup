package controllers

import (
	"github.com/Arjun259194/wecup-go/types"
	"github.com/Arjun259194/wecup-go/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ctrl *Controller) GetProfile(c *fiber.Ctx) error {
	ID := c.Locals("id").(primitive.ObjectID)
	user, err := ctrl.GetUserByID(ID)

	if err != nil {
		sendUserErrResponse(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(types.Response{
		Status:       fiber.StatusOK,
		ResponseData: user.GetResponse(),
	})
}

func (ctrl *Controller) GetUserController(c *fiber.Ctx) error {
	strID := c.Params("id")

  // MongoDB uses primitive.ObjectID type for ID
	ID, err := primitive.ObjectIDFromHex(strID)
	if err != nil {
		return utils.SendErrResponse(err, "ID not valid", fiber.StatusBadRequest, c)
	}

	user, err := ctrl.GetUserByID(ID)

	if err != nil {
		sendUserErrResponse(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(types.Response{
		Status:       fiber.StatusOK,
		ResponseData: user.GetResponse(),
	})
}

func sendUserErrResponse(err error, c *fiber.Ctx) error {
	var responseStatus, responseMessage = fiber.StatusInternalServerError, "Error while fetching data"

	if err == mongo.ErrNoDocuments {
		responseStatus, responseMessage = fiber.StatusNotFound, "Document not found"
	}

	return utils.SendErrResponse(err, responseMessage, responseStatus, c)
}
