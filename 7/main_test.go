package main

import "testing"

func TestConcat(t *testing.T) {
	cases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "simple",
			a:        1,
			b:        2,
			expected: 12,
		}, {
			name:     "more digits",
			a:        123,
			b:        32,
			expected: 12332,
		},
	}
	for _, c := range cases {
		actual := Concat(c.a, c.b)
		if actual != c.expected {
		}
	}
}
