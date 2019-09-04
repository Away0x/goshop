package context

import (
	"github.com/Away0x/validate"
	"github.com/labstack/echo/v4"
)

// 统一将各种错误都包装成 validate.Messages 类型
func wrapEchoHTTPError(err error) validate.Messages {
	switch typed := err.(type) {
	case *echo.HTTPError:
		return validate.Messages{
			"bind": {typed.Message.(string)},
		}
	case validate.Messages:
		return typed
	default:
		return validate.Messages{
			"error": {typed.Error()},
		}
	}
}

// ValidateStruct 验证 struct 字段
func (*Context) ValidateStruct(v validate.Validater) validate.Messages {
	err, ok := validate.Run(v)
	if ok {
		return nil
	}

	return err
}

// ValidateStructWithConfig 验证 struct 字段
func (*Context) ValidateStructWithConfig(v validate.Validater, config validate.Config) validate.Messages {
	err, ok := validate.RunWithConfig(v, config)
	if ok {
		return nil
	}

	return err
}

// BindAndValidate bind and validate
func (c *Context) BindAndValidate(v validate.Validater) validate.Messages {
	if err := c.Context.Bind(v); err != nil {
		return wrapEchoHTTPError(err)
	}

	if err := c.ValidateStruct(v); err != nil {
		return err
	}

	return nil
}

// BindAndValidateWithConfig bind and validate
func (c *Context) BindAndValidateWithConfig(v validate.Validater, config validate.Config) validate.Messages {
	if err := c.Context.Bind(v); err != nil {
		return wrapEchoHTTPError(err)
	}

	if err := c.ValidateStructWithConfig(v, config); err != nil {
		return err
	}

	return nil
}
