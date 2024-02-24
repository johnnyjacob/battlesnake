package service

import (
	"encoding/json"
	"johnnyjacob/battlesnake/models"
	"johnnyjacob/battlesnake/version"
	"net/http"
)

type MetaHandlers interface {
	HandleMeta(w http.ResponseWriter, req *http.Request)
}
type MetaService struct {
	// Logger Instance
	MetaHandlers
}

func NewMetaService() MetaHandlers {
	return &MetaService{}
}

func (service *MetaService) HandleMeta(w http.ResponseWriter, req *http.Request) {
	//FIXME if it is not / then return 404
	meta := GetSnakeMeta()
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
