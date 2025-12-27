package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Instruction struct {
	Type  int // 0 = mul, 1 = do, -1 = dont
	Start int
	Num1  int
	Num2  int
}

func getInputPartOne(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	return sanitizeInput(scanner)
}

func sanitizeInput(scanner *bufio.Scanner) int {
	pattern := `mul\(\d{1,3},\d{1,3}\)`
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal("Error compiling regex:", err)
		return -1
	}

	validMuls := []string{}

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		validMuls = append(validMuls, re.FindAllString(scanner.Text(), -1)...)
	}
	fmt.Println(validMuls)
	return sumMuls(validMuls)
}

func sumMuls(validMuls []string) int {
	sum := 0
	pattern := `\d+` // Matches sequences of one or more digits

	// Compile the regex
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return -1
	}

	for _, validMul := range validMuls {
		matches := strings.Split(validMul, ",")
		num1, err := strconv.Atoi(re.FindString(matches[0]))
		num2, err := strconv.Atoi(re.FindString(matches[1]))
		if err != nil {
			log.Fatal(err)
		}
		sum += num1 * num2
	}
	return sum
}

func partOne(filename string) int {
	sum := 0
	sum += getInputPartOne(filename)
	return sum
}

func getInputPartTwo(filename string) int {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return sanitizeInputWithKeywords(file)
}

func sanitizeInputWithKeywords(input []byte) int {
	pattern := `mul\(\d{1,3},\d{1,3}\)`
	valid := `do\(\)`
	invalid := `don't\(\)`
	num := `\d+`
	mulRex := regexp.MustCompile(pattern)
	doRex := regexp.MustCompile(valid)
	dontRex := regexp.MustCompile(invalid)
	numRex := regexp.MustCompile(num)

	instructions := []Instruction{}

	for _, match := range mulRex.FindAllSubmatchIndex(input, -1) {
		matches := strings.Split(string(input[match[0]:match[1]]), ",")
		num1, err := strconv.Atoi(numRex.FindString(matches[0]))
		num2, err := strconv.Atoi(numRex.FindString(matches[1]))
		if err != nil {
			log.Fatal(err)
		}
		instructions = append(instructions, Instruction{
			Type:  0,
			Start: match[0],
			Num1:  num1,
			Num2:  num2,
		})
	}

	for _, match := range doRex.FindAllSubmatchIndex(input, -1) {
		instructions = append(instructions, Instruction{
			Type:  1,
			Start: match[0],
			Num1:  0,
			Num2:  0,
		})
	}

	for _, match := range dontRex.FindAllSubmatchIndex(input, -1) {
		instructions = append(instructions, Instruction{
			Type:  -1,
			Start: match[0],
			Num1:  0,
			Num2:  0,
		})
	}

	sort.Slice(instructions, func(i, j int) bool { return instructions[i].Start < instructions[j].Start })

	return sumMulsWithKeywords(instructions)
}

func sumMulsWithKeywords(instructions []Instruction) int {
	sum := 0
	isValid := true

	for _, instruction := range instructions {
		switch instruction.Type {
		case 1:
			isValid = true
		case 0:
			if isValid {
				sum += instruction.Num1 * instruction.Num2
			}
		case -1:
			isValid = false
		}
	}
	return sum
}

func partTwo(filename string) int {
	sum := 0
	sum += getInputPartTwo(filename)
	return sum
}
