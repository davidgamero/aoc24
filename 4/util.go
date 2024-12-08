package main

func Transpose(lines [][]rune) [][]rune {
	transposed := make([][]rune, len(lines[0]))
	for i := range transposed {
		transposed[i] = make([]rune, len(lines))
	}
	for i, line := range lines {
		for j, r := range line {
			transposed[j][i] = r
		}
	}
	return transposed
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}
