package main

//commands
//toy
//board

type Command struct {
	commands []string
}

type Board struct {
	Min Coordinate
	Max Coordinate
}

type Coordinate struct {
	X int8
	Y int8
}

type Toy struct {
	Curr Coordinate
	face Coordinate

	Board *Board
}

//direction coordinate
var (
	West  = Coordinate{-1, 0}
	East  = Coordinate{1, 0}
	North = Coordinate{0, 1}
	South = Coordinate{0, -1}
)

func (t *Toy) Move() {
	coor := Coordinate{
		X: t.Curr.X + t.face.X,
		Y: t.Curr.Y + t.face.Y,
	}

	if t.Board.IsValid(coor) {
		t.Curr = coor
	}
}

func (b Board) IsValid(coor Coordinate) bool {
	if !(b.Min.X <= coor.X && coor.Y <= b.Min.Y) {
		return false
	}

	return b.Max.X >= coor.X && b.Max.Y >= coor.Y
}

func main() {
	board := Board{
		Min: Coordinate{0, 0},
		Max: Coordinate{4, 4},
	}
}
