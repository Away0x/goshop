package context

// 权限相关

import (
	"echo_shop/pkg/constants"
	"echo_shop/pkg/errno"
)

// AuthCheck 权限检查
func (a *AppContext) AuthCheck(status bool) *errno.Errno {
	if status {
		return nil
	}

	if constants.IsAPIRequest(a) {
		return errno.UnauthorizedErr.JSON()
	}

	return errno.UnauthorizedErr.HTML()
}

// AuthCheckByFunc 权限检查
func (a *AppContext) AuthCheckByFunc(fn func() bool) *errno.Errno {
	return a.AuthCheck(fn())
}
