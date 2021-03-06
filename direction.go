package main

// Direction represents a capture direction when a new piece is set on the board.
type Direction struct {
	dx, dy int
}

// Next returns the coordinates of the next cell along the given direction.
func (d *Direction) Next(x, y int) (int, int) {
	return x + d.dx, y + d.dy
}

// TODO: the order is not relevant: consider changing to a set
var directions = []*Direction{
	&Direction{-1, 0},  // up
	&Direction{-1, 1},  // up-right
	&Direction{0, 1},   // right
	&Direction{1, 1},   // down-right
	&Direction{1, 0},   // down
	&Direction{1, -1},  // down-left
	&Direction{0, -1},  // left
	&Direction{-1, -1}, // up-left
}
