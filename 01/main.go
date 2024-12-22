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

	l1 := []int{}
	l2 := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		n := strings.Split(line, "   ")
		if len(n) != 2 {
			log.Fatal("line doesn't have two numbers")
		}
		n1, err := strconv.Atoi(n[0])
		if err != nil {
			fmt.Println(line)
			log.Fatal("failed to convert n1 to int")
		}
		n2, err := strconv.Atoi(n[1])
		if err != nil {
			fmt.Println(line)
			log.Fatal("failed to convert n1 to int")
		}
		l1 = append(l1, n1)
		l2 = append(l2, n2)
	}
	fmt.Printf("parsed lists of len1=%d len2=%d\n", len(l1), len(l2))
	if len(l1) != len(l2) {
		log.Fatal("list lengths don't match")
	}

	fmt.Println("sorting list 1...")
	l1s := sort(l1)
	fmt.Println("sorting list 2...")
	l2s := sort(l2)

	fmt.Println("adding distances")
	total_distance := 0

	l2_counts := map[int]int{}
	for i := range l1s {
		n2 := l2s[i]
		l2_counts[n2] = l2_counts[n2] + 1 // count occurences
		d := l2s[i] - l1s[i]
		if d < 0 {
			d = -d // we have abs() at home
		}
		total_distance = total_distance + d
		fmt.Println("i=", i, " l1 ", l1s[i], " l2 ", l2s[i], " d=", d)
	}
	fmt.Println("------------------")
	fmt.Printf("total distance: %d\n", total_distance)

	similarity_score := 0
	for _, n1 := range l1s {
		n1_counts := l2_counts[n1]
		n1_sim := n1 * n1_counts
		similarity_score = similarity_score + n1_sim
	}
	fmt.Println("Similarity Score: ", similarity_score)
}

func sort(l []int) []int {
	if len(l) <= 1 {
		return l
	}
	i_split := len(l) / 2
	fmt.Println("spliting at i_split ", i_split)
	h1 := l[0:i_split]
	h1_sorted := sort(h1)

	h2 := l[i_split:]
	h2_sorted := sort(h2)

	l_sorted := []int{}
	for len(h1_sorted) > 0 || len(h2_sorted) > 0 {
		if len(h1_sorted) == 0 {
			l_sorted = append(l_sorted, h2_sorted...)
			break
		}
		if len(h2_sorted) == 0 {
			l_sorted = append(l_sorted, h1_sorted...)
			break
		}
		if h1_sorted[0] < h2_sorted[0] {
			l_sorted = append(l_sorted, h1_sorted[0])
			h1_sorted = h1_sorted[1:]
		} else {
			l_sorted = append(l_sorted, h2_sorted[0])
			h2_sorted = h2_sorted[1:]
		}
	}
	return l_sorted
}
