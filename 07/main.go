package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	Result   int
	Sequence []int
}

func line2eq(line string) (Equation, error) {
	eq := Equation{}
	colonsplit := strings.Split(line, ":")
	if len(colonsplit) != 2 {
		return eq, fmt.Errorf("parsing line, unexpected number of colon-delimited segments %d in line '%s', expected 2", len(colonsplit), line)
	}

	result, err := strconv.Atoi(colonsplit[0])
	if err != nil {
		return eq, fmt.Errorf("parsing line, unable to convert result string '%s' to int: %w", colonsplit[0], err)
	}

	sequencestrings := strings.Split(colonsplit[1], " ")

	sequence := []int{}

	for _, s := range sequencestrings {
		if len(s) == 0 {
			continue
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			return eq, fmt.Errorf("splitting sequence '%s' by whitespaces: %w", colonsplit[1], err)
		}
		sequence = append(sequence, i)
	}

	return Equation{
		Result:   result,
		Sequence: sequence,
	}, nil
}

func Concat(a int, b int) int {
	bdigits := len(strconv.Itoa(b))
	return (a * int(math.Pow10(bdigits))) + b
}

func CanBeCreated(e Equation, withConcat bool) (bool, error) {
	values := map[int]bool{e.Sequence[0]: true} // the values the operation could take at this point

	//fmt.Printf("testing %v\n", e)
	for _, nextnumberinsequence := range e.Sequence[1:] {
		//fmt.Printf("  values: %v\n", values)
		newpotentialvalues := map[int]bool{}

		for currentvalue := range values {
			addvalue := nextnumberinsequence + currentvalue
			if addvalue <= e.Result {
				newpotentialvalues[addvalue] = true
			}
			multvalue := nextnumberinsequence * currentvalue
			if multvalue <= e.Result {
				newpotentialvalues[multvalue] = true
			}
			concatvalue := Concat(currentvalue, nextnumberinsequence)
			if withConcat && concatvalue <= e.Result {
				newpotentialvalues[concatvalue] = true
			}
		}
		if len(newpotentialvalues) == 0 {
			return false, nil
		}
		// bump to new potentials for this entry
		values = newpotentialvalues
	}
	return values[e.Result], nil
}

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	equations := []Equation{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		eq, err := line2eq(line)
		if err != nil {
			panic(err)
		}
		equations = append(equations, eq)
	}

	fmt.Printf("loaded %d equations from %d lines\n", len(equations), len(lines))

	totalcalibrationresult := 0
	totalcalibrationresultconcat := 0
	for _, e := range equations {
		possible, err := CanBeCreated(e, false)
		possiblewithconcat, err := CanBeCreated(e, true)
		if err != nil {
			panic(err)
		}
		if possible {
			totalcalibrationresult += e.Result
		}
		if possiblewithconcat {
			totalcalibrationresultconcat += e.Result
		}
	}
	fmt.Printf("p1 total calibration result: %d\n", totalcalibrationresult)
	fmt.Printf("p2 total calibration result: %d\n", totalcalibrationresultconcat)
}
