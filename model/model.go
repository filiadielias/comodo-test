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

//Toy struct
type Robot struct {
	Current Coordinate
}
