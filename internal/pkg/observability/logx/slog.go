package logx

import (
	"context"
	"io"
	"log/slog"

	"github.com/samber/lo"
)

type logFn func(ctx context.Context, msg string, fields ...any)

var (
	logLevelMap = map[Level]slog.Level{
		LevelError: slog.LevelError,
		LevelWarn:  slog.LevelWarn,
		LevelInfo:  slog.LevelInfo,
		LevelDebug: slog.LevelDebug,
	}
)

type slogCore struct {
	l *slog.Logger
}

func newSlogCore(w io.Writer, l Level) core {
	return &slogCore{
		l: slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
			Level: logLevelMap[l],
		})),
	}
}

func (c *slogCore) log(ctx context.Context, l Level, msg string, fields ...slog.Attr) {
	c.method(l)(ctx, msg, slogAttrSliceToAnySlice(fields...)...)
}

func (c *slogCore) withGroup(g string) core {
	return &slogCore{
		l: c.l.WithGroup(g),
	}
}

func (c *slogCore) withAttrs(fields ...slog.Attr) core {
	return &slogCore{
		l: c.l.With(slogAttrSliceToAnySlice(fields...)...),
	}
}

func (c *slogCore) method(l Level) logFn {
	switch l {
	case LevelError:
		return c.l.ErrorContext
	case LevelWarn:
		return c.l.WarnContext
	case LevelInfo:
		return c.l.InfoContext
	case LevelDebug:
		return c.l.DebugContext
	default:
		return func(_ context.Context, _ string, _ ...any) {}
	}
}

func slogAttrSliceToAnySlice(fields ...slog.Attr) []any {
	return lo.Map(fields, func(item slog.Attr, _ int) any {
		return item
	})
}
