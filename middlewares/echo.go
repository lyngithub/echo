package middlewares

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
	"echo/common/logger"
	"echo/cons"
	"echo/models"
	"echo/services"
	"io"
	"runtime"
	"strings"
)

var UIDKey = "Con"

var DefaultBodyDumpConfig = middleware.BodyDumpConfig{
	Skipper: BodyDumpDefaultSkipper,
	Handler: func(context echo.Context, bytes []byte, bytes2 []byte) {
		if !strings.HasPrefix(context.Path(), "/api/") && !strings.HasPrefix(context.Path(), "/openapi/") {
			return
		}
		uid := context.Get(UIDKey).(string)
		logger.AdminLog.Info("[ADMIN] "+"end of request", zap.String("request uid", uid), zap.String("request ip", context.RealIP()), zap.String(context.Request().RequestURI, string(bytes2)))

	},
}

func BodyDumpDefaultSkipper(c echo.Context) bool {
	if !strings.HasPrefix(c.Path(), "/api/") && !strings.HasPrefix(c.Path(), "/openapi/") {
		return true
	}
	return false
}

func RequestLog() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					stack := make([]byte, 4<<10)
					length := runtime.Stack(stack, false)
					logger.AdminLog.Error("[ADMIN] "+"program crash", zap.String("crash log", string(stack[:length])))
				}
			}()
			if !strings.HasPrefix(context.Path(), "/api/") && !strings.HasPrefix(context.Path(), "/openapi/") {
				logger.AdminLog.Info("[ADMIN] "+"request to start", zap.Any(context.Request().RequestURI, "web request"))
				return handlerFunc(context)
			}
			uid := uuid.New().String()
			context.Set(UIDKey, uid)
			logger.AdminLog.Info("[ADMIN] "+"request to start", zap.Any(context.Request().RequestURI, uid))
			err := handlerFunc(context)
			return err
		}
	}
}

func LoggerMiddleware() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			bean := &models.Adminlog{
				Method:    c.Request().Method,
				Url:       c.Request().RequestURI,
				Params:    "",
				Ip:        c.RealIP(),
				UserAgent: c.Request().UserAgent(),
			}
			if c.Request().Body != nil {
				all, _ := io.ReadAll(c.Request().Body)
				c.Request().Body = io.NopCloser(bytes.NewBuffer(all))
				bean.Params = string(all)
			}
			if c.Get(cons.ECHO_USER) != nil {
				user := c.Get(cons.ECHO_USER).(*models.User)
				bean.CreateBy = user.ID
				bean.UserName = user.Username
			}
			go func() {
				services.AdminLogChan <- bean
			}()
			return handlerFunc(c)
		}
	}
}
