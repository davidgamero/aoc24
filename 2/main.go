package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	safeReports := [][]int{}
	unsafeReports := [][]int{}
	safeReportsDampened := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		reportStrings := strings.Split(line, " ")
		newReport := []int{}
		for _, s := range reportStrings {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal("unable to convert to integer: ", err)
			}
			newReport = append(newReport, i)
		}
		if reportIsSafe(newReport) {
			safeReports = append(safeReports, newReport)
		} else {
			unsafeReports = append(unsafeReports, newReport)
		}

		if reportIsSafeDampener(newReport) {
			safeReportsDampened = append(safeReportsDampened, newReport)
		}
	}

	fmt.Printf("%d safe reports\n", len(safeReports))
	fmt.Printf("%d safe reports (dampened)\n", len(safeReportsDampened))
}

func reportIsSafe(report []int) bool {
	increased := false
	decreased := false

	for i := 0; i < len(report)-1; i++ {
		n1 := report[i]
		n2 := report[i+1]

		if n1 == n2 {
			return false
		}

		// must be monotonically increasing or decreasing
		if n2 > n1 {
			increased = true
		}
		if n2 < n1 {
			decreased = true
		}
		if decreased && increased {
			return false
		}

		// can only increase or decrease by 1,2,3
		amount := n2 - n1
		if amount < 0 {
			amount = -amount
		}
		if amount > 3 || amount < 1 {
			return false
		}
	}

	return true
}

func reportIsSafeDampener(report []int) bool {
	fmt.Println("checking ", report)
	if reportIsSafe(report) {
		return true
	}
	fmt.Println("  not safe, checking dampened...")

	// greedy time there's not that many
	for i := range report {
		reportwithouti := []int{}
		// before i
		if i > 0 {
			reportwithouti = append(reportwithouti, report[0:i]...)
		}
		// after i
		if i < len(report)-1 {
			reportwithouti = append(reportwithouti, report[i+1:]...)
		}
		fmt.Println(" i=", i)
		fmt.Println(" ", reportwithouti)
		if reportIsSafe(reportwithouti) {
			return true
		}
		fmt.Println(" not safe, continuing")
	}
	return false
}
