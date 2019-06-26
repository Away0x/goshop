package context

import (
	"echo_shop/config"
	"net/http"
)

const (
	// 登录页 route name
	loginPageRouteName = "login.show"
	// 首页 route name
	rootPageRouteName = "root"
)

// RedirectByName 重定向
// routeName: 路由名
func (a *AppContext) RedirectByName(routeName string) error {
	return a.Context.Redirect(http.StatusFound, config.Application.Reverse(routeName))
}

// RedirectToLoginPage 重定向到登录页面
func (a *AppContext) RedirectToLoginPage() error {
	return a.Context.Redirect(http.StatusFound, config.Application.Reverse(loginPageRouteName))
}

// RedirectToRegisterPage 重定向到注册页面
func (a *AppContext) RedirectToRegisterPage() error {
	return a.Context.Redirect(http.StatusFound, config.Application.Reverse("register.show"))
}

// RedirectToHomePage 重定向到首页
func (a *AppContext) RedirectToHomePage() error {
	return a.Context.Redirect(http.StatusFound, config.Application.Reverse(rootPageRouteName))
}

// RenderHTML 渲染 html
func (a *AppContext) RenderHTML(tplName string, data ...map[string]interface{}) error {
	renderTypeExt := ".html"

	if len(data) != 0 {
		return a.Context.Render(http.StatusOK, tplName+renderTypeExt, data[0])
	}

	return a.Context.Render(http.StatusOK, tplName+renderTypeExt, map[string]interface{}{})
}
