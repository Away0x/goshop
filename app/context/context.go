package context

import (
	"echo_shop/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	// AppContext : 项目自定义的 context，包装了 echo.Context，提供更多方法
	AppContext struct {
		echo.Context
	}
	// AppHandlerFunc : 项目 handler type
	AppHandlerFunc = func(c *AppContext) error
)

// RouteRedirect 重定向
// routeName: 路由名
func (a *AppContext) RouteRedirect(routeName string) error {
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
