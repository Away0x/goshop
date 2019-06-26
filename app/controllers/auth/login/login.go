package login

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
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
		return c.RedirectToLoginPage()
	}

	user, err := models.GetUserByEmail(req.Email)
	if err != nil {
		c.ErrorFlash(validate.Messages{
			"email": {"该用户还没有注册过用户: " + err.Error()},
		})
		return c.RedirectToLoginPage()
	}

	if err := user.Compare(req.Password); err != nil {
		c.ErrorFlash(validate.Messages{
			"email": {"很抱歉，您的邮箱和密码不匹配: " + err.Error()},
		})
		return c.RedirectToLoginPage()
	}

	c.Login(user)
	c.FlashSuccessMessage("登录成功")
	return c.RedirectByName("root")
}
