package main

import (
	"github.com/berkatps/database"
	"github.com/berkatps/model"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	app.Use(logger.New())

	NewRouter(app)

	log.Fatal(app.Listen(":3002"))
}

