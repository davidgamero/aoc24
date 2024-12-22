package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	row int
	col int
}

type PositionSet map[Position]bool

// Merge toInclude into IntSet is
func (ps PositionSet) Merge(toInclude PositionSet) {
	for k, v := range toInclude {
		ps[k] = v
	}
}
func (ps PositionSet) Has(i Position) bool {
	return ps[i]
}
func (ps PositionSet) Put(i ...Position) {
	for _, pos := range i {
		ps[pos] = true
	}
}

// recursively search for summits reachable with elevation rise of 1
// look for adjacent space with elevation+1 value
func GetTrailHeadSummits(lines []string, row, col, targetElevation int) PositionSet {
	summits := PositionSet{}
	// zero value if we are outside the map
	if row < 0 || col < 0 || row >= len(lines) || col >= len(lines[row]) {
		return summits
	}
	currentElevation, err := strconv.Atoi(string(lines[row][col]))
	if err != nil {
		panic(err)
	}
	if currentElevation != targetElevation {
		return summits
	}
	if currentElevation == 9 {
		summits.Put(Position{row, col})
		return summits
	}

	summits.Merge(GetTrailHeadSummits(lines, row+1, col, targetElevation+1))
	summits.Merge(GetTrailHeadSummits(lines, row-1, col, targetElevation+1))
	summits.Merge(GetTrailHeadSummits(lines, row, col+1, targetElevation+1))
	summits.Merge(GetTrailHeadSummits(lines, row, col-1, targetElevation+1))
	return summits
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	totalScore := 0
	scores := [][]int{}
	for row, line := range lines {
		lineScores := []int{}
		for col := range line {
			summmits := GetTrailHeadSummits(lines, row, col, 0)
			lineScore := len(summmits)
			totalScore += lineScore
		}
		scores = append(scores, lineScores)
	}
	fmt.Println("")
	fmt.Printf("p1 total: %d\n", totalScore)

}
