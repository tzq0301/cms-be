package adaptor

import (
	"context"
	"sync"
)

type Adaptor interface {
	// Run will block the goroutine, until the error occurred.
	//
	// A non-nil ctx should control the lifetime of the Adaptor.
	Run(ctx context.Context) error
}

func Run(ctx context.Context, adaptors ...Adaptor) error {
	var (
		err   error
		errMu sync.Mutex
	)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // if error occurred, cancel all goroutines

	for _, adaptor := range adaptors {
		adaptor := adaptor
		go func() {
			defer cancel()
			adaptorErr := adaptor.Run(ctx)
			if adaptorErr != nil {
				errMu.Lock()
				defer errMu.Unlock()
				err = adaptorErr
				return
			}
		}()
	}

	select {
	case <-ctx.Done(): // if len(adaptors) == 0, it will also be blocking here
	}

	return err // capture the first error in the goroutines, and return it
}
