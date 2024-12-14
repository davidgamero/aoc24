package main

import (
	"fmt"
	"testing"
)

func TestGetGuardPosition(t *testing.T) {
	tests := []struct {
		name        string
		lines       []string
		expectedpos Position
		shoulderr   bool
	}{
		{
			name:        "start on edge",
			expectedpos: Position{X: 2, Y: 0},
			lines: []string{
				"..^",
				"...",
				"...",
			},
		},
		{
			name:      "err on missing",
			shoulderr: true,
			lines: []string{
				"...",
				"...",
				"...",
			},
		},
	}

	for _, test := range tests {
		actual, err := GetGuardPosition(test.lines)
		diderr := err != nil
		if diderr != test.shoulderr {
			t.Errorf("err: %v, expected: %v", diderr, test.shoulderr)
		}
		if actual != test.expectedpos {
			t.Errorf("got %v position, expected: %v", actual, test.expectedpos)
		}
	}
}

func TestGuardSim(t *testing.T) {
	tests := []struct {
		name       string
		lines      []string
		lenvisited int
		isloop     bool
		shoulderr  bool
	}{
		{
			name: "start on top",
			lines: []string{
				"..^",
				"...",
				"...",
			},
			lenvisited: 1,
		}, {
			name: "start on bottom",
			lines: []string{
				"...",
				"...",
				".^.",
			},
			lenvisited: 3,
		}, {
			name: "inserted objects respected",
			lines: []string{
				"...",
				".O.",
				".^.",
			},
			lenvisited: 2,
		}, {
			name: "loop",
			lines: []string{
				"####",
				"#..#",
				"#^.#",
				"####",
			},
			lenvisited: 4,
			isloop:     true,
		}, {
			name: "loop with insert",
			lines: []string{
				".#..",
				".^.O",
				"#...",
				"..#.",
			},
			lenvisited: 4,
			isloop:     true,
		}, {
			name: "start outside loop",
			lines: []string{
				".#..",
				"...O",
				"#...",
				".^#.",
			},
			lenvisited: 5,
			isloop:     true,
		},
	}

	for _, test := range tests {
		g, isloop, err := GuardSim(test.lines)
		diderr := err != nil
		if diderr != test.shoulderr {
			t.Errorf("%s err: %v, expected: %v", test.name, diderr, test.shoulderr)
		}
		if isloop != test.isloop {
			t.Errorf("%s got isloop=%v, expected: %v", test.name, isloop, test.isloop)
		}
		if len(g.Visited) != test.lenvisited {
			t.Errorf("%s visited %d, expected %d\n", test.name, len(g.Visited), test.lenvisited)
			for _, line := range g.ShowVisited(test.lines) {
				fmt.Println(line)
			}
			fmt.Printf("visited: %v\n", g.Visited)
		}
	}
}
