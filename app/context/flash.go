package context

import (
	"echo_shop/pkg/flash"
	"echo_shop/pkg/validate"
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
func (a *AppContext) Flash(flashType string, msg string) {
	switch flashType {
	case FLashTypeSuccess:
		flash.NewSuccessMessage(a, msg)
	case FLashTypeInfo:
		flash.NewInfoMessage(a, msg)
	case FLashTypeWarning:
		flash.NewWarningMessage(a, msg)
	case FLashTypeDanger:
		flash.NewDangerMessage(a, msg)
	default:
		flash.NewMessage(a, flashType, msg)
	}
}

// FlashSuccessMessage success message flash
func (a *AppContext) FlashSuccessMessage(msg string) {
	a.Flash(FLashTypeSuccess, msg)
}

// FlashInfoMessage info message flash
func (a *AppContext) FlashInfoMessage(msg string) {
	a.Flash(FLashTypeInfo, msg)
}

// FlashWarningMessage warning message flash
func (a *AppContext) FlashWarningMessage(msg string) {
	a.Flash(FLashTypeWarning, msg)
}

// FlashDangerMessage danger message flash
func (a *AppContext) FlashDangerMessage(msg string) {
	a.Flash(FLashTypeDanger, msg)
}

// ErrorFlash 错误信息的 flash
// 相比于 flash package 更易用
// 当错误类型为 validate.Messages，表示表单验证错误，调用 NewErrors
// 其他错误类型使用 NewDangerMessage
func (a *AppContext) ErrorFlash(err error) {
	switch typed := err.(type) {
	case validate.Messages:
		flash.NewErrors(a, typed)
	default:
		a.FlashDangerMessage(typed.Error())
	}
}
