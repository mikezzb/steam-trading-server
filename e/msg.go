package e

var CodeMsg = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",
}

func GetMsg(code int) string {
	msg, ok := CodeMsg[code]
	if ok {
		return msg
	}

	return CodeMsg[ERROR]
}
