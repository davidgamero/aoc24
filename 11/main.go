package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BlinkStone(i int) []int {
	if i == 0 {
		return []int{1}
	}
	s := strconv.Itoa(i)
	if len(s)%2 == 0 {
		left, err := strconv.Atoi(s[:len(s)/2])
		if err != nil {
			panic(err)
		}
		right, err := strconv.Atoi(s[(len(s) / 2):])
		if err != nil {
			panic(err)
		}
		return []int{left, right}
	}
	return []int{i * 2024}
}

func Blink(stones []int) []int {
	newStones := []int{}
	for _, s := range stones {
		newStones = append(newStones, BlinkStone(s)...)
	}
	return newStones
}

type StoneAndDepth struct {
	stone int
	depth int
}

var CountStoneCache = map[StoneAndDepth]int{}

func CountStones(stone int, depth int) int {
	if depth == 0 {
		return 1
	}
	count := 0
	for _, s := range BlinkStone(stone) {
		if c, ok := CountStoneCache[StoneAndDepth{s, depth - 1}]; ok {
			count += c
			continue
		}
		c := CountStones(s, depth-1)
		CountStoneCache[StoneAndDepth{s, depth - 1}] = c
		count += c
	}
	return count
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	strs := strings.Split(string(data), " ")
	nums := []int{}
	for _, s := range strs {
		s = strings.TrimSpace(s)
		if len(s) == 0 {
			continue
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		nums = append(nums, i)
	}

	p1Count := 0
	for _, s := range nums {
		p1Count += CountStones(s, 25)
	}
	fmt.Printf("p1: %d stones\n", p1Count)
	// 213625
	//fmt.Printf("p2: len to process: %d \n", len(nums))
	p2Count := 0
	for _, s := range nums {
		p2Count += CountStones(s, 75)
	}
	fmt.Printf("p2: %d stones\n", p2Count)
}
