package controllers

import (
	"github.com/agaUHO/aga/database"
	"github.com/agaUHO/aga/models"
	"github.com/gofiber/fiber/v2"
)

func PermissionGetHandler(c *fiber.Ctx) error {
	var systemPermission []models.SystemPermissions
	database.DB.Find(&systemPermission)
	return c.Status(fiber.StatusOK).JSON(systemPermission)
}

func UserPermissionGetHandler(c *fiber.Ctx) error {
	var userPermission []models.UserPermissions
	id := c.Params("id")
	database.DB.Where("user_id = ?", id).Find(&userPermission)
	return c.Status(fiber.StatusOK).JSON(userPermission)
}

func UserPermissionPutHandler(c *fiber.Ctx) error {
	userPermission := new(models.UserPermissions)
	id := c.Params("id")

	if err := c.BodyParser(userPermission); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.DB.Where("id = ?", id).Updates(&userPermission)
	return c.Status(fiber.StatusCreated).JSON(userPermission)
}

func UserPermissionPostHandler(c *fiber.Ctx) error {
	userPermission := new(models.UserPermissions)

	if err := c.BodyParser(userPermission); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.DB.Create(&userPermission)
	return c.Status(fiber.StatusCreated).JSON(userPermission)
}

func UserPermissionDeleteHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var userPermission models.UserPermissions

	result := database.DB.Delete(&userPermission, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendStatus(fiber.StatusOK)
}
