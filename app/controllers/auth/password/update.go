package password

import (
	"echo_shop/app/context"
	"echo_shop/app/request"

	"github.com/Away0x/validate"
)

type passwordUpdateForm struct {
	validate.BaseValidate
	Email                string
	Token                string
	Password             string
	PasswordConfirmation string
}

func (p *passwordUpdateForm) Plugins() validate.Plugins {
	return validate.Plugins{
		request.RegisterPasswordPlugin(p.Password, p.PasswordConfirmation),
	}
}

func (p *passwordUpdateForm) Validators() validate.Validators {
	return validate.Validators{
		"token": {
			validate.RequiredValidator(p.Token),
		},
	}
}

func (p *passwordUpdateForm) Messages() validate.Messages {
	return validate.Messages{
		"token": {
			"token 不能为空",
			"该 token 不存在",
		},
	}
}

// Update 重置密码
func Update(c *context.AppContext) error {
	return c.String(200, "password.update")
}
