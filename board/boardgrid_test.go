package board

import (
	"encoding/json"
	"fmt"
	"johnnyjacob/battlesnake/models"
	"testing"
)

func TestBoardGrid_SetState1(t *testing.T) {

}

func TestBoardGrid_SetState(t *testing.T) {
	type fields struct {
		grid [][]byte
	}
	type args struct {
		board models.Board
	}

	request1 := `{
		"height": 11,
		"width": 11,
		"food": [
		  {"x": 5, "y": 5},
		  {"x": 9, "y": 0},
		  {"x": 2, "y": 6}
		],
		"hazards": [
		  {"x": 3, "y": 2}
		],
		"snakes": [
		  {
			"id": "snake-508e96ac-94ad-11ea-bb37",
			"name": "My Snake",
			"health": 54,
			"body": [
			  {"x": 0, "y": 0},
			  {"x": 1, "y": 0},
			  {"x": 2, "y": 0}
			],
			"latency": "111",
			"head": {"x": 0, "y": 0},
			"length": 3,
			"shout": "why are we shouting??",
			"customizations":{
			  "color":"#FF0000",
			  "head":"pixel",
			  "tail":"pixel"
			}
		  },
		  {
			"id": "snake-b67f4906-94ae-11ea-bb37",
			"name": "Another Snake",
			"health": 16,
			"body": [
			  {"x": 5, "y": 4},
			  {"x": 5, "y": 3},
			  {"x": 6, "y": 3},
			  {"x": 6, "y": 2}
			],
			"latency": "222",
			"head": {"x": 5, "y": 4},
			"length": 4,
			"shout": "I'm not really sure...",
			"customizations":{
			  "color":"#26CF04",
			  "head":"silly",
			  "tail":"curled"
			}
		  }
		]
	  }
`
	var board1 models.Board
	err := json.Unmarshal([]byte(request1), &board1)
	if err != nil {
		fmt.Println("unmarshall error", err)
	}
	bg1 := NewBoardGrid()
	bg1.SetSize(board1.Height)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{name: "Good Json", args: args{board: board1}, fields: fields{grid: bg1.grid}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bg := &BoardGrid{
				grid: tt.fields.grid,
			}
			bg.SetState(tt.args.board)
			for _, row := range bg.grid {
				fmt.Println(row)
			}
		})
	}

}
