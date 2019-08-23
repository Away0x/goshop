// Package flash : old value 用于回填表单 (记忆上次表单提交的数据)
package flash

import (
	"net/http"

	"github.com/Away0x/flash"
	"github.com/labstack/echo/v4"
)

const (
	flashOldValueKeyName = "_flash_old_value"
)

// OldValueFlash -
type OldValueFlash struct {
	EchoContext echo.Context
	Data        flash.DataValue
}

// NewOldValueFlash -
func NewOldValueFlash(c echo.Context) *OldValueFlash {
	return &OldValueFlash{
		EchoContext: c,
		Data:        make(flash.DataValue),
	}
}

// Save -
func (o *OldValueFlash) Save(data flash.DataValue) {
	flash.NewFlash(flashOldValueKeyName, o.EchoContext.Request()).
		Set(data).
		Save(o.EchoContext.Response())
}

func (o *OldValueFlash) Read() flash.DataValue {
	return flash.NewFlash(flashOldValueKeyName, o.EchoContext.Request()).
		Read(o.EchoContext.Response())
}

// OldValueFlashMiddleware 中间件
func OldValueFlashMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()

			switch req.Method {
			case http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodTrace:
				return next(c)
			default:
				if req.Form == nil {
					req.ParseForm()
				}

				olaValue := make(flash.DataValue)
				for k, v := range req.Form {
					olaValue[k] = []string{v[0]}
				}
				NewOldValueFlash(c).Save(olaValue)

				return next(c)
			}
		}
	}
}
