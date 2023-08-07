package logx

import (
	"context"
	"log/slog"

	"github.com/samber/lo"
)

type slogLogger struct {
	l *slog.Logger
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
