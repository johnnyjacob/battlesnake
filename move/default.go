package move

import (
	"johnnyjacob/battlesnake/board"
	"johnnyjacob/battlesnake/models"
)

type DefaultMove struct {
	Move
}

func (m DefaultMove) Recommend(b *board.BoardGrid) models.Direction {
	return models.MOVE_UP
}

func NewDefaultMove() Move {
	return DefaultMove{}
}
