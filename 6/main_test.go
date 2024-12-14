package main

import (
	"fmt"
	"testing"
)

func TestGuardSim(t *testing.T) {
	tests := []struct {
		name       string
		lines      []string
		lenvisited int
	}{
		{
			name: "start on edge",
			lines: []string{
				"..^",
				"...",
				"...",
			},
			lenvisited: 1,
		}, {
			name: "start on edge",
			lines: []string{
				"...",
				"...",
				".^.",
			},
			lenvisited: 3,
		},
	}

	for _, test := range tests {
		g := GuardSim(test.lines)
		if len(g.Visited) != test.lenvisited {
			t.Errorf("%s: visited %d, expected %d\n", test.name, len(g.Visited), test.lenvisited)
			for _, line := range g.ShowVisited(test.lines) {
				fmt.Println(line)
			}
			fmt.Printf("visited: %v\n", g.Visited)
		}
	}
}
