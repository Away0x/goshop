package register

import (
	"echo_shop/app/context"
	"echo_shop/app/helpers/captcha"
)

// Show 展示注册页面
func Show(c *context.AppContext) error {
	return c.RenderHTML("auth/register", map[string]interface{}{
		"captcha": captcha.CreateCaptcha(),
	})
}
