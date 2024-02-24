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
	meta := GetSnakeMeta()
	j, err := json.Marshal(meta)
	if err != nil {
		//Fixme
		return
	}

	service.Log.Info(string(j))

	json.NewEncoder(w).Encode(meta)
}

func GetSnakeMeta() *models.SnakeMeta {

	return &models.SnakeMeta{
		ApiVersion: 0,
		Author:     "johnnyjacob",
		Color:      "#FF00FF",
		Head:       "default",
		Tail:       "default",
		Version:    version.Version,
	}
}
