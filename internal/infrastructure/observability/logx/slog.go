package logx

import (
	"context"
	"log/slog"
)

type logger struct {
	l *slog.Logger
}

func (l *logger) Error(msg string, fields ...slog.Attr) {
	l.l.Error(msg, fields)
}

func (l *logger) Warn(msg string, fields ...slog.Attr) {
	l.l.Warn(msg, fields)
}

func (l *logger) Info(msg string, fields ...slog.Attr) {
	l.l.Info(msg, fields)
}

func (l *logger) Debug(msg string, fields ...slog.Attr) {
	l.l.Debug(msg, fields)
}

func (l *logger) ErrorCtx(ctx context.Context, msg string, fields ...slog.Attr) {
	l.l.ErrorContext(ctx, msg, fields)
}

func (l *logger) WarnCtx(ctx context.Context, msg string, fields ...slog.Attr) {
	l.l.WarnContext(ctx, msg, fields)
}

func (l *logger) InfoCtx(ctx context.Context, msg string, fields ...slog.Attr) {
	l.l.InfoContext(ctx, msg, fields)
}

func (l *logger) DebugCtx(ctx context.Context, msg string, fields ...slog.Attr) {
	l.l.DebugContext(ctx, msg, fields)
}
