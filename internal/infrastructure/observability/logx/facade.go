package logx

import (
	"context"
	"log/slog"
)

type Logger interface {
	Error(msg string, fields ...slog.Attr)
	Warn(msg string, fields ...slog.Attr)
	Info(msg string, fields ...slog.Attr)
	Debug(msg string, fields ...slog.Attr)
	ErrorCtx(ctx context.Context, msg string, fields ...slog.Attr)
	WarnCtx(ctx context.Context, msg string, fields ...slog.Attr)
	InfoCtx(ctx context.Context, msg string, fields ...slog.Attr)
	DebugCtx(ctx context.Context, msg string, fields ...slog.Attr)
}
