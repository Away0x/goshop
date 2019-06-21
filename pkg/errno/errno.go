package errno

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// Errno 项目统一错误处理
type Errno struct {
	Summary  string        `json:"-"`                // 错误概述
	HTTPCode int           `json:"-"`                // http 状态码
	Code     int           `json:"code"`             // 自定义的状态码
	Message  string        `json:"msg"`              // 错误信息 (客户端可能会用来提示用户)
	Detail   *RenderDetail `json:"detail,omitempty"` // 错误详情和错误的渲染方式
}

// Error error interface
func (e Errno) Error() string {
	return e.Message
}

// WrapEchoHTTPError 包装 echo HTTPError 类型为 Errno 类型
func WrapEchoHTTPError(err *echo.HTTPError) *Errno {
	newErr := &Errno{
		Summary:  UnknowErr.Summary,
		HTTPCode: err.Code,
		Code:     UnknowErr.Code,
		Message:  UnknowErr.Message,
	}

	return newErr.SetErrorContent(err)
}

// SetSummary 设置 Errno 的错误概述
func (e *Errno) SetSummary(summary string) *Errno {
	return &Errno{
		Summary:  summary,
		HTTPCode: e.HTTPCode,
		Code:     e.Code,
		Message:  e.Message,
		Detail:   e.Detail,
	}
}

// SetMessage 设置 Errno 的错误信息
func (e *Errno) SetMessage(newMsg string) *Errno {
	return &Errno{
		Summary:  e.Summary,
		HTTPCode: e.HTTPCode,
		Code:     e.Code,
		Message:  newMsg,
		Detail:   e.Detail,
	}
}

// SetHTTPCode 设置 HTTPCode
func (e *Errno) SetHTTPCode(code int) *Errno {
	return &Errno{
		Summary:  e.Summary,
		HTTPCode: code,
		Code:     e.Code,
		Message:  e.Message,
		Detail:   e.Detail,
	}
}

// SetErrorContent 设置 Errno 的具体错误内容
func (e *Errno) SetErrorContent(content interface{}) *Errno {
	var errContent map[string]interface{}

	switch typed := content.(type) {
	case map[string]interface{}:
		errContent = typed
	case *echo.HTTPError: // echo error
		errContent = map[string]interface{}{
			"message":  typed.Message,
			"internal": typed.Internal,
		}
	case error:
		errContent = map[string]interface{}{
			"error": typed.Error(),
		}
	default:
		errContent = map[string]interface{}{
			"error": typed,
		}
	}

	var detail *RenderDetail
	if e.Detail != nil {
		if e.Detail.Type != "" {
			detail = NewRenderDetail(e.Detail.Type, errContent)
		}
	} else {
		detail = DefaultRenderDetail(errContent)
	}

	return &Errno{
		Summary:  e.Summary,
		HTTPCode: e.HTTPCode,
		Code:     e.Code,
		Message:  e.Message,
		Detail:   detail,
	}
}

// HideErrorDetail 隐藏错误详情
func (e *Errno) HideErrorDetail() *Errno {
	return &Errno{
		Summary:  e.Summary,
		HTTPCode: e.HTTPCode,
		Code:     e.Code,
		Message:  e.Message,
	}
}

// Wrap 基于某 Errno，设置一些自定义信息
func (e *Errno) Wrap(msg string, errContent map[string]interface{}) *Errno {
	err := e.SetMessage(msg)
	return err.SetErrorContent(errContent)
}

// ------------------------------------------------ render type function ------------------------------------------------

// JSON render json
func (e *Errno) JSON(data ...map[string]interface{}) *Errno {
	detail := WrapRenderContent(e.Detail, RenderDetailTypeJSON, data...)

	return &Errno{
		Summary:  e.Summary,
		HTTPCode: e.HTTPCode,
		Code:     e.Code,
		Message:  e.Message,
		Detail:   detail,
	}
}

// HTML render html
func (e *Errno) HTML(data ...map[string]interface{}) *Errno {
	// 填充到模板中的数据
	renderData := map[string]interface{}{
		"title":    e.Summary,
		"msg":      e.Message,
		"code":     e.Code,
		"back_url": "/",
	}

	for _, v1 := range data {
		for k2, v2 := range v1 {
			renderData[k2] = v2
		}
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

	detail := WrapRenderContent(e.Detail, RenderDetailTypeHTML, renderData)

	return &Errno{
		Summary:  e.Summary,
		HTTPCode: e.HTTPCode,
		Code:     e.Code,
		Message:  e.Message,
		Detail:   detail,
	}
}
