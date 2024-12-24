package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	// store connections going both ways
	connections := make(map[string]map[string]bool)
	startsWithT := []string{}
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		// split line into two computers
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			panic("invalid line")
		}
		computer1 := parts[0]
		computer2 := parts[1]
		if connections[computer1] == nil {
			connections[computer1] = make(map[string]bool)
		}
		if connections[computer2] == nil {
			connections[computer2] = make(map[string]bool)
		}
		connections[computer1][computer2] = true
		connections[computer2][computer1] = true
		if strings.HasPrefix(computer1, "t") {
			startsWithT = append(startsWithT, computer1)
		}
		if strings.HasPrefix(computer2, "t") {
			startsWithT = append(startsWithT, computer2)
		}
	}

	tConnectedTriplets := map[string]struct{}{}
	for _, tComputer := range startsWithT {
		tConnectedComputers := connections[tComputer]
		for tConnected1 := range tConnectedComputers {
			for tConnected2 := range tConnectedComputers {
				if tConnected1 == tConnected2 {
					continue
				}
				if connections[tConnected1][tConnected2] {
					computers := []string{tComputer, tConnected1, tConnected2}
					sort.Strings(computers)
					sorted := strings.Join(computers, "-")
					tConnectedTriplets[sorted] = struct{}{}
				}
			}
		}
	}
	fmt.Println("p1: number of triplets:", len(tConnectedTriplets))

	remainingNodes := make(map[string]bool)
	for node := range connections {
		remainingNodes[node] = true
	}
	largestClique := GetLargestClique(connections, map[string]bool{}, remainingNodes)
	fmt.Println("p2: largest clique:", len(largestClique))
}

func GetLargestClique(edges map[string]map[string]bool, currentClique map[string]bool, remainingNodes map[string]bool) map[string]bool {
	for node := range remainingNodes {
		for cliqueNode := range currentClique {
			if !edges[node][cliqueNode] {
				continue
			}
		}
	}
}
