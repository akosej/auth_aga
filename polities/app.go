package polities

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

// PolityAppAuthorize --> Comprobar si la app que hace la petición esta autorizada 3185da79077acfaff40445b4365b4574
func AppAuthorize(c *fiber.Ctx) error {
	// Obtención del token de autenticación de la aplicación
	appAuthorization := c.Get("Authorization")
	appOrigin := c.Get("Origin")

	fmt.Println("------------")
	fmt.Println(c.Get("User-Agent"))
	fmt.Println(c.Get("Origin"))
	fmt.Println("------------")
	// fmt.Println(c.Hostname(), "---", c.IP(), "---", c.IsFromLocal(), "---")

	if appAuthorization == "" {
		// comprobar si existe algun registro
		var app models.AuthorizedApp

		app.Origin = appOrigin
		app.Ipadress = c.IP()
		fmt.Println(c.IP(), c.Get("Origin"))
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
		// Verificación de campos requeridos
		var missingFields []string
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
		if err := database.DB.Where("origin = ? AND ipaddress = ?", appOrigin, app.Ipadress).First(&existingApp).Error; err != nil {
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
				Description: core.GetTextMessage("authorizedApp_ok_001", string(app.Origin)),
				App:         string(app.Origin),
			}
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"OK":      true,
				"message": core.GetTextMessage("authorizedApp_ok_001", string(app.Secret)),
			})

		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"OK":      false,
			"message": core.GetTextMessage("authorizedApp_err_003", "global_info_002"),
		})
	}

	token := strings.TrimPrefix(appAuthorization, "AGA ")
	// Verificación del token con la clave secreta de la aplicación
	var app models.AuthorizedApp
	if err := database.DB.Where("secret = ?", token).First(&app).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"OK":      false,
			"message": core.GetTextMessage("authorizedApp_err_004", "global_info_002"),
		})
	}
	if !app.Active {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"OK":      false,
			"message": core.GetTextMessage("authorizedApp_err_005", "global_info_002"),
		})
	}
	if app.Origin != appOrigin {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"OK":      false,
			"message": core.GetTextMessage("authorizedApp_err_007", "global_info_002"),
		})
	}
	// Si la aplicación está autorizada
	return c.Next()
}

func AppAGA(c *fiber.Ctx) error {
	for k, v := range c.GetReqHeaders() {
		fmt.Println(k, v)
	}
	fmt.Println("------------")
	fmt.Println(c.Get("User-Agent"))
	fmt.Println(c.Get("Origin"))
	fmt.Println("------------")
	fmt.Println(c.Hostname(), "---", c.IP(), "---", c.IsFromLocal(), "---")
	return c.JSON(fiber.Map{
		"hostname":       c.Hostname(),
		"ip":             c.IP(),
		"IsFromLocal":    c.IsFromLocal(),
		"IsProxyTrusted": c.IsProxyTrusted(),
		"Headers":        c.GetReqHeaders(),
	})
}
