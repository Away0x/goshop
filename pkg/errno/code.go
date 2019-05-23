package errno

var (
	ParamsErr = &Errno{HTTPCode: 200, Code: 1000, Summary: "参数错误", Message: "参数错误"}
)
