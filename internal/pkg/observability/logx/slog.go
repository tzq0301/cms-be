package logx

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/samber/lo"

	"cms-be/internal/pkg/runtimex/shutdown"
)

type slogLogger struct {
	l *slog.Logger
}

func (l *slogLogger) clone() *slogLogger {
	return &slogLogger{}
}

func newSlogLogger(handler slog.Handler, service ServiceConfig) *slogLogger {
	logger := slog.
		New(handler).
		With(slog.Any("service", service))

	return &slogLogger{
		l: logger,
	}
}

func newSlogConsoleLogger(config ConsoleAppenderConfig) (*slogLogger, error) {
	level, err := logxLevelToSlogLeveler(config.Level)
	if err != nil {
		return nil, errors.Join(err, errors.New("convert logx.Level to slog.Leveler"))
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	return newSlogLogger(handler, config.ServiceConfig), nil
}

func newSlogFileLogger(config FileAppenderConfig) (*slogLogger, error) {
	level, err := logxLevelToSlogLeveler(config.Level)
	if err != nil {
		return nil, errors.Join(err, errors.New("convert logx.Level to slog.Leveler"))
	}

	file, err := os.OpenFile(config.FilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("open file: %s", config.FilePath))
	}

	shutdown.AddHook(file.Close)

	handler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: level,
	})

	return newSlogLogger(handler, config.ServiceConfig), nil
}

func (l *slogLogger) Error(ctx context.Context, msg string, fields ...slog.Attr) {
	l.with().l.ErrorContext(ctx, msg, slogAttrSliceToAnySlice(fields...)...)
}

func (l *slogLogger) Warn(ctx context.Context, msg string, fields ...slog.Attr) {
	l.with().l.WarnContext(ctx, msg, slogAttrSliceToAnySlice(fields...)...)
}

func (l *slogLogger) Info(ctx context.Context, msg string, fields ...slog.Attr) {
	l.with().l.InfoContext(ctx, msg, slogAttrSliceToAnySlice(fields...)...)
}

func (l *slogLogger) Debug(ctx context.Context, msg string, fields ...slog.Attr) {
	l.with().l.DebugContext(ctx, msg, slogAttrSliceToAnySlice(fields...)...)
}

func (l *slogLogger) With(fields ...slog.Attr) Logger {
	cloned := l.clone()
	cloned.l = l.l.With(slogAttrSliceToAnySlice(fields...)...)
	return cloned
}

// TODO
func (l *slogLogger) with() *slogLogger {
	cloned := l.clone()
	cloned.l = l.l.WithGroup("fields")
	return cloned
}

func logxLevelToSlogLeveler(l Level) (slog.Leveler, error) {
	switch l {
	case LevelError:
		return slog.LevelError, nil
	case LevelWarn:
		return slog.LevelWarn, nil
	case LevelInfo:
		return slog.LevelInfo, nil
	case LevelDebug:
		return slog.LevelDebug, nil
	default:
		return nil, ErrInvalidLevel
	}
}

func slogAttrSliceToAnySlice(fields ...slog.Attr) []any {
	return lo.Map(fields, func(item slog.Attr, _ int) any {
		return item
	})
}
