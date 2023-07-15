package controllers

import (
	"github.com/Arjun259194/wecup-go/types"
	"github.com/Arjun259194/wecup-go/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ctrl *Controller) GetProfile(c *fiber.Ctx) error {
	ID := c.Locals("id").(primitive.ObjectID)
	user, err := ctrl.GetUserByID(ID)

	if err != nil {
		utils.SingleUserErrorHandler(err, c)
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
		return utils.NotValidIDHandler(c, err)
	}

	user, err := ctrl.GetUserByID(ID)

	if err != nil {
		utils.SingleUserErrorHandler(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(types.Response{
		Status:       fiber.StatusOK,
		ResponseData: user.GetResponse(),
	})
}

func (ctrl *Controller) UpdateUserController(c *fiber.Ctx) error {
	ID := c.Locals("id").(primitive.ObjectID)

	reqBody := new(types.UpdateRequest)
	if err := c.BodyParser(reqBody); err != nil {
		return utils.ReqBodyFailedToDecodeHandler(c, err)
	}

	if err := val.Struct(reqBody); err != nil {
		return utils.InvalidRequestBodyHandler(c, err)
	}

	if err := ctrl.Storage.FindByIDAndUpdateUser(ID, *reqBody); err != nil {
		return utils.SingleUserErrorHandler(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(types.Response{
		Status:       fiber.StatusOK,
		ResponseData: nil,
	})
}

// func (ctrl *Controller) FollowUserController(c *fiber.Ctx) error {
// 	ID, err := utils.GetIDFromParams(c)
// 	if err != nil {

// 	}
// }
