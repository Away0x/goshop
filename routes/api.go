package routes

import (
	"echo_shop/app/context"
	"echo_shop/app/controllers/api/auth"
	"echo_shop/app/controllers/api/user_address"
	"echo_shop/pkg/constants"
	"echo_shop/routes/wrapper"

	mymiddleware "echo_shop/routes/middleware"

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
		context.RegisterHandler(authRouter.POST, "/login", auth.Login).Name = "api.auth.login"
		// 注册
		context.RegisterHandler(authRouter.POST, "/register", auth.Register).Name = "api.auth.register"
		// 刷新 token
		context.RegisterHandler(authRouter.PUT, "/refresh_token", wrapper.GetToken(auth.RefreshToken)).Name = "api.auth.refresh_token"
		// 登出
		context.RegisterHandler(authRouter.DELETE, "/logout", wrapper.GetToken(auth.Logout)).Name = "api.auth.logout"
		// 获取用户信息
		context.RegisterHandler(authRouter.GET, "/user_info", wrapper.GetTokenAndUser(auth.UserInfo)).Name = "api.auth.user_info"
	}

	// ------------------------------------- user adress -------------------------------------
	userAddressRouter := ee.Group("/user_adress", mymiddleware.TokenAuth)
	{
    // 用户收货地址列表
		context.RegisterHandler(userAddressRouter.GET, "",
      wrapper.GetTokenAndUser(useraddress.Index)).Name = "api.user_addresses.index"
    // 新建收货地址
		context.RegisterHandler(userAddressRouter.POST, "",
      wrapper.GetTokenAndUser(useraddress.Store)).Name = "api.user_addresses.store"
    // 编辑收货地址
		context.RegisterHandler(userAddressRouter.PUT, "/:id",
			wrapper.GetTokenAndUser(useraddress.Update)).Name = "api.user_addresses.update"
    // 删除收货地址
		context.RegisterHandler(userAddressRouter.DELETE, "/:id",
			wrapper.GetTokenAndUser(useraddress.Destroy)).Name = "api.user_addresses.destroy"
	}
}
