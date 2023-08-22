package async

import (
	"sync"
	"testing"

	"github.com/pkg/errors"
)

func TestAsyncGo(t *testing.T) {
	expectError := errors.New("expect")

	var actualError error

	if SetErrLogger(func(err error) {
		actualError = err
	}) != nil {
		t.Fatal()
	}

	var wg sync.WaitGroup

	wg.Add(1)
	Go(func() {
		defer wg.Done()
	})
	wg.Wait()

	if actualError != nil {
		t.Fatal()
	}

	wg.Add(1)
	Go(func() {
		defer wg.Done()
		panic(expectError)
	})
	wg.Wait()

	if actualError.Error() != expectError.Error() {
		t.Fatal()
	}
}
