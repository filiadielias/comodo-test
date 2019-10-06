package command

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/filiadielias/comodo-test/model"
	"github.com/filiadielias/comodo-test/service"
)

type commandUsecaseImpl struct {
	robotUsecase service.RobotUsecase
}

func NewCommandUsecaseImpl(robotUsecase service.RobotUsecase) service.CommandUsecase {
	return &commandUsecaseImpl{robotUsecase}
}

func (c *commandUsecaseImpl) Run(command string, board model.Board) ([]string, error) {
	var robot *model.Robot
	var result []string

	command = strings.ToLower(command)
	command = removeDuplicateSpaces(command)

	commands := strings.Split(command, " ")

	//at least must have 2 commands (ex: PLACE 1,2,NORTH)
	if len(commands) <= 1 {
		return result, model.ErrInvalidCommand
	}

	var prevCommand string
	for _, cmd := range commands {
		if prevCommand == "place" && !isValidCommand(cmd) {

			//get PLACE parameter
			arr := strings.Split(cmd, ",")
			if len(arr) != 3 {
				//if error, skip this command
				prevCommand = cmd
				continue
			}

			//convert to int
			x, err := strconv.ParseInt(arr[0], 10, 10)
			if err != nil {
				prevCommand = cmd
				continue
			}
			y, err := strconv.ParseInt(arr[1], 10, 10)
			if err != nil {
				prevCommand = cmd
				continue
			}

			//validate face
			var face model.Coordinate
			switch strings.ToUpper(arr[2]) {
			case model.NORTH:
				face = model.FaceNorth
			case model.EAST:
				face = model.FaceEast
			case model.WEST:
				face = model.FaceWest
			case model.SOUTH:
				face = model.FaceSouth
			default:
				prevCommand = cmd
				continue
			}

			newRobot, err := c.robotUsecase.New(board, model.Coordinate{int(x), int(y)}, face)
			if err != nil {
				prevCommand = cmd
				continue
			}

			robot = &newRobot
		}
		prevCommand = cmd

		//if robot is nil (no valid place command) then skip it
		if robot == nil {
			continue
		}

		if cmd == "left" {
			c.robotUsecase.FaceLeft(robot)
		}
		if cmd == "right" {
			c.robotUsecase.FaceRight(robot)
		}
		if cmd == "move" {
			err := c.robotUsecase.Move(robot, board)
			if err != nil {
				continue
			}
		}
		if cmd == "report" {
			result = append(result, c.robotUsecase.Report(robot))
		}
	}

	return result, nil
}

//helper function
func isValidCommand(sub string) bool {
	s := []string{"place", "left", "right", "move", "report"}
	for _, str := range s {
		if str == sub {
			return true
		}
	}

	return false
}

func removeDuplicateSpaces(s string) string {
	space := regexp.MustCompile(`\s+`)
	return space.ReplaceAllString(s, " ")
}
