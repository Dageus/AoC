package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func getInput(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var unrolledProblems [][]string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		unrolledProblems = append(unrolledProblems, nums)
	}

	return unrolledProblems
}

func getInput2(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var problemMatrix [][]rune

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		problemMatrix = append(problemMatrix, []rune(scanner.Text()))
	}

	return problemMatrix
}

func calculateOperation(curNums []int, operator rune) int {
	res := 0
	switch operator {
	case '+':
		for _, num := range curNums {
			res += num
		}
	case '*':
		res = 1
		for _, num := range curNums {
			res *= num
		}
	default:
		panic("err with wrong operator")
	}
	return res
}

func partOne(filename string) int {
	sum := 0
	input := getInput(filename)

	rows := len(input)
	cols := len(input[0])

	for c := range cols {
		var curNums []int
		for r := range rows {
			if input[r][c] == "+" || input[r][c] == "*" {
				sum += calculateOperation(curNums, rune(input[r][c][0]))
			} else {
				num, _ := strconv.Atoi(input[r][c])
				curNums = append(curNums, num)
			}
		}
	}

	return sum
}

func getColumn(matrix [][]rune, col int, maxRow int) (string, rune, bool) {
	empty := true

	var sb strings.Builder

	for r := 0; r < maxRow-1; r++ {
		if matrix[r][col] != ' ' {
			empty = false
			sb.WriteRune(matrix[r][col])
		}
	}

	op := matrix[maxRow-1][col]

	return sb.String(), op, empty
}

func partTwo(filename string) int {
	sum := 0
	matrix := getInput2(filename)

	maxCol := len(matrix[0])
	maxRow := len(matrix)

	var numberStack []int
	var currOp rune

	for c := maxCol - 1; c >= 0; c-- {
		digits, opChar, empty := getColumn(matrix, c, maxRow)
		if empty {
			sum += calculateOperation(numberStack, rune(currOp))

			numberStack = []int{}
			currOp = -1
		} else {
			num, _ := strconv.Atoi(digits)
			numberStack = append(numberStack, num)

			if opChar != ' ' {
				currOp = opChar
			}
		}
	}

	sum += calculateOperation(numberStack, rune(currOp))

	return sum
}
