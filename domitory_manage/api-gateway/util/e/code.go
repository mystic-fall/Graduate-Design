package e

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	ErrorAuthCheckTokenFail    = 30001 //token 错误
	ErrorAuthCheckTokenTimeout = 30002 //token 过期
	ErrorAuthNotFound          = 30005
	ErrorDatabase              = 40001
)
