package login

import (
	"echo_shop/app/context"
	"echo_shop/app/models"
	"echo_shop/app/request"

	"github.com/Away0x/validate"
)

type loginForm struct {
	validate.Base
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

	user := new(models.User)
	if err := models.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.ErrorFlashWEM(err, "email", "该用户还没有注册过用户")
		return c.RedirectToLoginPage()
	}

	if err := user.Compare(req.Password); err != nil {
		c.ErrorFlashWEM(err, "email", "很抱歉，您的邮箱和密码不匹配")
		return c.RedirectToLoginPage()
	}

	c.Login(user)
	// 用户未激活
	if !user.IsActivated() {
		return c.RedirectToUserVerificationPage()
	}

	c.FlashSuccessMessage("欢迎回来！")
	return c.RedirectBack("root")
}
