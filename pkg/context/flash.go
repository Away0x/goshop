package context

import (
	"echo_shop/pkg/flash"

	"github.com/Away0x/validate"
)

const (
	// FLashTypeSuccess -
	FLashTypeSuccess = flash.FlashMessageSuccessType
	// FLashTypeInfo -
	FLashTypeInfo = flash.FlashMessageInfoType
	// FLashTypeWarning -
	FLashTypeWarning = flash.FlashMessageWarningType
	// FLashTypeDanger -
	FLashTypeDanger = flash.FlashMessageDangerType
)

// Flash falsh
func (c *Context) Flash(flashType string, msg string) {
	switch flashType {
	case FLashTypeSuccess:
		flash.NewSuccessMessage(c, msg)
	case FLashTypeInfo:
		flash.NewInfoMessage(c, msg)
	case FLashTypeWarning:
		flash.NewWarningMessage(c, msg)
	case FLashTypeDanger:
		flash.NewDangerMessage(c, msg)
	default:
		flash.NewMessage(c, flashType, msg)
	}
}

// FlashSuccessMessage success message flash
func (c *Context) FlashSuccessMessage(msg string) {
	c.Flash(FLashTypeSuccess, msg)
}

// FlashInfoMessage info message flash
func (c *Context) FlashInfoMessage(msg string) {
	c.Flash(FLashTypeInfo, msg)
}

// FlashWarningMessage warning message flash
func (c *Context) FlashWarningMessage(msg string) {
	c.Flash(FLashTypeWarning, msg)
}

// FlashDangerMessage danger message flash
func (c *Context) FlashDangerMessage(msg string) {
	c.Flash(FLashTypeDanger, msg)
}

// ErrorFlash 错误信息的 flash
// 相比于 flash package 更易用
// 当错误类型为 validate.Messages，表示表单验证错误，调用 NewErrors
// 其他错误类型使用 NewDangerMessage
func (c *Context) ErrorFlash(err error) {
	switch typed := err.(type) {
	case validate.Messages:
		flash.NewErrors(c, typed)
	default:
		c.FlashDangerMessage(typed.Error())
	}
}
