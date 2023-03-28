package constant

type BizError struct {
	Code string
	Msg  string
}

func (be *BizError) Error() string {
	return be.Msg
}

const (
	SUCCESS_CODE string = "0000"
	SUCCESS_MSG  string = ""
)

var (
	USER_NOT_EXIST   = &BizError{Code: "0001", Msg: "USER_NOT_EXIST"}
	PASSWORD_INVALID = &BizError{Code: "0002", Msg: "PASSWORD_INVALID"}

	INNER_SERVER_ERR = &BizError{Code: "1000", Msg: "服务异常 请稍后重试"}
)
