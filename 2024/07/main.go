package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var result int

func getInput(fileName string, f func(string)) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		f(scanner.Text())
	}
}

func partOne(fileName string) int {
	getInput(fileName, testLine)
	return result
}

func testLine(line string) {
	el := strings.Split(line, " ")

	target, err := strconv.Atoi(el[0][:len(el[0])-1])
	if err != nil {
		log.Fatal(err)
	}

	args := el[1:]

	// test permutations
	if canAchieveTarget(args, target, args[0], 1) {
		result += target
	}

}

func canAchieveTarget(nums []string, target int, current string, index int) bool {
	if index == len(nums) {
		curr, err := strconv.Atoi(current)
		if err != nil {
			log.Fatal(err)
		}
		return curr == target
	}

	curr, err := strconv.Atoi(current)
	if err != nil {
		log.Fatal(err)
	}

	nums_idx, err := strconv.Atoi(nums[index])
	if err != nil {
		log.Fatal(err)
	}

	// Try adding
	// fmt.Println("Testing", current, "+", nums[index], "==", target, "?")
	if canAchieveTarget(nums, target, strconv.Itoa(curr+nums_idx), index+1) {
		return true
	}

	// Try multiplying
	// fmt.Println("Testing", current, "*", nums[index], "==", target, "?")
	if canAchieveTarget(nums, target, strconv.Itoa(curr*nums_idx), index+1) {
		return true
	}

	// Try concatenating
	if canAchieveTarget(nums, target, current+nums[index], index+1) {
		return true
	}

	return false
}

func partTwo(fileName string) int {
	getInput(fileName, testLine)
	return 0
}

func main() {
	fmt.Println("Result: ", partOne("input"))
}
