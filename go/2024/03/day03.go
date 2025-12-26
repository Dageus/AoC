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

func getInputPartOne(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	sanitizeInput(scanner)
}

func sanitizeInput(scanner *bufio.Scanner) {
	pattern := `mul\(\d{1,3},\d{1,3}\)`
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal("Error compiling regex:", err)
		return
	}

	validMuls := []string{}

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		validMuls = append(validMuls, re.FindAllString(scanner.Text(), -1)...)
	}
	fmt.Println(validMuls)
	sumMuls(validMuls)
}

func sumMuls(validMuls []string) {
	sum := 0
	pattern := `\d+` // Matches sequences of one or more digits

	// Compile the regex
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
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
	fmt.Println(sum)
}

func partOne(fileName string) {
	getInputPartOne(fileName)
}

func getInputPartTwo(fileName string) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	sanitizeInputWithKeywords(file)
}

func sanitizeInputWithKeywords(input []byte) {
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

	sumMulsWithKeywords(instructions)
}

func sumMulsWithKeywords(instructions []Instruction) {
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
	fmt.Println(sum)
}

func partTwo(fileName string) {
	getInputPartTwo(fileName)
}

func main() {
	partTwo("input")
}
