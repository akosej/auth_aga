package core

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"

	"github.com/agaUHO/aga/database"
	"github.com/agaUHO/aga/models"
	"github.com/agaUHO/aga/system"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func LoadPluginAndCreatePermission(pluginPath string) (*plugin.Plugin, error) {
	pluginName := filepath.Base(pluginPath)
	loadedPlugin, err := plugin.Open(pluginPath)
	if err != nil {
		system.Log <- models.Logs{
			App:         "MAIN",
			Action:      "PLUGIN_LOAD_PERMISSION_ERROR",
			Login:       "aga",
			Uid:         "aga",
			Description: "Error al cargar los permisos de " + pluginName + " , " + err.Error(),
		}
		fmt.Println("Error loading plugin: ", err)
	}
	GetPermissions, err := loadedPlugin.Lookup("GetPermissions")
	if err != nil {
		fmt.Println("Error loading plugin: ", err)
	}
	permissions := GetPermissions.(func() []string)()
	for _, permission := range permissions {
		_, _ = database.WithDB(func(db *gorm.DB) interface{} {
			return db.Clauses(clause.OnConflict{UpdateAll: true, Columns: []clause.Column{{Name: "Name"}, {Name: "Module"}}}).Create(&models.SystemPermissions{Name: permission, Module: pluginName})
		})

	}
	return loadedPlugin, err
}

func ExtractFunctionsPlugins(pluginName, function string, args ...interface{}) interface{} {
	var staticPluginConfig models.StaticPluginConfig
	result, _ := database.WithDB(func(db *gorm.DB) interface{} {
		return db.Where("name = ?", pluginName).First(&staticPluginConfig)
	})

	if result.(*gorm.DB).RowsAffected == 0 {
		_, _ = database.WithDB(func(db *gorm.DB) interface{} {
			return db.Create(&models.StaticPluginConfig{Name: pluginName, Active: false})
		})
		return nil
	} else if !staticPluginConfig.Active {
		fmt.Println("Plugins " + pluginName + ".plugin not active")
		return nil
	}
	if _, err := os.Stat(system.Path + "/plugins/" + pluginName + ".plugin"); os.IsNotExist(err) {
		fmt.Println("Plugins " + pluginName + ".plugin not found")
		return nil
	} else {
		plug, err := LoadPluginAndCreatePermission(system.Path + "/plugins/" + pluginName + ".plugin")
		if err != nil {
			fmt.Println(err)
		}

		GetFunctions, err := plug.Lookup(function)
		if err != nil {
			fmt.Println(err)
		}

		return GetFunctions.(func(arg ...interface{}) interface{ any })(args...)
	}
}

func ExtractFunctionsPluginsWithPermissions(username, pluginName, function string, args ...interface{}) interface{} {
	var permissions []string
	database.DB.Model(models.SystemPermissions{}).Joins("left join user_permissions on user_permissions.permission_id = system_permissions.id").Where("user_id = ?", username).Select("CONCAT(name,'@',module) as fullPermissionName").Pluck("fullPermissionName", &permissions)
	var interfaceArgs []interface{}
	for _, v := range permissions {
		interfaceArgs = append(interfaceArgs, v)
	}
	newArgs := append(interfaceArgs, args...)
	return ExtractFunctionsPlugins(pluginName, function, newArgs...)
}

// Función TwoFactorPluginActive verifica si el plugin TwoFactor está activo o no
func TwoFactorPluginActive() bool {
	// Se declara una variable para almacenar la configuración del plugin
	var pluginConfig models.StaticPluginConfig
	// Se busca la configuración del plugin en la base de datos por su nombre
	_ = database.DB.Where("name = ?", "TwoFactor").First(&pluginConfig)
	// Si no se encuentra la configuración, se considera que el plugin está activo
	// Si se encuentra la configuración, se devuelve su estado de activación
	return pluginConfig.Active
}

// Función TwoFactorPluginActive verifica si el plugin TwoFactor está activo o no
func CredentialsPluginActive() bool {
	// Se declara una variable para almacenar la configuración del plugin
	var pluginConfig models.StaticPluginConfig
	// Se busca la configuración del plugin en la base de datos por su nombre
	_ = database.DB.Where("name = ?", "credentials").First(&pluginConfig)
	// Si no se encuentra la configuración, se considera que el plugin está activo
	// Si se encuentra la configuración, se devuelve su estado de activación
	return pluginConfig.Active
}
