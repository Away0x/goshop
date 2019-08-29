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

	DatabaseErr = &Errno{HTTPCode: http.StatusInternalServerError, Code: 5002, Summary: "Database error", Message: "Database error"}

	// ------------------------------------------------------------------------
	// 用户未登录
	LoginRequiredErr = &Errno{HTTPCode: http.StatusInternalServerError, Code: 6000, Summary: "Login Required", Message: "Login Required"}
	// 用户未激活
	UserActivateErr = &Errno{HTTPCode: http.StatusInternalServerError, Code: 6001, Summary: "User need to activate", Message: "User need to activate"}
	// jwt token error
	JWTTokenErr            = &Errno{HTTPCode: http.StatusInternalServerError, Code: 6002, Summary: "token error", Message: "token error"}
	JWTTokenExpireErr      = &Errno{HTTPCode: http.StatusInternalServerError, Code: 6003, Summary: "Token expired", Message: "Token expired"}
	JWTTokenRefreshErr     = &Errno{HTTPCode: http.StatusInternalServerError, Code: 6004, Summary: "Token has been refreshed", Message: "Token has been refreshed"}
	JWTTokenInBlackListErr = &Errno{HTTPCode: http.StatusInternalServerError, Code: 6005, Summary: "Token cannot be used", Message: "Token cannot be used"}
	JWTTokenMissingErr     = &Errno{HTTPCode: http.StatusInternalServerError, Code: 6006, Summary: "Token not found", Message: "Token not found"}
)
