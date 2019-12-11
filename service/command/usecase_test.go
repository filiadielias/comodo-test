package command

import (
	"testing"

	"github.com/filiadielias/toy-robot-simulator/model"
	"github.com/filiadielias/toy-robot-simulator/service/board"
	"github.com/filiadielias/toy-robot-simulator/service/robot"
	"github.com/stretchr/testify/assert"
)

// func Test_commandUsecaseImpl_New(t *testing.T) {
// 	b := board.NewBoardUsecaseImpl()
// 	r := robot.NewRobotUsecaseImpl(b)
// 	c := commandUsecaseImpl{r}

// 	tests := []struct {
// 		Command     string
// 		ExpectedErr error
// 	}{
// 		{"PLACE 0,1,NORTH", nil},
// 		{"PLACE 0,1,NORTH MOVE LEFT LEFT MOVE REPORT", nil},
// 		{"PLACES 0,1,NORTH", model.ErrInvalidCommand},
// 		{"PLACE MOVE", model.ErrInvalidCommand},
// 		{"1,2,NORTH PLACE", model.ErrInvalidCommand},
// 		{"PLACE 1,2,NORTHS", model.ErrInvalidCommand},
// 		{"REPORT PLACE 1,2,NORTH", model.ErrInvalidCommand},
// 		{"", model.ErrInvalidCommand},
// 		{"PLACE", model.ErrInvalidCommand},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.Command, func(t *testing.T) {
// 			_, err := c.New(tt.Command)

// 			assert.Equal(t, tt.ExpectedErr, err)
// 		})
// 	}
// }

func Test_commandUsecaseImpl_Run(t *testing.T) {
	b := board.NewBoardUsecaseImpl()
	r := robot.NewRobotUsecaseImpl(b)
	c := NewCommandUsecaseImpl(r)

	board, _ := b.CreateBoard(5, 5)
	var empty []string

	tests := []struct {
		Name           string
		Command        []string
		ExpectedResult []string
		ExpectedErr    error
	}{
		{"Valid", []string{"PLACE", "1,2,NORTH"}, empty, nil},
		{"Valid REPORT", []string{"PLACE", "1,2,NORTH", "REPORT"}, []string{"1,2,NORTH"}, nil},
		{"Invalid PLACE", []string{"PLACES", "1,2,NORTH", "REPORT"}, empty, nil},
		{"PLACE in middle", []string{"MOVE", "MOVE", "REPORT", "LEFT", "PLACE", "1,2,NORTH"}, empty, nil},
		{"PLACE in middle REPORT", []string{"MOVE", "MOVE", "REPORT", "LEFT", "PLACE", "1,2,NORTH", "REPORT"}, []string{"1,2,NORTH"}, nil},
		{"Out of board move", []string{"PLACE", "0,0,NORTH", "LEFT", "MOVE", "REPORT"}, []string{"0,0,WEST"}, nil},
		{"Double PLACE", []string{"PLACE", "0,0,NORTH", "LEFT", "MOVE", "PLACE", "4,4,SOUTH", "REPORT"}, []string{"4,4,SOUTH"}, nil},
		{"Double PLACE REPORT", []string{"PLACE", "0,0,NORTH", "LEFT", "MOVE", "REPORT", "PLACE", "4,4,SOUTH", "REPORT"}, []string{"0,0,WEST", "4,4,SOUTH"}, nil},
		{"Double PLACE with invalid", []string{"PLACE", "0,0,NORTH", "LEFT", "MOVE", "REPORT", "PLACES", "4,4,SOUTH", "REPORT"}, []string{"0,0,WEST", "0,0,WEST"}, nil},
		{"Double PLACE with invalid begin", []string{"PLACES", "0,0,NORTH", "LEFT", "MOVE", "REPORT", "PLACE", "4,4,SOUTH", "REPORT"}, []string{"4,4,SOUTH"}, nil},
		{"Invalid PLACE param", []string{"PLACE", "0,0,NORTHS", "REPORT"}, empty, nil},
		{"Invalid PLACE param begin", []string{"PLACE", "0,0,NORTHS", "REPORT", "PLACE", "0,0,SOUTH", "LEFT", "MOVE", "REPORT"}, []string{"1,0,EAST"}, nil},
		{"Empty command", empty, empty, model.ErrInvalidCommand},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := c.Run(tt.Command, board)

			assert.Equal(t, tt.ExpectedErr, err)
			assert.Equal(t, tt.ExpectedResult, result)
		})
	}
}
