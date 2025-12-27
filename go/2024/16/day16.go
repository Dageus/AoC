package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	TURN  = 1000
	MOVE  = 1
	WALL  = '#'
	START = 'S'
	END   = 'E'
	UP    = 0
	RIGHT = 1
	DOWN  = 2
	LEFT  = 3
)

var (
	maze           [][]rune
	startX, startY int
	endX, endY     int
	validPaths     []int
)

type Position struct {
	x, y int
}

type Direction struct {
	dx, dy int
}

type Node struct {
	idx Position
	dir Direction
}

type State struct {
	reindeer Node
	path     []Position
	score    int
}

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
		maze = append(maze, []rune(scanner.Text()))
		for c, char := range scanner.Text() {
			if char == START {
				startX = c
				startY = len(maze) - 1
			} else if char == END {
				endX = c
				endY = len(maze) - 1
			}
		}
	}
}

func BFS(startPos, endPos Position) (minScore int, bestSeatCount int) {
	minScore = math.MaxInt
	reindeer := Node{idx: startPos, dir: Direction{dx: 0, dy: 1}}
	queue := []State{
		{
			reindeer: reindeer,
			path:     []Position{startPos},
			score:    0,
		},
	}
	visited := make(map[Node]int)
	sizeToIndices := make(map[int][]Position)

	for len(queue) > 0 {
		currState := queue[0]
		queue = queue[1:]

		if currState.score > minScore {
			continue
		}

		if currState.reindeer.idx == endPos {
			if currState.score <= minScore {
				minScore = currState.score
				sizeToIndices[minScore] = append(sizeToIndices[minScore], currState.path...)
			}
			continue
		}

		for _, n := range getNeighbours(currState.reindeer) {
			if maze[n.idx.x][n.idx.y] == WALL {
				continue
			}
			score := currState.score + 1
			if currState.reindeer.dir != n.dir {
				score += 1000
			}
			if previous, has := visited[n]; has {
				if previous < score {
					continue
				}
			}
			visited[n] = score

			nPath := make([]Position, len(currState.path))
			copy(nPath, currState.path)

			queue = append(queue, State{
				reindeer: n,
				path:     append(nPath, n.idx),
				score:    score,
			})
		}
	}

	countMap := make(map[Position]bool)
	for _, index := range sizeToIndices[minScore] {
		countMap[index] = true
	}

	return minScore, len(countMap)
}

func getNeighbours(reindeer Node) (neighbours []Node) {
	neighbours = make([]Node, 0, 4)
	currDir, currIdx := reindeer.dir, reindeer.idx
	oppositeDir := Direction{dx: -currDir.dx, dy: -currDir.dy}

	for _, dir := range directions {
		if dir == oppositeDir {
			continue
		}
		nIdx := Position{x: currIdx.x + dir.dx, y: currIdx.y + dir.dy}
		neighbours = append(neighbours, Node{idx: nIdx, dir: dir})
	}

	return
}

func displayMaze(x, y int) {

	fmt.Println("-----------------------")

	for r, line := range maze {
		for c, char := range line {
			if r == y && c == x {
				fmt.Print("S")
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}
}

func partOne(filename string) int {
	getInput(filename)
	score, _ := BFS(Position{x: startX, y: startY}, Position{x: endX, y: endY})
	return score
}

func partTwo(filename string) int {
	getInput(filename)
	_, bestTile := BFS(Position{x: startX, y: startY}, Position{x: endX, y: endY})
	return bestTile
}
