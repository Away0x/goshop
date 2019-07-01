package bootstrap

import (
	"echo_shop/app/models"
	"echo_shop/database"

	"github.com/jinzhu/gorm"
)

// SetupDB 初始化数据库
func SetupDB() (*gorm.DB, error) {
	database.InitDB()
	db := database.DBManager()

	db.AutoMigrate(
		&models.User{},
		&models.UserAddress{},
	)

	return db, nil
}
