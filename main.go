package main

import (
	"context"
	"fmt"

	"johnnyjacob/battlesnake/logger"
)

func main() {
	logger, err := logger.NewLogger()
	if err != nil {
		fmt.Println("Unable to Initialize logger.")
		panic("unable to init logger")
	}
	ctx := context.Background()

	server := Server{}
	err = server.Run(ctx, logger)
	if err != nil {
		logger.Error("Error")
	}
}
