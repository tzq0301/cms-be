package logx

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/samber/lo"

	"cms-be/internal/pkg/runtimex/shutdownx"
)

type slogLogger struct {
	l *slog.Logger
}

func newSlogConsoleLogger(config ConsoleAppenderConfig) (*slogLogger, error) {
	level, err := logxLevelToSlogLeveler(config.Level)
	if err != nil {
		return nil, errors.Join(err, errors.New("fail to convert logx.Level to slog.Leveler"))
	}

	return &slogLogger{
		l: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})),
	}, nil
}

func newSlogFileLogger(config FileAppenderConfig) (*slogLogger, error) {
	level, err := logxLevelToSlogLeveler(config.Level)
	if err != nil {
		return nil, errors.Join(err, errors.New("fail to convert logx.Level to slog.Leveler"))
	}

	file, err := os.OpenFile(config.FilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("fail to open file: %s", config.FilePath))
	}

	shutdownx.AddHook(file.Close)

	return &slogLogger{
		l: slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{
			Level: level,
		})),
	}, nil
}

func (l *slogLogger) Error(ctx context.Context, msg string, fields ...slog.Attr) {
	l.l.ErrorContext(ctx, msg, lo.Map(fields, func(item slog.Attr, _ int) any {
		return item
	})...)
}

func (l *slogLogger) Warn(ctx context.Context, msg string, fields ...slog.Attr) {
	l.l.WarnContext(ctx, msg, lo.Map(fields, func(item slog.Attr, _ int) any {
		return item
	})...)
}

func (l *slogLogger) Info(ctx context.Context, msg string, fields ...slog.Attr) {
	l.l.InfoContext(ctx, msg, lo.Map(fields, func(item slog.Attr, _ int) any {
		return item
	})...)
}

func (l *slogLogger) Debug(ctx context.Context, msg string, fields ...slog.Attr) {
	l.l.DebugContext(ctx, msg, lo.Map(fields, func(item slog.Attr, _ int) any {
		return item
	})...)
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