package request

import (
	"echo_shop/app/request/validators"

	"github.com/Away0x/validate"
)

// 登录
type LoginForm struct {
	validate.Base
	Email    string
	Password string
}

func (*LoginForm) IsStrict() bool {
	return false
}

func (f *LoginForm) Plugins() validate.Plugins {
	return validate.Plugins{
		validators.EmailPlugin(f.Email),
		validators.PasswordPlugin(f.Password),
	}
}

// 注册
type RegisterForm struct {
	validate.Base
	Name                 string
	Email                string
	Password             string
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation"`
	Captcha              string
	CaptchaID            string `json:"captcha_id" form:"captcha_id"`

	APIForm bool `json:"-" form:"-"` // 是 api 请求，不需检查验证码
}

func (*RegisterForm) IsStrict() bool {
	return false
}

func (r *RegisterForm) Plugins() validate.Plugins {
	if r.APIForm {
		return validate.Plugins{
			validators.UserNamePlugin(r.Name),
			validators.EmailPlugin(r.Email),
			validators.PasswordPlugin(r.Password),
		}
	}
	return validate.Plugins{
		validators.UserNamePlugin(r.Name),
		validators.EmailPlugin(r.Email),
		validators.RegisterPasswordPlugin(r.Password, r.PasswordConfirmation),
		validators.CaptchaPlugin(r.CaptchaID, r.Captcha),
	}
}
