package main

import (
	"fmt"

	"internal/infrastructure/config"
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

	fmt.Printf("%+v\n", c)

	return nil
}
