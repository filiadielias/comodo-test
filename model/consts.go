package model

//face direction
const (
	NORTH string = "NORTH"
	SOUTH string = "SOUTH"
	WEST  string = "WEST"
	EAST  string = "EAST"
)

//direction coordinate
var (
	FaceNorth = Coordinate{0, 1}
	FaceEast  = Coordinate{1, 0}
	FaceSouth = Coordinate{0, -1}
	FaceWest  = Coordinate{-1, 0}
)
