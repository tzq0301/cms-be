package logx

import (
	"context"
	"log/slog"
)

var (
	discardL Logger = &discardLogger{}
)

type discardLogger struct {
}

func (_ discardLogger) Error(_ context.Context, _ string, _ ...slog.Attr) {
}

func (_ discardLogger) Warn(_ context.Context, _ string, _ ...slog.Attr) {
}

func (_ discardLogger) Info(_ context.Context, _ string, _ ...slog.Attr) {
}

func (_ discardLogger) Debug(_ context.Context, _ string, _ ...slog.Attr) {
}

func (l discardLogger) WithAttrs(_ ...slog.Attr) Logger {
	return l
}

func (l discardLogger) withGroup(_ string) Logger {
	return l
}
