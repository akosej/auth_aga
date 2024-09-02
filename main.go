// AGA
// Adaptar Gestionar Automatizar
package main

import (
	"context"
	"time"

	"github.com/agaUHO/aga/core"
	"github.com/agaUHO/aga/database"
	"github.com/agaUHO/aga/routes"
	"github.com/agaUHO/aga/system"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
)

// Initals func
func init() {
	// Create a base context
	system.AgaContext = context.Background()
	system.AgaContext = context.WithValue(system.AgaContext, "activeUser", "null")
	// Check system files and folders
	system.CheckFilesFolders()
	// Connect database
	database.Connect()
	// Initialize channel listening
	go core.InitializeChannelListening()
	// -- Remote Procedure Call (rpc PROTOCOL)
	go core.ServerRPC()
	// Watcher the plugins and modules folder
	go core.WatcherModulesPlugins()
	// Task scheduler
	// core.CreateCronJob()
}

func main() {
	app := fiber.New(fiber.Config{
		AppName:           "AGA v1.1",
		CaseSensitive:     false,
		EnablePrintRoutes: false,
		//GETOnly: true,
	})

	app.Use(favicon.New(favicon.Config{
		File: system.Path + "/aga.ico",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     system.AllowOrigins,
		AllowHeaders:     system.AllowHeaders,
		AllowCredentials: true,
	}))

	app.Static("/", system.Path+"/app", fiber.Static{
		Compress:      false,
		ByteRange:     true,
		Browse:        false,
		Index:         "index.html",
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	routes.Routes(app)

	//--------------------------------------------------------------------
	_ = app.Listen(":" + system.Port)
}
