package api

import (
	"github.com/gofiber/fiber/v2"
)

func setRoutes(server *fiber.App) {
	setAuthRoutes(server)
	setUserRoutes(server)
}

func setAuthRoutes(server *fiber.App) {
	auth := server.Group("/auth")
	auth.Post("/register", controller.RegisterController)
	auth.Post("/login", controller.LoginController)
	auth.Post("/logout", controller.LogoutController)
}

func setUserRoutes(server *fiber.App) {
	user := server.Group("/user", controller.JWTMiddleware)
	user.Get("/profile", controller.GetProfile)
	user.Get("/:id", controller.GetUserController)
	user.Put("/profile/update", controller.UpdateUserController)
}
