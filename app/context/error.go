package context

import (
	"echo_shop/bootstrap/config"
	"errors"

	"github.com/Away0x/validate"
)

type M = validate.Messages

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

func (a *AppContext) WE(err error, msg string) error {
	return a.WrapError(err, msg)
}

func (a *AppContext) WEM(err error, msg string) string {
	return a.WrapErrorMessage(err, msg)
}

// EM 包装一条参数错误信息
func (a *AppContext) EM(err error, key string, msg string) M {
	data := M{}
	data[key] = []string{a.WEM(err, msg)}
	return data
}

// ErrorFlashWEM 包装一条参数错误信息 (flash error 输出)
func (a *AppContext) ErrorFlashWEM(err error, key string, msg string) {
	ms := a.EM(err, key, msg)
	a.ErrorFlash(ms)
}

// ErrorFlashWM 包装一条参数错误信息 (flash error 输出)
func (a *AppContext) ErrorFlashWM(key string, msg string) {
	ms := M{}
	ms[key] = []string{msg}
	a.ErrorFlash(ms)
}
