package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"pdb/model"
	"pdb/utils"
)

var GormDB *gorm.DB

func SetupDB() {
	var err error

	utils.EnsureDirectoryExists("data")

	GormDB, err = gorm.Open(sqlite.Open("data/database.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Connected to database")

	if err := GormDB.AutoMigrate(&model.Patient{}); err != nil {
		fmt.Println(err.Error())
	}

	if err := GormDB.AutoMigrate(&model.Documentation{}); err != nil {
		fmt.Println(err.Error())
	}
}
