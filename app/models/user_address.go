package models

import (
	"echo_shop/pkg/serializer"
	"fmt"
	"time"
)

const (
	// UserAddressTableName table name
	UserAddressTableName = "user_addresses"
)

// UserAddress 收货地址 model
type UserAddress struct {
	BaseModel

	User   User
	UserID uint `sql:"not null; index"` // 该地址所属的用户

	Province     string     `sql:"not null"` // 省
	City         string     `sql:"not null"` // 市
	District     string     `sql:"not null"` // 区
	Address      string     `sql:"not null"` // 具体地址
	Zip          uint       `sql:"not null"` // 邮编
	ContactName  string     `sql:"not null"` // 联系人姓名
	ContactPhone string     `sql:"not null"` // 联系人电话
	LastUsedAt   *time.Time // 最后一次使用时间
}

// TableName 表名
func (UserAddress) TableName() string {
	return UserAddressTableName
}

// Serialize -
func (u *UserAddress) Serialize() serializer.Data {
	v := serializer.Data{
		"id":            u.ID,
		"province":      u.Province,
		"city":          u.City,
		"district":      u.District,
		"address":       u.Address,
		"zip":           u.Zip,
		"contact_name":  u.ContactName,
		"contact_phone": u.ContactPhone,
		"full_address":  u.GetFullAddressAttribute(),
	}

	return v
}

// GetFullAddressAttribute -
func (u *UserAddress) GetFullAddressAttribute() string {
	return fmt.Sprintf("%s%s%s%s",
		u.Province,
		u.City,
		u.District,
		u.Address,
	)
}
