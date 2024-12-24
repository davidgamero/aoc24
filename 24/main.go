package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operator string

const (
	AND Operator = "AND"
	OR  Operator = "OR"
	XOR Operator = "XOR"
)

type Rule struct {
	wire1    string
	wire2    string
	operator Operator
}

func GetValue(wire string, wires map[string]int, rules map[string]Rule) int {
	if value, ok := wires[wire]; ok {
		return value
	}
	rule := rules[wire]
	value1 := GetValue(rule.wire1, wires, rules)
	value2 := GetValue(rule.wire2, wires, rules)
	switch rule.operator {
	case AND:
		return value1 & value2
	case OR:
		return value1 | value2
	case XOR:
		return value1 ^ value2
	default:
		panic("unknown operator")
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	rules := make(map[string]Rule) // map from wire to rule required to produce a signal on that wire
	wires := make(map[string]int)  // map from wire to signal value

	for _, line := range lines {
		line := strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			wire := strings.TrimSpace(parts[0])
			valueString := strings.TrimSpace(parts[1])
			value, err := strconv.Atoi(valueString)
			if err != nil {
				panic(err)
			}
			wires[wire] = value
			continue
		}
		if strings.Contains(line, "->") {
			parts := strings.Split(line, " ")

			wire1 := strings.TrimSpace(parts[0])
			operator := strings.TrimSpace(parts[1])
			wire2 := strings.TrimSpace(parts[2])
			outWire := strings.TrimSpace(parts[4])

			var op Operator
			switch operator {
			case "AND":
				op = AND
			case "OR":
				op = OR
			case "XOR":
				op = XOR
			default:
				panic("unknown operator")
			}

			newRule := Rule{wire1, wire2, op}
			fmt.Printf("rule %s = %v\n", outWire, newRule)
			rules[outWire] = newRule
		}
	}

	// now we have all the rules and wires, we can start processing the rules
	for wire := range rules {
		// we only care about getting values for z wire
		if !strings.HasPrefix(wire, "z") {
			continue
		}

		// get the value for this wire
		v := GetValue(wire, wires, rules)
		wires[wire] = v
	}

	// build a string of all the values for wires starting with z
	zBase2String := ""
	for i := 99; i >= 0; i-- {
		numberWith2Digits := fmt.Sprintf("%02d", i)
		wire := "z" + numberWith2Digits
		if value, ok := wires[wire]; ok {
			zBase2String += strconv.Itoa(value)
		}
	}

	fmt.Printf("p1: zBase2String = %s\n", zBase2String)
	// convert the string to an int
	zBase2, err := strconv.ParseInt(zBase2String, 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("p1: zBase2 = %d\n", zBase2)
}
