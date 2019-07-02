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
	UserID uint `sql:"not null; index"`

	Province     string `sql:"not null"`
	City         string `sql:"not null"`
	District     string `sql:"not null"`
	Address      string `sql:"not null"`
	Zip          uint   `sql:"not null"`
	ContactName  string `sql:"not null"`
	ContactPhone string `sql:"not null"`
	LastUsedAt   *time.Time
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
