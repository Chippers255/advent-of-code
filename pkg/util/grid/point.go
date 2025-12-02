package grid

// Point represents a 2D coordinate.
type Point struct {
	X int
	Y int
}

// Add returns the vector sum of two points.
func (p Point) Add(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

// Neighbors4 returns the four cardinal neighbor offsets around the point.
func Neighbors4() []Point {
	return []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
}

// Neighbors8 returns the eight surrounding neighbor offsets.
func Neighbors8() []Point {
	return []Point{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}
}
