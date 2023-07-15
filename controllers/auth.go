package controllers

import (
	"encoding/json"

	"github.com/Arjun259194/wecup-go/types"
	"github.com/Arjun259194/wecup-go/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func (ctrl *Controller) RegisterController(c *fiber.Ctx) error {
	reqBody := c.Body()

	var reqBodyStruct types.RegisterRequest

	if err := json.Unmarshal(reqBody, &reqBodyStruct); err != nil {
		return utils.ReqBodyFailedToDecodeHandler(c, err)
	}

	if err := val.Struct(reqBodyStruct); err != nil {
		return utils.InvalidRequestBodyHandler(c, err)
	}

	hashPass, err := utils.HashPassword(reqBodyStruct.Password)
	if err != nil {
		return utils.SendErrResponse(err, "Error while hashing password", fiber.StatusInternalServerError, c)
	}

	newUser := types.NewUser(reqBodyStruct.Name, reqBodyStruct.Email, hashPass, reqBodyStruct.Gender)

	if _, err := ctrl.Storage.AddUser(newUser); err != nil {
		return utils.SendErrResponse(err, "Error while inserting into database", fiber.StatusInternalServerError, c)
	}

	return c.Status(fiber.StatusOK).JSON(types.Response{
		Status:       fiber.StatusOK,
		ResponseData: nil,
	})
}

func (ctrl *Controller) LoginController(c *fiber.Ctx) error {
	reqBody := new(types.LoginRequest)
	if err := c.BodyParser(&reqBody); err != nil {
		return utils.ReqBodyFailedToDecodeHandler(c, err)
	}

	if err := val.Struct(reqBody); err != nil {
		return utils.InvalidRequestBodyHandler(c, err)
	}

	filter := bson.M{"email": reqBody.Email}
	foundUser, err := ctrl.Storage.FindOneUser(filter)

	if err != nil {
		utils.SingleUserErrorHandler(err, c)
	}

	if err := utils.ComparePassword(reqBody.Password, foundUser.Password); err != nil {
		return utils.SendErrResponse(err, "incorrect password", fiber.StatusUnauthorized, c)
	}

	token, err := utils.GenerateToken(foundUser.ID)
	if err != nil {
		return utils.SendErrResponse(err, "Error while generating token", fiber.StatusInternalServerError, c)
	}

	cookie := utils.NewHTTPOnlyCookie(token)

	c.Cookie(cookie)

	return c.Status(fiber.StatusOK).JSON(types.Response{
		Status: fiber.StatusOK,
		ResponseData: fiber.Map{
			"token": token,
		},
	})
}

func (ctrl *Controller) LogoutController(c *fiber.Ctx) error {
	emptyCookie := utils.EmptyCookie()
	c.Cookie(emptyCookie)
	return c.SendStatus(fiber.StatusOK)
}
