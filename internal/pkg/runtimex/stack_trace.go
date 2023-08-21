package runtimex

import (
	"fmt"
	"path/filepath"
	"runtime"
)

const (
	skipTrace                 = 1
	skipTraceOfCaller         = 2
	skipTraceOfCallerOfCaller = 3
)

func StackTrace() string {
	return stackTrace(skipTrace)
}

func StackTraceOfCaller() string {
	return stackTrace(skipTraceOfCaller)
}

func StackTraceOfCallerOfCaller() string {
	return stackTrace(skipTraceOfCallerOfCaller)
}

func stackTrace(skip int) string {
	pc, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}

	funcName := filepath.Base(runtime.FuncForPC(pc).Name())

	return fmt.Sprintf("%s(...) at %s:%d", funcName, file, line)
}
