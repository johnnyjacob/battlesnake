package board

import (
	"johnnyjacob/battlesnake/models"
)

const (
	EMPTY        = 0
	FOOD         = 1
	HAZARD       = 2
	SNAKE        = 3
	OTHER_SNAKES = 4
)

type BoardGrid struct {
	Grid [][]byte
	Size int
}

func NewBoardGrid() *BoardGrid {
	return &BoardGrid{}
}

func (b *BoardGrid) SetSize(size int) {
	b.Size = size
	b.Grid = make([][]byte, size)
	for i := 0; i < size; i++ {
		b.Grid[i] = make([]byte, size)
	}
}

func (bg *BoardGrid) SetState(board models.Board) {
	for _, coord := range board.Food {
		bg.Grid[coord.X][coord.Y] = FOOD
	}

	for _, coord := range board.Hazards {
		bg.Grid[coord.X][coord.Y] = HAZARD
	}

	for i, snake := range board.Snakes {
		for _, coord := range snake.Body {
			bg.Grid[coord.X][coord.Y] = OTHER_SNAKES + byte(i)
		}
	}
}
