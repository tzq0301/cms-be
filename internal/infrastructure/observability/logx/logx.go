package logx

import (
	"context"
	"log/slog"
	"os"
)

func Init( /* TODO(TZQ) init by config */ ) (Logger, error) {
	consoleLogger := &slogLogger{
		l: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			// TODO(TZQ)
		})),
	}

	logger := newLoggerCollection(consoleLogger)

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
