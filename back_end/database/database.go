package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
)

func Open(source string) (e error) {
	database, e = gorm.Open(sqlite.Open(source))
	if e != nil {
		return
	}
	e = database.AutoMigrate(&UserModel{})
	if e != nil {
		return
	}
	e = database.AutoMigrate(&TeamModel{})
	if e != nil {
		return
	}
	e = database.AutoMigrate(&ItemModel{})
	if e != nil {
		return
	}
	return
}
