package service

import (
	"github.com/filiadielias/comodo-test/model"
)

//interface contract
type (
	BoardUsecase interface {
		CreateBoard(int, int) (model.Board, error)
		IsValidCoordinate(model.Board, model.Coordinate) bool
	}
	RobotUsecase interface {
		New(model.Board, model.Coordinate, model.Coordinate) (model.Robot, error)
		Move(*model.Robot, model.Board) error
		FaceLeft(*model.Robot)
		FaceRight(*model.Robot)

		Report(*model.Robot) string
	}

	CommandUsecase interface {
		Run([]string, model.Board) ([]string, error)
	}
)
