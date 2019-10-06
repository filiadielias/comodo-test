package main

import (
	"fmt"
	"os"

	board_service "github.com/filiadielias/comodo-test/service/board"
	command_service "github.com/filiadielias/comodo-test/service/command"
	robot_service "github.com/filiadielias/comodo-test/service/robot"
)

func main() {
	params := os.Args[1:]

	// params = "PLACES 1,2,NORTH MOVE LEFT REPORT PLACE 0,0,EAST MOVE MOVE MOVE MOVE REPORT RIGHT MOVE MOVE REPORT"

	boardUsecase := board_service.NewBoardUsecaseImpl()
	board, err := boardUsecase.CreateBoard(5, 5)
	if err != nil {
		panic(err)
	}

	robotUsecase := robot_service.NewRobotUsecaseImpl(boardUsecase)

	commandUsecase := command_service.NewCommandUsecaseImpl(robotUsecase)
	result, err := commandUsecase.Run(params, board)
	if err != nil {
		panic(err)
	}

	for _, r := range result {
		fmt.Println(r)
	}
}
