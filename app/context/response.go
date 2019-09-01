package context

import (
	"echo_shop/pkg/errno"
	"echo_shop/pkg/serializer"
	"net/http"
)

type G = serializer.Data

// RenderHTML 渲染 html
func (a *AppContext) RenderHTML(tplName string, data ...serializer.Data) error {
	renderTypeExt := ".html"

	if len(data) != 0 {
		return a.Context.Render(http.StatusOK, tplName+renderTypeExt, data[0].ToJSON())
	}

	return a.Context.Render(http.StatusOK, tplName+renderTypeExt, map[string]interface{}{})
}

// RenderOKJSON render ok json
func (a *AppContext) RenderOKJSON(data serializer.Data) error {
	return a.JSON(http.StatusOK, map[string]interface{}(data))
}

// RenderErrorJSON render error json
func (a *AppContext) RenderErrorJSON(err *errno.Errno) error {
	resp := make(map[string]interface{})

	resp["code"] = err.Code
	resp["msg"] = err.Message

	if err.Detail != nil && err.Detail.Content != nil {
		errLen := len(err.Detail.Content)
		// detail.context 中只有一个 kv，提取出来
		if errLen == 1 {
			for _, v := range err.Detail.Content {
				resp["errors"] = v
			}
		} else {
			resp["errors"] = err.Detail.Content
		}
	}

	return a.JSON(err.HTTPCode, resp)
}
