package logx

import (
	"context"
	"log/slog"
)

var l Logger

func Init() {
	l = &logger{
		l: &slog.Logger{},
	}
}

func Error(msg string, fields ...slog.Attr) {
	l.Error(msg, fields...)
}

func Warn(msg string, fields ...slog.Attr) {
	l.Warn(msg, fields...)
}

func Info(msg string, fields ...slog.Attr) {
	l.Info(msg, fields...)
}

func Debug(msg string, fields ...slog.Attr) {
	l.Debug(msg, fields...)
}

func ErrorCtx(ctx context.Context, msg string, fields ...slog.Attr) {
	l.ErrorCtx(ctx, msg, fields...)
}

func WarnCtx(ctx context.Context, msg string, fields ...slog.Attr) {
	l.WarnCtx(ctx, msg, fields...)
}

func InfoCtx(ctx context.Context, msg string, fields ...slog.Attr) {
	l.InfoCtx(ctx, msg, fields...)
}

func DebugCtx(ctx context.Context, msg string, fields ...slog.Attr) {
	l.DebugCtx(ctx, msg, fields...)
}
