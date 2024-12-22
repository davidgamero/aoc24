package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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
	xmasCount := CountAllDirections(lines, "XMAS")
	fmt.Println("p1: ", xmasCount)

	p2Count := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if IsXMAS(lines, i, j) {
				p2Count++
			}
		}
	}
	fmt.Println("p2: ", p2Count)
}

// CountAllDirections the number of times a word appears in a 2D array of runes in any direction
func CountAllDirections(lines [][]rune, word string) int {
	countHorizontal := CountInRows(lines, word)
	fmt.Printf("Horizontal: %d\n", countHorizontal)
	countVertical := CountInRows(Transpose(lines), word)
	fmt.Printf("Vertical: %d\n", countVertical)

	se, sw := GetDiags(lines)
	countDiagSE := CountInRows(se, word)
	fmt.Printf("Diag SE: %d\n", countDiagSE)
	countDiagSW := CountInRows(sw, word)
	fmt.Printf("Diag SW: %d\n", countDiagSW)

	return countHorizontal + countVertical + countDiagSE + countDiagSW
}

// Count the number of times a word appears in a 2D array of runes in a row including backwards
func CountInRows(lines [][]rune, word string) int {
	count := 0
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			if i+len(word) > len(line) {
				continue
			}
			matchstring := line[i : i+len(word)]
			if string(matchstring) == word || string(matchstring) == Reverse(word) {
				count++
			}
		}
	}

	return count
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
		// se digonals index across row first, then down left side col
		// 4x4 ex (value is diagonal containing the position):
		// 0 1 2 3
		// 6 0 1 2  4 =
		// 5 6 0 1  5 = 2-1+4-0, 4 = 2-1+4-1
		// 4 5 6 0
		//
		// sw diag
		// 3 2 1 0
		// 2 1 0 6  4 = i-1+len(row)-j = 1-1+4-0
		// 1 0 6 5  5 = 2-1+4-0, 4 = 2-1+4-1
		// 0 6 5 4

		for j := 0; j < cols; j++ {
			sediagindex := Mod(i-j, len(sediags))
			swdiagindex := Mod((cols-i-1)-j, len(swdiags))

			sediags[sediagindex] = append(sediags[sediagindex], lines[i][j])
			swdiags[swdiagindex] = append(swdiags[swdiagindex], lines[i][j])
		}

	}
	return swdiags, sediags
}

type DiagonalDirection string

const SE DiagonalDirection = "SE"
const SW DiagonalDirection = "SW"

func GetDiagIndex(data [][]rune, i int, j int, dir DiagonalDirection) int {
	cols := len(data[0])
	rows := len(data)
	numdiags := cols + rows - 1
	if dir == SE {
		return Mod(i-j, numdiags)
	}
	if dir == SW {
		return Mod((cols-i-1)-j, numdiags)
	}
	panic("Invalid DiagonalDirection")
}

// Do both SE and SW diagonals have the word MAS forwards or backwards?
func IsXMAS(data [][]rune, i int, j int) bool {
	if i+2 >= len(data) || j+2 >= len(data[i]) {
		return false
	}
	sw := string(data[i][j]) + string(data[i+1][j+1]) + string(data[i+2][j+2])
	se := string(data[i][j+2]) + string(data[i+1][j+1]) + string(data[i+2][j])
	if sw != "MAS" && sw != "SAM" {
		return false
	}
	if se != "MAS" && se != "SAM" {
		return false
	}
	return true
}
