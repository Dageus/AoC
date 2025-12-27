package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	WHITE = 'w'
	BLUE  = 'u'
	BLACK = 'b'
	RED   = 'r'
	GREEN = 'g'
)

var cache = make(map[string]uint64)

func getInput(filename string) (towels []string, patterns []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	towels = strings.Split(scanner.Text(), ", ")

	scanner.Scan()
	for scanner.Scan() {
		patterns = append(patterns, scanner.Text())
	}

	return towels, patterns
}

func patternIsPossible(towels []string, pattern []rune) uint64 {
	patternKey := string(pattern)

	if len(pattern) == 0 {
		return uint64(1)
	}

	if result, exists := cache[patternKey]; exists {
		return result
	}

	numPos := uint64(0)

	for _, towel := range towels {
		if len(towel) > len(pattern) {
			continue
		}
		match := true
		for i := 0; i < len(towel); i++ {
			if rune(towel[i]) != pattern[i] {
				match = false
				break
			}
		}

		if match {
			numPos += patternIsPossible(towels, pattern[len(towel):])
		}
	}
	cache[patternKey] = numPos
	return numPos
}

func countPossiblePatterns(towels, patterns []string) uint64 {
	sum := uint64(0)
	for _, pattern := range patterns {
		sum += patternIsPossible(towels, []rune(pattern))
	}
	return sum
}

func partOne(filename string) int {
	towels, patterns := getInput(filename)
	return int(countPossiblePatterns(towels, patterns))
}
