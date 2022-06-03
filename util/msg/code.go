package msg

const (
	SUCCESS        = 200
	REDIRECT       = 203
	ERROR          = 500
	INVALID_PARAMS = 400

	ERR_MSG                = 1000
	ERROR_USER_NOT_EXISTS  = 1003
	ERROR_TOKEN_NOT_EXISTS = 1004
	ERROR_TOKEN_OVERDUE    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
)

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	REDIRECT:       "redirect",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERR_MSG:                "错误",
	ERROR_USER_NOT_EXISTS:  "用户不存在",
	ERROR_TOKEN_NOT_EXISTS: "TOKEN不存在",
	ERROR_TOKEN_OVERDUE:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN错误",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
