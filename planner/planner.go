package planner

import (
	"johnnyjacob/battlesnake/board"
	"johnnyjacob/battlesnake/models"
)

type Planner interface {
	Recommend(b *board.BoardGrid) models.Direction
}
