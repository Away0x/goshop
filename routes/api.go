package routes

import (
	"echo_shop/app/context"
	"echo_shop/app/controllers/api/auth"
	"echo_shop/pkg/constants"

	"github.com/labstack/echo/v4"
)

const (
	apiVersion = "v1"
)

var (
	// APIPrefix -
	APIPrefix = constants.RestfulAPIPrefix + "/" + apiVersion
)

func registerAPI(e *echo.Echo) {
	ee := e.Group(APIPrefix)

	// ------------------------------------- Auth -------------------------------------
	authRouter := ee.Group("/auth")
	{
		// 登录
		context.RegisterHandler(authRouter.POST, "/login", auth.Login)
		// 注册
		context.RegisterHandler(authRouter.POST, "/register", auth.Register)
		// 刷新 token
		context.RegisterHandler(authRouter.PUT, "/refresh_token", auth.RefreshToken)
		// 登出
		context.RegisterHandler(authRouter.DELETE, "/logout", auth.Logout)
	}
}
