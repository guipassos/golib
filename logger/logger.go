package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	cfg     = zap.NewProductionConfig()
	logs    *zap.Logger
	logsApp *zap.Logger
	err     error
)

func init() {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig = encoderCfg
	cfg.DisableStacktrace = true
	logs, err = cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	logsApp, err = cfg.Build(zap.AddCallerSkip(2))
	if err != nil {
		panic(err)
	}
	defer logs.Sync()
}

func Info(msg ...interface{}) {
	logs.Info(fmt.Sprint(msg...))
}

func Error(msg ...interface{}) {
	logs.Error(fmt.Sprint(msg...))
}

func ErrorApp(msg ...interface{}) {
	logsApp.Error(fmt.Sprint(msg...))
}

func InfoApp(msg ...interface{}) {
	logsApp.Info(fmt.Sprint(msg...))
}

func Fatal(msg ...interface{}) {
	logs.Fatal(fmt.Sprint(msg...))
}
