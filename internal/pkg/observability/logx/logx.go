package logx

import (
	"context"
	"errors"
	"log/slog"

	"cms-be/internal/pkg/runtimex"
)

const (
	defaultGroupName = "fields"

	keyStackTrace = "stack"
)

func Init(config Config) (Logger, error) {
	var collection []Logger

	if config.ConsoleAppenderConfig != nil {
		consoleLogger, err := newSlogConsoleLogger(*config.ConsoleAppenderConfig)
		if err != nil {
			return nil, errors.Join(err, errors.New("create slogConsoleLogger instance"))
		}

		collection = append(collection, consoleLogger)
	}

	for _, c := range config.FileAppenderConfigs {
		fileLogger, err := newSlogFileLogger(c)
		if err != nil {
			return nil, errors.Join(err, errors.New("create slogFileLogger instance"))
		}

		collection = append(collection, fileLogger)
	}

	logger := newLoggerCollection(collection...)

	return logger, nil
}

func Error(ctx context.Context, msg string, fields ...slog.Attr) {
	log(ctx).Error(ctx, msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...slog.Attr) {
	log(ctx).Warn(ctx, msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...slog.Attr) {
	log(ctx).Info(ctx, msg, fields...)
}

func Debug(ctx context.Context, msg string, fields ...slog.Attr) {
	log(ctx).Debug(ctx, msg, fields...)
}

func log(ctx context.Context) Logger {
	l := LoggerFromContext(ctx).
		WithAttrs(slog.String(keyStackTrace, runtimex.StackTraceOfCallerOfCaller())).
		withGroup(defaultGroupName)

	return l
}
