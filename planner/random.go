package planner

import (
	"johnnyjacob/battlesnake/board"
	"johnnyjacob/battlesnake/models"
	"math/rand"
)

type RandomPlanner struct {
	Planner
}

func (m RandomPlanner) Recommend(b *board.BoardGrid, snake *models.Snake) models.Direction {
	dirs := []models.Direction{models.MOVE_DOWN, models.MOVE_LEFT, models.MOVE_RIGHT, models.MOVE_UP}
	selection := rand.Intn(4)

	return dirs[selection]
}

func NewRandomPlanner() Planner {
	return RandomPlanner{}
}
