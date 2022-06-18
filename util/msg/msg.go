package msg

const (
	SUCCESS  = 200
	REDIRECT = 203
	ERROR    = 500

	//ERR_MSG = 1000
)

var Flags = map[int]string{
	SUCCESS:  "ok",
	REDIRECT: "redirect",
	ERROR:    "request fail",
}

func GetMsg(code int) string {
	msg, ok := Flags[code]
	if ok {
		return msg
	}
	return Flags[ERROR]
}
