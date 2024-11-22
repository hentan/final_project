package logger

import (
	"log/slog"
)

type LoggerConfig struct {
	Level      slog.Level
	OutputPath string
	Format     string
}

func NewDefaultConfig() *LoggerConfig {
	return &LoggerConfig{
		Level:      slog.LevelInfo,
		OutputPath: "",
		Format:     "json",
	}
}
