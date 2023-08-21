package logx

import (
	"context"
	"log/slog"
)

type Logger interface {
	Error(ctx context.Context, msg string, fields ...slog.Attr)
	Warn(ctx context.Context, msg string, fields ...slog.Attr)
	Info(ctx context.Context, msg string, fields ...slog.Attr)
	Debug(ctx context.Context, msg string, fields ...slog.Attr)
	WithAttrs(fields ...slog.Attr) Logger
	withGroup(g string) Logger
}
