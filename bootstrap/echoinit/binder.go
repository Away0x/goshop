package echoinit

import (
	"github.com/labstack/echo/v4"
)

// CustomBinder 自定义 binder
type CustomBinder struct{}

// Bind : bind params and validate
func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	d := new(echo.DefaultBinder)
	if err = d.Bind(i, c); err != echo.ErrUnsupportedMediaType {
		return
	}

	return nil
}

// SetupBinder -
func SetupBinder(e *echo.Echo) {
	e.Binder = &CustomBinder{}
}
