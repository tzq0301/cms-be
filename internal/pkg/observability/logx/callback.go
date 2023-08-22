package logx

import (
	"log/slog"
	"maps"
)

type Callback func(msg string, fields ...slog.Attr)

type callbacks map[Level][]Callback

func (c callbacks) clone() callbacks {
	return maps.Clone(c)
}
