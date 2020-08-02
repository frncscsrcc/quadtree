package quadtree

import (
	"testing"
)

func TestNewRectangle(t *testing.T) {
	r1, err1 := NewRectangle(0, 0, 10, 10)
	if err1 != nil {
		t.Errorf("Rectangle %v should be created", r1)
	}

	_, err2 := NewRectangle(10, 0, 0, 10)
	if err2 == nil {
		t.Errorf("An error schould be reported if x1 > x2 ")
	}
	err2Expected := "invalid rectangle (10,0,0,10) x1 must be >= x2"
	if err2.Error() != err2Expected {
		t.Errorf("Correct error is returned if x1 > x2")
	}

	_, err3 := NewRectangle(0, 10, 0, 0)
	if err3 == nil {
		t.Errorf("An error schould be reported if y1 > y2 ")
	}
	err3Expected := "invalid rectangle (0,10,0,0) y1 must be >= y2"
	if err3.Error() != err3Expected {
		t.Errorf("Correct error is returned if y1 > y2")
	}
}

func TestRectangleString(t *testing.T) {
	r, _ := NewRectangle(1, 2, 3, 4)
	expected := "R(1, 2, 3, 4)"
	if r.String() != expected {
		t.Errorf("r.String() incorrect format. It is %s, but should be %s", r.String(), expected)
	}
}

func TestHelperMethods(t *testing.T) {
	r, _ := NewRectangle(1, 2, 3, 4)
	if r.X1() != 1 {
		t.Errorf("X1() returned invalid value %d instead of 1", r.X1())
	}
	if r.Y1() != 2 {
		t.Errorf("Y1() returned invalid value %d instead of 2", r.Y1())
	}
	if r.X2() != 3 {
		t.Errorf("X2() returned invalid value %d instead of 3", r.X2())
	}
	if r.Y2() != 4 {
		t.Errorf("Y2() returned invalid value %d instead of 4", r.Y2())
	}
}

func TestIntersects(t *testing.T) {
	getRectangle := func(x1, y1, x2, y2 int) Rectangle {
		rectangle, err := NewRectangle(x1, y1, x2, y2)
		if err != nil {
			t.Errorf("Rectangle %v should be valid", rectangle)
		}
		return rectangle
	}

	x1 := 10
	y1 := 10
	x2 := 20
	y2 := 20

	r := getRectangle(x1, y1, x2, y2)

	other1 := getRectangle(0, 0, 100, 100)
	if !r.Intersect(other1) {
		t.Errorf("Rectangle %v should intersect rectangle %v", r, other1)
	}

	same := getRectangle(0, 0, 100, 100)
	if !r.Intersect(other1) {
		t.Errorf("Rectangle %v should intersect rectangle %v", r, same)
	}

	notItersectingRectangles := []Rectangle{
		getRectangle(x2+1, y1, x2+2, y2),
		getRectangle(x1, y2+1, x2, y2+2),
		getRectangle(x1-2, y1, x1-1, y2),
		getRectangle(x1, y1-2, x2, y1-1),
	}

	for _, other2 := range notItersectingRectangles {
		if r.Intersect(other2) {
			t.Errorf("Rectangle %v should not intersect rectangle %v", r, other2)
		}
	}
}
