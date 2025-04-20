package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	MODULO            = 16777216
	NUMBERS_GENERATED = 2000
)

var buyers []int
var globalResults = make(map[[4]int]int)

func getInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		buyers = append(buyers, num)
	}
}

func nextRandomNumber(n int) int {
	temp := n * 64
	n ^= temp
	n = n % MODULO

	temp = int(math.Floor(float64(n / 32)))
	n ^= temp
	n = n % MODULO

	temp = n * 2048
	n ^= temp
	n = n % MODULO

	return n
}

func partOne(filename string) {
	getInput(filename)
	sum := 0
	for _, buyer := range buyers {
		secret := buyer
		for range NUMBERS_GENERATED {
			secret = nextRandomNumber(secret)
		}
		sum += secret
	}
	fmt.Println("sum:", sum)
}

func getPrice(secret int) int {
	return secret % 10
}

func partTwo(filename string) {
	getInput(filename)

	seen := make(map[[4]int]bool)
	for _, buyer := range buyers {
		secret0 := buyer
		secret1 := nextRandomNumber(secret0)
		secret2 := nextRandomNumber(secret1)
		secret3 := nextRandomNumber(secret2)

		for i := 3; i < NUMBERS_GENERATED; i++ {
			secret := nextRandomNumber(secret3)
			price := getPrice(secret)
			sequence := [4]int{getPrice(secret0) - getPrice(secret1), getPrice(secret1) - getPrice(secret2), getPrice(secret2) - getPrice(secret3), getPrice(secret3) - getPrice(secret)}
			if _, exists := seen[sequence]; !exists {
				seen[sequence] = true
				if _, exists := globalResults[sequence]; exists {
					// fmt.Println("price was", globalResults[sequence], "| is now", globalResults[sequence]+price)
					globalResults[sequence] += price
				} else {
					globalResults[sequence] = price
				}
			}
			secret0, secret1, secret2, secret3 = secret1, secret2, secret3, secret
		}
		seen = make(map[[4]int]bool)
	}
	p2 := math.MinInt
	for _, v := range globalResults {
		if v > p2 {
			p2 = v
		}
	}
	fmt.Println("max bananas:", p2)
}

func main() {
	// partOne("input")
	partTwo("input")
}
