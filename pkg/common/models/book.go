package models

import "gorm.io/gorm"

type Book struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func MigrateDB(db *gorm.DB) error {
	// Auto Migrate will ONLY create tables, missing columns and missing indexes,
	// and WON'T change existing column's type or delete unused columns to protect your data.
	if err := db.AutoMigrate(&Book{}); err != nil {
		return err
	}
	return nil
}
