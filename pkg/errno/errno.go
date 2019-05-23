package errno

import (
	"github.com/labstack/echo/v4"
)

// Errno -
type Errno struct {
	Summary  string      `json:"-"`                // 错误概述
	HTTPCode int         `json:"-"`                // http 状态码
	Code     int         `json:"code"`             // 自定义的状态码
	Message  string      `json:"msg"`              // 错误信息 (客户端可能会用来提示用户)
	Errors   interface{} `json:"errors,omitempty"` // 具体错误信息 (用于定位错误)
	Template string      `json:"-"`                // 有该值则会渲染指定的模板页面
}

// Error -
func (e Errno) Error() string {
	return e.Message
}

// New -
func (e *Errno) New(msg string, es interface{}) *Errno {
	err := e.SetMessage(msg)
	return err.SetErrors(es)
}

// SetMessage -
func (e *Errno) SetMessage(msg string) *Errno {
	err := &Errno{
		Summary:  e.Summary,
		HTTPCode: e.HTTPCode,
		Code:     e.Code,
		Template: e.Template,
		Errors:   e.Errors,
		Message:  msg,
	}

	return err
}

// SetErrors -
func (e *Errno) SetErrors(es interface{}) *Errno {
	data := es
	switch typed := es.(type) {
	// echo error
	case *echo.HTTPError:
		// echo bind error
		if e.Code == ParamsErr.Code {
			data = map[string]map[string]interface{}{
				"bind": {
					"message":  typed.Message,
					"internal": typed.Internal,
				},
			}
		}
	default:
	}

	err := &Errno{
		Summary:  e.Summary,
		HTTPCode: e.HTTPCode,
		Code:     e.Code,
		Template: e.Template,
		Errors:   data,
		Message:  e.Message,
	}

	return err
}
