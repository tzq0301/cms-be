package logx

import (
	"context"
	"log/slog"
)

type loggerCollection struct {
	loggers []Logger
}

func newLoggerCollection(loggers ...Logger) Logger {
	return &loggerCollection{
		loggers: loggers,
	}
}

func (l *loggerCollection) Error(ctx context.Context, msg string, fields ...slog.Attr) {
	for _, l := range l.loggers {
		l.Error(ctx, msg, fields...)
	}
}

func (l *loggerCollection) Warn(ctx context.Context, msg string, fields ...slog.Attr) {
	for _, l := range l.loggers {
		l.Warn(ctx, msg, fields...)
	}
}

func (l *loggerCollection) Info(ctx context.Context, msg string, fields ...slog.Attr) {
	for _, l := range l.loggers {
		l.Info(ctx, msg, fields...)
	}
}

func (l *loggerCollection) Debug(ctx context.Context, msg string, fields ...slog.Attr) {
	for _, l := range l.loggers {
		l.Debug(ctx, msg, fields...)
	}
}

func (l *loggerCollection) With(fields ...slog.Attr) Logger {
	var loggers []Logger

	for _, logger := range l.loggers {
		loggers = append(loggers, logger.With(fields...))
	}

	return &loggerCollection{
		loggers: loggers,
	}
}
