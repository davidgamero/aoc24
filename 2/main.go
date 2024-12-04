package main

import (
  "os"
  "bufio"
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		report := strings.Split(line, " ")
  }
}

func reportIsSafe([]int report) bool{
  increased := false
  decreased := false

  for i := 0; i < len(report) - 1; i++{
    n1 = report[i]
    n2 = report[i+1]

    if n1 == n2{
      return false
    }

    // must be monotonically increasing or decreasing
    if n2 > n1{
      increased = true
    }
    if n2 < n1{
      decreased = true
    }
    if decreased && increased{
      return false
    }

    // can only increase or decrease by 1,2,3
    amount := n2 - n1
    if amount < 0{
      amount = -amount
    }
    if amount > 3 || amount < 1{
      return false
    }
  }

  return true
}
