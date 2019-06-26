// Package flash : errors 一般用来回显参数验证的错误
package flash

import (
	"github.com/Away0x/validate"

	"github.com/labstack/echo/v4"
)

const (
	flashErrorsKeyName = "_flash_errors"
)

type errorsFlash struct {
	EchoContext echo.Context
	Data        flashDataValue
}

// NewErrorsFlash -
func NewErrorsFlash(c echo.Context) *errorsFlash {
	return &errorsFlash{
		EchoContext: c,
		Data:        make(flashDataValue),
	}
}

func (e *errorsFlash) Save(err error) {
	data, ok := err.(validate.Messages)
	if !ok {
		data = validate.Messages{"unknown": {err.Error()}}
	}

	NewFlashData(flashErrorsKeyName, e.EchoContext.Request()).
		Set(flashDataValue(data)).
		Save(e.EchoContext.Response())
}

func (e *errorsFlash) Read() flashDataValue {
	return NewFlashData(flashErrorsKeyName, e.EchoContext.Request()).
		Read(e.EchoContext.Response())
}

// NewErrors -
func NewErrors(c echo.Context, err error) {
	NewErrorsFlash(c).Save(err)
}

// GetAllErrors 获取平铺的错误信息列表
func GetAllErrors(data flashDataValue) []string {
	flatErrors := make([]string, 0)

	if data == nil {
		return flatErrors
	}

	for _, v := range data {
		for _, v2 := range v {
			flatErrors = append(flatErrors, v2)
		}
	}

	return flatErrors
}
