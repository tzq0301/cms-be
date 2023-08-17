package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"

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
	c, err := config.Load()
	if err != nil {
		return errors.Join(err, errors.New("load config"))
	}

	{
		// TODO(TZQ) delete

		fmt.Printf("%+v\n", c)
		fmt.Printf("%+v\n", *c.Log.Console)
		fmt.Println()
	}

	runtimeEnvironment, err := runtimex.Load()
	if err != nil {
		return errors.Join(err, errors.New("load runtime environment"))
	}

	logger, err := initLogger(c, runtimeEnvironment)
	if err != nil {
		return errors.Join(err, errors.New("init logger"))
	}

	{
		// TODO(TZQ) delete

		ctx := logx.ContextWithLogger(context.Background(), logger)
		logx.Debug(ctx, "test", slog.String("hello", "world"))
		logx.Info(ctx, "test", slog.String("hello", "world"))
		logx.Warn(ctx, "test", slog.String("hello", "world"))
		logx.Error(ctx, "test", slog.String("hello", "world"))

		var wg sync.WaitGroup

		for i := 0; i < 20; i++ {
			logx.Info(ctx, "concurrent log")
		}

		wg.Wait()
	}

	errLogger := func(err error) {
		ctx := logx.ContextWithLogger(context.TODO(), logger)
		logx.Error(ctx, err.Error())
	}

	err = async.SetErrLogger(errLogger)
	if err != nil {
		return errors.Join(err, errors.New("set logger for async"))
	}

	{
		// TODO(TZQ) delete

		async.Go(func() {
			panic("test panic & recover")
		})
	}

	err = shutdown.SetErrLogger(errLogger)
	if err != nil {
		return errors.Join(err, errors.New("set logger for shutdown"))
	}

	return nil
}

func initLogger(c config.Config, re runtimex.RuntimeEnvironment) (logx.Logger, error) {
	logConfig := logx.Config{}

	serviceConfig := logx.ServiceConfig{
		Name: c.Service.Name,
		IP: logx.IPConfig{
			V4: re.IP.V4,
			V6: re.IP.V6,
		},
	}

	if c.Log.Console != nil {
		consoleLogConfig := c.Log.Console
		level, err := logx.LevelFromString(consoleLogConfig.Level)
		if err != nil {
			return nil, errors.Join(err, fmt.Errorf("convert the level field of console: level=%s", consoleLogConfig.Level))
		}

		logConfig.ConsoleAppenderConfig = &logx.ConsoleAppenderConfig{
			CommonAppenderConfig: logx.CommonAppenderConfig{
				Level:         level,
				ServiceConfig: serviceConfig,
			},
		}
	}

	for _, fileLogConfig := range c.Log.File {
		level, err := logx.LevelFromString(fileLogConfig.Level)
		if err != nil {
			return nil, errors.Join(err, fmt.Errorf("convert the level field of file: level=%s, filepath=%s", fileLogConfig.Level, fileLogConfig.FilePath))
		}

		logConfig.FileAppenderConfigs = append(logConfig.FileAppenderConfigs, logx.FileAppenderConfig{
			CommonAppenderConfig: logx.CommonAppenderConfig{
				Level:         level,
				ServiceConfig: serviceConfig,
			},
			FilePath: fileLogConfig.FilePath,
		})
	}

	logger, err := logx.Init(logConfig)
	if err != nil {
		return nil, errors.Join(err, errors.New("init logger"))
	}

	return logger, nil
}
