package controllers

import (
	"context"

	"github.com/Arjun259194/wecup-go/types"
	"github.com/Arjun259194/wecup-go/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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

func (ctrl *Controller) GetPostsController(c *fiber.Ctx) error {
	cur, err := ctrl.Storage.PostModel.Find(context.Background(), bson.M{})
	if err != nil {
		return utils.DatabaseFetchHandler(c, err)
	}
	defer cur.Close(context.Background())

	posts := new(types.Posts)
	if err := cur.All(context.Background(), posts); err != nil {
		return utils.FailedToDecodeDataHandler(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(types.Response{
		Status:       fiber.StatusOK,
		ResponseData: posts,
	})
}

func (ctrl *Controller) UpdatePostController(c *fiber.Ctx) error {
	ID, err := utils.GetIDFromParams(c)
	if err != nil {
		return utils.NotValidIDHandler(c, err)
	}

	reqBody := new(types.UpdatePostRequest)
	if err := c.BodyParser(reqBody); err != nil {
		return utils.ReqBodyFailedToDecodeHandler(c, err)
	}

	clientID := c.Locals("id").(primitive.ObjectID)

	post, err := ctrl.Storage.FindOnePostByID(ID)
	if err != nil {
		return utils.SingleUserErrorHandler(err, c)
	}

	if post.UserID != clientID {
		return utils.SendErrResponse(err, "User not authorized", fiber.StatusUnauthorized, c)
	}

	if err := ctrl.Storage.FindOnePostByIDAndUpdate(ID, reqBody); err != nil {
		return utils.SingleUserErrorHandler(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(types.Response{
		Status:       fiber.StatusOK,
		ResponseData: nil,
	})
}

func (ctrl *Controller) DeletePostController(c *fiber.Ctx) error {
	ID, err := utils.GetIDFromParams(c)
	if err != nil {
		return utils.NotValidIDHandler(c, err)
	}

	clientID := c.Locals("id").(primitive.ObjectID)

	post, err := ctrl.Storage.FindOnePostByID(ID)
	if err != nil {
		return utils.SingleUserErrorHandler(err, c)
	}

	if post.UserID != clientID {
		return utils.SendErrResponse(err, "User not authorized", fiber.StatusUnauthorized, c)
	}

	if err := ctrl.Storage.FindOnePostByIDAndDelete(ID); err != nil {
		return utils.SingleUserErrorHandler(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(types.Response{
		Status:       fiber.StatusOK,
		ResponseData: nil,
	})
}

func (ctrl *Controller) LikePostController(c *fiber.Ctx) error {

	clientID := c.Locals("id").(primitive.ObjectID)

	ID, err := utils.GetIDFromParams(c)
	if err != nil {
		return utils.NotValidIDHandler(c, err)
	}

	post, err := ctrl.Storage.FindOnePostByID(ID)
	if err != nil {
		return utils.SingleUserErrorHandler(err, c)
	}

	isLiked := false

	for _, likerID := range post.Likes {
		if likerID == clientID {
			isLiked = true
		}
	}

	err = ctrl.Storage.LikeOrUnlikePost(ID, clientID, isLiked)

	if err != nil {
		return utils.SingleUserErrorHandler(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(types.Response{
		Status:       fiber.StatusOK,
		ResponseData: nil,
	})
}
