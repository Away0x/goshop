package login

import (
	"echo_shop/app/context"
	"echo_shop/app/request"
	"github.com/Away0x/validate"
)

type loginForm struct {
	validate.BaseValidate
	Email    string
	Password string
}

func (*loginForm) IsStrict() bool {
	return false
}

func (f *loginForm) Plugins() validate.Plugins {
	return validate.Plugins{
		request.EmailPlugin(f.Email),
		request.PasswordPlugin(f.Password),
	}
}

// Login 登录
func Login(c *context.AppContext) error {
	req := new(loginForm)

	if err := c.BindAndValidate(req); err != nil {
		c.ErrorFlash(err)
		return c.RedirectByName("login.show")
	}

	c.FlashSuccessMessage("登录成功")
	return c.RedirectByName("login.show")
}
