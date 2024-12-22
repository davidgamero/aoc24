package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	row int
	col int
}
type EdgeInsideDirection string

const LEFT EdgeInsideDirection = "l"
const RIGHT EdgeInsideDirection = "r"
const UP EdgeInsideDirection = "u"
const DOWN EdgeInsideDirection = "d"

// Edge represents an edge left or above a row/col coord
type Edge struct {
	row       int
	col       int
	direction EdgeInsideDirection
}

type Set[T comparable] map[T]struct{}

func (ps Set[T]) Merge(otherSet Set[T]) {
	for p := range otherSet {
		ps[p] = struct{}{}
	}
}
func (ps Set[T]) Has(p T) bool {
	_, ok := ps[p]
	return ok
}
func (ps Set[T]) Put(p T) {
	ps[p] = struct{}{}
}

// FloodFindRegion returns the perimeter and map of positions and edges (left or top of position)
func FloodFindRegion(lines []string, row, col int, plantType rune, mapped Set[Position], entryEdge Edge) (int, Set[Position], Set[Edge]) {
	perimeter := 0
	areaSet := Set[Position]{} // set of positions within the region
	edgeSet := Set[Edge]{}
	thisPosition := Position{row, col}

	// if we have exited the map
	if row < 0 || col < 0 || row >= len(lines) || col >= len(lines[row]) {
		edgeSet.Put(entryEdge)
		return 1, areaSet, edgeSet
	}
	// if we have exited our region
	if lines[row][col] != byte(plantType) {
		edgeSet.Put(entryEdge)
		return 1, areaSet, edgeSet
	}

	// this position matches our plantType
	if mapped.Has(thisPosition) {
		return 0, areaSet, edgeSet
	}
	areaSet.Put(thisPosition)
	mapped.Put(thisPosition)

	lp, la, le := FloodFindRegion(lines, row, col-1, plantType, mapped, Edge{row, col, LEFT})
	edgeSet.Merge(le)
	perimeter += lp
	areaSet.Merge(la)
	rp, ra, re := FloodFindRegion(lines, row, col+1, plantType, mapped, Edge{row, col + 1, RIGHT})
	edgeSet.Merge(re)
	perimeter += rp
	areaSet.Merge(ra)
	up, ua, ue := FloodFindRegion(lines, row-1, col, plantType, mapped, Edge{row, col, UP})
	edgeSet.Merge(ue)
	perimeter += up
	areaSet.Merge(ua)
	dp, da, de := FloodFindRegion(lines, row+1, col, plantType, mapped, Edge{row + 1, col, DOWN})
	edgeSet.Merge(de)
	perimeter += dp
	areaSet.Merge(da)

	return perimeter, areaSet, edgeSet
}

func NextEdge(e Edge) Edge {
	if e.direction == UP || e.direction == DOWN {
		return Edge{e.row, e.col + 1, e.direction}
	}
	if e.direction == LEFT || e.direction == RIGHT {
		return Edge{e.row + 1, e.col, e.direction}
	}
	panic("unrecognized direction")
}
func PrevEdge(e Edge) Edge {
	if e.direction == UP || e.direction == DOWN {
		return Edge{e.row, e.col - 1, e.direction}
	}
	if e.direction == LEFT || e.direction == RIGHT {
		return Edge{e.row - 1, e.col, e.direction}
	}
	panic("unrecognized direction")
}

func CountSides(edges Set[Edge]) int {
	numberOfSides := 0
	countedEdges := Set[Edge]{}
	for edge := range edges {
		if countedEdges.Has(edge) {
			continue
		}
		for prevEdge := edge; edges.Has(prevEdge); prevEdge = PrevEdge(prevEdge) {
			countedEdges.Put(prevEdge)
		}
		for nextEdge := edge; edges.Has(nextEdge); nextEdge = NextEdge(nextEdge) {
			countedEdges.Put(nextEdge)
		}
		numberOfSides += 1
	}
	return numberOfSides
}

func main() {
	file := ""
	if len(os.Args) < 2 {
		file = "input.txt"
	} else {
		file = os.Args[1]
	}
	fmt.Printf("reading input file: %s\n", file)
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")

	mapped := Set[Position]{}
	totalPrice := 0

	totalDiscountPrice := 0
	for row, line := range lines {
		for col := range line {
			if mapped.Has(Position{row, col}) {
				continue
			}
			p, areaSet, edgeSet := FloodFindRegion(lines, row, col, rune(lines[row][col]), mapped, Edge{row, col, DOWN})
			//fmt.Printf("edgeset : %v\n", edgeSet)
			mapped.Merge(areaSet)

			area := len(areaSet)
			fmt.Printf("region '%c' area=%d perimeter=%d\n", rune(lines[row][col]), area, p)
			price := p * area
			totalPrice += price

			sides := CountSides(edgeSet)
			fmt.Printf(" sides=%d\n", sides)
			discountPrice := sides * area
			totalDiscountPrice += discountPrice
		}
	}

	fmt.Printf("p1 total price: %d\n", totalPrice)
	fmt.Printf("p2 total discount price: %d\n", totalDiscountPrice)
}

