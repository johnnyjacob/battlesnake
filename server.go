package main

import (
	"context"
	"johnnyjacob/battlesnake/board"
	"johnnyjacob/battlesnake/logger"
	"johnnyjacob/battlesnake/service"
	"net/http"
)

type Server struct {
}

func (server *Server) Run(ctx context.Context, log logger.Logger) error {
	err := server.startServer(ctx, log)
	if err != nil {
		log.Error(err.Error())
	}
	return err
}

func (snake *Server) startServer(ctx context.Context, log logger.Logger) error {
	log.Info("starting servers...")

	mux := http.NewServeMux()

	metaService := service.NewMetaService(log)
	mux.HandleFunc("GET /", metaService.HandleMeta)

	board := board.NewBoardGrid()
	gameService := service.NewGameService(log, board)
	mux.HandleFunc("POST /start", gameService.HandleStart)
	mux.HandleFunc("POST /move", gameService.HandleMove)
	mux.HandleFunc("POST /end", gameService.HandleEnd)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	log.Info("ready.")

	return nil
}
