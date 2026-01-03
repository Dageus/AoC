package main

import (
	"bufio"
	"os"
)

const (
	PAPER = '@'
	EMPTY = '.'
)

type Position struct {
	X, Y int
}

func getInput(filename string) map[Position]bool {
	var paperMap map[Position]bool
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	paperMap = make(map[Position]bool)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	j := 0

	for scanner.Scan() {
		text := scanner.Text()

		for i, r := range []rune(text) {
			if r == PAPER {
				paperMap[Position{X: j, Y: i}] = true
			}
		}
		j++
	}
	return paperMap
}

func validPaper(pos Position, paperMap map[Position]bool) bool {
	neighs := []Position{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	adjPapers := 0

	for _, d := range neighs {
		neighborPos := Position{
			X: pos.X + d.X,
			Y: pos.Y + d.Y,
		}

		if paperMap[neighborPos] {
			adjPapers++
		}
	}
	return adjPapers < 4
}

func partOne(filename string) int {
	paperMap := getInput(filename)

	sum := 0

	for pos := range paperMap {
		if validPaper(pos, paperMap) {
			sum++
		}
	}
	return sum
}

func partTwo(filename string) int {
	paperMap := getInput(filename)

	sum := 0

	for {
		var toRemove []Position

		for pos := range paperMap {
			if validPaper(pos, paperMap) {
				toRemove = append(toRemove, pos)
			}
		}

		if len(toRemove) == 0 {
			break
		}

		sum += len(toRemove)
		for _, pos := range toRemove {
			delete(paperMap, pos)
		}
	}
	return sum
}
