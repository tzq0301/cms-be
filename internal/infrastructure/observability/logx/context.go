package logx

import "context"

var logxKey = contextKey{}

type contextKey struct{}

func ContextWithLogger(ctx context.Context, logger Logger) context.Context {
	if logger == nil {
		logger = discardL
	}

	return context.WithValue(ctx, logxKey, logger)
}

func LoggerFromContext(ctx context.Context) Logger {
	val := ctx.Value(logxKey)
	if logger, ok := val.(Logger); ok {
		return logger
	}
	return discardL
}
