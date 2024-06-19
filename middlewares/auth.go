package middlewares

import (
	"echo/common/logger"
	"echo/cons"
	"echo/daos/mysql"
	"echo/models/resp"
	"echo/systemSetting"
	"echo/utils/jwt_auth"
	"github.com/labstack/echo"
	"strings"
)

func AuthLogin() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				logger.AdminLog.Error("[ADMIN] " + "no token")
				return resp.Auth(c)
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				logger.AdminLog.Error("[ADMIN] " + "token authentication failed")
				return resp.Auth(c)
			}

			jwt, err := jwt_auth.ParseToken(parts[1])
			if err != nil {
				logger.AdminLog.Error("[ADMIN] " + "token authentication failed")
				return resp.Auth(c)
			}
			//if jwt.Loca != cons.LOGIN_ADMIN {
			//	logger.AdminLog.Error("[APP] " + "token authentication failed")
			//	return resp.Auth(c)
			//}
			// 添加状态校验
			u, err := mysql.User.GetBeanByLoginName(jwt.Username)
			if u == nil {
				logger.AdminLog.Error("[ADMIN] " + "token authentication failed")
				return resp.Auth(c)
			}
			if cons.STATUS_YES != u.Status {
				return resp.AuthMsg(c, "abnormal user status")
			}
			if jwt.Uuid != u.ID {
				return resp.AuthMsg(c, "token has expired")
			}
			// 验证T出
			_, has := systemSetting.LoginAdmin[u.ID]
			if !has {
				//return resp.AuthMsg(c, "token has expired")
			}
			c.Set(cons.ECHO_USER, u)

			return handlerFunc(c)
		}
	}
}
