package session

import (
	sess "echo_shop/pkg/echohelper/session/echo-session"

	"github.com/labstack/echo/v4"
)

// Set 设置 session
func Set(c echo.Context, key string, val string) {
	s := sess.Default(c)
	if s == nil {
		return
	}
	s.Set(key, val)
	s.Save()
}

// Get 获取 session
func Get(c echo.Context, key string) string {
	s := sess.Default(c)
	if s == nil {
		return ""
	}
	v := s.Get(key)
	if v == nil {
		return ""
	}

	if str, ok := v.(string); ok {
		return str
	}

	return ""
}

// Delete 删除 session
func Delete(c echo.Context, key string) {
	s := sess.Default(c)
	s.Delete(key)
	s.Save()
}
