package controllers

import (
	"time"

	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/agaUHO/aga/core"
	"github.com/agaUHO/aga/database"
	"github.com/agaUHO/aga/models"
	"github.com/agaUHO/aga/system"
	"github.com/go-ldap/ldap/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

func authenticateUser(c *fiber.Ctx, resultSearch ldap.SearchResult, username string, password string) (models.ActiveUser, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = username
	claims["Name"] = resultSearch.Entries[0].GetAttributeValue("cn")
	claims["Admin"] = resultSearch.Entries[0].GetAttributeValue("businessCategory")
	claims["Password"] = password
	claims["Expires"] = time.Now().Add(time.Hour * 8).Unix() // El token expira en 4 horas
	tokenString, err := token.SignedString([]byte(system.SecretKeyAGA))
	if err != nil {
		return models.ActiveUser{}, errors.New("InvalidSecretKey")
	}

	cookie := fiber.Cookie{
		Name:     system.UserKeyAGA,
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 4),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	// -- Create overlapping
	// CreateAccessCard(resultSearch.Entries[0])

	system.Log <- models.Logs{
		Module:      "AGA",
		Action:      "LOGIN_SUCCESS",
		Login:       username,
		Uid:         username,
		Description: core.GetTextMessage("login_ok_001"),
		App:         core.GetApp(c),
	}
	activeUser := GetUserInfo(resultSearch.Entries[0], username)
	database.DB.Model(&models.User{}).Where("username = ?", claims["ID"].(string)).Update("accept_system_policies", true)
	system.AgaContext = context.WithValue(system.AgaContext, system.ClaimsKey, claims)
	system.AgaContext = context.WithValue(system.AgaContext, system.ActiveUserKey, activeUser)
	system.AgaContext = context.WithValue(system.AgaContext, system.LoggedInUserKey, claims["Name"].(string)+" ["+claims["ID"].(string)+"]")

	c.Context().SetUserValue("claims", claims)
	c.Context().SetUserValue("activeUser", activeUser)
	c.Context().SetUserValue("loggedInUser", claims["Name"].(string)+" ["+claims["ID"].(string)+"]")

	return activeUser, nil
}

func GetClaims(c *fiber.Ctx) (jwt.MapClaims, error) {
	cookie := c.Cookies(system.UserKeyAGA)
	// Verificar el token JWT
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		// Verificar el método de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Método de firma inválido")
		}
		// Devolver la clave secreta
		return []byte(system.SecretKeyAGA), nil
	})
	if err != nil {
		return nil, err // Devuelve el error
	}

	// Obtener los datos del usuario del token JWT
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("No se pudieron obtener los datos del usuario del token JWT")
	}
	return claims, nil
}

func GetUserInfo(user *ldap.Entry, username string) models.ActiveUser {
	accountState := "FALSE"

	var personalInformation models.PersonalInformation
	var accountInfo models.AccountInfo

	resultHow, days := system.HowManyDaysAgo(user.GetAttributeValue("userPasswordSet"))
	accountState = user.GetAttributeValue("accountState")
	personalInformation.Dni = user.GetAttributeValue("dni")
	personalInformation.Cn = user.GetAttributeValue("cn")
	personalInformation.GivenName = user.GetAttributeValue("givenName")
	personalInformation.Sn = user.GetAttributeValue("sn")
	personalInformation.PersonalPhoto = system.ImgToB64(user.GetAttributeValue("dni"), "peronal_photo")
	personalInformation.Overlapping = system.ImgToB64(user.GetAttributeValue("dni"), "overlappin")

	accountInfo.UserType = user.GetAttributeValue("userType")
	accountInfo.CreateUser = user.GetAttributeValue("createUser")
	createData := strings.Split(user.GetAttributeValue("createDate"), ".")
	accountInfo.CreateDate = createData[0]

	accountInfo.ModifyUser = user.GetAttributeValue("modifyUser")
	modifyData := strings.Split(user.GetAttributeValue("modifyData"), ".")
	accountInfo.ModifyData = modifyData[0]

	passwordSet := strings.Split(user.GetAttributeValue("userPasswordSet"), ".")
	accountInfo.Password.UserPasswordSet = passwordSet[0]
	accountInfo.Password.PassValid = resultHow
	accountInfo.Password.PassSet = strconv.Itoa(days)

	var systemUser models.User
	database.DB.Model(models.User{}).Where("username = ?", username).FirstOrCreate(&systemUser, models.User{Username: username, AcceptSystemPolicies: false})
	accountInfo.AcceptSystemPolicies = systemUser.AcceptSystemPolicies

	// statusFa = user.GetAttributeValue("securityCode")
	//----- Permissions

	var permissions []models.SystemPermissions
	database.DB.Model(models.SystemPermissions{}).Joins("left join user_permissions on user_permissions.permission_id = system_permissions.id").Where("user_id = ?", username).Find(&permissions)

	return models.ActiveUser{
		Status:              200,
		AccountState:        accountState,
		Uid:                 username,
		PersonalInformation: personalInformation,
		AccountInfo:         accountInfo,
	}

}

func AuthLogin(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	username := strings.ToLower(data["username"])
	password := data["password"]

	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		system.Log <- models.Logs{
			Module:      "AGA",
			Action:      "LOGIN_ERROR",
			Login:       username,
			Uid:         username,
			Description: core.GetTextMessage("login_err_001"),
			App:         core.GetApp(c),
		}
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{
			"OK":      false,
			"message": core.GetTextMessage("login_err_001"),
		})
	}
	if username == "" || password == "" {
		var missingFields []string
		if username == "" {
			missingFields = append(missingFields, core.GetTextMessage("login_info_001"))
		}
		if password == "" {
			missingFields = append(missingFields, core.GetTextMessage("login_info_002"))
		}
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"OK":      false,
			"message": fmt.Sprintf("%s: %v.", core.GetTextMessage("global_info_001"), strings.Join(missingFields, ", ")),
		})
	}

	if strings.Contains(username, "@") {
		cutUser := strings.Split(username, "@")
		username = cutUser[0]
	}

	resultLogin := core.ExtractFunctionsPlugins("ldap", "Login", username, password)

	if !resultLogin.(bool) {
		c.Status(fiber.StatusUnauthorized)
		system.Log <- models.Logs{
			Module:      "AGA",
			Action:      "LOGIN_ERROR",
			Login:       username,
			Uid:         username,
			Description: core.GetTextMessage("login_err_002"),
			App:         core.GetApp(c),
		}
		return c.JSON(fiber.Map{
			"OK":      false,
			"message": core.GetTextMessage("login_err_002"),
		})
	}

	result := core.ExtractFunctionsPlugins("ldap", "Search", "(&(uid="+username+"))")
	bytes, _ := json.Marshal(&result)
	var resultSearch ldap.SearchResult
	_ = json.Unmarshal(bytes, &resultSearch)

	activeUser, err := authenticateUser(c, resultSearch, username, password)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"OK":      false,
			"message": core.GetTextMessage("login_err_005"),
		})
	}

	return c.JSON(fiber.Map{
		"OK":         true,
		"message":    core.GetTextMessage("login_ok_001"),
		"activeUser": activeUser,
	})
}

func AuthLogout(c *fiber.Ctx) error {
	claims, err := GetClaims(c)
	if err == nil {
		system.Log <- models.Logs{
			Module:      "AGA",
			Action:      "LOGOUT",
			Login:       claims["ID"].(string),
			Uid:         claims["ID"].(string),
			Description: "Cerró sesión.",
			App:         core.GetApp(c),
		}
		cookie := fiber.Cookie{
			Name:     system.UserKeyAGA,
			Value:    "",
			Expires:  time.Now().Add(-(time.Hour * 2)),
			HTTPOnly: true,
		}
		c.Cookie(&cookie)
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  200,
			"message": "Cerró sesión",
		})
	}
	c.Status(fiber.StatusForbidden)
	return c.JSON(fiber.Map{
		"status":  403,
		"message": "No esta logeado",
	})
}

func AuthActiveUser(c *fiber.Ctx) error {
	claims, err := GetClaims(c)
	if err != nil {
		c.Status(fiber.StatusForbidden)
		return c.JSON(fiber.Map{
			"message": "error",
		})
	}
	activeUser := c.Context().Value("activeUser")

	if system.AgaContext.Value("activeUser") == "null" {
		result := core.ExtractFunctionsPlugins("ldap", "Search", "(&(uid="+claims["ID"].(string)+"))")
		bytes, _ := json.Marshal(&result)
		var resultSearch ldap.SearchResult
		_ = json.Unmarshal(bytes, &resultSearch)

		if len(resultSearch.Entries) > 0 {
			system.AgaContext = context.WithValue(system.AgaContext, system.ClaimsKey, claims)
			system.AgaContext = context.WithValue(system.AgaContext, system.ActiveUserKey, GetUserInfo(resultSearch.Entries[0], claims["ID"].(string)))
			system.AgaContext = context.WithValue(system.AgaContext, system.LoggedInUserKey, claims["Name"].(string)+" ["+claims["ID"].(string)+"]")
		}

	}
	return c.JSON(fiber.Map{
		"activeUser": activeUser,
	})
}
