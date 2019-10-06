package board

import (
	"github.com/filiadielias/comodo-test/model"
	"github.com/filiadielias/comodo-test/service"
)

//contract implementation
type (
	boardUsecaseImpl struct{}
)

//NewBoardUsecaseImpl create new usecase implementation
func NewBoardUsecaseImpl() service.BoardUsecase {
	return &boardUsecaseImpl{}
}

func (b *boardUsecaseImpl) CreateBoard(width, height int) (model.Board, error) {
	//coordinate is zero-based, decrease value by one
	width--
	height--

	if width < 0 || height < 0 {
		return model.Board{}, model.ErrInvalidBoardSize
	}

	board := model.Board{
		Min: model.Coordinate{0, 0},
		Max: model.Coordinate{width, height},
	}

	return board, nil
}

func (b *boardUsecaseImpl) IsValidCoordinate(board model.Board, coor model.Coordinate) bool {
	switch {
	case coor.X < board.Min.X:
		return false
	case coor.X > board.Max.X:
		return false
	case coor.Y < board.Min.Y:
		return false
	case coor.Y > board.Max.Y:
		return false
	}
	return true
}
