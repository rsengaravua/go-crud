package models

import "gorm.io/gorm"

type book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"Author"`
	Description string `json:"Description"`
}
