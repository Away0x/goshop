package models

import (
	"echo_shop/database"
	"echo_shop/pkg/serializer"
	"echo_shop/pkg/utils"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

// ModelError model error (会在统一错误处理时转换为 errno.DatabaseErr)
type ModelError struct {
	Message string
}

func (m *ModelError) Error() string {
	return m.Message
}

func newError(err error) *ModelError {
	if err == nil {
		return nil
	}
	return &ModelError{
		Message: err.Error(),
	}
}

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

// DB : db
func DB() *gorm.DB {
	return database.DBManager()
}

// TinyBool -> bool
func TinyBool(i uint) bool {
	if i == TrueTinyint {
		return true
	}

	return false
}

// IDString 获取字符串形式的 id
func (m *BaseModel) IDString() string {
	return strconv.Itoa(int(m.ID))
}

// Delete : db delete
func Delete(v interface{}) *ModelError {
	e := database.DBManager().Delete(v).Error
	return newError(e)
}

// Save : db save
func Save(v interface{}) *ModelError {
	e := database.DBManager().Save(v).Error
	return newError(e)
}

// Create : db create
func Create(v interface{}) *ModelError {
	e := database.DBManager().Create(v).Error
	return newError(e)
}

// Update : db update
func Update(v interface{}) *ModelError {
	e := database.DBManager().Model(v).Updates(v).Error
	return newError(e)
}

// AssignAndCreate 批量赋值，并且创建 model
// force: 会过滤掉空值
func AssignAndCreate(force bool, model interface{}, other interface{}) (err *ModelError) {
	var e error
	if e = utils.BatchAssign(force, model, other); err != nil {
		return newError(e)
	}

	if e = Create(model); err != nil {
		return newError(e)
	}

	return nil
}

// AssignAndUpdate 批量赋值，并且更新或创建 model
// force: 会过滤掉空值
func AssignAndUpdate(force bool, model interface{}, other interface{}) (err *ModelError) {
	var e error
	if e = utils.BatchAssign(force, model, other); err != nil {
		return newError(e)
	}

	if e = Save(model); err != nil {
		return newError(e)
	}

	return nil
}

// WhereFirst where+first
func WhereFirst(v interface{}, query string, args ...interface{}) *ModelError {
	var e error
	if e = database.DBManager().Where(query, args...).First(v).Error; e != nil {
		return newError(e)
	}

	return nil
}

// WhereFind where+find
func WhereFind(v interface{}, query string, args ...interface{}) *ModelError {
	var e error
	if e = database.DBManager().Where(query, args...).Find(v).Error; e != nil {
		return newError(e)
	}

	return nil
}
