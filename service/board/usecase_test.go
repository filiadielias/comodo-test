package board

import (
	"fmt"
	"testing"

	"github.com/filiadielias/comodo-test/model"
	"github.com/stretchr/testify/assert"
)

func Test_boardUsecaseImpl_IsValidCoordinate(t *testing.T) {

	b := &boardUsecaseImpl{}
	board := model.Board{
		Min: model.Coordinate{0, 0},
		Max: model.Coordinate{4, 4},
	}

	tests := []struct {
		X, Y     int
		Expected bool
	}{
		{0, 0, true},
		{-1, 0, false},
		{0, -1, false},
		{5, 0, false},
		{0, 5, false},
		{5, 5, false},
		{0, 4, true},
		{4, 4, true},
		{4, 1, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("{%d,%d}", tt.X, tt.Y), func(t *testing.T) {
			result := b.IsValidCoordinate(board, model.Coordinate{tt.X, tt.Y})
			assert.Equal(t, tt.Expected, result)
		})
	}
}

func Test_boardUsecaseImpl_CreateBoard(t *testing.T) {
	b := &boardUsecaseImpl{}

	emptyBoard := model.Board{}

	tests := []struct {
		X, Y          int
		ExpectedBoard model.Board
		ExpectedErr   error
	}{
		{0, 0, emptyBoard, model.ErrInvalidBoardSize},
		{-1, 0, emptyBoard, model.ErrInvalidBoardSize},
		{0, -1, emptyBoard, model.ErrInvalidBoardSize},
		{5, 0, emptyBoard, model.ErrInvalidBoardSize},
		{0, 5, emptyBoard, model.ErrInvalidBoardSize},
		{5, 5, model.Board{Min: model.Coordinate{0, 0}, Max: model.Coordinate{4, 4}}, nil},
		{4, 1, model.Board{Min: model.Coordinate{0, 0}, Max: model.Coordinate{3, 0}}, nil},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("{%d,%d}", tt.X, tt.Y), func(t *testing.T) {
			board, err := b.CreateBoard(tt.X, tt.Y)

			assert.Equal(t, tt.ExpectedErr, err)
			assert.Equal(t, tt.ExpectedBoard, board)
		})
	}
}
