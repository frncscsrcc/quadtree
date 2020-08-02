package quadtree

import (
	"fmt"
)

type Point struct {
	x    int
	y    int
	data interface{}
}

func NewPoint(x int, y int, data interface{}) *Point {
	return &Point{x, y, data}
}

func (p *Point) String() string {
	return fmt.Sprintf("P(%d, %d: %+v)", p.x, p.y, p.data)
}

func (p *Point) ContainedIn(r Rectangle) bool {
	return p.x >= r.X1() && p.x < r.X2() && p.y >= r.Y1() && p.y < r.Y2()
}
