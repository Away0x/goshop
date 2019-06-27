package context

import (
	"echo_shop/config"
	"echo_shop/pkg/constants"
	"echo_shop/pkg/session"
	"net/http"
)

const (
	// 登录页 route name
	loginPageRouteName = "login.show"
	// 首页 route name
	rootPageRouteName = "root"
)

// RedirectTo 重定向
func (a *AppContext) RedirectTo(path string) error {
	return a.Context.Redirect(http.StatusFound, path)
}

// RedirectBack 返回之前的页面
func (a *AppContext) RedirectBack(defaultRoute ...string) error {
	referrer := a.Request().Header.Get("referer")
	if referrer != "" {
		return a.RedirectTo(referrer)
	}

	previousURL := session.Get(a, constants.PreviousURL)
	if previousURL != "" {
		return a.RedirectTo(previousURL)
	}

	if len(defaultRoute) != 0 || defaultRoute[0] != "" {
		return a.RedirectByName(defaultRoute[0])
	}

	return a.RedirectToHomePage()
}

// RedirectByName 重定向到指定路由 (routeName: 路由名)
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

// RedirectToUserVerificationPage 重定向到用户激活页面
func (a *AppContext) RedirectToUserVerificationPage() error {
	return a.Context.Redirect(http.StatusFound, config.Application.Reverse("verification.show_link_form"))
}
