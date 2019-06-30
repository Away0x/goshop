package context

import (
	"net/http"
)

// RenderHTML 渲染 html
func (a *AppContext) RenderHTML(tplName string, data ...map[string]interface{}) error {
	renderTypeExt := ".html"

	if len(data) != 0 {
		return a.Context.Render(http.StatusOK, tplName+renderTypeExt, data[0])
	}

	return a.Context.Render(http.StatusOK, tplName+renderTypeExt, map[string]interface{}{})
}

// RenderJSON resp json
func (a *AppContext) RenderJSON(data map[string]interface{}) error {
	return a.JSON(http.StatusOK, data)
}
