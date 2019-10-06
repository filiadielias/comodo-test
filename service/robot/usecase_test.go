package robot

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/filiadielias/comodo-test/model"
	"github.com/filiadielias/comodo-test/service/board"
	"github.com/stretchr/testify/assert"
)

func Test_robotUsecaseImpl_New(t *testing.T) {
	b := board.NewBoardUsecaseImpl()
	r := &robotUsecaseImpl{b}

	board5x5 := model.Board{
		Min: model.Coordinate{0, 0},
		Max: model.Coordinate{4, 4},
	}

	robotEmpty := model.Robot{}
	robot0x0 := model.Robot{model.Coordinate{0, 0}, model.FaceNorth}
	robot4x0 := model.Robot{model.Coordinate{4, 0}, model.FaceNorth}
	robot0x4 := model.Robot{model.Coordinate{0, 4}, model.FaceNorth}
	robot3x3 := model.Robot{model.Coordinate{3, 3}, model.FaceNorth}
	robot3x3South := model.Robot{model.Coordinate{3, 3}, model.FaceSouth}

	tests := []struct {
		Name          string
		X, Y          int
		Face          model.Coordinate
		ExpectedRobot model.Robot
		ExpectedErr   error
	}{
		{"valid", 0, 0, model.FaceNorth, robot0x0, nil},
		{"invalid coordinate", -1, 0, model.FaceNorth, robotEmpty, model.ErrCoordinateNotExists},
		{"invalid coordinate", 0, -1, model.FaceNorth, robotEmpty, model.ErrCoordinateNotExists},
		{"invalid coordinate", 5, 0, model.FaceNorth, robotEmpty, model.ErrCoordinateNotExists},
		{"invalid coordinate", 0, 5, model.FaceNorth, robotEmpty, model.ErrCoordinateNotExists},
		{"valid", 4, 0, model.FaceNorth, robot4x0, nil},
		{"valid", 0, 4, model.FaceNorth, robot0x4, nil},
		{"valid", 3, 3, model.FaceNorth, robot3x3, nil},
		{"invalid coordinate", 6, 6, model.FaceNorth, robotEmpty, model.ErrCoordinateNotExists},
		{"invalid face", 3, 3, model.FaceNorth, robot3x3, nil},
		{"face south", 3, 3, model.FaceSouth, robot3x3South, nil},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s {%d,%d}", tt.Name, tt.X, tt.Y), func(t *testing.T) {
			robot, err := r.New(board5x5, model.Coordinate{tt.X, tt.Y}, tt.Face)

			assert.Equal(t, tt.ExpectedErr, err)
			assert.Equal(t, tt.ExpectedRobot, robot)
		})
	}
}

func Test_robotUsecaseImpl_Move(t *testing.T) {
	b := board.NewBoardUsecaseImpl()
	r := &robotUsecaseImpl{b}

	board5x5 := model.Board{
		Min: model.Coordinate{0, 0},
		Max: model.Coordinate{4, 4},
	}

	coor0x0 := model.Coordinate{0, 0}
	coor4x4 := model.Coordinate{4, 4}
	coor3x4 := model.Coordinate{3, 4}
	coor3x3 := model.Coordinate{3, 3}

	coor1x0 := model.Coordinate{1, 0}
	coor0x1 := model.Coordinate{0, 1}
	coor4x3 := model.Coordinate{4, 3}
	coor2x3 := model.Coordinate{2, 3}
	coor3x2 := model.Coordinate{3, 2}

	tests := []struct {
		Name         string
		Current      model.Coordinate
		Face         model.Coordinate
		ExpectedCoor model.Coordinate
		ExpectedErr  error
	}{
		{"valid move north", coor0x0, model.FaceNorth, coor0x1, nil},
		{"valid move east", coor0x0, model.FaceEast, coor1x0, nil},
		{"invalid move south", coor0x0, model.FaceSouth, coor0x0, model.ErrInvalidMove},
		{"invalid move west", coor0x0, model.FaceWest, coor0x0, model.ErrInvalidMove},
		{"invalid move north", coor4x4, model.FaceNorth, coor4x4, model.ErrInvalidMove},
		{"invalid move east", coor4x4, model.FaceEast, coor4x4, model.ErrInvalidMove},
		{"valid move west", coor4x4, model.FaceWest, coor3x4, nil},
		{"valid move south", coor4x4, model.FaceSouth, coor4x3, nil},
		{"invalid move north", coor3x4, model.FaceNorth, coor3x4, model.ErrInvalidMove},
		{"valid move east", coor3x4, model.FaceEast, coor4x4, nil},
		{"invalid move west", coor0x1, model.FaceWest, coor0x1, model.ErrInvalidMove},
		{"valid move south", coor0x1, model.FaceSouth, coor0x0, nil},
		{"valid move north", coor3x3, model.FaceNorth, coor3x4, nil},
		{"valid move south", coor3x3, model.FaceSouth, coor3x2, nil},
		{"valid move west", coor3x3, model.FaceWest, coor2x3, nil},
		{"valid move east", coor3x3, model.FaceEast, coor4x3, nil},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s {%d,%d}", tt.Name, tt.Current.X, tt.Current.Y), func(t *testing.T) {
			robot := model.Robot{
				Current: tt.Current,
				Face:    tt.Face,
			}

			err := r.Move(&robot, board5x5)

			assert.Equal(t, tt.ExpectedErr, err)
			assert.Equal(t, tt.ExpectedCoor, robot.Current)
		})
	}
}

func Test_robotUsecaseImpl_FaceLeft(t *testing.T) {
	r := &robotUsecaseImpl{}

	coor0x0 := model.Coordinate{0, 0}

	tests := []struct {
		Name         string
		Face         model.Coordinate
		ExpectedFace model.Coordinate
	}{
		{"north left", model.FaceNorth, model.FaceWest},
		{"west left", model.FaceWest, model.FaceSouth},
		{"south left", model.FaceSouth, model.FaceEast},
		{"east left", model.FaceEast, model.FaceNorth},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			robot := model.Robot{
				Current: coor0x0,
				Face:    tt.Face,
			}
			r.FaceLeft(&robot)

			assert.Equal(t, tt.ExpectedFace, robot.Face)
		})
	}
}

func Test_robotUsecaseImpl_FaceRight(t *testing.T) {
	r := &robotUsecaseImpl{}

	coor0x0 := model.Coordinate{0, 0}

	tests := []struct {
		Name         string
		Face         model.Coordinate
		ExpectedFace model.Coordinate
	}{
		{"north right", model.FaceNorth, model.FaceEast},
		{"west right", model.FaceWest, model.FaceNorth},
		{"south right", model.FaceSouth, model.FaceWest},
		{"east right", model.FaceEast, model.FaceSouth},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			robot := model.Robot{
				Current: coor0x0,
				Face:    tt.Face,
			}
			r.FaceRight(&robot)

			assert.Equal(t, tt.ExpectedFace, robot.Face)
		})
	}
}

func Test_robotUsecaseImpl_Report(t *testing.T) {
	r := &robotUsecaseImpl{}

	tests := []struct {
		Name     string
		X, Y     int
		Face     model.Coordinate
		Expected string
	}{
		{"0,0,NORTH", 0, 0, model.FaceNorth, "0,0,NORTH"},
		{"0,0,WEST", 0, 0, model.FaceWest, "0,0,WEST"},
		{"0,0,EAST", 0, 0, model.FaceEast, "0,0,EAST"},
		{"0,0,SOUTH", 0, 0, model.FaceSouth, "0,0,SOUTH"},
	}

	//add random test
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < 5; i++ {
		x := r1.Intn(5)
		y := r1.Intn(5)
		f := r1.Intn(4)
		var face model.Coordinate
		var faceStr string
		switch f {
		case 0:
			face = model.FaceNorth
			faceStr = model.NORTH
		case 1:
			face = model.FaceWest
			faceStr = model.WEST
		case 2:
			face = model.FaceSouth
			faceStr = model.SOUTH
		case 3:
			face = model.FaceEast
			faceStr = model.EAST
		}

		test := struct {
			Name     string
			X, Y     int
			Face     model.Coordinate
			Expected string
		}{
			fmt.Sprintf("Random %d,%d,%s", x, y, faceStr),
			x, y, face, fmt.Sprintf("%d,%d,%s", x, y, faceStr),
		}

		tests = append(tests, test)
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			robot := model.Robot{
				Current: model.Coordinate{tt.X, tt.Y},
				Face:    tt.Face,
			}
			result := r.Report(&robot)

			assert.Equal(t, tt.Expected, result)
		})
	}
}
