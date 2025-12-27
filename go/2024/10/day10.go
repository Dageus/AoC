package main

import (
	"bufio"
	"os"
)

type Coordinate struct {
	Value, R, C int
}

type Point struct {
	R, C int
}

var trailheads []Point
var grid [][]int
var dimension int
var directions = []Point{
	{R: -1, C: 0},
	{R: 1, C: 0},
	{R: 0, C: 1},
	{R: 0, C: -1},
}

func getInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	grid = [][]int{}

	r := 0
	for scanner.Scan() {
		// construct grid and map
		row := []int{}
		for c, val := range scanner.Text() {
			num := int(val - '0')
			row = append(row, num)
			if num == 0 {
				trailheads = append(trailheads, Point{R: r, C: c})
			}
		}
		grid = append(grid, row)
		r++
	}
	dimension = len(grid)
}

func findPaths(trailhead Point) int {
	directions := []Point{
		{R: -1, C: 0},
		{R: 1, C: 0},
		{R: 0, C: 1},
		{R: 0, C: -1},
	}

	nines := make(map[Point]bool)

	visited := make(map[Point]bool)
	queue := []Point{trailhead}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if visited[cur] {
			continue
		}

		visited[cur] = true

		if grid[cur.R][cur.C] == 9 {
			nines[cur] = true
		}

		for _, direction := range directions {
			newVal := Point{R: cur.R + direction.R, C: cur.C + direction.C}
			if insideGrid(newVal) && !visited[newVal] {
				if grid[newVal.R][newVal.C] == grid[cur.R][cur.C]+1 {
					queue = append(queue, newVal)
				}
			}
		}
	}
	return len(nines)
}

func insideGrid(point Point) bool {
	if point.R >= 0 && point.R < dimension && point.C >= 0 && point.C < dimension {
		return true
	}
	return false
}

func partOne(filename string) int {
	getInput(filename)
	sum := 0
	for _, trailhead := range trailheads {
		sum += findPaths(trailhead)
	}
	return sum
}

func findDistinctPaths(point Point, visited map[Point]bool) int {

	if grid[point.R][point.C] == 9 {
		return 1
	}

	visited[point] = true
	// this basically represents a function call that puts visited at false so we can backtrack
	defer func() { visited[point] = false }()

	paths := 0

	for _, direction := range directions {
		newVal := Point{R: point.R + direction.R, C: point.C + direction.C}
		if insideGrid(newVal) && !visited[newVal] &&
			grid[newVal.R][newVal.C] == grid[point.R][point.C]+1 {
			paths += findDistinctPaths(newVal, visited)
		}
	}
	return paths
}

func partTwo(filename string) int {
	getInput(filename)
	totalScore := 0

	visited := make(map[Point]bool)

	for _, trailhead := range trailheads {
		totalScore += findDistinctPaths(trailhead, visited)
	}
	return totalScore
}
