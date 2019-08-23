// Package flash : errors 一般用来回显参数验证的错误
package flash

import (
	"github.com/Away0x/flash"
	"github.com/Away0x/validate"

	"github.com/labstack/echo/v4"
)

const (
	flashErrorsKeyName = "_flash_errors"
)

type ErrorsFlash struct {
	EchoContext echo.Context
	Data        flash.DataValue
}

// NewErrorsFlash -
func NewErrorsFlash(c echo.Context) *ErrorsFlash {
	return &ErrorsFlash{
		EchoContext: c,
		Data:        make(flash.DataValue),
	}
}

func (e *ErrorsFlash) Save(err error) {
	data, ok := err.(validate.Messages)
	if !ok {
		data = validate.Messages{"unknown": {err.Error()}}
	}

	flash.NewFlash(flashErrorsKeyName, e.EchoContext.Request()).
		Set(flash.DataValue(data)).
		Save(e.EchoContext.Response())
}

func (e *ErrorsFlash) Read() flash.DataValue {
	return flash.NewFlash(flashErrorsKeyName, e.EchoContext.Request()).
		Read(e.EchoContext.Response())
}

// NewErrors -
func NewErrors(c echo.Context, err error) {
	NewErrorsFlash(c).Save(err)
}

// GetAllErrors 获取平铺的错误信息列表
func GetAllErrors(data flash.DataValue) []string {
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
