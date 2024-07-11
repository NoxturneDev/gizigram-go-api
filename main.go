package main

import (
	"github.com/berkatps/database"
	"github.com/berkatps/model"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDatabase()

	database.DB.AutoMigrate(&model.Users{}, &model.Parent{}, &model.Children{})

	NewRouter(app)

	app.Listen(":3002")
}
