package main

import (
	"echo/bootstrap"
	"echo/common/logger"
	"echo/conf"
	"echo/router"
	"go.uber.org/zap"
)

var (
	//cfg = "/code/config.toml"
	cfg = "config.toml"
	err error
)

func init() {
	conf.InitializeConfig(cfg)
	bootstrap.Bootstrap()
}

func main() {
	errs := make(chan error)
	router.RunHttpServer(conf.Config.Echo.Bind)
	select {
	case err := <-errs:
		logger.AdminLog.Fatal("goldenShieldadmin: Run Server failed, err: %v", zap.Error(err))
	}
}
