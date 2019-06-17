package routes

import (
	"echo_shop/app/controllers/page"
	"echo_shop/pkg/constants"

	mymiddleware "echo_shop/routes/middleware"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"echo_shop/app/controllers/auth/login"
	"echo_shop/app/controllers/auth/password"
	"echo_shop/app/controllers/auth/register"
	"echo_shop/app/controllers/auth/verification"
	"echo_shop/pkg/captcha"
)

func registerWeb(e *echo.Echo, middlewares ...echo.MiddlewareFunc) {
	ee := e.Group("")

	// session
	ee.Use(session.MiddlewareWithConfig(session.Config{
		Store: sessions.NewCookieStore([]byte("secret")),
	}))

	// csrf
	ee.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		Skipper:        middleware.DefaultSkipper,
		TokenLength:    32,
		TokenLookup:    "header:" + constants.CsrfValueName,
		ContextKey:     constants.CsrfContexntName,
		CookieName:     constants.CsrfCookieName,
		CookiePath:     "/",
		CookieHTTPOnly: true,
		CookieMaxAge:   86400,
	}))

	ee.GET("/", page.Root).Name = "root"
	ee.GET("captcha/:id", captcha.Handler).Name = "captcha" // 验证码

	// ------------------------------------- Auth -------------------------------------
	// +++++++++++++++ 用户身份验证相关的路由 +++++++++++++++
	// 展示登录页面
	ee.GET("login", login.Show, mymiddleware.Guest).Name = "login.show"
	// 登录
	ee.POST("login", login.Login, mymiddleware.Guest).Name = "login"
	// 登出
	ee.DELETE("logout", login.Logout).Name = "logout"

	// +++++++++++++++ 用户注册相关路由 +++++++++++++++
	// 展示注册页面
	ee.GET("register", register.Show, mymiddleware.Guest).Name = "register.show"
	// 注册
	ee.POST("register", register.Register, mymiddleware.Guest).Name = "register"

	// +++++++++++++++ 密码重置相关路由 +++++++++++++++
	pwdRouter := ee.Group("/password", mymiddleware.Guest)
	{
		// 展示发送重置密码链接 email 的页面
		pwdRouter.GET("/reset", password.ShowLinkForm).Name = "password.show_link_form"
		// 发送重置密码链接的 email
		pwdRouter.POST("/email", password.Email).Name = "password.email"
		// 展示重置密码的页面
		pwdRouter.GET("/reset/:token", password.ShowResetForm).Name = "password.show_reset_form"
		// 重置密码
		pwdRouter.POST("/reset", password.Update).Name = "password.update"
	}

	// +++++++++++++++ Email 认证相关路由 +++++++++++++++
	verificationRouter := ee.Group("/verification", mymiddleware.Auth)
	{
		// 展示发送激活用户链接邮件的页面
		verificationRouter.GET("/verify", verification.ShowLinkForm).Name = "verification.show_link_form"
		// 激活用户
		verificationRouter.GET("/verify", verification.Verify).Name = "verification.verify"
		// 重新发送激活用户链接
		verificationRouter.GET("/verify", verification.Resend).Name = "verification.resend"
	}
}
