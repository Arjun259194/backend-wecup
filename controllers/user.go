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

	return utils.SendOKResponse(c, user.GetResponse())
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

	return utils.SendOKResponse(c, user.GetResponse())
}

func (ctrl *Controller) UpdateUserController(c *fiber.Ctx) error {
	ID := c.Locals("id").(primitive.ObjectID)

	reqBody := new(types.UserUpdateRequest)
	if err := c.BodyParser(reqBody); err != nil {
		return utils.ReqBodyFailedToDecodeHandler(c, err)
	}

	if err := val.Struct(reqBody); err != nil {
		return utils.InvalidRequestBodyHandler(c, err)
	}

	if err := ctrl.Storage.FindByIDAndUpdateUser(ID, *reqBody); err != nil {
		return utils.SingleUserErrorHandler(err, c)
	}

	return utils.SendOKResponse(c, nil)
}

func (ctrl *Controller) FollowUserController(c *fiber.Ctx) error {
	ID, err := utils.GetIDFromParams(c)
	if err != nil {
		utils.NotValidIDHandler(c, err)
	}
	clientID := c.Locals("id").(primitive.ObjectID)

	// Client is the user who is sending this request
	client, err := ctrl.GetUserByID(clientID)
	if err != nil {
		utils.SingleUserErrorHandler(err, c)
	}

	isFollowing := false

	for _, followingID := range client.Following {
		if followingID == ID {
			isFollowing = true
		}
	}

	if err := ctrl.Storage.FindByIDAndFollowOrUnfollow(ID, clientID, isFollowing); err != nil {
		return utils.SingleUserErrorHandler(err, c)
	}

	return utils.SendOKResponse(c, nil)
}
