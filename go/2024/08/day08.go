package main

import (
	"bufio"
	"log"
	"os"
)

type Coordinate struct {
	R, C int
}

var dimensions Coordinate
var antenas = make(map[rune][]Coordinate)
var antinodes = make(map[Coordinate]bool)

func outsideMap(point Coordinate) bool {
	return point.C < 0 || point.R < 0 || point.C >= dimensions.C || point.R >= dimensions.R
}

func getInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	r := 0
	c := 0
	for scanner.Scan() {
		char_arr := []rune(scanner.Text())
		c = 0
		for _, char := range char_arr {
			if char == '.' {
				c++
				continue
			}

			coordinate := Coordinate{R: r, C: c}
			antenas[char] = append(antenas[char], coordinate)
			// NOTE: remove this for partOne
			antinodes[coordinate] = true
			c++
		}
		r++
	}
	dimensions = Coordinate{R: r, C: c}
}

func calculateAntiNodes(calculateOutside bool) {
	for _, coordinates := range antenas {

		start_idx := 0
		for start_idx != len(coordinates)-1 {
			for idx := start_idx + 1; idx < len(coordinates); idx++ {
				computeAntiNodes(coordinates[start_idx], coordinates[idx], calculateOutside)
			}
			start_idx++
		}
	}
}

func computeAntiNodes(point1 Coordinate, point2 Coordinate, calculateOutside bool) {
	// NOTE: int conversions go brrrr
	vector := Coordinate{
		R: point2.R - point1.R,
		C: point2.C - point1.C,
	}

	antinode1 := Coordinate{
		R: point1.R - vector.R,
		C: point1.C - vector.C,
	}
	antinode2 := Coordinate{
		R: point2.R + vector.R,
		C: point2.C + vector.C,
	}

	if calculateOutside {
		for !outsideMap(antinode1) {
			antinodes[antinode1] = true
			antinode1 = Coordinate{
				R: antinode1.R - vector.R,
				C: antinode1.C - vector.C,
			}
		}

		for !outsideMap(antinode2) {
			antinodes[antinode2] = true
			antinode2 = Coordinate{
				R: antinode2.R + vector.R,
				C: antinode2.C + vector.C,
			}
		}
	}
}

func partOne(filename string) int {
	getInput(filename)
	calculateAntiNodes(false)
	return len(antinodes)
}

func partTwo(filename string) int {
	getInput(filename)
	calculateAntiNodes(true)
	return len(antinodes)
}
