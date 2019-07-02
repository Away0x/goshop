package context

import (
	"echo_shop/pkg/serializer"
	"net/http"
)

// RenderHTML 渲染 html
func (a *AppContext) RenderHTML(tplName string, data ...serializer.Data) error {
	renderTypeExt := ".html"

	if len(data) != 0 {
		return a.Context.Render(http.StatusOK, tplName+renderTypeExt, data[0].ToJSON())
	}

	return a.Context.Render(http.StatusOK, tplName+renderTypeExt, map[string]interface{}{})
}

// RenderJSON resp json
func (a *AppContext) RenderJSON(data serializer.Data) error {
	return a.JSON(http.StatusOK, data.ToJSON())
}
