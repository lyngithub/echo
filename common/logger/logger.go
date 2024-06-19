package logger

import (
	"github.com/natefinch/lumberjack"
	"github.com/robfig/cron"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	AdminLog   *zap.Logger
	AppLog     *zap.Logger
	OpenApiLog *zap.Logger
)

func InitAdminLog(filepath string) {
	encoder := getEncoder()
	writeSyncer := getLogWriter(filepath)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	//consoleDebug := zapcore.Lock(os.Stdout)
	//consoleEncodeer := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	//p := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
	//	return level >= zapcore.DebugLevel
	//})
	var allcode []zapcore.Core
	allcode = append(allcode, core)
	//allcode = append(allcode, zapcore.NewCore(consoleEncodeer, consoleDebug, p))
	c := zapcore.NewTee(allcode...)
	//zap.AddCaller() //添加将调用函数信息记录到日志中的功能。
	AdminLog = zap.New(c, zap.AddCaller())
}

func InitAppLog(filepath string) {
	encoder := getEncoder()
	writeSyncer := getLogWriter(filepath)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	//consoleDebug := zapcore.Lock(os.Stdout)
	//consoleEncodeer := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	//p := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
	//	return level >= zapcore.DebugLevel
	//})
	var allcode []zapcore.Core
	allcode = append(allcode, core)
	//allcode = append(allcode, zapcore.NewCore(consoleEncodeer, consoleDebug, p))
	c := zapcore.NewTee(allcode...)
	//zap.AddCaller() //添加将调用函数信息记录到日志中的功能。
	AppLog = zap.New(c, zap.AddCaller())
}

func InitOpenApiLog(filepath string) {
	encoder := getEncoder()
	writeSyncer := getLogWriter(filepath)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	//consoleDebug := zapcore.Lock(os.Stdout)
	//consoleEncodeer := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	//p := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
	//	return level >= zapcore.DebugLevel
	//})
	var allcode []zapcore.Core
	allcode = append(allcode, core)
	//allcode = append(allcode, zapcore.NewCore(consoleEncodeer, consoleDebug, p))
	c := zapcore.NewTee(allcode...)
	//zap.AddCaller() //添加将调用函数信息记录到日志中的功能。
	OpenApiLog = zap.New(c, zap.AddCaller())
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// NewConsoleEncoder 打印更符合观察的方式
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(filepath string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath,
		MaxSize:    10240,
		MaxAge:     7,
		Compress:   true,
		LocalTime:  true,
		MaxBackups: 4,
	}

	c := cron.New()
	c.AddFunc("0 0 0 1/1 * ?", func() {
		lumberJackLogger.Rotate()
	})
	c.Start()
	return zapcore.AddSync(lumberJackLogger)
}

// 数据库日志
var Xlogger *XormLogger

type XormLogger struct {
}

func (this *XormLogger) Write(p []byte) (n int, err error) {
	AppLog.Info("数据库操作", zap.String("数据库", string(p)))
	return len(p), nil
}

// echo日志
var EchoLog *EchoLogger

type EchoLogger struct {
}

func (this *EchoLogger) Write(p []byte) (n int, err error) {
	AdminLog.Info("ECHO", zap.String("请求", string(p)))
	return len(p), nil
}
