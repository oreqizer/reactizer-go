package utils

import "github.com/kataras/iris"

const (
	noAuthHeader     = "auth.no_auth_header"
	invalidToken     = "auth.invalid_token"
	invalidPassword  = "auth.invalid_password"
	passwordTooShort = "auth.password_too_short"
	passwordTooLong  = "auth.password_too_long"
	passwordNoNumber = "auth.password_no_number"
	passwordNoUpper  = "auth.password_no_upper"
	passwordNoLower  = "auth.password_no_lower"
)

func Error(c iris.Context, text string, statusCode int) {
	c.StatusCode(statusCode)
	c.WriteString(GetT(c)(text))
	c.StopExecution()
}
