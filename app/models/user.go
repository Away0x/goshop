package models

import (
	"crypto/md5"
	"echo_shop/database"
	"echo_shop/pkg/serializer"
	"echo_shop/pkg/utils"
	"encoding/hex"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

const (
	// UserTableName : user table name
	UserTableName = "users"
)

// User 用户 Model
type User struct {
	BaseModel
	Name     string `sql:"not null; index"`
	Email    string `sql:"unique; index"`
	Password string `sql:"not null"`
	Avatar   string `sql:"not null"`

	// 用户激活
	ActivationToken string
	Activated       uint       `gorm:"type:tinyint(1);default:0"`
	EmailVerifiedAt *time.Time // 激活时间

	RememberToken string `gorm:"size:100"` // 用于实现记住我功能，存入 cookie 中，下次带上时，即可直接登录

	UserAddress []UserAddress
}

// TableName 表名
func (User) TableName() string {
	return UserTableName
}

// Serialize -
func (u *User) Serialize() serializer.Data {
	return serializer.Data{
		"id":     u.ID,
		"name":   u.Name,
		"email":  u.Email,
		"avatar": u.Avatar,
	}
}

// BeforeCreate - hook
func (u *User) BeforeCreate(scope *gorm.Scope) (err error) {
	if u.Password != "" {
		if !passwordEncrypted(u.Password) {
			pwd, err := u.Encrypt()
			if err != nil {
				return errors.New("User Model 创建失败: passwordEncrypted")
			}
			scope.SetColumn("password", pwd)
		}
	}

	// 生成用户 remember_token
	if u.RememberToken == "" {
		scope.SetColumn("remember_token", string(utils.RandomCreateBytes(10)))
	}

	// 生成用户激活 token
	if u.ActivationToken == "" && !u.IsActivated() {
		scope.SetColumn("activation_token", string(utils.RandomCreateBytes(30)))
	}

	// 生成用户头像
	if u.Avatar == "" {
		hash := md5.Sum([]byte(u.Email))
		avatar := "http://www.gravatar.com/avatar/" + hex.EncodeToString(hash[:])
		scope.SetColumn("avatar", avatar)
	}

	return err
}

// BeforeUpdate - hook
func (u *User) BeforeUpdate(scope *gorm.Scope) (err error) {
	if !passwordEncrypted(u.Password) {
		pwd, err := u.Encrypt()
		if err != nil {
			return errors.New("User Model 更新失败")
		}
		scope.SetColumn("password", pwd)
	}

	return
}

// 判断密码是否加密过了
func passwordEncrypted(pwd string) (status bool) {
	return len(pwd) == 60 // 长度等于 60 说明加密过了
}

// Encrypt 对密码进行加密
func (u *User) Encrypt() (pwd string, err error) {
	pwd, err = utils.Encrypt(u.Password)
	u.Password = pwd
	return
}

// Compare 验证用户密码
func (u *User) Compare(pwd string) (err error) {
	err = utils.Compare(u.Password, pwd)
	return
}

// GetUser 根据 id 获取 user
func GetUser(id uint) (user *User, err error) {
	user = new(User)
	err = database.DBManager().First(&user, id).Error
	return
}

// IsActivated 是否已激活
func (u *User) IsActivated() bool {
	return u.Activated == TrueTinyint
}

// Addresses 用户收货地址
func (u *User) Addresses() (addresses []*UserAddress) {
	addresses = make([]*UserAddress, 0)
	database.DBManager().Model(&u).Order("updated_at desc").Related(&addresses)
	return
}
