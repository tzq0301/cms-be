package ginadaptor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"cms-be/internal/adaptor"
)

type Adaptor struct {
	httpServer *http.Server
}

type Options struct {
	middlewares []gin.HandlerFunc
	routers     map[string]RouterFunc
}

type Option func(o *Options)

func WithMiddlewares(middlewares ...gin.HandlerFunc) Option {
	return func(o *Options) {
		o.middlewares = append(o.middlewares, middlewares...)
	}
}

func WithRouter(version string, router RouterFunc) Option {
	return func(o *Options) {
		o.routers[version] = router
	}
}

func New(_ context.Context, port int, urlPrefix string, options ...Option) adaptor.Adaptor {
	gin.SetMode(gin.ReleaseMode)

	o := Options{
		routers: make(map[string]RouterFunc),
	}

	for _, option := range options {
		option(&o)
	}

	engine := gin.New()

	routerGroup := engine.Group(urlPrefix, o.middlewares...)

	for version, routerFunc := range o.routers {
		group := routerGroup.Group(version)
		routerFunc(group)
	}

	return &Adaptor{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: engine,
		},
	}
}

func (a *Adaptor) Run(ctx context.Context) error {
	select {
	case err := <-a.run():
		return err
	case <-ctx.Done():
		a.shutdown(ctx)
		return nil
	}
}

func (a *Adaptor) run() <-chan error {
	ch := make(chan error, 1)

	go func() {
		defer close(ch)
		ch <- a.httpServer.ListenAndServe()
	}()

	return ch
}

func (a *Adaptor) shutdown(ctx context.Context) {
	_ = a.httpServer.Shutdown(ctx)
}
