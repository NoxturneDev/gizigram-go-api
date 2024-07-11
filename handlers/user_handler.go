package handlers

import (
	"github.com/berkatps/model"
	"github.com/berkatps/services"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(model.Users)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	if services.CreateUser(user) != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Couldn't create user", "data": nil})
	}
	return c.JSON(&fiber.Map{"status": "success", "message": "Created user", "data": user})
}

func GetUser(c *fiber.Ctx) error {
	users, err := services.GetUser()
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Couldn't get users", "data": err})
	}
	return c.JSON(&fiber.Map{"status": "success", "message": "All users", "data": users})
}

func DeleteUser(c *fiber.Ctx) error {
	paramsInt, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Review your input", "data": err})

	}
	id := paramsInt
	if err := services.DeleteUser(id); err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Couldn't delete user", "data": err})
	}
	return c.JSON(&fiber.Map{"status": "success", "message": "Deleted user", "data": nil})
}
