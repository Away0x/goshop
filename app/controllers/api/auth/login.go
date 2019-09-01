package auth

import (
	"echo_shop/app/context"
	"echo_shop/app/helpers/jwt"
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

func Login(c *context.AppContext) error {
	req := new(loginForm)

	if err := c.BindAndValidate(req); err != nil {
		return err
	}

	user := new(models.User)
	if err := models.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return err
	}
	if err := user.Compare(req.Password); err != nil {
		return err
	}

	tokenInfo, err := jwt.Sign(user)
	if err != nil {
		return err
	}

	return c.RenderOKJSON(context.G{
		"info":  req,
		"token": tokenInfo,
	})
}
