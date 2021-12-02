package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToDB(filepath string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(filepath), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}
