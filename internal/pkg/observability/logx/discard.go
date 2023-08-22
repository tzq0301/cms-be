package logx

import (
	"context"
	"log/slog"
)

var (
	discardC core = &discardCore{}
)

type discardCore struct {
}

func (_ discardCore) log(_ context.Context, _ Level, _ string, _ ...slog.Attr) {
}

func (c discardCore) withGroup(_ string) core {
	return c
}

func (c discardCore) withAttrs(_ ...slog.Attr) core {
	return c
}
