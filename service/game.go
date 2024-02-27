package service

import (
	"encoding/json"
	"fmt"
	"io"
	"johnnyjacob/battlesnake/board"
	"johnnyjacob/battlesnake/logger"
	"johnnyjacob/battlesnake/models"
	"net/http"
)

type GameHandlers interface {
	HandleStart(w http.ResponseWriter, req *http.Request)
	HandleMove(w http.ResponseWriter, req *http.Request)
	HandleEnd(w http.ResponseWriter, req *http.Request)
}

type GameService struct {
	Log   logger.Logger
	Board *board.BoardGrid
	GameHandlers
}

func NewGameService(l logger.Logger, b *board.BoardGrid) GameHandlers {
	return &GameService{Log: l, Board: b}
}

func (service *GameService) HandleStart(w http.ResponseWriter, req *http.Request) {
	service.Log.Info("Start request")
}

func (service *GameService) HandleMove(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal the JSON into a User struct
	var board1 models.Board
	err = json.Unmarshal([]byte(body), &board1)
	if err != nil {
		fmt.Println("unmarshall error", err)
	}
	service.Log.Info(string(body))

	//TODO : Calculate next move

	// TODO : Send response

}

func (service *GameService) HandleEnd(w http.ResponseWriter, req *http.Request) {
	service.Log.Info("End request")
}
