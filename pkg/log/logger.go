package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"sync"
)

var setupOnce sync.Once
var Logger *zap.Logger

// SetupLogger - Ensures log setup runs only once
func SetupLogger(environment string, serviceName string) {
	setupOnce.Do(func() {
		writerSyncer := getLogWriter(serviceName)

		core := getCore(environment, writerSyncer)

		Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	})
}

/**
Specifies where the logs will be written to
*/
func getLogWriter(serviceName string) zapcore.WriteSyncer {
	lumberJackLogger := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./" + serviceName + ".log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})
	return lumberJackLogger
}

/**
Specifies how the logs will be written and which ones to log
*/
func getCore(environment string, writerSyncer zapcore.WriteSyncer) zapcore.Core {
	var core zapcore.Core

	if environment == "dev" {
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
		jsonFileEncoder := zapcore.NewJSONEncoder(encoderConfig)
		core = zapcore.NewTee(
			zapcore.NewCore(jsonFileEncoder, writerSyncer, zapcore.DebugLevel),
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel))
	} else {
		encoderConfig := zap.NewProductionEncoderConfig()
		jsonFileEncoder := zapcore.NewJSONEncoder(encoderConfig)
		core = zapcore.NewCore(jsonFileEncoder, writerSyncer, zapcore.InfoLevel)
	}

	return core
}
