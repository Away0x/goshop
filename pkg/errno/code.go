package errno

import (
	"net/http"
)

// 只有特定情况下的返回非 200 的 http code，具体错误通过 Errno.Code 说明
var (
	// ------------------------------- 400x -------------------------------
	ParamsErr          = &Errno{HTTPCode: http.StatusBadRequest, Code: 4000, Summary: "Params Error", Message: "Params Error"}
	UnauthorizedErr    = &Errno{HTTPCode: http.StatusUnauthorized, Code: 4001, Summary: "Unauthorized", Message: "Unauthorized"}
	ForbiddenErr       = &Errno{HTTPCode: http.StatusForbidden, Code: 4002, Summary: "Forbidden", Message: "Forbidden"}
	NotFoundErr        = &Errno{HTTPCode: http.StatusNotFound, Code: 4003, Summary: "Not Found", Message: "Not Found"}
	TooManyRequestsErr = &Errno{HTTPCode: http.StatusTooManyRequests, Code: 4004, Summary: "Too Many Requests", Message: "Too Many Requests"}

	// ------------------------------- 500x -------------------------------
	UnknowErr = &Errno{HTTPCode: http.StatusInternalServerError, Code: 5000, Summary: "Unknown Error", Message: "Unknown Error"}
	// 还未定义的接口可以使用该错误响应
	NotImplemented = &Errno{HTTPCode: http.StatusNotImplemented, Code: 5001, Summary: "Not Implemented", Message: "Not Implemented"}
)
