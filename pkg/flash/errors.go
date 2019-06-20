// Package flash : errors 一般用来回显参数验证的错误
package flash

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

const (
	flashErrorsKeyName = "_flash_errors"
)

type errorsFlash struct {
	EchoContext echo.Context
	Data        flashDataValue
}

type errorsData = flashDataValue

func (errorsData) Error() string {
	return "flash erros"
}

// NewErrorsFlash -
func NewErrorsFlash(c echo.Context) *errorsFlash {
	return &errorsFlash{
		EchoContext: c,
		Data:        make(flashDataValue),
	}
}

func (e *errorsFlash) Save(err error) {
	data, ok := err.(errorsData)
	if !ok {
		data = errorsData{
			"error_msg": {err.Error()},
		}
	}

	flatErrors := make([]string, 0)
	for _, v := range data {
		for _, v2 := range v {
			flatErrors = append(flatErrors, v2)
		}
	}
	fmt.Println(data)
	data["flat_errors"] = flatErrors // 存储平铺的错误信息列表

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
