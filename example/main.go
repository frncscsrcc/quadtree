package main

import (
	"fmt"
	"github.com/frncscsrcc/quadtree/pkg/quadtree"
)

func main() {
	area, err := quadtree.NewRectangle(0, 0, 200, 200)
	if err != nil {
		panic(err)
	}

	qt, err := quadtree.NewQuadTree(5, area)
	if err != nil {
		panic(err)
	}

	qt.Add(quadtree.NewPoint(10, 10, "A"))
	qt.Add(quadtree.NewPoint(20, 20, "B"))
	qt.Add(quadtree.NewPoint(30, 30, "C"))
	qt.Add(quadtree.NewPoint(40, 40, "D"))
	qt.Add(quadtree.NewPoint(50, 50, "E"))
	qt.Add(quadtree.NewPoint(160, 160, "F"))
	qt.Add(quadtree.NewPoint(161, 160, "G"))
	qt.Add(quadtree.NewPoint(162, 160, "H"))
	qt.Add(quadtree.NewPoint(163, 160, "I"))
	qt.Add(quadtree.NewPoint(164, 160, "L"))
	qt.Add(quadtree.NewPoint(165, 160, "M"))
	qt.Add(quadtree.NewPoint(110, 10, "N"))

	searchArea, err := quadtree.NewRectangle(20, 20, 60, 60)
	if err != nil {
		panic(err)
	}

	searchResults := qt.Search(searchArea)

	fmt.Printf("%v", searchResults)
}
