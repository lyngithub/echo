package middlewares

import (
	"echo/common/logger"
	"echo/models/resp"
	"github.com/labstack/echo"
	"strings"
)

func AuthAppLogin() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				logger.AdminLog.Error("[APP] " + "no token")
				return resp.Auth(c)
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				logger.AdminLog.Error("[APP] " + "token authentication failed")
				return resp.Auth(c)
			}

			//jwt, err := jwt_auth.ParseToken(parts[1])
			//if err != nil {
			//	logger.AdminLog.Error("[APP] " + "token authentication failed")
			//	return resp.Auth(c)
			//}
			//if jwt.Loca != cons.LOGIN_APP {
			//	logger.AdminLog.Error("[APP] " + "token authentication failed")
			//	return resp.Auth(c)
			//}
			//// 添加状态校验
			//err, u := mysql.Merchant.GetDataByUsername(jwt.Username)
			//if u == nil {
			//	logger.AdminLog.Error("[APP] " + "token authentication failed")
			//	return resp.Auth(c)
			//}
			//if cons.STATUS_YES != u.Status {
			//	return resp.AuthMsg(c, "abnormal user status")
			//}
			//if jwt.Uuid != u.ID {
			//	return resp.AuthMsg(c, "token has expired")
			//}
			//c.Set(cons.ECHO_APP_USER, u)

			return handlerFunc(c)
		}
	}
}

// 判断用户是否被禁用
func AuthAppIsDisabled() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				logger.AdminLog.Error("[APP] " + "no token")
				return resp.Auth(c)
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				logger.AdminLog.Error("[APP] " + "token authentication failed")
				return resp.Auth(c)
			}

			//jwt, err := jwt_auth.ParseToken(parts[1])
			//if err != nil {
			//	logger.AdminLog.Error("[APP] " + "token authentication failed")
			//	return resp.Auth(c)
			//}
			//if jwt.Loca != cons.LOGIN_APP {
			//	logger.AdminLog.Error("[APP] " + "token authentication failed")
			//	return resp.Auth(c)
			//}
			// 添加状态校验
			//err, u := mysql.Merchant.GetDataByUsername(jwt.Username)
			//if u == nil {
			//	logger.AdminLog.Error("[APP] " + "token authentication failed")
			//	return resp.Auth(c)
			//}
			//if cons.STATUS_YES != u.Status {
			//	return resp.AuthMsg(c, "abnormal user status")
			//}
			//if jwt.Uuid != u.ID {
			//	return resp.AuthMsg(c, "token has expired")
			//}
			//if u.IsDisabled == cons.STATUS_NO {
			//	return resp.AuthMsg(c, "your account is abnormal and temporarily unavailable.")
			//}
			return handlerFunc(c)
		}
	}
}
