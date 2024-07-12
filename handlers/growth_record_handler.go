package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gizigram-go-api/database"
	"gizigram-go-api/model"
	"gizigram-go-api/services"
	"gorm.io/gorm"
)

func CreateGrowthRecordHandler(c *fiber.Ctx) error {
	var children model.Children
	if err := c.BodyParser(&children); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		return services.CreateGrowthRecord(tx, &children)
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Growth record created successfully"})
}

func CreateNewGrowthRecordWithoutChildren(c *fiber.Ctx) error {
	var growthRecord model.GrowthRecord
	if err := c.BodyParser(&growthRecord); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := services.CreateGrowthRecordWithoutChildren(&growthRecord)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Growth record created successfully"})
}

func ShowGrowthRecordByChildrenIDHandler(c *fiber.Ctx) error {
	paramsInt, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	id := paramsInt
	growthRecords, err := services.ShowGrowthRecordByChildrenID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Growth records", "data": growthRecords})
}
