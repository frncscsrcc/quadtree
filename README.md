QuadTree (data structure)
===
[![Go Report Card](https://goreportcard.com/badge/github.com/frncscsrcc/quadtree)](https://goreportcard.com/report/github.com/frncscsrcc/quadtree)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-99%25-brightgreen.svg?longCache=true&style=flat)</a>

Implement a quadtree search structure, inspired from the great code challenge https://thecodingtrain.com/CodingChallenges/098.1-quadtree.html

```
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

	// [(20, 20: B) (30, 30: C) (40, 40: D) (50, 50: E)]
}
```