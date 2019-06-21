package controllers

import (
	"echo_shop/config"
	"echo_shop/pkg/captcha"
)

// CreateCaptcha 生成验证码
func CreateCaptcha() captcha.Captcha {
	return captcha.New(func(id string) string {
		return config.Application.Reverse("captcha", id) // 生成验证码 url
	})
}
