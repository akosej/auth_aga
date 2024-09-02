package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/agaUHO/aga/core"
	"github.com/agaUHO/aga/database"
	"github.com/agaUHO/aga/models"
	"github.com/agaUHO/aga/system"
	"github.com/gofiber/fiber/v2"
)

/**
 * Func: AuthorizeAppRegisterHandler is for ...
 *
 * @author GoCommnets
 *
 * @params ...
 * @return
 */

func AuthorizeAppRegisterHandler(c *fiber.Ctx) error {

	// Decodificación de la información de la aplicación
	var app models.AuthorizedApp
	if err := c.BodyParser(&app); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": core.GetTextMessage("authorizedApp_err_001"),
		})
	}

	// Generar un SecretKey aleatorio
	const randomByteSize = 16
	randomBytes := make([]byte, randomByteSize)
	if _, err := rand.Read(randomBytes); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": core.GetTextMessage("authorizedApp_err_006"),
		})
	}
	app.Secret = hex.EncodeToString(randomBytes)
	app.Origin = c.Get("Origin")
	app.Ipadress = c.IP()
	// Verificación de campos requeridos
	var missingFields []string
	if app.Name == "" {
		missingFields = append(missingFields, core.GetTextMessage("authorizedApp_info_001"))
	}
	if app.Origin == "" {
		missingFields = append(missingFields, core.GetTextMessage("authorizedApp_info_002"))
	}
	if len(missingFields) > 0 {
		errorMessage := fmt.Sprintf("%s: %v.", core.GetTextMessage("global_info_001"), strings.Join(missingFields, ", "))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": errorMessage,
		})
	}

	// Buscar la aplicación en la base de datos utilizando cualquiera de los campos especificados
	var existingApp models.AuthorizedApp
	if err := database.DB.Where("name = ? OR origin = ?", app.Name, app.Origin).First(&existingApp).Error; err != nil {
		// si no se encuentra la aplicación, crearla
		if err := database.DB.Create(&app).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"OK":      false,
				"message": core.GetTextMessage("authorizedApp_err_002"),
			})
		}
		system.Log <- models.Logs{
			Module:      "AGA",
			Action:      "AUTHAPP_ADD",
			Login:       "aga",
			Uid:         "aga",
			Description: core.GetTextMessage("authorizedApp_ok_001", app.Name),
		}
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"OK":      false,
			"message": core.GetTextMessage("authorizedApp_info_004"),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"OK":      true,
		"message": core.GetTextMessage("authorizedApp_ok_001", "Su llave es: "+app.Secret),
	})
}
