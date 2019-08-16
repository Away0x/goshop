package models

import (
	"echo_shop/database"
	"echo_shop/pkg/serializer"
	"echo_shop/pkg/utils"
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

// Delete : db delete
func Delete(v interface{}) error {
	return database.DBManager().Delete(v).Error
}

// DB : db
func DB() *gorm.DB {
	return database.DBManager()
}

// Save : db save
func Save(v interface{}) error {
	return database.DBManager().Save(v).Error
}

// Create : db create
func Create(v interface{}) error {
	return database.DBManager().Create(v).Error
}

// Update : db update
func Update(v interface{}) error {
	return database.DBManager().Model(v).Updates(v).Error
}

// Where : db where
func Where(query string, args ...interface{}) *gorm.DB {
	return database.DBManager().Where(query, args...)
}

// AssignAndCreate 批量赋值，并且创建 model
// force: 会过滤掉空值
func AssignAndCreate(force bool, model interface{}, other interface{}) (err error) {
	if err = utils.BatchAssign(force, model, other); err != nil {
		return err
	}

	if err = Create(model); err != nil {
		return err
	}

	return nil
}

// AssignAndUpdate 批量赋值，并且更新 model
// force: 会过滤掉空值
func AssignAndUpdate(force bool, model interface{}, other interface{}) (err error) {
	if err = utils.BatchAssign(force, model, other); err != nil {
		return err
	}

	if err = Save(model); err != nil {
		return err
	}

	return nil
}
