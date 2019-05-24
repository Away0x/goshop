package errno

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// Errno -
type Errno struct {
	Summary   string      `json:"-"`                // 错误概述
	HTTPCode  int         `json:"-"`                // http 状态码
	Code      int         `json:"code"`             // 自定义的状态码
	Message   string      `json:"msg"`              // 错误信息 (客户端可能会用来提示用户)
	Errors    interface{} `json:"errors,omitempty"` // 具体错误信息 (用于定位错误)
	RenderTpl string      `json:"-"`                // 有值则渲染页面 (默认是返回 json)
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
		Errors:   data,
		Message:  e.Message,
	}

	return err
}

// Render 渲染模板
func (e *Errno) Render(data map[string]interface{}) *Errno {
	renderData := map[string]interface{}{
		"title":    e.Summary,
		"msg":      e.Message,
		"code":     e.Code,
		"back_url": "/",
	}

	for k, v := range data {
		renderData[k] = v
	}

	// 目前 img 只支持 403、404、500、503 4个 httpcode
	imgCode := 500
	switch e.HTTPCode {
	case 404:
		imgCode = 404
	case 503:
		imgCode = 503
	case 403, 429:
		imgCode = 403
	default:
		imgCode = 500
	}
	renderData["img"] = "/public/svg/" + strconv.Itoa(imgCode) + ".svg"

	err := &Errno{
		Summary:   e.Summary,
		HTTPCode:  e.HTTPCode,
		Code:      e.Code,
		Errors:    renderData,
		Message:   e.Message,
		RenderTpl: "errors/error.html",
	}

	return err
}

// RenderNoContent 渲染模板
func (e *Errno) RenderNoContent() *Errno {
	return e.Render(map[string]interface{}{})
}
