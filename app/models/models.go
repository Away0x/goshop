package models

import (
	"echo_shop/database"
	"echo_shop/pkg/viewmodel"
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

// Serialize viewmodel
func (m *BaseModel) Serialize() viewmodel.ViewModelSerialize {
	return viewmodel.ViewModelSerialize{
		"id": m.ID,
	}
}

// IDString 获取字符串形式的 id
func (m *BaseModel) IDString() string {
	return strconv.Itoa(int(m.ID))
}

// NewRecord model 是否已创建
func (m *BaseModel) NewRecord() bool {
	return m.ID <= 0
}

// Destroy 删除 model
func (m *BaseModel) Destroy() error {
	err := database.DBManager().Delete(&m).Error
	return err
}

// Updates 更新 model
func (m *BaseModel) Updates(d map[string]interface{}) error {
	err := database.DBManager().Model(&m).Updates(d).Error
	return err
}

// IsDeleted model 是否已被删除了
func (m *BaseModel) IsDeleted() bool {
	return m.DeletedAt != nil
}

// Where : db where
func Where(query string, args ...interface{}) *gorm.DB {
	return database.DBManager().Where(query, args...)
}

// Save : db save
func Save(v interface{}) error {
	return database.DBManager().Save(v).Error
}
