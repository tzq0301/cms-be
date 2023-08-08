package shutdownx

import (
	"errors"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/samber/lo"
)

var (
	defaultManager shutdownHookManager
)

var (
	ErrNilLogger = errors.New("param Logger is nil")
)

type HookFn func() error

type shutdownHookManager struct {
	hooks []HookFn
	mu    sync.Mutex

	errLogger func(error)
}

// AddHook hooks will be executed in reverse order of their addition
func (m *shutdownHookManager) addHook(hook HookFn) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.hooks = append(m.hooks, hook)
}

func (m *shutdownHookManager) setErrLogger(errLogger func(error)) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if errLogger == nil {
		return ErrNilLogger
	}

	m.errLogger = errLogger

	return nil
}

func (m *shutdownHookManager) listen() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	<-signalCh

	var errs []error

	for _, hook := range lo.Reverse(m.hooks) {
		err := hook()
		if err != nil {
			errs = append(errs, err)
		}
	}

	if errs != nil {
		errMessages := lo.Map(errs, func(err error, _ int) string {
			return err.Error()
		})
		err := errors.New(strings.Join(errMessages, "; "))
		m.errLogger(err)
	}

	os.Exit(0)
}

func init() {
	go defaultManager.listen()
}

// AddHook hooks will be executed in reverse order of their addition
func AddHook(hook HookFn) {
	defaultManager.addHook(hook)
}

func SetErrLogger(errLogger func(error)) error {
	err := defaultManager.setErrLogger(errLogger)
	if err != nil {
		return errors.Join(err, errors.New("fail to set error logger for package shutdownx"))
	}

	return nil
}
