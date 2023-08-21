package runtimex

import (
	"testing"
)

func TestTrace(t *testing.T) {
	wrapperOfwrapper := func() string {
		return func() string {
			return StackTraceOfCallerOfCaller()
		}()
	}

	wrapper := func() string {
		return StackTraceOfCaller()
	}

	// t.Log(wrapper())

	if wrapper() != StackTrace() && wrapperOfwrapper() != StackTrace() {
		t.FailNow()
	}
}
