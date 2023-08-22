package logx

import (
	"errors"
	"strings"

	"github.com/samber/lo"
)

const (
	LevelError Level = "ERROR"
	LevelWarn  Level = "WARN"
	LevelInfo  Level = "INFO"
	LevelDebug Level = "DEBUG"
)

var (
	ErrInvalidLevel = errors.New("invalid logx Level")
)

type Level string

func LevelFromString(s string) (Level, error) {
	switch strings.ToUpper(s) {
	case strings.ToUpper(string(LevelError)):
		return LevelError, nil
	case strings.ToUpper(string(LevelWarn)):
		return LevelWarn, nil
	case strings.ToUpper(string(LevelInfo)):
		return LevelInfo, nil
	case strings.ToUpper(string(LevelDebug)):
		return LevelDebug, nil
	default:
		return lo.Empty[Level](), ErrInvalidLevel
	}
}
