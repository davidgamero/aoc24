package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type PageRule []int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pageRules := []PageRule{}
	updates := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "|") {
			l := strings.Split(line, "|")
			if len(l) != 2 {
				log.Fatalf("invalid rule line with len>2: %s", line)
			}
			lineints := []int{}
			for _, s := range l {
				i, err := strconv.Atoi(s)
				if err != nil {
					log.Fatalf("invalid rule line: %s", line)
				}
				lineints = append(lineints, i)
			}
			pageRules = append(pageRules, lineints)
			continue
		}

		if strings.Contains(line, ",") {
			l := strings.Split(line, ",")

			lineints := []int{}
			for _, s := range l {
				i, err := strconv.Atoi(s)
				if err != nil {
					log.Fatalf("invalid update line: %s", line)
				}
				lineints = append(lineints, i)
			}

			updates = append(updates, lineints)
			continue
		}
		log.Fatalf("invalid line: %s", line)
	}

	validUpdates := [][]int{}
	invalidUpdates := [][]int{}
	for _, update := range updates {
		if UpdateFollowsRules(pageRules, update) {
			validUpdates = append(validUpdates, update)
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}
	fmt.Printf("valid: %d, invalid: %d\n", len(validUpdates), len(invalidUpdates))

	middleSum := 0
	for _, update := range validUpdates {
		middleSum += update[(len(update) / 2)]
	}

	fmt.Println("p1: ", middleSum)
}

func UpdateFollowsRules(rules []PageRule, update []int) bool {
	page2i := map[int]int{}
	for i, v := range update {
		page2i[v] = i
	}

	for _, rule := range rules {
		if len(rule) != 2 {
			log.Fatalf("invalid rule: %v", rule)
		}
		if _, ok := page2i[rule[0]]; !ok {
			//fmt.Printf("skipping rule %v for update %v as %d is not present\n", rule, update, rule[0])
			continue
		}
		if _, ok := page2i[rule[1]]; !ok {
			//fmt.Printf("skipping rule %v for update %v as %d is not present\n", rule, update, rule[1])
			continue
		}

		if page2i[rule[0]] > page2i[rule[1]] {
			fmt.Printf("rule %v failed for update %v\n", rule, update)
			return false
		}
	}
	return true
}
