package xerr

var codeDesc = map[int]string{
	SERVER_COMMON_ERROR: "服务器异常，稍后再尝试",
	REQUEST_PARAMERROR:  "请求参数有误 ",
	DB_ERROR:            "数据库错误",
}

func ErrMsg(errcode int) string {
	if msg, ok := codeDesc[errcode]; ok {
		return msg
	}
	return "default ERROR!"
}
