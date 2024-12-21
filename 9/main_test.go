package main

import (
	"reflect"
	"testing"
)

const E = EMPTYID

func TestExpandDiskMap(t *testing.T) {
	tests := []struct {
		name     string
		rawMap   string
		expected []int
	}{
		{
			name:     "simple",
			rawMap:   "11111",
			expected: []int{0, E, 1, E, 2},
		},
		{
			name:     "given",
			rawMap:   "12345",
			expected: []int{0, E, E, 1, 1, 1, E, E, E, E, 2, 2, 2, 2, 2},
		},
	}

	for _, test := range tests {
		actual, err := ExpandDiskMap(test.rawMap)
		if err != nil {
			t.Errorf("error expanding disk map")
		}
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("ExpandRawMap(%v) expected(%v) != actual(%v)", test.rawMap, test.expected, actual)
		}
	}
}

func TestCompactDiskMap(t *testing.T) {
	tests := []struct {
		name        string
		expandedMap []int
		expected    []int
	}{
		{
			name:        "simple",
			expandedMap: []int{1, E, 2, E, 3},
			expected:    []int{1, 3, 2, E, E},
		},
		{
			name:        "given",
			expandedMap: []int{0, E, E, 1, 1, 1, E, E, E, E, 2, 2, 2, 2, 2},
			expected:    []int{0, 2, 2, 1, 1, 1, 2, 2, 2, E, E, E, E, E, E},
		},
	}

	for _, test := range tests {
		actual := CompactExpandedMap(test.expandedMap)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("CompactExpandedMap(%v) expected(%v) != actual(%v)", test.expandedMap, test.expected, actual)
		}
	}
}
func TestCompactWholeFiles(t *testing.T) {
	tests := []struct {
		name        string
		expandedMap []int
		expected    []int
	}{
		{
			name:        "simple",
			expandedMap: []int{1, E, 2, E, 3},
			expected:    []int{1, 3, 2, E, E},
		},
		{
			name:        "given",
			expandedMap: []int{0, E, E, 1, 1, E, E, E, E, 2, 2, 2},
			expected:    []int{0, 1, 1, E, E, 2, 2, 2, E, E, E, E},
		},
	}

	for _, test := range tests {
		actual := CompactWholeFiles(test.expandedMap)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("CompactWholeFiles(%v) expected(%v) != actual(%v)", test.expandedMap, test.expected, actual)
		}
	}
}

func TestSwapLen(t *testing.T) {
	tests := []struct {
		name     string
		s        []int
		i        int
		j        int
		n        int
		expected []int
	}{
		{
			name:     "simple",
			s:        []int{1, 2, 3, 4},
			i:        0,
			j:        2,
			n:        2,
			expected: []int{3, 4, 1, 2},
		}, {
			name:     "longer",
			s:        []int{1, 2, 3, 4, 5, 6},
			i:        1,
			j:        4,
			n:        2,
			expected: []int{1, 5, 6, 4, 2, 3},
		}, {
			name:     "not last",
			s:        []int{1, 2, 3, 4, 5, 6},
			i:        1,
			j:        3,
			n:        2,
			expected: []int{1, 4, 5, 2, 3, 6},
		},
	}

	for _, test := range tests {
		actual, err := SwapLen(test.s, test.i, test.j, test.n)
		if err != nil {
			t.Errorf("%s: error swapping", test.name)
		}
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("SwapLen(%v) expected(%v) != actual(%v)", test.s, test.expected, actual)
		}
	}
}

func TestGetFirstEmptySpace(t *testing.T) {
	tests := []struct {
		name        string
		expandedMap []int
		spaceLength int
		expected    int
	}{
		{
			name:        "simple",
			expandedMap: []int{1, E, 2, E, 3},
			spaceLength: 1,
			expected:    1,
		},
		{
			name:        "given",
			expandedMap: []int{0, E, E, 1, 1, 1, E, E, E, E, 2, 2, 2, 2, 2},
			spaceLength: 3,
			expected:    6,
		},
	}

	for _, test := range tests {
		actual, found := GetFirstEmptySpace(test.expandedMap, test.spaceLength)
		if !found {
			t.Errorf("%s: expected to find empty space", test.name)
		}
		if actual != test.expected {
			t.Errorf("GetFirstEmptySpace(%v) expected(%v) != actual(%v)", test.expandedMap, test.expected, actual)
		}
	}
}
