package api

import "github.com/gofiber/fiber/v2"

func setRoutes(server *fiber.App) {
	authRoutes(server)
}

func authRoutes(server *fiber.App) {
	server.Post("/auth/register", controller.RegisterController)
}
