package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Program struct {
	A, B, C            uint64
	instructionPointer int
}

func getCombo(program *Program, operand int) uint64 {
	var val uint64
	switch operand {
	case 4:
		val = program.A
	case 5:
		val = program.B
	case 6:
		val = program.C
	default:
		val = uint64(operand)
	}
	return val
}

func getInput(filename string) (Program, []int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	program := Program{}
	instructionStack := []int{}

	scanner.Scan()
	tempA, _ := strconv.Atoi(strings.Split(scanner.Text(), ": ")[1])
	program.A = uint64(tempA)

	scanner.Scan()
	tempB, _ := strconv.Atoi(strings.Split(scanner.Text(), ": ")[1])
	program.B = uint64(tempB)

	scanner.Scan()
	tempC, _ := strconv.Atoi(strings.Split(scanner.Text(), ": ")[1])
	program.C = uint64(tempC)

	program.instructionPointer = 0

	scanner.Scan()
	scanner.Scan()
	nums := strings.Split(strings.Split(scanner.Text(), ": ")[1], ",")
	for _, num := range nums {
		val, _ := strconv.Atoi(num)
		instructionStack = append(instructionStack, val)
	}

	return program, instructionStack
}

func execProgram(program *Program, instructions []int) []int {
	program.instructionPointer = 0
	output := []int{}
	for program.instructionPointer < len(instructions) {
		instruction := instructions[program.instructionPointer]
		operand := instructions[program.instructionPointer+1]

		program.instructionPointer += 2
		switch instruction {
		case 0:
			program.A >>= getCombo(program, operand)
		case 1:
			program.B ^= uint64(operand)
		case 2:
			program.B = getCombo(program, operand) % 8
		case 3:
			if program.A != 0 {
				program.instructionPointer = operand
			}
		case 4:
			program.B ^= program.C
		case 5:
			val := getCombo(program, operand) % 8
			output = append(output, int(val))
		case 6:
			program.B = program.A >> getCombo(program, operand)
		case 7:
			program.C = program.A >> getCombo(program, operand)
		}
	}
	return output
}

func partOne(filename string) {
	program, instructions := getInput(filename)
	fmt.Println(program.A)
	fmt.Println(program.B)
	fmt.Println(program.C)
	fmt.Println(instructions)

	output := execProgram(&program, instructions)

	s := ""
	for i, v := range output {
		if i != 0 {
			s += fmt.Sprintf(",%d", v)
		} else {
			s += fmt.Sprintf("%d", v)
		}
	}
	fmt.Printf("%s\n", s)
}

func calc(program *Program, p []int, n []int, testA uint64, l int) uint64 {
	if len(n) == 0 {
		return testA
	}

	for k := uint64(0); k < 8; k++ {
		a := (k << ((3 * l) + 7)) | testA

		program.A = a >> (3 * l)
		program.B = 0
		program.C = 0
		n1 := execProgram(program, p)
		if n1[0] == n[0] {
			r := calc(program, p, n[1:], a, l+1)
			if r > 0 {
				return r
			}
		}
	}
	return 0
}

func partTwo(filename string) {
	program, instructions := getInput(filename)
	testA := uint64(0)
	for k := uint64(0); k < 128; k++ {
		a := k
		r := calc(&program, instructions, instructions, a, 0)
		if r > 0 {
			testA = r
			break
		}

	}
	fmt.Printf("%d\n", testA)
}

func main() {
	partOne("input")
	partTwo("input")
}
