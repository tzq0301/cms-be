package runtimex

import (
	"testing"
)

func TestTrace(t *testing.T) {
	wrapper := func() string {
		return TraceOfCaller()
	}

	t.Log(wrapper())

	if wrapper() != Trace() {
		t.FailNow()
	}
}
