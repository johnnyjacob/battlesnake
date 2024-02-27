package service

import (
	"encoding/json"
	"fmt"
	"io"
	"johnnyjacob/battlesnake/board"
	"johnnyjacob/battlesnake/logger"
	"johnnyjacob/battlesnake/models"
	"johnnyjacob/battlesnake/planner"
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
	service.Log.Info("Handling Move...")
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
	service.Log.Debug(string(body))
	// Create a new grid representation
	b := board.NewBoardGrid()
	b.SetSize(board1.Height)

	// Recommend move
	planner := planner.NewRandomPlanner()
	dir := planner.Recommend(b)
	service.Log.Info(string(dir))

	move := models.MoveResponse{Move: dir}
	json.NewEncoder(w).Encode(move)
}

func (service *GameService) HandleEnd(w http.ResponseWriter, req *http.Request) {
	service.Log.Info("End request")
}
