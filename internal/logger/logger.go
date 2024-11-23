package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type LoggerConfig struct {
	Level          slog.Level
	OutputPath     string
	HandlerFactory func(output *os.File) (slog.Handler, error)
}

func (cfg *LoggerConfig) CreateHandler() (slog.Handler, error) {
	var output *os.File
	var err error

	if cfg.OutputPath == "" {
		output = os.Stdout
	} else {
		output, err = os.OpenFile(cfg.OutputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
	}

	if cfg.HandlerFactory == nil {
		return nil, fmt.Errorf("HadlerFactory отсутствует в LoggerConfig")
	}
	return cfg.HandlerFactory(output)
}

func NewConfigWithFormat(format string) (*LoggerConfig, error) {
	var factory func(output *os.File) (slog.Handler, error)

	switch format {
	case "json":
		factory = func(output *os.File) (slog.Handler, error) {
			return slog.NewJSONHandler(output, nil), nil
		}
	case "text":
		factory = func(output *os.File) (slog.Handler, error) {
			return slog.NewTextHandler(output, nil), nil
		}
	default:
		factory = func(output *os.File) (slog.Handler, error) {
			return slog.NewJSONHandler(output, nil), nil
		}
	}

	return &LoggerConfig{
		Level:          slog.LevelInfo,
		OutputPath:     "",
		HandlerFactory: factory,
	}, nil
}

func ParseLogLevel(level string) (slog.Level, error) {
	switch level {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return slog.LevelInfo, fmt.Errorf("invalid log level: %s", level)
	}
}

func NewLogger(cfg *LoggerConfig) (*slog.Logger, error) {
	handler, err := cfg.CreateHandler()
	if err != nil {
		return nil, err
	}
	return slog.New(handler), nil
}

var globalLogger *slog.Logger

func InitGlobalLogger(cfg *LoggerConfig) error {
	var err error
	globalLogger, err = NewLogger(cfg)
	if err != nil {
		return err
	}
	return nil
}

func GetLogger() *slog.Logger {
	if globalLogger == nil {
		globalLogger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}
	return globalLogger
}
