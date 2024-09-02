package polities

import (
	"context"
	"encoding/json"

	"github.com/agaUHO/aga/controllers"
	"github.com/agaUHO/aga/core"
	"github.com/agaUHO/aga/models"
	"github.com/agaUHO/aga/system"
	"github.com/go-ldap/ldap/v3"
	"github.com/gofiber/fiber/v2"
)

// LoggedIn --> Comprobar si el usuario tiene la sección activa
func LoggedIn(c *fiber.Ctx) error {
	cookie := c.Cookies(system.UserKeyAGA)
	if cookie == "" {
		return c.SendStatus(fiber.StatusForbidden)
	}
	return c.Next()
}

// LoggedInfo --> Comprobar la información de la sección
func LoggedInfo(c *fiber.Ctx) error {
	claims, _ := controllers.GetClaims(c)
	// ---------------------------------------------
	result := core.ExtractFunctionsPlugins("ldap", "Search", "(&(uid="+claims["ID"].(string)+"))")
	bytes, _ := json.Marshal(&result)
	var resultSearch ldap.SearchResult
	_ = json.Unmarshal(bytes, &resultSearch)
	var activeUser models.ActiveUser
	if len(resultSearch.Entries) > 0 {
		activeUser = controllers.GetUserInfo(resultSearch.Entries[0], claims["ID"].(string))
		system.AgaContext = context.WithValue(system.AgaContext, system.ActiveUserKey, activeUser)
		c.Context().SetUserValue("activeUser", activeUser)
	}
	//---------------------------------------------
	return c.Next()
}

// LoggedIsAdmin --> Comprobar si la sección es de un administrador
func LoggedIsAdmin(c *fiber.Ctx) error {
	claims, _ := controllers.GetClaims(c)
	if claims["ID"].(string) != "admin" {
		return c.SendStatus(fiber.StatusForbidden)
	}
	return c.Next()
}
