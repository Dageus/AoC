package main

import (
	"bufio"
	"os"
)

func getInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
}

func partOne(filename string) {
	getInput(filename)
}

func main() {
	partOne("input")
}
