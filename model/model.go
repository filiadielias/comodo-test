package model

//Coordinate struct
type Coordinate struct {
	X int
	Y int
}

//Board contains Min & Max to define board size
type Board struct {
	Min Coordinate
	Max Coordinate
}

//Robot struct
type Robot struct {
	Current Coordinate
	Face    Coordinate
}
