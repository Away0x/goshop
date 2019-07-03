package context

import (
	"echo_shop/config"
	"echo_shop/pkg/errno"
	"echo_shop/pkg/serializer"
	"net/http"
)

// CommonResponse 标准响应格式
type CommonResponse struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Debug interface{} `json:"debug,omitempty"`
}

// RenderHTML 渲染 html
func (a *AppContext) RenderHTML(tplName string, data ...serializer.Data) error {
	renderTypeExt := ".html"

	if len(data) != 0 {
		return a.Context.Render(http.StatusOK, tplName+renderTypeExt, data[0].ToJSON())
	}

	return a.Context.Render(http.StatusOK, tplName+renderTypeExt, map[string]interface{}{})
}

// RenderJSON resp json
func (a *AppContext) RenderJSON(code int, data CommonResponse) error {
	r := map[string]interface{}{}

	r["code"] = data.Code
	r["data"] = data.Data
	r["msg"] = data.Msg
	if data.Debug != nil {
		r["debug"] = data.Debug
	}

	return a.JSON(code, r)
}

// RenderOKJSON render ok json
func (a *AppContext) RenderOKJSON(data serializer.Data) error {
	return a.RenderJSON(http.StatusOK, CommonResponse{
		Code: 0,
		Msg:  "OK",
		Data: data,
	})
}

// RenderErrorJSON render error json
func (a *AppContext) RenderErrorJSON(err *errno.Errno) error {
	// 隐藏错误详情 (默认开启)
	if !config.Bool("APP.SHOW_ERROR_DETAIL") {
		err = err.HideErrorDetail()
	}

	return a.RenderJSON(err.HTTPCode, CommonResponse{
		Code:  err.Code,
		Msg:   err.Message,
		Debug: err.Detail,
	})
}
