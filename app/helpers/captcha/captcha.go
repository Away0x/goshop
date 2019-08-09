package captcha

import (
	"echo_shop/bootstrap/app"
	captchaUtil "echo_shop/pkg/captcha"
)

// CreateCaptcha 生成验证码
func CreateCaptcha() captchaUtil.Captcha {
	return captchaUtil.New(func(id string) string {
		return app.Application.Reverse("captcha", id) // 生成验证码 url
	})
}

// Verify 验证
func Verify(captchaID, captchaVal string) bool {
	return captchaUtil.Verify(captchaID, captchaVal)
}
