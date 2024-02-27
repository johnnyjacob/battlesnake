package move

import (
	"johnnyjacob/battlesnake/board"
	"johnnyjacob/battlesnake/models"
	"math/rand"
)

type RandomMove struct {
	Move
}

func (m RandomMove) Recommend(b *board.BoardGrid) models.Direction {
	dirs := []models.Direction{models.MOVE_DOWN, models.MOVE_LEFT, models.MOVE_RIGHT, models.MOVE_UP}
	selection := rand.Intn(4)

	return dirs[selection]
}

func NewRandomMove() Move {
	return RandomMove{}
}
