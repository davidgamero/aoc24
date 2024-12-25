package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	currentSchematic := []string{}
	keys := [][]string{}
	locks := [][]string{}
	for _, line := range strings.Split(string(data), "\n") {
		if len(strings.TrimSpace(line)) == 0 {
			if currentSchematic[0][0] == '.' {
				keys = append(keys, currentSchematic)
			} else {
				locks = append(locks, currentSchematic)
			}
			currentSchematic = []string{}
		} else {
			currentSchematic = append(currentSchematic, line)
		}
	}

	fitCount := 0
	for _, key := range keys {
		for _, lock := range locks {
			if !Interfere(key, lock) {
				fitCount++
			}
		}
	}
	fmt.Println("p1: fitCount:", fitCount)
}

func Interfere(key, lock []string) bool {
	for i := 0; i < len(key); i++ {
		for j := 0; j < len(key[i]); j++ {
			if key[i][j] == '#' && lock[i][j] == '#' {
				return true
			}
		}
	}
	// check for interference
	return false
}
