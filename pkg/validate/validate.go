package validate

import (
	"github.com/gookit/validate"
)

// BaseForm -
type BaseForm struct{}

// Translates 翻译字段名
func (BaseForm) Translates() map[string]string {
	return validate.MS{
		"Name":  "用户名",
		"Email": "邮箱",
	}
}

// Messages 翻译错误信息
func (BaseForm) Messages() map[string]string {
	return validate.MS{
		"minLength": "{field}不能小于 %d 个字符",
		"maxLength": "{field}不能大于 %d 个字符",
		"required":  "邮箱不能为空",
		"email":     "{field}格式有误",
	}
}
