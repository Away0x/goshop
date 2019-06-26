package request

import (
	"echo_shop/app/helpers/captcha"

	"github.com/Away0x/validate"
)

// CaptchaValidator 验证码验证
func CaptchaValidator(captchaID, captchaVal string) validate.ValidatorFunc {
	return func() (msg string) {
		if ok := captcha.Verify(captchaID, captchaVal); ok {
			return ""
		}
		return "验证码错误"
	}
}
