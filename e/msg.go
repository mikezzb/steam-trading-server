package e

var CodeMsg = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	SERVER_ERROR:   "server error",
	INVALID_PARAMS: "invalid params",

	ERROR_AUTH_CHECK_TOKEN_EXPIRED: "token expired",
	ERROR_INVALID_AUTH_HEADER:      "invalid auth header",
	ERROR_USER_NOT_EXIST:           "user not exist",
	ERROR_USER_WRONG_PWD:           "wrong password",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "check token fail",
	ERROR_AUTH_CHECK_ROLE_FAIL:     "check role fail",
}

func GetMsg(code int) string {
	msg, ok := CodeMsg[code]
	if ok {
		return msg
	}

	return CodeMsg[ERROR]
}
