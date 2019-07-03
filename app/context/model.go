package context

import (
	"echo_shop/database"
	"echo_shop/pkg/errno"
)

// ModelById 通过 id 查找 model，未找到会进入 404
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
