package quadtree

import (
	"fmt"
)

// Rectangle defines a simple 2D rectangle object
type Rectangle struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

// NewRectangle validates the coordinates and returns a new Rectangle object or an error.
func NewRectangle(x1, y1, x2, y2 int) (Rectangle, error) {
	if x1 > x2 {
		return Rectangle{}, fmt.Errorf("invalid rectangle (%d,%d,%d,%d) x1 must be >= x2", x1, y1, x2, y2)
	}
	if y1 > y2 {
		return Rectangle{}, fmt.Errorf("invalid rectangle (%d,%d,%d,%d) y1 must be >= y2", x1, y1, x2, y2)
	}
	return Rectangle{x1, y1, x2, y2}, nil
}

// Intersect checks if a rectangle intersects with another rectangle
func (r Rectangle) Intersect(other Rectangle) bool {
	return !(r.X2() < other.X1() || r.X1() > other.X2() || r.Y2() < other.Y1() || r.Y1() > other.Y2())
}

// X1 returns x1
func (r Rectangle) X1() int {
	return r.x1
}

// X2 returns x2
func (r Rectangle) X2() int {
	return r.x2
}

// Y1 returns y1
func (r Rectangle) Y1() int {
	return r.y1
}

// Y2 returns y2
func (r Rectangle) Y2() int {
	return r.y2
}

// String formats a nice string for each rectangle object
func (r Rectangle) String() string {
	return fmt.Sprintf("R(%d, %d, %d, %d)", r.x1, r.y1, r.x2, r.y2)
}
