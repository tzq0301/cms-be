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

func (c *loggerCollection) Error(ctx context.Context, msg string, fields ...slog.Attr) {
	for _, l := range c.loggers {
		l.Error(ctx, msg, fields...)
	}
}

func (c *loggerCollection) Warn(ctx context.Context, msg string, fields ...slog.Attr) {
	for _, l := range c.loggers {
		l.Warn(ctx, msg, fields...)
	}
}

func (c *loggerCollection) Info(ctx context.Context, msg string, fields ...slog.Attr) {
	for _, l := range c.loggers {
		l.Info(ctx, msg, fields...)
	}
}

func (c *loggerCollection) Debug(ctx context.Context, msg string, fields ...slog.Attr) {
	for _, l := range c.loggers {
		l.Debug(ctx, msg, fields...)
	}
}
