package main

import (
	"context"
	"encoding/json"
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

func (snake *Snake) handleSnakeMeta(w http.ResponseWriter, req *http.Request) {
	//FIXME if it is not / then return 404
	meta := GetSnakeMeta()
	json.NewEncoder(w).Encode(meta)
}

func (snake *Snake) startServer(ctx context.Context, logger *zap.Logger) error {
	logger.Info("starting servers...")

	mux := http.NewServeMux()
	mux.HandleFunc("POST /start", snake.handleStart)
	mux.HandleFunc("GET /", snake.handleSnakeMeta)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Info("ready.")

	return nil
}
