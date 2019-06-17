package errno

import (
	"net/http"
)

// 只有特定情况下的返回非 200 的 http code，具体错误通过 Errno.Code 说明
var (
	// 400
	ParamsErr          = &Errno{HTTPCode: http.StatusBadRequest, Code: 4000, Summary: "参数错误", Message: "参数错误"}
	UnauthorizedErr    = &Errno{HTTPCode: http.StatusUnauthorized, Code: 4001, Summary: "unauthorized", Message: "unauthorized"}
	ForbiddenErr       = &Errno{HTTPCode: http.StatusForbidden, Code: 4002, Summary: "forbidden", Message: "forbidden"}
	NotFoundErr        = &Errno{HTTPCode: http.StatusNotFound, Code: 4003, Summary: "not found", Message: "not found"}
	TooManyRequestsErr = &Errno{HTTPCode: http.StatusTooManyRequests, Code: 4004, Summary: "too many requests", Message: "too many requests"}

	// 500
	UnknowErr      = &Errno{HTTPCode: http.StatusInternalServerError, Code: 5000, Summary: "unk", Message: "未知错误"}
	NotImplemented = &Errno{HTTPCode: http.StatusNotImplemented, Code: 5001, Summary: "not implemented", Message: "not implemented"}
)
