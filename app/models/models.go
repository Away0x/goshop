package models

import (
	"echo_shop/database"
	"echo_shop/pkg/serializer"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

const (
	// TrueTinyint true
	TrueTinyint = 1
	// FalseTinyint false
	FalseTinyint = 0
)

// BaseModel model 基类
type BaseModel struct {
	ID uint `sql:"primary_key; auto_increment; not null"`
	// MySQL的DATE/DATATIME类型可以对应Golang的time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	// 有 DeletedAt(类型需要是 *time.Time) 即支持 gorm 软删除
	DeletedAt *time.Time `sql:"index"`
}

// Serialize -
func (m *BaseModel) Serialize() serializer.Data {
	return serializer.Data{
		"id": m.ID,
	}
}

// IDString 获取字符串形式的 id
func (m *BaseModel) IDString() string {
	return strconv.Itoa(int(m.ID))
}

// Save : db save
func Save(v interface{}) error {
	return database.DBManager().Save(v).Error
}

// Create : db create
func Create(v interface{}) error {
	return database.DBManager().Create(v).Error
}

// Where : db where
func Where(query string, args ...interface{}) *gorm.DB {
	return database.DBManager().Where(query, args...)
}
