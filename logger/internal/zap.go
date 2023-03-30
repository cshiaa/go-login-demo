package internal

import (

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/cshiaa/go-login-demo/global"

)

func GetEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func GetLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   global.RY_CONFIG.Zap.Filename,
		MaxSize:    global.RY_CONFIG.Zap.MaxSize,
		MaxBackups: global.RY_CONFIG.Zap.MaxBackups,
		MaxAge:     global.RY_CONFIG.Zap.MaxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
