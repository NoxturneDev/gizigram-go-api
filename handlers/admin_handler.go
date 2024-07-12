package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gizigram-go-api/database"
	"gizigram-go-api/model"
	"gizigram-go-api/services"
	"gorm.io/gorm"
)

type ParentPayload struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

func CreateParent(c *fiber.Ctx) error {
	payload := new(ParentPayload)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	user := services.GetUserByPhoneNumber(payload.PhoneNumber)
	if user != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "User already exists", "data": user})
	}

	var parent model.Parent
	parent.Name = payload.Name
	if services.CreateParent(&parent, payload.PhoneNumber) != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Couldn't create parent", "data": nil})
	}

	return c.JSON(&fiber.Map{"status": "success", "message": "Created parent", "data": payload})
}

func ShowAllParent(c *fiber.Ctx) error {
	parents, err := services.ShowAllParrent()
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Couldn't get parents", "data": err})
	}
	return c.JSON(&fiber.Map{"status": "success", "message": "All parents", "data": parents})
}

func ShowParrentByID(c *fiber.Ctx) error {
	paramsInt, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	id := paramsInt
	parent, err := services.ShowParrentByID(id)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Couldn't get parent", "data": err})
	}
	return c.JSON(&fiber.Map{"status": "success", "message": "parent", "data": parent})
}

func CreateChildrenHandler(c *fiber.Ctx) error {
	var children model.Children
	if err := c.BodyParser(&children); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		return services.CreateChildren(tx, &children)
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "message": "Couldn't create children", "data": err.Error()})
	}

	return c.JSON(&fiber.Map{"status": "success", "message": "Created children", "data": children})
}

func GetChildrenMatchByParentID(c *fiber.Ctx) error {
	paramsInt, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	id := paramsInt
	children, err := services.GetChildrenMatchByParentID(id)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Couldn't get children", "data": err})
	}
	return c.JSON(&fiber.Map{"status": "success", "message": "children", "data": children})
}

func ShowAllChildren(c *fiber.Ctx) error {
	children, err := services.ShowAllChildren()
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Couldn't get children", "data": err})
	}
	return c.JSON(&fiber.Map{"status": "success", "message": "All children", "data": children})
}

func ShowChildrenByID(c *fiber.Ctx) error {
	paramsInt, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	id := paramsInt
	children, err := services.ShowChildrenByID(id)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Couldn't get children", "data": err})
	}
	return c.JSON(&fiber.Map{"status": "success", "message": "children", "data": children})
}

func DeleteChidren(c *fiber.Ctx) error {
	paramsInt, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	id := paramsInt
	if err := services.DeleteChildren(id); err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Couldn't delete children", "data": err})
	}
	return c.JSON(&fiber.Map{"status": "success", "message": "Deleted children", "data": nil})
}

func DeleteParent(c *fiber.Ctx) error {
	paramsInt, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	id := paramsInt
	if err := services.DeleteParent(id); err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Couldn't delete parent", "data": err})
	}
	return c.JSON(&fiber.Map{"status": "success", "message": "Deleted parent", "data": nil})
}
