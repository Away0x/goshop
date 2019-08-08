package context

import (
	"echo_shop/config"
	"errors"
)

// WrapErrorMessage 包装输出到客户端的错误信息
// 开发模式才会输出真实 error message 到客户端
func (a *AppContext) WrapErrorMessage(err error, msg string) string {
	if config.IsDev() {
		return msg + ": " + err.Error()
	}

	return msg
}

// WrapError 包装输出到客户端的错误信息
func (a *AppContext) WrapError(err error, msg string) error {
	if config.IsDev() {
		return errors.New(msg)
	}

	return errors.New(msg + ": " + err.Error())
}
