package request

import (
	"github.com/Away0x/validate"
)

// EmailPlugin 项目中 email 字段请求参数的验证
func EmailPlugin(email string) validate.PluginFunc {
	return func() (string, []validate.ValidatorFunc, []string) {
		return "email", []validate.ValidatorFunc{
				validate.RequiredValidator(email),
				validate.EmailValidator(email),
				validate.MaxLengthValidator(email, 255),
			}, []string{
				"邮箱不能为空",
				"邮箱格式错误",
				"邮箱长度不能大于 255 个字符",
			}
	}
}

// PasswordPlugin 项目中 password 字段请求参数的验证
func PasswordPlugin(password string) validate.PluginFunc {
	return func() (string, []validate.ValidatorFunc, []string) {
		return "password", []validate.ValidatorFunc{
				validate.RequiredValidator(password),
			}, []string{
				"密码不能为空",
			}
	}
}
