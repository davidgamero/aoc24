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

type PositionWithDirection struct {
	Position
	Direction Direction
}

type Guard struct {
	Position
	direction Direction
	// dont @ me im giving up ergonomics for 1byte savings on empty struct pattern
	Visited map[Position]bool
}

func NewGuard(p Position, d Direction) Guard {
	g := Guard{
		Position:  p,
		direction: UP,
		Visited:   map[Position]bool{},
	}
	g.Visit(p)
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

	g, _, err := GuardSim(lines)
	if err != nil {
		panic(err)
	}
	fmt.Printf("spaces visited: %d\n", len(g.Visited))

	loopplacements := 0
	gstart, err := GetGuardPosition(lines)
	if err != nil {
		panic(err)
	}

	fmt.Printf("checking %d potential object positions\n", len(g.Visited))
	potentialpositioncount := len(g.Visited)
	donecount := 0
	for p := range g.Visited {
		if p == gstart {
			fmt.Printf(" %d skipping %v\n", donecount*100/potentialpositioncount, p)
			continue
		}
		fmt.Printf(" %d testing %v\n", (donecount*100)/potentialpositioncount, p)
		lineswithobject := InsertItem(lines, p)
		_, isloop, err := GuardSim(lineswithobject)
		if err != nil {
			panic(err)
		}
		if isloop {
			loopplacements += 1
		}
		donecount += 1
	}

	fmt.Printf("places that make a loop: %d\n", loopplacements)
}

func InsertItem(lines []string, p Position) []string {
	newlines := []string{}
	for y := range lines {
		l := ""
		for x := range len(lines[y]) {
			if x == p.X && y == p.Y {
				l = l + "O"
			} else {
				l = l + string(lines[y][x])
			}
		}
		newlines = append(newlines, l)
	}
	return newlines
}

func GetGuardPosition(lines []string) (Position, error) {
	for y := range len(lines) {
		for x := range len(lines[y]) {
			if lines[y][x] == '^' {
				return Position{X: x, Y: y}, nil
			}
		}
	}
	return Position{}, fmt.Errorf("guard not found in lines")
}

// run simulation, returns Guard, loops
func GuardSim(lines []string) (Guard, bool, error) {
	var guard Guard

	pos, err := GetGuardPosition(lines)
	if err != nil {
		return guard, false, fmt.Errorf("getting guard position: %w", err)
	}
	guard = NewGuard(pos, UP)

	looping := map[PositionWithDirection]bool{
		{
			Position:  guard.Position,
			Direction: guard.direction,
		}: true,
	}
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
		case 'O':
			fallthrough // user-placed object
		case '#':
			guard.direction = turnright[guard.direction]
		}

		if looping[PositionWithDirection{Position: guard.Position, Direction: guard.direction}] {
			return guard, true, nil
		}
		looping[PositionWithDirection{Position: guard.Position, Direction: guard.direction}] = true
	}

	return guard, false, nil
}
