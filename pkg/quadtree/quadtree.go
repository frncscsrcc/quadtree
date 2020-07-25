package quadtree

import (
	"errors"
	"fmt"
)

type QuadTree struct {
	level    int
	capacity int
	points   []*Point
	r        Rectangle
	x1       int
	y1       int
	x2       int
	y2       int
	subTrees [4]*QuadTree
}

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
			subTree_points := subTree.Search(r)
			if len(subTree_points) > 0 {
				for _, p := range subTree_points {
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
	ne_r, _ := NewRectangle(getMiddle(x1, x2), y1, x2, getMiddle(y1, y2))
	qt.subTrees[0] = newChildQuadTree(qt.capacity, ne_r, qt.level+1)

	// SE
	se_r, _ := NewRectangle(getMiddle(x1, x2), getMiddle(y1, y2), x2, y2)
	qt.subTrees[1] = newChildQuadTree(qt.capacity, se_r, qt.level+1)

	// NW
	nw_r, _ := NewRectangle(x1, y1, getMiddle(x1, x2), getMiddle(y1, y2))
	qt.subTrees[2] = newChildQuadTree(qt.capacity, nw_r, qt.level+1)

	// SW
	sw_r, _ := NewRectangle(x1, getMiddle(y1, y2), getMiddle(x1, x2), qt.y2)
	qt.subTrees[3] = newChildQuadTree(qt.capacity, sw_r, qt.level+1)
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

func (qt *QuadTree) String() string {

	space := func(n int) string {
		str := ""
		for i := 0; i < n; i++ {
			str += "   "
		}
		return str
	}

	str := space(qt.level)
	for _, p := range qt.points {
		str += fmt.Sprintf("%v ", p)
	}
	if qt.subTrees[0] != nil {
		str += "\n"
		if len(qt.ne().points) > 0 {
			str += space(qt.level) + "NE:\n" + fmt.Sprintf("%v\n", qt.ne())
		}
		if len(qt.se().points) > 0 {
			str += space(qt.level) + "SE:\n" + fmt.Sprintf("%v\n", qt.se())
		}
		if len(qt.sw().points) > 0 {
			str += space(qt.level) + "SW:\n" + fmt.Sprintf("%v\n", qt.sw())
		}
		if len(qt.nw().points) > 0 {
			str += space(qt.level) + "NW:\n" + fmt.Sprintf("%v\n", qt.nw())
		}
	}
	return str
}
