package quadtree

import (
	"errors"
	"fmt"
)

type Rectangle struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func NewRectangle(x1, y1, x2, y2 int) (Rectangle, error) {
	if x1 > x2 {
		return Rectangle{}, errors.New(fmt.Sprintf("invalid rectangle (%d,%d,%d,%d) x1 must be >= x2", x1, y1, x2, y2))
	}
	if y1 > y2 {
		return Rectangle{}, errors.New(fmt.Sprintf("invalid rectangle (%d,%d,%d,%d) y1 must be >= y2", x1, y1, x2, y2))
	}
	return Rectangle{x1, y1, x2, y2}, nil
}

func (r Rectangle) Intersect(other Rectangle) bool {
	return !(r.X2() < other.X1() || r.X1() > other.X2() || r.Y2() < other.Y1() || r.Y1() > other.Y2())
}

func (r Rectangle) X1() int {
	return r.x1
}

func (r Rectangle) X2() int {
	return r.x2
}

func (r Rectangle) Y1() int {
	return r.y1
}

func (r Rectangle) Y2() int {
	return r.y2
}
