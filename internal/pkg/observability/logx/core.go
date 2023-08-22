package logx

import (
	"context"
	"log/slog"
)

type core interface {
	log(ctx context.Context, l Level, msg string, fields ...slog.Attr)
	withGroup(g string) core
	withAttrs(fields ...slog.Attr) core
}
