package routes

import (
	"github.com/gofiber/fiber/v2"
	"gizigram-go-api/handlers"
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

	api.Get("/parents", handlers.ShowAllParent)
	api.Get("/parent/:id", handlers.ShowParrentByID)
	api.Get("/parent-options", handlers.ShowParentOptions)
	api.Post("/parent/create", handlers.CreateParent)

	api.Post("/children/create", handlers.CreateChildrenHandler)
	api.Get("/children/match/:id", handlers.GetChildrenMatchByParentID)
	api.Get("/childrens", handlers.ShowAllChildren)
	api.Get("/children/:id", handlers.ShowChildrenByID)

	api.Delete("/children/:id", handlers.DeleteChidren)
	api.Delete("/parent/:id", handlers.DeleteParent)

	api.Post("/login", handlers.LoginUser)
	api.Post("/logout", handlers.LogoutUser)

	api.Post("/growth-add", handlers.CreateNewGrowthRecordWithoutChildren)
	api.Post("/growth/create", handlers.CreateGrowthRecordHandler)
	api.Get("/growth/:id", handlers.ShowGrowthRecordByChildrenIDHandler)

	api.Post("/ai/scanner", handlers.AiScanner)
}
