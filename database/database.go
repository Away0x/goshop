package database

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitDB() {

}

func DBManager() *gorm.DB {
	return db
}
