package main

import (
	"bufio"
	"os"
)

const (
	START    = 'S'
	SPLITTER = '^'
	EMPTY    = '.'
)

var (
	max_depth, max_width int
)

func getInput(filename string) ([][]rune, int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var matrix [][]rune

	var start_x int

	for scanner.Scan() {
		text := scanner.Text()
		line := []rune{}
		for i, r := range text {
			line = append(line, r)
			if r == START {
				start_x = i
			}
		}
		matrix = append(matrix, line)
	}

	return matrix, start_x
}

func partOne(filename string) int {
	matrix, start_x := getInput(filename)

	max_depth = len(matrix)
	max_width = len(matrix[0])
	splits := 0

	activeBeams := make(map[int]bool)
	activeBeams[start_x] = true

	for y := range len(matrix) {
		nextBeams := make(map[int]bool)

		for x := range activeBeams {

			if x < 0 || x >= len(matrix[0]) {
				continue
			}

			cell := matrix[y][x]

			switch cell {
			case SPLITTER:
				splits++
				nextBeams[x-1] = true
				nextBeams[x+1] = true

			case START, EMPTY:
				nextBeams[x] = true

			default:
				nextBeams[x] = true
			}
		}

		activeBeams = nextBeams

		if len(activeBeams) == 0 {
			break
		}
	}

	return splits
}

func partTwo(filename string) int {
	matrix, start_x := getInput(filename)

	max_depth = len(matrix)
	max_width = len(matrix[0])
	timelines := 1

	activeBeams := make(map[int]int)
	activeBeams[start_x] = 1

	for y := range len(matrix) {
		nextBeams := make(map[int]int)

		for x, count := range activeBeams {

			if x < 0 || x >= len(matrix[0]) {
				continue
			}

			cell := matrix[y][x]

			switch cell {
			case SPLITTER:
				timelines += count

				nextBeams[x-1] += count
				nextBeams[x+1] += count

			case START, EMPTY:
				nextBeams[x] += count

			default:
				nextBeams[x] += count
			}
		}

		activeBeams = nextBeams

		if len(activeBeams) == 0 {
			break
		}
	}

	return timelines
}

func main() {
	res := partTwo("input")
	println(res)
}
