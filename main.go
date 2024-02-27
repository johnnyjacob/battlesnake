package main

import (
	"context"
	"fmt"

	"johnnyjacob/battlesnake/logger"
	"johnnyjacob/battlesnake/version"
)

func main() {
	logger, err := logger.NewLogger()
	if err != nil {
		fmt.Println("Unable to Initialize logger.")
		panic("unable to init logger")
	}
	logger.Info(version.Version)
	ctx := context.Background()

	server := Server{}
	err = server.Run(ctx, logger)
	if err != nil {
		logger.Error("Error")
	}
}
