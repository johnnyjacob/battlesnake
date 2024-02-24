package main

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("Unable to Initialize logger.")
		panic("unable to init logger")
	}
	ctx := context.Background()

	snake := Snake{}
	err = snake.Run(ctx, logger)
	if err != nil {
		logger.Error("Error")
	}
}
