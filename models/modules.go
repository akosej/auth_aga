package models

import (
	"gorm.io/gorm"
)

type Oup struct {
	Name string
	Num  int
}

type ModuleConfig struct {
	gorm.Model
	Name   string `gorm:"primaryKey"`
	Active bool
}

type SystemPermissions struct {
	gorm.Model
	Name   string `gorm:"uniqueIndex:idx_name_module,length:255"`
	Module string `gorm:"uniqueIndex:idx_name_module,length:255"`
}

type UserPermissions struct {
	gorm.Model
	UserId            string
	PermissionId      string
	SystemPermissions SystemPermissions `gorm:"foreignKey:PermissionId"`
}
