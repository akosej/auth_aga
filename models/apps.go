package models

import "gorm.io/gorm"

type AuthorizedApp struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex:idx_name,length:255" json:"name"`
	Description string `json:"description"`
	Secret      string `gorm:"uniqueIndex:idx_secret,length:255" json:"secret"`
	Origin      string `gorm:"uniqueIndex:idx_origin_ipadress,length:255" json:"origin"`
	Ipadress    string `gorm:"uniqueIndex:idx_origin_ipadress,length:255" json:"ipadress"`
	Mail        string `json:"mail"`
	Active      bool   `json:"active"`
}
