package runtimex

import (
	"fmt"
	"path/filepath"
	"runtime"
)

const (
	skipTrace         = 1
	skipTraceOfCaller = 2
)

func Trace() string {
	return trace(skipTrace)
}

func TraceOfCaller() string {
	return trace(skipTraceOfCaller)
}

func trace(skip int) string {
	pc, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}

	funcName := filepath.Base(runtime.FuncForPC(pc).Name())

	return fmt.Sprintf("%s(...) at %s:%d\n", funcName, file, line)
}
