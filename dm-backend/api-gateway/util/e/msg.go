package e

var MsgFlags = map[uint]string{
	SUCCESS:                    "ok",
	ERROR:                      "fail",
	InvalidParams:              "请求参数错误",
	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorDatabase:              "数据库操作出错,请重试",
	ErrorAuthNotFound:          "用户未登录",
}

// GetMsg 获取状态码对应信息
func GetMsg(code uint) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
