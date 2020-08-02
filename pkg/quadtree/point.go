package quadtree

import (
	"fmt"
)

// Point reppresent a simple 2D point in an integer grid, data
// can be anything
type Point struct {
	x    int
	y    int
	data interface{}
}

// NewPoint creates a reference to a new Point object
func NewPoint(x int, y int, data interface{}) *Point {
	return &Point{x, y, data}
}

// String formats a nice string
func (p *Point) String() string {
	return fmt.Sprintf("P(%d, %d: %+v)", p.x, p.y, p.data)
}

// ContainedIn checks if the point is contained in a Rectangle object
func (p *Point) ContainedIn(r Rectangle) bool {
	return p.x >= r.X1() && p.x < r.X2() && p.y >= r.Y1() && p.y < r.Y2()
}
