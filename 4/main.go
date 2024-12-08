package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := [][]rune{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		newLine := make([]rune, 0)
		for _, r := range line {
			newLine = append(newLine, r)
		}
		lines = append(lines, newLine)
	}

	test := [][]rune{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}}
	//fmt.Println(Mod(-2, 5))
	fmt.Println(GetDiags(test))

}

func GetDiags(lines [][]rune) (sw, se [][]rune) {
	cols := len(lines[0])
	rows := len(lines)
	numdiags := cols + rows - 1
	swdiags := make([][]rune, numdiags)
	sediags := make([][]rune, numdiags)

	// build sw diags
	for i := 0; i < rows; i++ {
		// i is offset index for getting diagonals
		// we assume non-jagged array here
		// sw digonals index across row first, then down left side col
		// 4x4 ex (value is diagonal containing the position):
		// 0 1 2 3
		// 6 0 1 2  4 =
		// 5 6 0 1  5 = 2-1+4-0, 4 = 2-1+4-1
		// 4 5 6 0
		//
		// nw diag
		// 3 2 1 0
		// 2 1 0 6  4 = i-1+len(row)-j = 1-1+4-0
		// 1 0 6 5  5 = 2-1+4-0, 4 = 2-1+4-1
		// 0 6 5 4

		for j := 0; j < cols; j++ {
			swdiagindex := Mod(i-j, len(swdiags))
			sediagindex := Mod((cols-i-1)-(rows-j-1), len(sediags))

			swdiags[swdiagindex] = append(swdiags[swdiagindex], lines[i][j])
			sediags[sediagindex] = append(sediags[sediagindex], lines[i][j])
		}

	}
	return swdiags, sediags
}
