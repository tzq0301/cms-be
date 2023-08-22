package logx

import (
	"github.com/samber/lo"

	"cms-be/internal/pkg/stringutil"
)

type Option func(l *Logger)

func WithGroup(name string) Option {
	return func(l *Logger) {
		if stringutil.IsBlank(name) {
			return
		}
		l.group = name
	}
}

func WithCallback(callback Callback, levels ...Level) Option {
	return func(l *Logger) {
		for _, level := range lo.Uniq(levels) {
			l.callbacks[level] = append(l.callbacks[level], callback)
		}
	}
}
