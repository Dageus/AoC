package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Vertex struct {
	x, y     int
	distance int
}

type Cheat struct {
	startX, startY int
	endX, endY     int
}

var (
	startX, startY int
	endX, endY     int
	racetrack      [][]rune
	uniqueCheats   = make(map[Cheat]bool)
	optimal_paths  int
)

const (
	TRACK      = '.'
	WALL       = '#'
	START      = 'S'
	END        = 'E'
	SAVED_TIME = 100
)

func getInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	r := 0
	for scanner.Scan() {
		line := []rune{}
		for c, char := range scanner.Text() {
			if char == START {
				startX, startY = c, r
			}
			if char == END {
				endX, endY = c, r
			}
			line = append(line, char)
		}
		racetrack = append(racetrack, line)
		r++
	}
}

func findPath() (minScore int) {
	rows, cols := len(racetrack), len(racetrack[0])
	visited := make(map[[2]int]bool)

	directions := []struct{ dx, dy int }{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
	minScore = math.MaxInt

	queue := []Vertex{
		{startX, startY, 0},
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.x == endX && curr.y == endY {
			minScore = min(minScore, curr.distance)
			continue
		}
		visited[[2]int{curr.x, curr.y}] = true

		for _, d := range directions {
			nx, ny := curr.x+d.dx, curr.y+d.dy
			if nx >= 0 && nx < cols && ny >= 0 && ny < rows && !visited[[2]int{nx, ny}] {
				if racetrack[ny][nx] != WALL {
					queue = append(queue, Vertex{nx, ny, curr.distance + 1})
				}
			}
		}
	}
	return minScore
}

func canRemoveWall(r, c int) bool {
	if racetrack[r][c] != WALL {
		return false
	}
	return r > 0 && r < len(racetrack)-1 && c > 0 && c < len(racetrack[0])-1 &&
		((racetrack[r-1][c] == TRACK && racetrack[r+1][c] == TRACK) ||
			(racetrack[r][c-1] == TRACK && racetrack[r][c+1] == TRACK))
}

func partOne(filename string) {
	getInput(filename)
	default_time := findPath()
	fmt.Println(default_time)
	for r, line := range racetrack {
		for c, wall := range line {
			if wall == WALL {
				racetrack[r][c] = TRACK
				if time := findPath(); default_time-time >= SAVED_TIME {
					// fmt.Println("this path saves", default_time-time, "picoseconds")
					optimal_paths++
				}
				racetrack[r][c] = WALL
			}
		}
	}
	fmt.Println("possible paths:", optimal_paths)
}

func manhattanDistance(point1, point2 Vertex) int {
	return int(math.Abs(float64(point1.x)-float64(point2.x))) + int(math.Abs(float64(point1.y)-float64(point2.y)))
}

func partTwo(filename string) {
	getInput(filename)
	default_time := findPath()
	fmt.Println(default_time)
	for r, line := range racetrack {
		for c, wall := range line {
			if wall == WALL {
				racetrack[r][c] = TRACK
				if time := findPath(); default_time-time >= SAVED_TIME {
					// fmt.Println("this path saves", default_time-time, "picoseconds")
					optimal_paths++
				}
				racetrack[r][c] = WALL
			}
		}
	}
	fmt.Println("possible paths:", optimal_paths)
}

// func main() {
// 	partOne("input")
// 	partTwo("input")
// }
