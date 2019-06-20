package echoinit

import (
	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

// CustomValidator 自定义验证器
type CustomValidator struct{}

// Validate -
func (cv *CustomValidator) Validate(i interface{}) error {
	v := validate.Struct(i)
	if v.Validate() {
		return nil
	}

	return v.Errors
}

// SetupValidate -
func SetupValidate(e *echo.Echo) {
	e.Validator = &CustomValidator{}
}
