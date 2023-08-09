package async

import (
	"errors"
	"fmt"
	"sync"
)

var (
	errLogger = func(err error) {}
	mu        sync.Mutex
)

var (
	ErrNilLogger = errors.New("param Logger is nil")
)

func SetErrLogger(l func(err error)) error {
	if l == nil {
		return ErrNilLogger
	}

	mu.Lock()
	defer mu.Unlock()

	errLogger = l

	return nil
}

func Go(f func()) {
	defer func() {
		if r := recover(); r != nil {
			errLogger(fmt.Errorf("%v", r))
		}
	}()

	f()
}
