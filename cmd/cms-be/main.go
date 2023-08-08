package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"cms-be/internal/infrastructure/config"
	"cms-be/internal/pkg/observability/logx"
	"cms-be/internal/pkg/runtimex/shutdownx"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	c, err := config.Load()
	if err != nil {
		return err
	}

	{
		// TODO(TZQ) delete

		fmt.Printf("%+v\n", c)
		fmt.Printf("%+v\n", *c.Log.Console)
		fmt.Println()
	}

	logger, err := initLogger(c.Log)
	if err != nil {
		return errors.Join(err, errors.New("fail to init logger"))
	}

	{
		// TODO(TZQ) delete

		ctx := logx.ContextWithLogger(context.Background(), logger)
		logx.Debug(ctx, "test", slog.String("hello", "world"))
		logx.Info(ctx, "test", slog.String("hello", "world"))
		logx.Warn(ctx, "test", slog.String("hello", "world"))
		logx.Error(ctx, "test", slog.String("hello", "world"))
	}

	err = shutdownx.SetErrLogger(func(err error) {
		ctx := logx.ContextWithLogger(context.TODO(), logger)
		logx.Error(ctx, err.Error())
	})
	if err != nil {
		return errors.Join(err, errors.New("fail to set logger for shutdownx"))
	}

	return nil
}

func initLogger(c config.Log) (logx.Logger, error) {
	logConfig := logx.Config{}

	if c.Console != nil {
		consoleLogConfig := c.Console
		level, err := logx.LevelFromString(consoleLogConfig.Level)
		if err != nil {
			return nil, errors.Join(err, fmt.Errorf("fail to convert the level field of console: level=%s", consoleLogConfig.Level))
		}

		logConfig.ConsoleAppenderConfig = &logx.ConsoleAppenderConfig{
			CommonAppenderConfig: logx.CommonAppenderConfig{
				Level: level,
			},
		}
	}

	for _, fileLogConfig := range c.File {
		level, err := logx.LevelFromString(fileLogConfig.Level)
		if err != nil {
			return nil, errors.Join(err, fmt.Errorf("fail to convert the level field of file: level=%s, filepath=%s", fileLogConfig.Level, fileLogConfig.FilePath))
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
		return nil, errors.Join(err, errors.New("fail to init logger"))
	}

	return logger, nil
}
