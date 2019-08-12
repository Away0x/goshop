package context

import (
	"echo_shop/database"
	"echo_shop/pkg/errno"
)

// ModelByID 通过 id 查找 model，未找到会返回 404 错误
func (a *AppContext) ModelByID(idkey string, i interface{}) *errno.Errno {
	id, err := a.IntParam(idkey)
	if err != nil {
		return errno.NotFoundErr
	}

	if err = database.DBManager().First(i, id).Error; err != nil {
		return errno.NotFoundErr
	}

	return nil
}
