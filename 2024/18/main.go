package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	SIMULATION_LIMIT = 1024
	// SIMULATION_LIMIT = 12
	DIMENSION = 71
	// DIMENSION = 7
	CORRUPTED = 'O'
	FREE      = '.'
	UP        = 0
	RIGHT     = 1
	DOWN      = 2
	LEFT      = 3
)

type Position struct {
	x, y int
}

type Direction struct {
	dx, dy int
}

type State struct {
	user  Position
	path  []Position
	score int
}

var memory [][]rune
var corrupted []Position

var directions = []struct {
	dx, dy int
}{
	{1, 0},  // East
	{0, -1}, // North
	{-1, 0}, // West
	{0, 1},  // South
}

func getInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		corrupted = append(corrupted, Position{x, y})
	}
}

func BFS(startPos, endPos Position) (minScore int) {
	minScore = math.MaxInt
	queue := []Position{
		startPos,
	}
	dist := make(map[Position]int)

	for len(queue) > 0 {
		currState := queue[0]
		queue = queue[1:]
		for _, n := range getNeighbours(currState) {
			if outsideMemory(n) || memory[n.x][n.y] == CORRUPTED {
				continue
			}
			if _, has := dist[n]; has {
				continue
			}
			dist[n] = dist[currState] + 1
			if n == endPos {
				return dist[n]
			}

			queue = append(queue, n)
		}
	}

	return minScore
}

func outsideMemory(pos Position) bool {
	return pos.x < 0 || pos.y < 0 || pos.x >= DIMENSION || pos.y >= DIMENSION
}

func getNeighbours(position Position) (neighbours []Position) {
	neighbours = make([]Position, 0, 4)

	for _, dir := range directions {
		nIdx := Position{x: position.x + dir.dx, y: position.y + dir.dy}
		neighbours = append(neighbours, nIdx)
	}

	return
}

func partOne(filename string) {
	getInput(filename)

	for i := range SIMULATION_LIMIT {
		pos := corrupted[i]
		memory[pos.y][pos.x] = CORRUPTED
	}
	goal := Position{DIMENSION - 1, DIMENSION - 1}
	start := Position{0, 0}
	path := BFS(start, goal)
	fmt.Println("smallest path:", path)
}

func partTwo(filename string) {
	getInput(filename)
	goal := Position{DIMENSION - 1, DIMENSION - 1}
	start := Position{0, 0}
	for i := SIMULATION_LIMIT; i < len(corrupted); i++ {
		pos := corrupted[i]
		memory[pos.y][pos.x] = CORRUPTED
		if BFS(start, goal) == math.MaxInt {
			fmt.Println("first position that breaks path:", pos)
			break
		}
	}
}

func main() {
	for range DIMENSION {
		memory = append(memory, make([]rune, DIMENSION))
	}
	partOne("input")

	partTwo("input")
}
