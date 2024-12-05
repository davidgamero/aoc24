package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mulsum := 0

	mulsumCondition := 0
	doing := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		for i := range line {
			fmt.Printf("%s\n\n", line[i:])
			mulsum = mulsum + getMulValue(line[i:])
			if doing {
				mulsumCondition = mulsumCondition + getMulValue(line[i:])
			}

			if strings.HasPrefix(line[i:], "do()") {
				doing = true
			}
			if strings.HasPrefix(line[i:], "don't()") {
				doing = false
			}
		}
	}

	fmt.Printf("%d mul sum\n", mulsum)
	fmt.Printf("%d mul sum conditional\n", mulsumCondition)
}

// get the first mul value from a string
func getMulValue(s string) int {
	// min len 7 ex: mul(X,Y)
	if len(s) < 8 {
		fmt.Println(" len < 8")
		return 0
	}
	if !strings.HasPrefix(s, "mul(") {
		fmt.Println(" missing 'mul(' prefix")
		return 0
	}

	// value1
	if !unicode.IsDigit(rune(s[4])) {
		fmt.Println(" 4th isn't digit")
		return 0
	}
	v1, err := strconv.Atoi(string(s[4]))
	if err != nil {
		log.Fatal("cant convert digit to int: ", err)
	}
	fmt.Println(" v1=", v1)
	// parse first value optional digits
	icomma := -1
	for _, i := range []int{5, 6} {
		if i >= len(s) {
			fmt.Println(" out of chars at i=", i)
			return 0
		}
		// break on comma
		if s[i] == ',' {
			icomma = i
			break
		}
		// read digits into v1
		if unicode.IsDigit(rune(s[i])) {
			v1 = v1 * 10
			vnew, err := strconv.Atoi(string(s[i]))
			if err != nil {
				log.Fatal("cant convert digit to int: ", err)
			}
			v1 = v1 + vnew
			fmt.Println(" v1=", v1)
			continue
		}
		fmt.Printf(" invalid char '%c' at i=%d\n", rune(s[i]), i)
		return 0
	}
	if icomma == -1 {
		// no comma found, must be 3 digit v1 or invalid
		if 7 < len(s) && s[7] == ',' {
			icomma = 7
		} else {
			fmt.Println(" no comma found")
			return 0
		}
	}

	// parse second value
	iv2 := icomma + 1 // index to start reading value 2
	if iv2 >= len(s) {
		fmt.Println(" out of digits for iv2 at i=", iv2)
		return 0
	}
	if !unicode.IsDigit(rune(s[iv2])) {
		fmt.Println(" not digit at i=", iv2, " '", s[iv2], "'")
		return 0
	}
	v2, err := strconv.Atoi(string(s[iv2]))
	if err != nil {
		log.Fatal("cant convert digit to int for v2: ", err)
	}
	fmt.Println(" v2=", v2)
	iclose := -1
	// parse first value optional digits
	for _, i := range []int{iv2 + 1, iv2 + 2} {
		if i >= len(s) {
			fmt.Println(" out of chars for iv2 at i=", iv2)
			return 0
		}
		// break on close paren
		if s[i] == ')' {
			iclose = i
			break
		}
		// read digits into v2
		if unicode.IsDigit(rune(s[i])) {
			v2 = v2 * 10
			vnew, err := strconv.Atoi(string(s[i]))
			if err != nil {
				log.Fatal("cant convert digit to int: ", err)
			}
			v2 = v2 + vnew
			fmt.Println(" v2=", v2)
			continue
		}
		fmt.Printf(" invalid char '%s' at i=%d\n", string(s[i]), i)
		return 0
	}
	if iclose != -1 {
		return v1 * v2
	}
	if iv2+3 < len(s) && s[iv2+3] == ')' {
		return v1 * v2
	}
	fmt.Println(" closing paren missing")
	return 0
}
