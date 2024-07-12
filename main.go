package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gizigram-go-api/database"
	"gizigram-go-api/model"
	"gizigram-go-api/routes"
	"log"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	database.ConnectDatabase()
	database.DB.AutoMigrate(&model.Users{}, &model.Parent{}, &model.Children{}, &model.GrowthRecord{}, &model.GrowthResult{})

	routes.NewRouter(app)
	log.Println("Server is running on port 8080")
	app.Listen(":8080")
}
