package models

import "time"

const (
	// UserAddressTableName table name
	UserAddressTableName = "user_addresses"
)

// UserAddress 收货地址 model
type UserAddress struct {
	BaseModel

	User   User `gorm:"foreignkey:UserRefer"`
	UserID uint `sql:"not null; index"`

	Province     string
	City         string
	District     string
	Address      string
	Zip          uint
	ContactName  string
	ContactPhone string
	LastUsedAt   *time.Time
}

// TableName 表名
func (UserAddress) TableName() string {
	return UserAddressTableName
}
