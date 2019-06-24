package login

import (
	"echo_shop/app/controllers"
	"echo_shop/pkg/flash"
	"echo_shop/pkg/validate"

	"github.com/labstack/echo/v4"
)

type loginForm struct {
	validate.BaseForm
	Email    string `json:"email" form:"email" validate:"required|maxLen:255|email"`
	Password string `json:"password" form:"password" validate:"required"`
}

// Login 登录
func Login(c echo.Context) error {
	req := new(loginForm)

	if err := c.Bind(req); err != nil {
		flash.NewDangerMessage(c, err.Error())
		return controllers.Redirect(c, "login.show")
	}

	if err := c.Validate(req); err != nil {
		flash.NewErrors(c, err)
		return controllers.Redirect(c, "login.show")
	}

	flash.NewSuccessMessage(c, "登录成功")
	return controllers.Redirect(c, "login.show")
}
