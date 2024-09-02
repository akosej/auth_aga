package core

import "github.com/gofiber/fiber/v2"

func GetApp(c *fiber.Ctx) string {
	appOrigin := c.Get("Origin")
	return ternary(appOrigin != "", appOrigin, "AGA").(string)
}

func ternary(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}
