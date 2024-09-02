package database

import (
	"fmt"

	"github.com/agaUHO/aga/system"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open(system.MysqlCredentials), &gorm.Config{})
	if err != nil {
		fmt.Println("Could not connect to the database", err)
	}
	DB = db
	// _ = DB.AutoMigrate(&models.StaticPluginConfig{}, &models.PluginConfig{}, &models.ModuleConfig{}, &models.SystemPermissions{}, &models.UserPermissions{})
	// _ = DB.AutoMigrate(&models.Logs{}, models.AuthorizedApp{})
	// _ = DB.AutoMigrate(&models.User{})

}

func WithDB(fn func(db *gorm.DB) interface{}) (interface{}, error) {
	var db *gorm.DB
	var err error

	db, err = gorm.Open(mysql.Open(system.MysqlCredentials), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %s", err.Error())
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Printf("failed to get DB instance: %s\n", err.Error())
			return
		}
		sqlDB.Close()
	}()
	return fn(db), nil
}
