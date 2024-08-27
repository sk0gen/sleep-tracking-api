package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	LogLevel        string `env:"LOG_LEVEL" ,envDefault:"info"`
	DevelopmentMode bool   `env:"DEVELOPMENT_MODE" ,envDefault:"false"`
}

const (
	encodingConsole = "console"
	encodingJSON    = "json"
)

func encoding(cfg Config) string {
	if cfg.DevelopmentMode {
		return encodingConsole
	}
	return encodingJSON
}

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
		Encoding:         encoding(cfg),
		EncoderConfig:    zapEncoderConfig,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	return zapConfig.Build()
}
