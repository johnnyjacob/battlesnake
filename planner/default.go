package planner

import (
	"johnnyjacob/battlesnake/board"
	"johnnyjacob/battlesnake/models"
)

type DefaultPlanner struct {
	Planner
}

func (m DefaultPlanner) defaultMove(b *board.BoardGrid, snake *models.Snake) models.Direction {
	return models.MOVE_UP
}

func (m DefaultPlanner) Recommend(move *models.MoveRequest) models.Direction {
	// Create a new grid representation
	b := board.NewBoardGrid()
	b.SetSize(move.Board.Height)
	b.SetState(move.Board)

	return m.defaultMove(b, &move.You)
}

func NewDefaultPlanner() Planner {
	return DefaultPlanner{}
}
