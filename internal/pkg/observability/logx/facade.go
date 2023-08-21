package logx

import (
	"context"
	"log/slog"
)

type Logger interface {
	With(fields ...slog.Attr) Logger
	Error(ctx context.Context, msg string, fields ...slog.Attr)
	Warn(ctx context.Context, msg string, fields ...slog.Attr)
	Info(ctx context.Context, msg string, fields ...slog.Attr)
	Debug(ctx context.Context, msg string, fields ...slog.Attr)
}
