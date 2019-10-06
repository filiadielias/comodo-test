package service

import (
	"github.com/filiadielias/comodo-test/model"
)

//interface contract
type (
	CoordinateUsecase interface {
		Add(model.Coordinate) model.Coordinate
	}
	BoardUsecase interface {
		CreateBoard(int, int) (model.Board, error)
		IsValidCoordinate(model.Board, model.Coordinate) bool
	}
	RobotUsecase interface {
		New(model.Board, model.Coordinate) (model.Robot, error)
		Move(*model.Robot, model.Board) error
		FaceLeft(*model.Robot) error
		FaceRight(*model.Robot) error

		Report(*model.Robot) string
	}
)
