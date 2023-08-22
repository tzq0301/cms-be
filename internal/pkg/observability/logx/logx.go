package logx

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
)

const (
	defaultGroupName = "fields"
)

func Init(config Config, options ...Option) (*Logger, error) {
	var cores []core

	if config.ConsoleAppenderConfig != nil {
		consoleCore := newSlogCore(os.Stdout, config.ConsoleAppenderConfig.Level)
		cores = append(cores, consoleCore)
	}

	for _, c := range config.FileAppenderConfigs {
		file, err := os.OpenFile(c.FilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, errors.Join(err, fmt.Errorf("open file: %s", c.FilePath))
		}

		fileCore := newSlogCore(file, c.Level)
		cores = append(cores, fileCore)
	}

	logger := Logger{
		core:      newMultiCore(cores...),
		callbacks: make(callbacks),
		group:     defaultGroupName,
		service:   config.ServiceConfig,
	}

	for _, option := range options {
		option(&logger)
	}

	return &logger, nil
}

func Error(ctx context.Context, msg string, fields ...slog.Attr) {
	LoggerFromContext(ctx).log(ctx, LevelError, msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...slog.Attr) {
	LoggerFromContext(ctx).log(ctx, LevelWarn, msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...slog.Attr) {
	LoggerFromContext(ctx).log(ctx, LevelInfo, msg, fields...)
}

func Debug(ctx context.Context, msg string, fields ...slog.Attr) {
	LoggerFromContext(ctx).log(ctx, LevelDebug, msg, fields...)
}
