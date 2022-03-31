package error

var MsgMap = map[int]string{
	SUCCESS: "ok",
	ERROR:   "error out",
}

func GetCodeMessage(code int) string {
	msg, ok := MsgMap[code]
	if ok {
		return msg
	}
	return MsgMap[ERROR]
}
