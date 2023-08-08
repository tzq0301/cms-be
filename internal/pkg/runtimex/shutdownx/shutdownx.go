package shutdownx

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/samber/lo"
)

type HookFn func() error

var hooks []HookFn
var mu sync.Mutex

func init() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalCh

		var errs error

		for _, hook := range lo.Reverse(hooks) {
			err := hook()
			if err != nil {
				errs = errors.Join(errs, err)
			}
		}

		fmt.Printf("\n%s\n", errs.Error())

		os.Exit(0)
	}()
}

// AddHook hooks will be executed in reverse order of their addition
func AddHook(hook HookFn) {
	mu.Lock()
	defer mu.Unlock()

	hooks = append(hooks, hook)
}
