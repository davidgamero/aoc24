package main

import (
	"reflect"
	"testing"
)

func TestGetAllAntiNodes(t *testing.T) {
	tests := []struct {
		name      string
		positions []Position
		lines     []string
		expected  map[Position]bool
	}{
		{
			name:      "two nodes",
			positions: []Position{{X: 2, Y: 1}, {X: 2, Y: 2}},
			lines: []string{
				"....",
				"..a.",
				"..a.",
				"....",
			},
			expected: map[Position]bool{{X: 2, Y: 0}: true, {X: 2, Y: 3}: true},
		},
	}

	for _, test := range tests {
		actual := GetAllAntiNodes(test.positions, test.lines)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("%s: actual(%v) != expected(%v)", test.name, actual, test.expected)
		}
	}
}
