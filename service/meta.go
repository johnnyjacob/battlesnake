package service

import (
	"encoding/json"
	"johnnyjacob/battlesnake/logger"
	"johnnyjacob/battlesnake/models"
	"johnnyjacob/battlesnake/version"
	"net/http"
)

type MetaHandlers interface {
	HandleMeta(w http.ResponseWriter, req *http.Request)
}
type MetaService struct {
	Log logger.Logger
	MetaHandlers
}

func NewMetaService(l logger.Logger) MetaHandlers {
	return &MetaService{Log: l}
}

func (service *MetaService) HandleMeta(w http.ResponseWriter, req *http.Request) {
	//FIXME if it is not / then return 404
	service.Log.Info("Handling meta request")
	meta := getSnakeMeta()
	j, err := json.Marshal(meta)
	if err != nil {
		service.Log.Error("Fatal : Unable to marshal meta for response.")
		return
	}

	service.Log.Info("Sending : " + string(j))

	json.NewEncoder(w).Encode(meta)
}

func getSnakeMeta() *models.SnakeMeta {

	return &models.SnakeMeta{
		ApiVersion: "1",
		Author:     "johnnyjacob",
		Color:      "#FF00FF",
		Head:       "default",
		Tail:       "default",
		Version:    version.Version,
	}
}
