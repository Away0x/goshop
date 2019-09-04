package validators

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
				validate.MinLengthValidator(password, 6),
			}, []string{
				"密码不能为空",
				"密码长度不能小于 6 个字符",
			}
	}
}

// RegisterPasswordPlugin 注册用户时的 password 字段验证
func RegisterPasswordPlugin(password, confirm string) validate.PluginFunc {
	return validate.AssignPlugin(
		PasswordPlugin(password),
		[]validate.ValidatorFunc{
			validate.MinLengthValidator(password, 6),
			validate.EqualValidator(password, confirm),
		},
		[]string{
			"密码长度不能小于 6 个字符",
			"两次输入的密码不一致",
		},
	)
}

// UserNamePlugin 用户名字段请求参数的验证
func UserNamePlugin(name string) validate.PluginFunc {
	return func() (string, []validate.ValidatorFunc, []string) {
		return "name", []validate.ValidatorFunc{
				validate.RequiredValidator(name),
				validate.MaxLengthValidator(name, 50),
			}, []string{
				"名称不能为空",
				"名称长度不能大于 50 个字符",
			}
	}
}

// CaptchaPlugin 验证码验证
func CaptchaPlugin(captchaID, captchaVal string) validate.PluginFunc {
	return func() (string, []validate.ValidatorFunc, []string) {
		return "captcha", []validate.ValidatorFunc{
				validate.RequiredValidator(captchaVal),
				CaptchaValidator(captchaID, captchaVal),
			}, []string{
				"验证码不能为空",
				"验证码错误",
			}
	}
}
