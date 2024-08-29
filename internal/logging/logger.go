package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Config struct {
	LogLevel        string `env:"LOG_LEVEL" ,envDefault:"info"`
	DevelopmentMode bool   `env:"DEVELOPMENT_MODE" ,envDefault:"false"`
}

const (
	encodingJSON = "json"
)

func InitZap(cfg Config) (*zap.Logger, error) {
	level := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	switch cfg.LogLevel {
	case "debug":
		level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "info":
		level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warn":
		level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case "fatal":
		level = zap.NewAtomicLevelAt(zapcore.FatalLevel)
	case "panic":
		level = zap.NewAtomicLevelAt(zapcore.PanicLevel)
	}

	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./logs/sleep-tracking-api.logs",
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     28,
	})

	zapEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	zapConfig := zap.Config{
		Level:            level,
		Development:      cfg.DevelopmentMode,
		Encoding:         encodingJSON,
		EncoderConfig:    zapEncoderConfig,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	fileCore := zapcore.NewCore(zapcore.NewJSONEncoder(zapConfig.EncoderConfig), writeSyncer, level)
	consoleCore := zapcore.NewCore(zapcore.NewConsoleEncoder(zapConfig.EncoderConfig), zapcore.AddSync(os.Stdout), level)

	core := zapcore.NewTee(fileCore, consoleCore)
	logger := zap.New(core)
	return logger, nil
}
