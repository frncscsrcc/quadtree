package quadtree

import (
	"testing"
)

func TestPoint(t *testing.T) {
	p := NewPoint(10, 20, "ABCD")

	pString := p.String()
	expectedString := "P(10, 20: ABCD)"
	if pString != expectedString {
		t.Errorf("String() returns incorrect format (%s, expected %s)", pString, expectedString)
	}

	r1, _ := NewRectangle(0, 0, 50, 50)
	if !p.ContainedIn(r1) {
		t.Errorf("Point %v should be contained in %v", p, r1)
	}

	r2, _ := NewRectangle(50, 50, 100, 100)
	if p.ContainedIn(r2) {
		t.Errorf("Point %v should not be contained in %v", p, r2)
	}
}
