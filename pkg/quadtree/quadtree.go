package quadtree

import (
	"errors"
	"fmt"
	"strings"
)

// QuadTree reppresent quadtree object that can contains Points objects
type QuadTree struct {
	level    int
	capacity int
	points   []*Point
	r        Rectangle
	subTrees [4]*QuadTree
}

// NewQuadTree returns a new pointer to a Quadtree object, defining the capacity and
// the size of the searching space, as rectangle.
func NewQuadTree(capacity int, r Rectangle) (*QuadTree, error) {
	if capacity <= 0 {
		return &QuadTree{}, errors.New("capacity need to be > 0")
	}
	qt := QuadTree{
		level:    0,
		capacity: capacity,
		r:        r,
		points:   make([]*Point, 0, capacity),
	}
	return &qt, nil
}

func newChildQuadTree(capacity int, r Rectangle, level int) *QuadTree {
	qt := QuadTree{
		level:    level,
		capacity: capacity,
		r:        r,
		points:   make([]*Point, 0, capacity),
	}
	return &qt
}

// Add adds a point to the quadtree
func (qt *QuadTree) Add(point *Point) bool {
	if !point.ContainedIn(qt.r) {
		return false
	}
	if len(qt.points) < qt.capacity {
		qt.points = append(qt.points, point)
		return true
	}
	if qt.subTrees[0] == nil {
		qt.divide()
	}
	for _, subTree := range qt.subTrees {
		if subTree.Add(point) {
			return true
		}
	}
	return false
}

// Search returns the point incuded in a rectangle search area.
func (qt *QuadTree) Search(r Rectangle) []*Point {
	points := make([]*Point, 0)
	if !qt.intersect(r) {
		return points
	}
	for _, p := range qt.points {
		if p.ContainedIn(r) {
			points = append(points, p)
		}
	}
	if qt.subTrees[0] != nil {
		for _, subTree := range qt.subTrees {
			subTreePoints := subTree.Search(r)
			if len(subTreePoints) > 0 {
				for _, p := range subTreePoints {
					if p.ContainedIn(r) {
						points = append(points, p)
					}
				}
			}
		}
	}
	return points
}

func (qt *QuadTree) intersect(r Rectangle) bool {
	return qt.r.Intersect(r)
}

func (qt *QuadTree) divide() {
	getMiddle := func(a int, b int) int {
		return int((b - a) / 2)
	}

	x1 := qt.r.X1()
	y1 := qt.r.Y1()
	x2 := qt.r.X2()
	y2 := qt.r.Y2()

	// NE
	neRectangle, _ := NewRectangle(getMiddle(x1, x2)+1, y1, x2, getMiddle(y1, y2))
	qt.subTrees[0] = newChildQuadTree(qt.capacity, neRectangle, qt.level+1)

	// SE
	seRectangle, _ := NewRectangle(getMiddle(x1, x2)+1, getMiddle(y1, y2)+1, x2, y2)
	qt.subTrees[1] = newChildQuadTree(qt.capacity, seRectangle, qt.level+1)

	// NW
	nwRectangle, _ := NewRectangle(x1, y1, getMiddle(x1, x2), getMiddle(y1, y2))
	qt.subTrees[2] = newChildQuadTree(qt.capacity, nwRectangle, qt.level+1)

	// SW
	swRectangle, _ := NewRectangle(x1, getMiddle(y1, y2)+1, getMiddle(x1, x2), y2)
	qt.subTrees[3] = newChildQuadTree(qt.capacity, swRectangle, qt.level+1)
}

func (qt *QuadTree) ne() *QuadTree {
	return qt.subTrees[0]
}
func (qt *QuadTree) se() *QuadTree {
	return qt.subTrees[1]
}
func (qt *QuadTree) nw() *QuadTree {
	return qt.subTrees[2]
}
func (qt *QuadTree) sw() *QuadTree {
	return qt.subTrees[3]
}

// String formats the QuadTree object as a "nice" string
func (qt *QuadTree) String() string {
	if qt == nil {
		return "nil"
	}
	str := fmt.Sprintf("QT:%d(", qt.level)
	str += "PS:["
	pointAsStrings := make([]string, 0)
	for _, point := range qt.points {
		pointAsStrings = append(pointAsStrings, point.String())
	}
	str += strings.Join(pointAsStrings, ", ")
	str += "], "
	str += "NE:" + qt.ne().String() + ", "
	str += "SE:" + qt.se().String() + ", "
	str += "SW:" + qt.sw().String() + ", "
	str += "NW:" + qt.nw().String()
	str += ")"
	return str
}
