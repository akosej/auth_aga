package models

import "gorm.io/gorm"

type PluginConfig struct {
	gorm.Model
	Name   string `gorm:"primaryKey"`
	Active bool
}

type StaticPluginConfig struct {
	gorm.Model
	Name   string `gorm:"primaryKey"`
	Active bool
}
