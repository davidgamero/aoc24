package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type PageRule []int // rule of [a, b] where a must come before b in update

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
		if ValidateUpdate(pageRules, update) {
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

	p2MiddleSum := 0
	for _, update := range invalidUpdates {
		orderedUpdate := SortUpdate(pageRules, update)
		if !ValidateUpdate(pageRules, orderedUpdate) {
			panic("invalid update after sorting")
		}
		p2MiddleSum += orderedUpdate[(len(orderedUpdate) / 2)]
	}

	fmt.Println("p2: ", p2MiddleSum)
}

func ValidateUpdate(rules []PageRule, update []int) bool {
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

// Get map of int a to []int that must be present before a in an update
// *non recursive*
func BuildRequiresList(rules []PageRule) map[int][]int {
	requires := map[int][]int{}
	for _, rule := range rules {
		if len(rule) != 2 {
			log.Fatalf("invalid rule: %v", rule)
		}
		if _, ok := requires[rule[1]]; !ok {
			requires[rule[1]] = []int{}
		}
		requires[rule[1]] = append(requires[rule[1]], rule[0])
	}
	return requires
}

// Get list of ints required before a page
// *recursive*
func GetRequiredPages(a int, requiresList map[int][]int, required []int, requiredcontains map[int]bool, updatecontains map[int]bool) ([]int, map[int]bool) {

	//fmt.Println("a = ", a)
	reqlist := requiresList[a]
	//fmt.Printf("reqlist[%d]= %v\n", a, reqlist)

	for _, dep := range reqlist {
		//fmt.Println("dep = ", dep)
		if requiredcontains[dep] {
			continue
		}
		if !updatecontains[dep] {
			continue
		}
		required, requiredcontains = GetRequiredPages(dep, requiresList, required, requiredcontains, updatecontains)

	}
	if !requiredcontains[a] {
		//fmt.Printf("appending %d to required\n", a)
		required = append(required, a)
		requiredcontains[a] = true
	} else {
		fmt.Printf("skipping appending %d as required already has it required=%v reqcontains=%v\n", a, required, requiredcontains)
	}
	fmt.Printf("required a=%d %v\n", a, required)
	return required, requiredcontains
}

func Contains(update []int, i int) bool {
	for _, v := range update {
		if i == v {
			return true
		}
	}
	return false
}

func SortUpdate(rules []PageRule, update []int) []int {

	reqlist := BuildRequiresList(rules)
	updatecontains := map[int]bool{}

	for _, v := range update {
		updatecontains[v] = true
	}

	sorted := []int{}
	for _, i := range update {
		irequired, _ := GetRequiredPages(i, reqlist, sorted, map[int]bool{}, updatecontains)
		for _, j := range irequired {
			if !Contains(sorted, j) && Contains(update, j) {
				sorted = append(sorted, j)
			}
		}
	}

	if !ValidateUpdate(rules, sorted) {
		panic("failed validation on sorted")
	}
	return sorted
}
