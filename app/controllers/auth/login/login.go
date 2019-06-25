package login

import (
	"echo_shop/app/context"
	"echo_shop/pkg/validate"
)

type loginForm struct {
	validate.BaseValidate
	Email    string
	Password string
}

func (f *loginForm) Validators() validate.Validators {
	return validate.Validators{
		"email": {
			validate.RequiredValidator(f.Email),
			validate.EmailValidator(f.Email),
			validate.MaxLengthValidator(f.Email, 255),
		},
		"password": {
			validate.RequiredValidator(f.Password),
		},
	}
}

func (f *loginForm) Messages() validate.Messages {
	return validate.Messages{
		"email": {
			"邮箱不能为空",
			"邮箱格式错误",
			"邮箱长度不能大于 255 个字符",
		},
		"password": {
			"密码不能为空",
		},
	}
}

// Login 登录
func Login(c *context.AppContext) error {
	req := new(loginForm)

	if err := c.BindAndValidate(req); err != nil {
		c.ErrorFlash(err)
		return c.RouteRedirect("login.show")
	}

	c.FlashSuccessMessage("登录成功")
	return c.RouteRedirect("login.show")
}
