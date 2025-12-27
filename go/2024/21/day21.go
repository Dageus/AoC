package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	NUMPAD = [][]rune {
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{'-', '0', 'A'},
}
	KEYPAD = [][]rune{
	{'-', '^', 'A'},
	{'<', 'v', '>'},
	}
)

func getInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	codes := []string{}

	for scanner.Scan() {
		codes = append(codes, scanner.Text())
	}
	return codes
}

func findShortestPath(code string) int {

}

func partOne(filename string) int {
	codes := getInput(filename)

	complexity := 0

	for _, code := range codes {
		len := findShortestPath(code)
		complexity += len * ?
	}
	return complexity
}
