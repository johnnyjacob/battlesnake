package planner

import (
	"johnnyjacob/battlesnake/board"
	"johnnyjacob/battlesnake/models"
)

type DefaultPlanner struct {
	Planner
}

func (m DefaultPlanner) Recommend(b *board.BoardGrid, snake *models.Snake) models.Direction {
	return models.MOVE_UP
}

func NewDefaultPlanner() Planner {
	return DefaultPlanner{}
}
