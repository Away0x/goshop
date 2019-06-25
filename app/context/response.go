package context

import (
	"echo_shop/config"
	"net/http"
)

// RedirectByName 重定向
// routeName: 路由名
func (a *AppContext) RedirectByName(routeName string) error {
	return a.Context.Redirect(http.StatusFound, config.Application.Reverse(routeName))
}

// RenderHTML 渲染 html
func (a *AppContext) RenderHTML(tplName string, data ...map[string]interface{}) error {
	renderTypeExt := ".html"

	if len(data) != 0 {
		return a.Context.Render(http.StatusOK, tplName+renderTypeExt, data[0])
	}

	return a.Context.Render(http.StatusOK, tplName+renderTypeExt, map[string]interface{}{})
}
