package logx

import (
	"log/slog"
	"testing"
)

func TestWithGroup(t *testing.T) {
	l := Logger{}
	WithGroup("group")(&l)
	if l.group != "group" {
		t.FailNow()
	}
}

func TestWithCallback(t *testing.T) {
	l := Logger{
		callbacks: make(callbacks),
	}

	WithCallback(func(msg string, fields ...slog.Attr) {
	}, LevelInfo, LevelWarn, LevelWarn)(&l)

	if len(l.callbacks[LevelError]) != 0 {
		t.FailNow()
	}

	if len(l.callbacks[LevelWarn]) != 1 {
		t.FailNow()
	}

	if len(l.callbacks[LevelInfo]) != 1 {
		t.FailNow()
	}

	if len(l.callbacks[LevelDebug]) != 0 {
		t.FailNow()
	}
}
