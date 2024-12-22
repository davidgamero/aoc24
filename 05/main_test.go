package main

import (
	"reflect"
	"testing"
)

func TestGetRequiredPages(t *testing.T) {
	tests := []struct {
		a        int
		rules    []PageRule
		expected []int
	}{
		{
			a:        5,
			rules:    []PageRule{{1, 5}, {2, 1}},
			expected: []int{2, 1, 5},
		},
	}

	for _, test := range tests {
		reqList := BuildRequiresList(test.rules)
		actual, _ := GetRequiredPages(test.a, reqList, []int{}, map[int]bool{})
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("GetRequiredPages(%d, ... ) = %d; expected %d", test.a, actual, test.expected)
			t.Errorf("  reqlist = %v", reqList)
		}
	}
}

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
			pageRules: []PageRule{{1, 3}, {3, 4}},
			update:    []int{4, 1},
			expected:  []int{1, 4},
		},
		{
			pageRules: []PageRule{{1, 3}, {3, 4}},
			update:    []int{1, 2, 3, 4},
			expected:  []int{1, 2, 3, 4},
		},
		{
			pageRules: []PageRule{{11, 12}, {13, 14}, {12, 13}},
			update:    []int{14, 13, 12, 11},
			expected:  []int{11, 12, 13, 14},
		},
	}

	for _, test := range tests {
		actual := SortUpdate(test.pageRules, test.update)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("reorderToFollowRules(%v, %v) = %v; expected %v", test.pageRules, test.update, actual, test.expected)
		}

		if !ValidateUpdate(test.pageRules, actual) {

			t.Errorf("sorted update: %v does not follow rules: %v", actual, test.pageRules)
		}
	}
}
