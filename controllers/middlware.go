package controllers

import (
	"github.com/Arjun259194/wecup-go/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ctrl *Controller) JWTMiddleware(c *fiber.Ctx) error {
	tokenCookie := c.Cookies("accessToken")
	if tokenCookie == "" {
		return utils.SendErrResponse(nil, "Token not found", fiber.StatusUnauthorized, c)
	}

	claim, err := utils.VerifyToken(tokenCookie)
	if err != nil {
		return utils.SendErrResponse(err, "User not authorized", fiber.StatusUnauthorized, c)
	}

	id := claim["id"].(string)

	// Here we are converting string Id into mongoDB objectID because to make sure that it's a valid objectID and because our Go driver only uses objectID type
	userObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return utils.SendErrResponse(err, "User id not valid", fiber.StatusBadRequest, c)
	}

	c.Locals("id", userObjectID)

	return c.Next()
}
