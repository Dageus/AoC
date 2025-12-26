package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BLINKS = 25
const BIG_BLINKS = 75

type memoKey struct {
	num    int
	blinks int
}

var cache = make(map[memoKey][]int)

var stones []int

func getInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	if scanner.Scan() {
		populateList(scanner.Text())
	}
}

func populateList(text string) {
	els := strings.Split(text, " ")
	for _, el := range els {
		num, err := strconv.Atoi(el)
		if err != nil {
			panic(err)
		}
		stones = append(stones, num)
	}
}

func processStone(num int, blinks int) []int {
	if blinks == 0 {
		return []int{num}
	}

	// Check cacheization map
	key := memoKey{num: num, blinks: blinks}
	if result, exists := cache[key]; exists {
		return result
	}

	// Apply transformation rules
	var result []int
	if num == 0 {
		result = processStone(1, blinks-1)
	} else {
		strNum := strconv.Itoa(num)
		if len(strNum)%2 == 0 {
			// Split number into two halves
			left, _ := strconv.Atoi(strNum[:len(strNum)/2])
			right, _ := strconv.Atoi(strNum[len(strNum)/2:])
			result = append(result, processStone(left, blinks-1)...)
			result = append(result, processStone(right, blinks-1)...)
		} else {
			// Multiply by 2024
			newNum := num * 2024
			result = processStone(newNum, blinks-1)
		}
	}

	// Store result in cacheization map
	cache[key] = result
	return result
}

func partOne(filename string) {
	getInput(filename)
	result := []int{}
	for _, stone := range stones {
		result = append(result, processStone(stone, BLINKS)...)
	}
	fmt.Println("length:", len(result))
}

func processLine() {
	result := []int{}
	for _, stone := range stones {
		result = append(result, processStone(stone, BIG_BLINKS)...)
	}
	fmt.Println("length:", len(result))
}

func partTwo(filename string) {
	getInput(filename)
	processLine()
}

func main() {
	// partOne("input2")
	partTwo("input")
}
