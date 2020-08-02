package quadtree

import (
	"testing"
)

func TestNewQuadtree(t *testing.T) {
	r, _ := NewRectangle(0, 0, 100, 100)
	qt, err := NewQuadTree(10, r)
	if err != nil {
		t.Errorf("No error should be reported for capacity > 0")
	}

	if qt.level != 0 {
		t.Errorf("Default level should be 0")
	}
	if qt.capacity != 10 {
		t.Errorf("Capacity should be 10")
	}
	if qt.r != r {
		t.Errorf("QuadTree rectangle should be %v", r)
	}
	if len(qt.points) != 0 {
		t.Errorf("QuadTree points should be empty")
	}

	for _, capacity := range []int{-1, 0} {
		_, err := NewQuadTree(capacity, r)
		if err == nil {
			t.Errorf("Error should be reported for capacity = %d", capacity)
		}
	}
}

func TestString(t *testing.T) {
	r, _ := NewRectangle(0, 0, 100, 100)
	qt, _ := NewQuadTree(1, r)

	type formatChecker struct {
		newPoint *Point
		expectedString string
	}

	tests := []formatChecker{
		formatChecker{
			nil,
			"QT:0(PS:[], NE:nil, SE:nil, SW:nil, NW:nil)",
		},
		formatChecker{
			NewPoint(10,10,"P1"),
			"QT:0(PS:[P(10, 10: P1)], NE:nil, SE:nil, SW:nil, NW:nil)",
		},
		formatChecker{
			NewPoint(90,10,"P2"),
			"QT:0(PS:[P(10, 10: P1)], NE:QT:1(PS:[P(90, 10: P2)], NE:nil, SE:nil, SW:nil, NW:nil), SE:QT:1(PS:[], NE:nil, SE:nil, SW:nil, NW:nil), SW:QT:1(PS:[], NE:nil, SE:nil, SW:nil, NW:nil), NW:QT:1(PS:[], NE:nil, SE:nil, SW:nil, NW:nil))",
		},
	}

	for testNumber, test := range(tests){
		if test.newPoint != nil {
			qt.Add(test.newPoint)
		}
		returnedString := qt.String()
		if  returnedString != test.expectedString {
			t.Errorf("Qrong String for qt in case %d. Expected %s, returned %s", testNumber, test.expectedString, returnedString)
		}
	}


}


func TestNewChildQuadTree(t *testing.T) {
	r, _ := NewRectangle(0, 0, 100, 100)
	qt := newChildQuadTree(10, r, 2)
	if qt.level != 2 {
		t.Errorf("Child level should be 2")
	}
}

func TestAdd(t *testing.T) {

	pointsArrayContains := func(qt *QuadTree, p *Point) bool {
		if qt == nil {
			return false
		}
		for _, point := range qt.points {
			if p == point {
				return true
			}
		}
		return false
	}
	isContainedInNE := func(qt *QuadTree, p *Point) bool {
		return pointsArrayContains(qt.ne(), p)
	}
	isContainedInSE := func(qt *QuadTree, p *Point) bool {
		return pointsArrayContains(qt.se(), p)
	}
	isContainedInNW := func(qt *QuadTree, p *Point) bool {
		return pointsArrayContains(qt.nw(), p)
	}
	isContainedInSW := func(qt *QuadTree, p *Point) bool {
		return pointsArrayContains(qt.sw(), p)
	}

	r, _ := NewRectangle(0, 0, 100, 100)
	qt, _ := NewQuadTree(2, r)

	p1 := NewPoint(1, 1, "P1")
	qt.Add(p1)
	if !pointsArrayContains(qt, p1) {
		t.Errorf("%v should be contained in qt.points", p1)
	}
	if isContainedInNE(qt, p1) {
		t.Errorf("%v should not be in NE", p1)
	}
	if isContainedInSE(qt, p1) {
		t.Errorf("%v should not be in SE", p1)
	}
	if isContainedInNW(qt, p1) {
		t.Errorf("%v should not be in NW", p1)
	}
	if isContainedInSW(qt, p1) {
		t.Errorf("%v should not be in SW", p1)
	}

	p2 := NewPoint(2, 2, "P2")
	qt.Add(p2)
	if !pointsArrayContains(qt, p2) {
		t.Errorf("%v should be contained in qt.points", p2)
	}
	if isContainedInNE(qt, p2) {
		t.Errorf("%v should not be in NE", p2)
	}
	if isContainedInSE(qt, p2) {
		t.Errorf("%v should not be in SE", p2)
	}
	if isContainedInNW(qt, p2) {
		t.Errorf("%v should not be in NW", p2)
	}
	if isContainedInSW(qt, p2) {
		t.Errorf("%v should not be in SW", p2)
	}

	p3 := NewPoint(3, 3, "P3")
	qt.Add(p3)
	if pointsArrayContains(qt, p3) {
		t.Errorf("%v should not be contained in qt.points", p3)
	}
	if isContainedInNE(qt, p3) {
		t.Errorf("%v should not be in NE", p3)
	}
	if isContainedInSE(qt, p3) {
		t.Errorf("%v should not be in SE", p3)
	}
	if !isContainedInNW(qt, p3) {
		t.Errorf("%v should be in NW", p3)
	}
	if isContainedInSW(qt, p3) {
		t.Errorf("%v should not be in SW", p3)
	}
}

func TestDivide(t *testing.T) {
	capacity := 2
	r, _ := NewRectangle(0, 0, 100, 100)
	qt, _ := NewQuadTree(capacity, r)

	if qt.ne() != nil || qt.se() != nil || qt.nw() != nil || qt.sw() != nil {
		t.Errorf("No subtrees should exist now")
	}

	qt.divide()

	if qt.ne() == nil || qt.se() == nil || qt.nw() == nil || qt.sw() == nil {
		t.Errorf("All subtrees should exist now")
	}

	if qt.ne().capacity != capacity ||
		qt.se().capacity != capacity ||
		qt.nw().capacity != capacity ||
		qt.sw().capacity != capacity {
		t.Errorf("All subtrees should have the same capacity of the parent")
	}

	if qt.ne().level != 1 ||
		qt.se().level != 1 ||
		qt.nw().level != 1 ||
		qt.sw().level != 1 {
		t.Errorf("All subtrees should have level 1")
	}

	if qt.nw().r.String() != "R(0, 0, 50, 50)" {
		t.Errorf("NW rectangle %v is not correct", qt.nw().r)
	}
	if qt.sw().r.String() != "R(0, 51, 50, 100)" {
		t.Errorf("SW rectangle %v is not correct", qt.sw().r)
	}
	if qt.ne().r.String() != "R(51, 0, 100, 50)" {
		t.Errorf("NW rectangle %v is not correct", qt.ne().r)
	}
	if qt.se().r.String() != "R(51, 51, 100, 100)" {
		t.Errorf("SW rectangle %v is not correct", qt.se().r)
	}
}


func TestSearch(t *testing.T) {

	allPointsAreFound := func(qt *QuadTree, r Rectangle, expected []*Point) bool {
		result := qt.Search(r)
		if len(result) != len(expected) {
			return false
		}
		if len(result) == 0 {
			return true
		}

		for _, expectedPoint := range(expected) {
			found := false
			for _, resultPoint := range(result){
				if resultPoint == expectedPoint {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}

		return true
	}

	capacity := 1
	r, _ := NewRectangle(0, 0, 100, 100)
	qt, _ := NewQuadTree(capacity, r)

	P1 := NewPoint(10,10,"P1")
	P2 := NewPoint(40,40,"P2")
	P3 := NewPoint(60,60,"P3")
	P4 := NewPoint(90,90,"P4")

	qt.Add(P1)
	qt.Add(P2)
	qt.Add(P3)
	qt.Add(P4)

	r1, _ := NewRectangle(0, 0, 100, 100)
	if !allPointsAreFound(qt, r1, []*Point{P1, P2, P3, P4}){
		t.Errorf("Not all points returned for searching box %v", r1)
	}

	r2, _ := NewRectangle(0, 0, 50, 50)
	if !allPointsAreFound(qt, r2, []*Point{P1, P2}){
		t.Errorf("Not all points returned for searching box %v", r2)
	}

	r3, _ := NewRectangle(30, 30, 70, 70)
	if !allPointsAreFound(qt, r3, []*Point{P2, P3}){
		t.Errorf("Not all points returned for searching box %v", r3)
	}

	r4, _ := NewRectangle(95, 95, 0, 0)
	if !allPointsAreFound(qt, r4, []*Point{}){
		t.Errorf("Not all points returned for searching box %v", r4)
	}

}
