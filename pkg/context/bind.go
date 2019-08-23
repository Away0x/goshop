package context

import (
	"github.com/Away0x/validate"
)

// ValidateStruct 验证 struct 字段
func (*Context) ValidateStruct(v validate.Validater) error {
	err, ok := validate.Run(v)
	if ok {
		return nil
	}

	return err
}

// ValidateStructWithConfig 验证 struct 字段
func (*Context) ValidateStructWithConfig(v validate.Validater, config validate.Config) error {
	err, ok := validate.RunWithConfig(v, config)
	if ok {
		return nil
	}

	return err
}

// BindAndValidate bind and validate
func (c *Context) BindAndValidate(v validate.Validater) error {
	if err := c.Context.Bind(v); err != nil {
		return err
	}

	if err := c.ValidateStruct(v); err != nil {
		return err
	}

	return nil
}

// BindAndValidateWithConfig bind and validate
func (c *Context) BindAndValidateWithConfig(v validate.Validater, config validate.Config) error {
	if err := c.Context.Bind(v); err != nil {
		return err
	}

	if err := c.ValidateStructWithConfig(v, config); err != nil {
		return err
	}

	return nil
}
