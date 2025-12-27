package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Operation struct {
	var1, var2 string
	operation  func(int, int) int
	output     string
	typ        int
}

const (
	OPERATOR = " -> "
	AND      = 0
	OR       = 1
	XOR      = 2
)

var (
	variables  = make(map[string]int)
	operations []Operation
)

func getInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	operations := false
	for scanner.Scan() {
		if scanner.Text() == "" {
			operations = true
			continue
		}

		if operations {
			parseOperation(scanner.Text())
		} else {
			parseVariable(scanner.Text())
		}
	}
}

func parseVariable(text string) {
	strs := strings.Split(text, ": ")
	variable := strs[0]
	val, _ := strconv.Atoi(strs[1])
	variables[variable] = val
}

func parseOperation(text string) {
	expression := strings.Split(text, OPERATOR)
	output := expression[1]
	input := expression[0]
	vars := strings.Split(input, " ")
	operation := Operation{}
	operation.var1 = vars[0]
	operation.var2 = vars[2]
	if operation.var2 > operation.var1 {
		operation.var1, operation.var2 = operation.var2, operation.var1
	}
	operation.output = output
	switch vars[1] {
	case "OR":
		operation.operation = or
		operation.typ = OR
	case "XOR":
		operation.operation = xor
		operation.typ = XOR
	case "AND":
		operation.operation = and
		operation.typ = AND
	}
	operations = append(operations, operation)
}

func xor(var1, var2 int) int {
	return var1 ^ var2
}

func or(var1, var2 int) int {
	return var1 | var2
}

func and(var1, var2 int) int {
	return var1 & var2
}

func getValue(variable string) int {
	if val, exists := variables[variable]; exists {
		return val
	}
	panic("Variable not found: " + variable)
}

func evaluate(operation Operation) {
	val1 := resolveValue(operation.var1)
	val2 := resolveValue(operation.var2)
	variables[operation.output] = operation.operation(val1, val2)
}

func resolveValue(varName string) int {
	if val, exists := variables[varName]; exists {
		return val
	}

	for _, op := range operations {
		if op.output == varName {
			val1 := resolveValue(op.var1)
			val2 := resolveValue(op.var2)

			result := op.operation(val1, val2)
			variables[varName] = result
			return result
		}
	}

	// Handle the case where varName is a number
	if num, err := strconv.Atoi(varName); err == nil {
		return num
	}

	panic("Variable not found: " + varName)
}

func simulate() {
	for _, operation := range operations {
		evaluate(operation)
	}
}

func getOutput() int {
	result := ""
	for i := 0; ; i++ {
		wire := fmt.Sprintf("z%02d", i)
		if val, exists := variables[wire]; exists {
			result = strconv.Itoa(val) + result
		} else {
			break
		}
	}
	decimal, _ := strconv.ParseInt(result, 2, 64)
	return int(decimal)
}

func partOne(filename string) int {
	getInput(filename)
	simulate()
	return getOutput()
}

func partTwo(filename string) int {
	getInput(filename)

	var edges []string
	var andID, orID, xorID int
	for _, gate := range operations {
		var gateName string
		var gateID int

		switch gate.typ {
		case AND:
			gateID = andID
			gateName = "AND"
			andID++
		case OR:
			gateID = orID
			orID++
			gateName = "OR"
		case XOR:
			gateID = xorID
			xorID++
			gateName = "XOR"
		}

		edges = append(edges, fmt.Sprintf("%s --> %s%d(\"%s\") --> %s", gate.var1, gateName, gateID, gateName, gate.output))
		edges = append(edges, fmt.Sprintf("%s --> %s%d(\"%s\")", gate.var2, gateName, gateID, gateName))
	}

	slices.SortFunc(edges, func(a, b string) int {
		if a > b {
			return -1
		}
		return 1
	})
	for _, e := range edges {
		fmt.Println(e)
	}

	// Solution was inspired by r/adventofcode
	// Visualize solution in https://mermaid.live/edit
}
