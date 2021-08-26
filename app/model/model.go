package model

import (
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DBMigrate migrates the postgres database
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Tornado{})
	return db
}
