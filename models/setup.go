// models/setup.go

package models

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		// DEBUG Mode: Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed to connect database")
	}

	err = database.AutoMigrate(&Book{})
	if err != nil {
		return
	}

	DB = database
}
