package resp

import (
	"echo/utils"
	"net/http"

	"github.com/labstack/echo"
)

type Result struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// Resp API Response
func Resp(c echo.Context, success bool, code int, msg string, data interface{}) error {
	res := &Result{
		Success: success,
		Code:    code,
		Msg:     msg,
		Data:    data,
	}
	return c.JSON(http.StatusOK, res)
}

func OK(c echo.Context, data interface{}) error {
	res := &Result{
		Success: true,
		Code:    1000,
		Msg:     "success",
		Data:    data,
	}
	return c.JSON(http.StatusOK, res)
}

func Fail(c echo.Context, msg string) error {
	res := &Result{
		Success: false,
		Code:    5000,
		Msg:     msg,
		Data:    struct{}{},
	}
	return c.JSON(http.StatusOK, res)
}

func NotFindOrder(c echo.Context, msg string) error {
	res := &Result{
		Success: false,
		Code:    1001,
		Msg:     msg,
		Data:    struct{}{},
	}
	return c.JSON(http.StatusOK, res)
}

func Auth(c echo.Context) error {
	res := &Result{
		Success: false,
		Code:    401,
		Msg:     "token authentication failed",
		Data:    struct{}{},
	}
	return c.JSON(http.StatusOK, res)
}

func AuthSign(c echo.Context) error {
	res := &Result{
		Success: false,
		Code:    1002,
		Msg:     "Signature verification failed",
		Data:    struct{}{},
	}
	return c.JSON(http.StatusOK, res)
}

func AuthIP(c echo.Context, ip string) error {
	res := &Result{
		Success: false,
		Code:    1003,
		Msg:     utils.ConnectStr(ip, " is not in the whitelist"),
		Data:    struct{}{},
	}
	return c.JSON(http.StatusOK, res)
}

func AuthMenu(c echo.Context) error {
	res := &Result{
		Success: false,
		Code:    431,
		Msg:     "Insufficient permissions",
		Data:    struct{}{},
	}
	return c.JSON(http.StatusOK, res)
}

func AuthMsg(c echo.Context, msg string) error {
	res := &Result{
		Success: false,
		Code:    401,
		Msg:     msg,
		Data:    struct{}{},
	}
	return c.JSON(http.StatusOK, res)
}
