package robot

import (
	"fmt"

	"github.com/filiadielias/toy-robot-simulator/model"
	"github.com/filiadielias/toy-robot-simulator/service"
)

type robotUsecaseImpl struct {
	board service.BoardUsecase
}

//NewRobotUsecaseImpl new robot usecase implementation
func NewRobotUsecaseImpl(board service.BoardUsecase) service.RobotUsecase {
	return &robotUsecaseImpl{board}
}

func (r *robotUsecaseImpl) New(b model.Board, coor model.Coordinate, face model.Coordinate) (model.Robot, error) {
	//validate coordinate with board size
	if r.board.IsValidCoordinate(b, coor) == false {
		return model.Robot{}, model.ErrCoordinateNotExists
	}

	return model.Robot{Current: coor, Face: face}, nil
}

func (r *robotUsecaseImpl) Move(robot *model.Robot, b model.Board) error {

	newCoor := model.Coordinate{
		X: robot.Current.X + robot.Face.X,
		Y: robot.Current.Y + robot.Face.Y,
	}

	if r.board.IsValidCoordinate(b, newCoor) == false {
		return model.ErrInvalidMove
	}
	robot.Current = newCoor

	return nil
}

func (r *robotUsecaseImpl) FaceLeft(robot *model.Robot) {
	switch robot.Face {
	case model.FaceNorth:
		robot.Face = model.FaceWest
	case model.FaceWest:
		robot.Face = model.FaceSouth
	case model.FaceSouth:
		robot.Face = model.FaceEast
	case model.FaceEast:
		robot.Face = model.FaceNorth
	}
}

func (r *robotUsecaseImpl) FaceRight(robot *model.Robot) {
	switch robot.Face {
	case model.FaceNorth:
		robot.Face = model.FaceEast
	case model.FaceWest:
		robot.Face = model.FaceNorth
	case model.FaceSouth:
		robot.Face = model.FaceWest
	case model.FaceEast:
		robot.Face = model.FaceSouth
	}
}

func (r *robotUsecaseImpl) Report(robot *model.Robot) string {
	faceStr := ""

	switch robot.Face {
	case model.FaceNorth:
		faceStr = model.NORTH
	case model.FaceWest:
		faceStr = model.WEST
	case model.FaceSouth:
		faceStr = model.SOUTH
	case model.FaceEast:
		faceStr = model.EAST
	}

	return fmt.Sprintf("%d,%d,%s", robot.Current.X, robot.Current.Y, faceStr)
}
