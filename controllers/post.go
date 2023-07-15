package controllers

import (
	"github.com/Arjun259194/wecup-go/types"
	"github.com/Arjun259194/wecup-go/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ctrl *Controller) CreatePostController(c *fiber.Ctx) error {
	ID := c.Locals("id").(primitive.ObjectID)

	reqBody := new(types.CreatePostRequest)
	if err := c.BodyParser(reqBody); err != nil {
		return utils.ReqBodyFailedToDecodeHandler(c, err)
	}

	if err := val.Struct(reqBody); err != nil {
		return utils.InvalidRequestBodyHandler(c, err)
	}

	newPost := types.NewPost(ID, reqBody.Content)

	if err := ctrl.Storage.CreateNewPost(*newPost); err != nil {
		return utils.DatabaseInsertionHandler(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(types.Response{
		Status:       fiber.StatusOK,
		ResponseData: nil})
}

func (ctrl *Controller) GetPostController(c *fiber.Ctx) error {
	ID, err := utils.GetIDFromParams(c)
	if err != nil {
		return utils.NotValidIDHandler(c, err)
	}

	post, err := ctrl.Storage.FindOnePostByID(ID)
	if err != nil {
		return utils.SingleUserErrorHandler(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(types.Response{
		Status:       fiber.StatusOK,
		ResponseData: post})
}
