package planner

import (
	"johnnyjacob/battlesnake/models"
)

type Planner interface {
	Recommend(m *models.MoveRequest) models.Direction
}
