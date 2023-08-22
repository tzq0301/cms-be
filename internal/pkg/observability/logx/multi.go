package logx

import (
	"context"
	"log/slog"
)

type multiCore struct {
	cores []core
}

func newMultiCore(cores ...core) core {
	return &multiCore{
		cores: cores,
	}
}

func (c *multiCore) log(ctx context.Context, l Level, msg string, fields ...slog.Attr) {
	for _, core := range c.cores {
		core.log(ctx, l, msg, fields...)
	}
}

func (c *multiCore) withGroup(g string) core {
	var cores []core

	for _, core := range c.cores {
		cores = append(cores, core.withGroup(g))
	}

	return &multiCore{
		cores: cores,
	}
}

func (c *multiCore) withAttrs(fields ...slog.Attr) core {
	var cores []core

	for _, core := range c.cores {
		cores = append(cores, core.withAttrs(fields...))
	}

	return &multiCore{
		cores: cores,
	}
}
