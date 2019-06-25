package context

import (
	"echo_shop/pkg/validate"
)

// ValidateStruct 验证 struct 字段
func (a *AppContext) ValidateStruct(v validate.Validater) error {
	err, ok := validate.Run(v)
	if ok {
		return nil
	}

	return err
}

// BindAndValidate bind and validate
func (a *AppContext) BindAndValidate(v validate.Validater) error {
	if err := a.Context.Bind(v); err != nil {
		return err
	}

	if err := a.ValidateStruct(v); err != nil {
		return err
	}

	return nil
}
