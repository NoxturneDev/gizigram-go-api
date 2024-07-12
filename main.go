package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gizigram-go-api/database"
	"gizigram-go-api/model"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	database.ConnectDatabase()
	database.DB.AutoMigrate(&model.Users{}, &model.Parent{}, &model.Children{}, &model.GrowthRecord{})

	NewRouter(app)

	app.Listen(":3002")
}
