package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	X int
	Y int
}

// Get AntiNodes from a pair of Positions
func GetAntiNodes(p1, p2 Position) []Position {
	if p1 == p2 {
		return []Position{}
	}
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y

	return []Position{
		{X: p1.X - dx, Y: p1.Y - dy},
		{X: p2.X + dx, Y: p2.Y + dy},
	}
}

func GetResonantNodes(p1, p2 Position, lines []string) []Position {
	if p1 == p2 {
		return []Position{}
	}
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y

	nodePositions := []Position{p1, p2}
	// decreasing
	for i := 0; ; i += 1 {
		newnode := Position{X: p1.X - i*dx, Y: p1.Y - i*dy}
		if !IsValidNode(newnode, lines) {
			break
		}
		nodePositions = append(nodePositions, newnode)

	}
	// increasing
	for i := 0; ; i += 1 {
		newnode := Position{X: p1.X + i*dx, Y: p1.Y + i*dy}
		if !IsValidNode(newnode, lines) {
			break
		}
		nodePositions = append(nodePositions, newnode)

	}
	return nodePositions
}

func GetAllAntiNodes(positions []Position, lines []string) map[Position]bool {
	thisFrequencyAntiNodesSet := map[Position]bool{}
	for _, p1 := range positions {
		for _, p2 := range positions {
			antiNodes := GetAntiNodes(p1, p2)
			for _, node := range antiNodes {
				if !IsValidNode(node, lines) {
					fmt.Println("skipping invalid node: ", node)
					continue
				}
				thisFrequencyAntiNodesSet[node] = true
			}
		}
	}

	return thisFrequencyAntiNodesSet
}

func GetAllResonantNodes(positions []Position, lines []string) map[Position]bool {
	thisFrequencyAntiNodesSet := map[Position]bool{}
	for _, p1 := range positions {
		for _, p2 := range positions {
			antiNodes := GetResonantNodes(p1, p2, lines)
			for _, node := range antiNodes {
				thisFrequencyAntiNodesSet[node] = true
			}
		}
	}

	return thisFrequencyAntiNodesSet
}

func IsValidNode(p Position, lines []string) bool {
	return p.Y >= 0 && p.X >= 0 && p.Y < len(lines) && p.X < len(lines[p.Y])
}

func GetAntennaFrequencyToPositions(lines []string) map[rune][]Position {
	positions := map[rune][]Position{}
	for y, line := range lines {
		for x, char := range line {
			if char == '.' {
				continue
			}
			positions[char] = append(positions[char], Position{X: x, Y: y})
		}
	}
	return positions
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(dat), "\n")
	frequencyToPositions := GetAntennaFrequencyToPositions(lines)

	uniqueAntiNodes := map[Position]bool{}
	uniqueResonantNodes := map[Position]bool{}
	for _, positions := range frequencyToPositions {
		antiNodes := GetAllAntiNodes(positions, lines)
		for node := range antiNodes {
			uniqueAntiNodes[node] = true
		}

		resonantNodes := GetAllResonantNodes(positions, lines)
		for node := range resonantNodes {
			uniqueResonantNodes[node] = true
		}
	}

	fmt.Printf("number of unique antinodes: %d\n", len(uniqueAntiNodes))
	fmt.Printf("number of unique resonant nodes: %d\n", len(uniqueResonantNodes))
}
