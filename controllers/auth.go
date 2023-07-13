package controllers

import (
	"encoding/json"

	"github.com/Arjun259194/wecup-go/types"
	"github.com/Arjun259194/wecup-go/utils"
	"github.com/gofiber/fiber/v2"
)

func (ctrl *Controller) RegisterController(c *fiber.Ctx) error {
	var status int
	reqBody := c.Body()

	var reqBodyStruct registerRequestBody

	if err := json.Unmarshal(reqBody, &reqBodyStruct); err != nil {
    // make this error handling DRY
		status = fiber.StatusBadRequest
		return c.Status(status).JSON(errorResponse{
			Status:  status,
			Message: "Please check request body",
			Error:   err.Error(),
		})
	}

	hashPass, err := utils.HashPassword(reqBodyStruct.Password)

	if err != nil {
		status = fiber.StatusInternalServerError
		return c.Status(status).JSON(errorResponse{
			Status:  status,
			Message: "Error while inserting into database",
			Error:   err.Error(),
		})
	}

	newUser := types.NewUser(reqBodyStruct.Username, reqBodyStruct.Email, hashPass, reqBodyStruct.Gender)

	if _, err := ctrl.Storage.AddUser(newUser); err != nil {
		status = fiber.StatusInternalServerError
		return c.Status(status).JSON(errorResponse{
			Status:  status,
			Message: "Error while inserting into database",
			Error:   err.Error(),
		})
	}

	status = fiber.StatusOK

	return c.Status(status).JSON(reponse{
		Status:      status,
		Message:     "New user created",
		ReponseData: nil,
	})
}
