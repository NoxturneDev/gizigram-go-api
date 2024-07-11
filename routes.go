package main

import (
	"github.com/berkatps/handlers"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App) {
	api := app.Group("/api")

	// health check
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api.Post("/users", handlers.CreateUser)
	api.Get("/users", handlers.GetUser)
	api.Delete("/user/:id", handlers.DeleteUser)

	api.Post("/parent/create", handlers.CreateParent)
	api.Get("/parents", handlers.ShowAllParent)
	api.Get("/parent/:id", handlers.ShowParrentByID)
	api.Post("/children/create", handlers.CreateChildren)
	api.Get("/children/match/:id", handlers.GetChildrenMatchByParentID)
	api.Get("/childrens", handlers.ShowAllChildren)
	api.Get("/children/:id", handlers.ShowChildrenByID)
	api.Delete("/children/:id", handlers.DeleteChidren)
	api.Delete("/parent/:id", handlers.DeleteParent)

	api.Post("/login", handlers.LoginUser)
	api.Post("/logout", handlers.LogoutUser)

}
