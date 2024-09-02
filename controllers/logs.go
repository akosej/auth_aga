package controllers

import (
	"fmt"
	"strings"

	"github.com/agaUHO/aga/database"
	"github.com/agaUHO/aga/models"
	"github.com/gofiber/fiber/v2"
)

func LogsActivityMy(c *fiber.Ctx) error {
	//-------------------
	claims, _ := GetClaims(c)
	//-------------------
	var activity = map[string]int{}
	var logs []models.Logs
	database.DB.Where("Login = ?", claims["ID"].(string)).Find(&logs)

	date := ""
	numActivity := 0
	for k, value := range logs {
		cutDate := strings.Split(value.CreatedAt.String(), " ")
		if cutDate[0] == date {
			numActivity += 1
		} else {
			if date == "" {
				date = cutDate[0]
				numActivity = 1
			} else {
				activity[date] = numActivity
				date = cutDate[0]
				numActivity = 1
			}

		}
		if len(logs) == k+1 {
			activity[date] = numActivity + 1

		}
	}
	return c.JSON(fiber.Map{
		"values": activity,
	})
}

func MyActivityChart(c *fiber.Ctx) error {
	// Obtiene los claims del token JWT almacenado en la cookie de sesión
	claims, _ := GetClaims(c)
	// Define un slice vacío de modelos de actividades
	var logCounts []models.Activities

	// Define una subconsulta que cuenta los registros de la tabla "logs" agrupados por fecha y ordenados por cantidad

	maxCountQuery := fmt.Sprintf(`
	SELECT COUNT(*)
	FROM logs
	WHERE login = '%s' AND (action = 'LOGIN_SUCCESS' OR action = 'CHANGE_PASSWORD')
	GROUP BY DATE_FORMAT(created_at, '%%Y-%%m-%%d')
	ORDER BY COUNT(*) DESC
	LIMIT 1;
	`, claims["ID"].(string))

	var maxCount int
	database.DB.Model(models.Logs{}).Raw(maxCountQuery).Scan(&maxCount)
	firstQuarter := float64(maxCount) * 0.25
	secondQuarter := float64(maxCount) * 0.5
	thirdQuarter := float64(maxCount) * 0.75
	// Ejecuta una consulta en la tabla "logs" y selecciona la cantidad de registros, la fecha formateada y un nivel calculado
	database.DB.Model(models.Logs{}).
		// Agrega una cláusula WHERE para seleccionar los registros del usuario actual donde la acción es "LOGIN_SUCCESS" o "CHANGE_PASSWORD"
		Where("login = ? AND (action = 'LOGIN_SUCCESS' OR action = 'CHANGE_PASSWORD')", claims["ID"].(string)).
		// Selecciona la cantidad de registros, la fecha formateada y un nivel calculado
		Select(`
			COUNT(*) AS count,
			DATE_FORMAT(created_at, '%Y-%m-%d') AS date,
			CASE
				WHEN COUNT(*) = 0 THEN 0
				WHEN COUNT(*) > 0 AND COUNT(*) < @firstQuarter THEN 1
				WHEN COUNT(*) >= @firstQuarter AND COUNT(*) < @secondQuarter THEN 2
				WHEN COUNT(*) >= @secondQuarter AND COUNT(*) < @thirdQuarter THEN 3
				ELSE 4
			END AS level
		`, map[string]interface{}{"firstQuarter": firstQuarter, "secondQuarter": secondQuarter, "thirdQuarter": thirdQuarter}).
		// Agrega una cláusula WHERE para excluir los registros con fecha nula
		Where("DATE_FORMAT(created_at, '%Y-%m-%d') IS NOT NULL").
		// Agrupa los registros por fecha formateada
		Group("DATE_FORMAT(created_at, '%Y-%m-%d')").
		// Ejecuta la consulta y almacena los resultados en la variable "logCounts"
		Find(&logCounts)

	// Devuelve una respuesta JSON que contiene los resultados de la consulta
	return c.JSON(logCounts)
}
