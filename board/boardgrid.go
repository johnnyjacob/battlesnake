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
	grid [][]byte
}

func NewBoardGrid() *BoardGrid {
	return &BoardGrid{}
}

func (b *BoardGrid) SetSize(size int) {
	b.grid = make([][]byte, size)
	for i := 0; i < size; i++ {
		b.grid[i] = make([]byte, size)
	}
}

func (bg *BoardGrid) SetState(board models.Board) {
	for _, coord := range board.Food {
		bg.grid[coord.X][coord.Y] = FOOD
	}

	for _, coord := range board.Hazards {
		bg.grid[coord.X][coord.Y] = HAZARD
	}

	for i, snake := range board.Snakes {
		for _, coord := range snake.Body {
			bg.grid[coord.X][coord.Y] = OTHER_SNAKES + byte(i)
		}
	}
}
