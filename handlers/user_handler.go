package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
	"gizigram-go-api/model"
	"gizigram-go-api/services"
	"log"
)

var store = session.New()

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

func LoginUser(c *fiber.Ctx) error {
	var userInput struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	ID := userInput.ID
	username := userInput.Username
	password := userInput.Password

	user, err := services.LoginUser(username, password)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Couldn't login user", "data": err})
	}
	sess := store.Get(c)
	log.Printf("session: %v", sess)

	sess.Set("user_id", ID)
	sess.Save()
	return c.JSON(&fiber.Map{"status": "success", "message": "Logged in user", "data": user})
}

func LogoutUser(c *fiber.Ctx) error {
	sess := store.Get(c)

	sess.Destroy()
	return c.JSON(&fiber.Map{"status": "success", "message": "Logged out user", "data": nil})
}
