package models

import "time"

const (
	// PasswordResetTableName : password_reset table name
	PasswordResetTableName = "password_resets"
)

// PasswordReset 密码重置 model
type PasswordReset struct {
	Email     string `sql:"not null; index"`
	Token     string `sql:"not null; index"`
	CreatedAt time.Time
}

// TableName 表名
func (PasswordReset) TableName() string {
	return PasswordResetTableName
}
