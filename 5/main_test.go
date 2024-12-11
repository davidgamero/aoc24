package main

import (
	"reflect"
	"testing"
)

func TestReorderToFollowRules(t *testing.T) {
	// test cases
	tests := []struct {
		pageRules []PageRule
		update    []int
		expected  []int
	}{
		{
			pageRules: []PageRule{{1, 2}, {3, 4}},
			update:    []int{1, 2, 3, 4},
			expected:  []int{1, 2, 3, 4},
		},
		{
			pageRules: []PageRule{{1, 2}, {3, 4}, {2, 3}},
			update:    []int{4, 3, 2, 1},
			expected:  []int{1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		actual := ReorderToFollowRules(test.pageRules, test.update)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("reorderToFollowRules(%v, %v) = %v; expected %v", test.pageRules, test.update, actual, test.expected)
		}
	}
}
