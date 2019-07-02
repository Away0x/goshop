package context

import (
	"echo_shop/pkg/errno"
	"strconv"
)

// IntParam 获取 int param
func (a *AppContext) IntParam(key string) (int, error) {
	i, err := strconv.Atoi(a.Param(key))
	if err != nil {
		return 0, errno.ParamsErr
	}

	return i, nil
}
