package models

import "gorm.io/gorm"

type Logs struct {
	gorm.Model
	Module      string `json:"module"`
	Action      string `json:"action"`
	Login       string `json:"login"`
	Uid         string `json:"uid"`
	Description string `json:"description"`
	App         string `json:"app"`
}

type Activities struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
	Level int    `json:"level"`
}
