package main

import (
	"context"
	"johnnyjacob/battlesnake/service"
	"net/http"

	"go.uber.org/zap"
)

type Snake struct {
}

func (snake *Snake) Run(ctx context.Context, logger *zap.Logger) error {
	err := snake.startServer(ctx, logger)
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}

func (snake *Snake) handleStart(w http.ResponseWriter, req *http.Request) {

}

func (snake *Snake) startServer(ctx context.Context, logger *zap.Logger) error {
	logger.Info("starting servers...")

	mux := http.NewServeMux()

	metaService := service.NewMetaService()
	mux.HandleFunc("GET /", metaService.HandleMeta)

	mux.HandleFunc("POST /start", snake.handleStart)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Info("ready.")

	return nil
}
