package api

import (
	"github.com/gofiber/fiber/v2"
)

func setRoutes(server *fiber.App) {
	setAuthRoutes(server)
	setUserRoutes(server)
	setPostRoutes(server)
	server.Get("/feed", controller.GetPostsController)
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
	user.Post("/:id/follow", controller.FollowUserController)
}

func setPostRoutes(server *fiber.App) {
	post := server.Group("/posts", controller.JWTMiddleware)
	post.Post("/", controller.CreatePostController)
	post.Put("/:id", controller.UpdatePostController)
	post.Get("/:id", controller.GetPostController)
}
