// Package routes define las rutas y controladores para la aplicación.
// Las rutas se agregan al objeto de aplicación de Fiber y se agrupan por funcionalidad.
// Cada ruta tiene un controlador definido en otro archivo.
// Este archivo contiene la definición de las rutas y su agrupación.
package routes

import (
	"github.com/agaUHO/aga/controllers"
	"github.com/agaUHO/aga/polities"

	"github.com/gofiber/fiber/v2"
)

// Routes agrega las rutas al objeto de aplicación de Fiber.
func Routes(app *fiber.App) {
	// [Middleware] global para verificar la autorización de la aplicación
	app.Use(polities.AppAuthorize)

	// Rutas para autenticación
	app.Post("/login", controllers.AuthLogin).Name("AuthLogin")
	app.Post("/logout", controllers.AuthLogout).Name("AuthLogout")

	// [Rutas]para recursos estáticos
	app.Use(polities.LoggedIn, polities.LoggedInfo)
	app.Post("/active", controllers.AuthActiveUser).Name("AuthActiveUser")
	app.Get("/activity", controllers.LogsActivityMy).Name("LogsActivityMy")
}
