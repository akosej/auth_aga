package controllers

import (
	"github.com/agaUHO/aga/database"
	"github.com/agaUHO/aga/models"
	"github.com/gofiber/fiber/v2"
)

func PluginConfigGetHandler(c *fiber.Ctx) error {
	var pluginsConfig []models.PluginConfig
	database.DB.Find(&pluginsConfig)
	return c.Status(200).JSON(pluginsConfig)
}

func PluginConfigPutHandler(c *fiber.Ctx) error {
	pluginConfig := new(models.PluginConfig)
	id := c.Params("id")

	if err := c.BodyParser(pluginConfig); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.DB.Where("id = ?", id).Updates(&pluginConfig)
	return c.Status(200).JSON(pluginConfig)
}

func StaticPluginConfigGetHandler(c *fiber.Ctx) error {
	var staticPluginsConfig []models.StaticPluginConfig
	database.DB.Find(&staticPluginsConfig)
	return c.Status(200).JSON(staticPluginsConfig)
}

func StaticPluginConfigPutHandler(c *fiber.Ctx) error {
	staticPluginConfig := new(models.StaticPluginConfig)
	id := c.Params("id")

	if err := c.BodyParser(staticPluginConfig); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.DB.Where("id = ?", id).Updates(&staticPluginConfig)
	return c.Status(200).JSON(staticPluginConfig)
}
