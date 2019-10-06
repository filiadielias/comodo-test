package model

import "errors"

//Error variables
var (
	ErrInvalidBoardSize = errors.New("invalid board size")

	ErrCoordinateNotExists = errors.New("coordinate not exists")

	ErrInvalidMove = errors.New("invalid move")
)
