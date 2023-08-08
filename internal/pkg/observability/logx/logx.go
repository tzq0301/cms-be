package logx

import (
	"context"
	"errors"
	"log/slog"
)

func Init(config Config) (Logger, error) {
	var collection []Logger

	if config.ConsoleAppenderConfig != nil {
		consoleLogger, err := newSlogConsoleLogger(*config.ConsoleAppenderConfig)
		if err != nil {
			return nil, errors.Join(err, errors.New("fail to create slogConsoleLogger instance"))
		}

		collection = append(collection, consoleLogger)
	}

	for _, c := range config.FileAppenderConfigs {
		fileLogger, err := newSlogFileLogger(c)
		if err != nil {
			return nil, errors.Join(err, errors.New("fail to create slogFileLogger instance"))
		}

		collection = append(collection, fileLogger)
	}

	logger := newLoggerCollection(collection...)

	return logger, nil
}

func Error(ctx context.Context, msg string, fields ...slog.Attr) {
	LoggerFromContext(ctx).Error(ctx, msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...slog.Attr) {
	LoggerFromContext(ctx).Warn(ctx, msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...slog.Attr) {
	LoggerFromContext(ctx).Info(ctx, msg, fields...)
}

func Debug(ctx context.Context, msg string, fields ...slog.Attr) {
	LoggerFromContext(ctx).Debug(ctx, msg, fields...)
}
