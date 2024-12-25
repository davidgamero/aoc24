package main

import (
	"fmt"
	"os"
	"strings"
)

type SchematicType string

const KEY SchematicType = "key"
const LOCK SchematicType = "lock"

type Heights struct {
	Values    []int
	Type      SchematicType
	MaxHeight int
}

func ReadSchematic(schematic []string) Heights {
	isKey := schematic[0][0] == '.'
	values := []int{}
	numRows := len(schematic)
	maxHeight := numRows - 2
	for col := 0; col < len(schematic[0]); col++ {
		numDotsInHeight := 0
		for row := 1; row < numRows-1; row++ {
			if schematic[row][col] == '.' {
				numDotsInHeight++
			}
		}
		if isKey {
			values = append(values, maxHeight-numDotsInHeight)
			fmt.Println("key numDotsInHeight:", numDotsInHeight, "maxHeight:", maxHeight)
		} else {
			fmt.Println("lock numDotsInHeight:", numDotsInHeight, "maxHeight:", maxHeight)
			values = append(values, maxHeight-numDotsInHeight)
		}
	}
	h := Heights{
		Values:    values,
		MaxHeight: maxHeight,
	}
	if isKey {
		h.Type = KEY
	} else {
		h.Type = LOCK
	}

	return h

}

func CanFit(key Heights, lock Heights) bool {
	if key.MaxHeight != lock.MaxHeight {
		panic("key and lock must have the same max height")
	}
	if key.Type != KEY || lock.Type != LOCK {
		panic("key must be a key and lock must be a lock")
	}
	for i := 0; i < len(key.Values); i++ {
		if key.Values[i]+lock.Values[i] > key.MaxHeight {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <input>")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	currentSchematic := []string{}

	keys := []Heights{}
	locks := []Heights{}
	for _, line := range strings.Split(string(data), "\n") {
		if len(strings.TrimSpace(line)) == 0 {
			heights := ReadSchematic(currentSchematic)
			fmt.Println("heights:", heights, "type:", heights.Type)
			if heights.Type == KEY {
				keys = append(keys, heights)
			} else {
				locks = append(locks, heights)
			}
			currentSchematic = []string{}
		} else {
			currentSchematic = append(currentSchematic, line)
		}
	}

	fmt.Println("keys:", keys)
	fmt.Println("locks:", locks)

	fitCount := 0
	for _, key := range keys {
		for _, lock := range locks {
			if CanFit(key, lock) {
				fitCount++
				fmt.Println("key:", key, "fits lock:", lock)
			} else {
				fmt.Println("key:", key, "does not fit lock:", lock)
			}
		}
	}
	fmt.Println("p1: fitCount:", fitCount)
}
