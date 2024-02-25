package service

import (
	"johnnyjacob/battlesnake/logger"
	"net/http"
)

type GameHandlers interface {
	HandleStart(w http.ResponseWriter, req *http.Request)
	HandleMove(w http.ResponseWriter, req *http.Request)
}
type GameService struct {
	Log logger.Logger
	GameHandlers
}

func NewGameService(l logger.Logger) GameHandlers {
	return &GameService{Log: l}
}

func (service *GameService) HandleStart(w http.ResponseWriter, req *http.Request) {

}

func (service *GameService) HandleMove(w http.ResponseWriter, req *http.Request) {

}
