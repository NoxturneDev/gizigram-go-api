package handlers

import (
	"github.com/berkatps/database"
	"github.com/berkatps/model"
	"github.com/berkatps/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateParent(c *fiber.Ctx) error {
	parrent := new(model.Parent)
	if err := c.BodyParser(parrent); err != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	if services.CreateParent(parrent) != nil {
		return c.Status(500).JSON(&fiber.Map{"status": "error", "message": "Couldn't create parent", "data": nil})
	}
	return c.JSON(&fiber.Map{"status": "success", "message": "Created parent", "data": parrent})
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
