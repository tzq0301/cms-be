package main

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"cms-be/internal/adaptor"
	ginadaptor "cms-be/internal/adaptor/gin"
	"cms-be/internal/infrastructure/config"
	"cms-be/internal/pkg/async"
	"cms-be/internal/pkg/observability/logx"
	"cms-be/internal/pkg/runtimex"
	"cms-be/internal/pkg/runtimex/shutdown"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	ctx := context.Background()

	c, err := config.Load()
	if err != nil {
		return errors.Wrap(err, "load config")
	}

	{
		// TODO(TZQ) delete

		fmt.Printf("%+v\n", c)
		fmt.Printf("%+v\n", *c.Log.Console)
		fmt.Println()
	}

	runtimeEnvironment, err := runtimex.Load()
	if err != nil {
		return errors.Wrap(err, "load runtime environment")
	}

	logger, err := initLogger(c, runtimeEnvironment)
	if err != nil {
		return errors.Wrap(err, "init logger")
	}

	{
		// TODO(TZQ) delete

		ctx = logx.ContextWithLogger(ctx, logger)

		logx.Debug(ctx, "test", slog.String("hello", "world"))
		logx.Info(ctx, "test", slog.String("hello", "world"))
		logx.Warn(ctx, "test", slog.String("hello", "world"))
		logx.Error(ctx, "test", slog.String("hello", "world"))

		{
			var wg sync.WaitGroup
			wg.Add(20)

			for i := 0; i < 20; i++ {
				logx.Info(ctx, "concurrent log")
				wg.Done()
			}

			wg.Wait()
		}

		{
			enhancedLogger := logger.With(slog.String("enhance", "yes"))
			enhancedCtx := logx.ContextWithLogger(context.Background(), enhancedLogger)
			logx.Info(enhancedCtx, "enhance message", slog.String("one", "1"))
		}

		logx.Info(ctx, "no enhance message", slog.String("one", "1"))
	}

	errLogger := func(err error) {
		ctx := logx.ContextWithLogger(context.TODO(), logger)
		logx.Error(ctx, err.Error())
	}

	err = async.SetErrLogger(errLogger)
	if err != nil {
		return errors.Wrap(err, "set logger for async")
	}

	err = shutdown.SetErrLogger(errLogger)
	if err != nil {
		return errors.Wrap(err, "set logger for shutdown")
	}

	err = initAdaptor(ctx, c.Adaptor, logger)
	if err != nil {
		return errors.Wrap(err, "init adaptor")
	}

	return nil
}

func initLogger(c config.Config, re runtimex.RuntimeEnvironment) (*logx.Logger, error) {
	logConfig := logx.Config{
		ServiceConfig: logx.ServiceConfig{
			Name: c.Service.Name,
			IP: logx.IPConfig{
				V4: re.IP.V4,
				V6: re.IP.V6,
			},
		},
	}

	if c.Log.Console != nil {
		consoleLogConfig := c.Log.Console
		level, err := logx.LevelFromString(consoleLogConfig.Level)
		if err != nil {
			return nil, errors.Wrapf(err, "convert the level field of console: level=%s", consoleLogConfig.Level)
		}

		logConfig.ConsoleAppenderConfig = &logx.ConsoleAppenderConfig{
			CommonAppenderConfig: logx.CommonAppenderConfig{
				Level: level,
			},
		}
	}

	for _, fileLogConfig := range c.Log.File {
		level, err := logx.LevelFromString(fileLogConfig.Level)
		if err != nil {
			return nil, errors.Wrapf(err, "convert the level field of file: level=%s, filepath=%s", fileLogConfig.Level, fileLogConfig.FilePath)
		}

		logConfig.FileAppenderConfigs = append(logConfig.FileAppenderConfigs, logx.FileAppenderConfig{
			CommonAppenderConfig: logx.CommonAppenderConfig{
				Level: level,
			},
			FilePath: fileLogConfig.FilePath,
		})
	}

	logger, err := logx.Init(logConfig)
	if err != nil {
		return nil, errors.Wrap(err, "init logger")
	}

	return logger, nil
}

func initAdaptor(ctx context.Context, c config.Adaptor, logger *logx.Logger) error {
	v1 := ginadaptor.V1()

	ginAdaptor := ginadaptor.New(ctx, c.Gin.Port, c.Gin.UrlPrefix,
		ginadaptor.WithRouter("v1", v1),
		ginadaptor.WithMiddlewares(
			gin.Recovery(),
			ginadaptor.InjectLogx(logger),
			ginadaptor.Log()))

	err := adaptor.Run(ctx, ginAdaptor)
	if err != nil {
		return errors.Wrap(err, "run adaptors")
	}

	return nil
}
