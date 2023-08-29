package logx

import (
	"context"
	"log/slog"

	"cms-be/internal/pkg/runtimex"
)

const (
	keyStackTrace = "stack"
)

type Logger struct {
	core      core
	callbacks callbacks
	group     string
	service   ServiceConfig
}

func (logger *Logger) With(fields ...slog.Attr) *Logger {
	cloned := logger.Clone()
	cloned.core = cloned.core.withAttrs(fields...)
	return cloned
}

func (logger *Logger) Clone() *Logger {
	return &Logger{
		core:      logger.core,
		callbacks: logger.callbacks.clone(),
		group:     logger.group,
		service:   logger.service,
	}
}

func (logger *Logger) log(ctx context.Context, l Level, msg string, fields ...slog.Attr) {
	var attrs = []slog.Attr{
		slog.String(keyStackTrace, runtimex.StackTraceOfCallerOfCaller()),
		slog.Any("service", logger.service),
	}

	logger.core.
		withAttrs(attrs...).
		withGroup(logger.group).
		log(ctx, l, msg, fields...)

	for _, callback := range logger.callbacks[l] {
		callback(msg, fields...)
	}
}
