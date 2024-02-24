package main

import (
	"context"
	"johnnyjacob/battlesnake/logger"
	"johnnyjacob/battlesnake/service"
	"net/http"
)

type Snake struct {
}

func (snake *Snake) Run(ctx context.Context, log logger.Logger) error {
	err := snake.startServer(ctx, log)
	if err != nil {
		log.Error(err.Error())
	}
	return err
}

func (snake *Snake) handleStart(w http.ResponseWriter, req *http.Request) {

}

func (snake *Snake) startServer(ctx context.Context, log logger.Logger) error {
	log.Info("starting servers...")

	mux := http.NewServeMux()

	metaService := service.NewMetaService(log)
	mux.HandleFunc("GET /", metaService.HandleMeta)

	mux.HandleFunc("POST /start", snake.handleStart)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	log.Info("ready.")

	return nil
}
