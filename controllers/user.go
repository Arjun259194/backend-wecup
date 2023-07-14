package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (ctrl *Controller) GetUserController(c *fiber.Ctx) error {
  fmt.Println("Route endpoint reached")
	return c.SendString("User route running")
}
