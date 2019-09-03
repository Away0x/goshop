package auth

import (
	"echo_shop/app/auth/jwt"
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
		return err
	}

	user := new(models.User)
	if err := models.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.EM(err, "email", "该邮箱还未注册")
	}
	if err := user.Compare(req.Password); err != nil {
		return c.EM(err, "password", "密码错误")
	}

	// 签发 token
	tokenInfo, err := jwt.Sign(user)
	if err != nil {
		return err
	}

	return c.RenderJSON(context.G{
		"user":  user.Serialize(),
		"token": tokenInfo.Token,
	})
}
