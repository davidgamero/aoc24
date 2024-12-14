package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction string

const UP Direction = "up"
const DOWN Direction = "down"
const LEFT Direction = "left"
const RIGHT Direction = "right"

var turnright = map[Direction]Direction{
	UP:    RIGHT,
	RIGHT: DOWN,
	DOWN:  LEFT,
	LEFT:  UP,
}

type Position struct {
	X int // column, going right
	Y int // row, going down
}

type Guard struct {
	Position
	direction Direction
	// dont @ me im giving up ergonomics for 1byte savings on empty struct pattern
	Visited map[Position]bool
}

func NewGuard(x int, y int, d Direction) Guard {
	g := Guard{
		Position:  Position{x, y},
		direction: UP,
		Visited:   map[Position]bool{},
	}
	g.Visit(Position{x, y})
	return g
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (g *Guard) Visit(p Position) {
	hammingdist := Abs(p.X-g.X) + Abs(p.Y-g.Y)
	if hammingdist > 1 {
		panic("guard cannot move more than a hamming dist of 1")
	}
	g.Visited[p] = true
}

func (g *Guard) ShowVisited(lines []string) []string {
	visitedlines := []string{}
	for y := range lines {
		l := ""
		for x := range len(lines[y]) {
			if g.Visited[Position{x, y}] {
				l = l + "X"
			} else {
				l = l + string(lines[y][x])
			}
		}
		visitedlines = append(visitedlines, l)
	}
	return visitedlines
}

func main() {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(dat), "\n")
	fmt.Printf("read %d lines\n", len(lines))

	g := GuardSim(lines)
	fmt.Printf("spaces visited: %d\n", len(g.Visited))
}

func GuardSim(lines []string) Guard {
	var guard Guard
	for y := range len(lines) {
		for x := range len(lines[y]) {
			if lines[y][x] == '^' {
				guard = NewGuard(x, y, UP)
			}
		}
	}

	guard.Visit(guard.Position)

	for {
		var nextPosition Position
		switch guard.direction {
		case UP:
			nextPosition = Position{
				X: guard.X,
				Y: guard.Y - 1,
			}
		case DOWN:
			nextPosition = Position{
				X: guard.X,
				Y: guard.Y + 1,
			}
		case RIGHT:
			nextPosition = Position{
				X: guard.X + 1,
				Y: guard.Y,
			}
		case LEFT:
			nextPosition = Position{
				X: guard.X - 1,
				Y: guard.Y,
			}
		default:
			panic("unhandled guard direction")
		}

		// detect guard exit
		if nextPosition.Y < 0 ||
			nextPosition.Y >= len(lines) ||
			nextPosition.X < 0 ||
			nextPosition.X >= len(lines[nextPosition.Y]) {
			break
		}

		nextRune := lines[nextPosition.Y][nextPosition.X]
		switch nextRune {
		case '^':
			fallthrough // this is also a valid place for the guard to stand
		case '.':
			guard.Visit(nextPosition)
			guard.Position = nextPosition
		case '#':
			guard.direction = turnright[guard.direction]
		}
	}

	return guard
}
