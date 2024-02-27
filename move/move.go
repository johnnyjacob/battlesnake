package move

import (
	"johnnyjacob/battlesnake/board"
	"johnnyjacob/battlesnake/models"
)

type Move interface {
	Recommend(b *board.BoardGrid) models.Direction
}
