package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y int
}

type Robot struct {
	Position Coordinate
	Velocity Coordinate
}

type State struct {
	Time           int
	SafetyFactor   int
	RobotPositions []Coordinate
}

var robots []Robot

var states []State

const WIDTH = 101

// const WIDTH = 11

const HEIGHT = 103

// const HEIGHT = 7

const TIME = 100

func getInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		parseRobot(scanner.Text())
	}
}

func parseRobot(text string) {
	attributes := strings.Split(text, " ")
	position := attributes[0]
	velocity := attributes[1]

	positionValues := strings.Split(position, "=")
	velocityValues := strings.Split(velocity, "=")

	xyPosition := strings.Split(positionValues[1], ",")
	xyVelocity := strings.Split(velocityValues[1], ",")

	xPosition, err := strconv.Atoi(xyPosition[0])
	yPosition, err := strconv.Atoi(xyPosition[1])

	if err != nil {
		panic(err)
	}

	xVelocity, err := strconv.Atoi(xyVelocity[0])
	yVelocity, err := strconv.Atoi(xyVelocity[1])

	if err != nil {
		panic(err)
	}

	robots = append(robots, Robot{Coordinate{xPosition, yPosition}, Coordinate{xVelocity, yVelocity}})
}

func testNewPosition(robot Robot) Coordinate {
	// fmt.Println("previous Position:", robot.Position)
	newPosition := Coordinate{robot.Position.X + robot.Velocity.X, robot.Position.Y + robot.Velocity.Y}
	// fmt.Println("new Position:", newPosition)

	if newPosition.X >= WIDTH {
		newPosition.X = newPosition.X - WIDTH
	} else if newPosition.X < 0 {
		newPosition.X = newPosition.X + WIDTH
	}

	if newPosition.Y >= HEIGHT {
		newPosition.Y = newPosition.Y - HEIGHT
	} else if newPosition.Y < 0 {
		newPosition.Y = newPosition.Y + HEIGHT
	}

	return newPosition
}

func simulateMovement() {
	for i := range robots {
		robots[i].Position = testNewPosition(robots[i])
	}
}

func calculateSafetyFactor() int {
	Xmid := int(math.Ceil(WIDTH / 2))
	Ymid := int(math.Ceil(HEIGHT / 2))

	fmt.Println("Xmid:", Xmid)
	fmt.Println("Ymid:", Ymid)

	quadrant_1 := 0
	quadrant_2 := 0
	quadrant_3 := 0
	quadrant_4 := 0

	for _, robot := range robots {
		x := robot.Position.X
		y := robot.Position.Y

		if x < Xmid && y < Ymid { // first quadrant
			quadrant_1++
		} else if x > Xmid && x < WIDTH && y < Ymid { // second quadrant
			quadrant_2++
		} else if x < Xmid && y > Ymid && y < HEIGHT { // third quadrant
			quadrant_3++
		} else if x > Xmid && x < WIDTH && y > Ymid && y < HEIGHT { // fourth quadrant
			quadrant_4++
		}
	}

	fmt.Printf("Quadrant counts: Q1=%d, Q2=%d, Q3=%d, Q4=%d\n", quadrant_1, quadrant_2, quadrant_3, quadrant_4)
	return quadrant_1 * quadrant_2 * quadrant_3 * quadrant_4
}

func partOne(filename string) int {
	getInput(filename)
	for range TIME {
		print(".")
		simulateMovement()
	}
	fmt.Println("\n----------")
	for _, robot := range robots {
		fmt.Println(robot)
	}
	return calculateSafetyFactor()
}

func simulateAndTrackStates() {
	for t := 0; t < WIDTH*HEIGHT; t++ {
		simulateMovement()
		safetyFactor := calculateSafetyFactor()

		var positions []Coordinate
		for _, robot := range robots {
			positions = append(positions, robot.Position)
		}
		states = append(states, State{t, safetyFactor, positions})
	}
}

func writeTopStatesToFiles() {
	// Sort states by safety factor
	sort.Slice(states, func(i, j int) bool {
		return states[i].SafetyFactor < states[j].SafetyFactor
	})

	// Ensure the output directory exists
	outputDir := "output"
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to create output directory: %v\n", err)
		return
	}

	// Write the top 20 states to individual files
	for i := 0; i < 30 && i < len(states); i++ {
		state := states[i]
		filename := filepath.Join(outputDir, fmt.Sprintf("state_%02d_time_%d.txt", i+1, state.Time))

		// Write the grid to the file
		file, err := os.Create(filename)
		if err != nil {
			fmt.Printf("Failed to create file %s: %v\n", filename, err)
			continue
		}
		defer file.Close()

		// Write grid state and metadata
		grid := generateGrid(state.RobotPositions)
		_, err = file.WriteString(fmt.Sprintf("Time: %d\nSafety Factor: %d\n\n", state.Time, state.SafetyFactor))
		if err != nil {
			fmt.Printf("Failed to write to file %s: %v\n", filename, err)
			continue
		}

		for _, row := range grid {
			_, err = file.WriteString(string(row) + "\n")
			if err != nil {
				fmt.Printf("Failed to write to file %s: %v\n", filename, err)
				break
			}
		}
	}
	fmt.Println("Top 20 safety states have been written to the 'output' directory.")
}

func generateGrid(positions []Coordinate) [][]rune {
	grid := make([][]rune, HEIGHT)
	for i := range grid {
		grid[i] = make([]rune, WIDTH)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for _, pos := range positions {
		grid[pos.Y][pos.X] = '#'
	}

	return grid
}

func writeFilteredStatesToFiles() int {
	// Ensure the output directory exists
	outputDir := "filtered_output"
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to create output directory: %v\n", err)
		return -1
	}

	// Iterate through states and check for horizontal streaks
	fileCount := 0
	for _, state := range states {
		grid := generateGrid(state.RobotPositions)
		if hasHorizontalStreak(grid, 10) {
			fileCount++
			filename := filepath.Join(outputDir, fmt.Sprintf("state_%02d_time_%d.txt", fileCount, state.Time))

			// Write the grid to the file
			file, err := os.Create(filename)
			if err != nil {
				fmt.Printf("Failed to create file %s: %v\n", filename, err)
				continue
			}
			defer file.Close()

			// Write grid state and metadata
			_, err = file.WriteString(fmt.Sprintf("Time: %d\nSafety Factor: %d\n\n", state.Time, state.SafetyFactor))
			if err != nil {
				fmt.Printf("Failed to write to file %s: %v\n", filename, err)
				continue
			}

			for _, row := range grid {
				_, err = file.WriteString(string(row) + "\n")
				if err != nil {
					fmt.Printf("Failed to write to file %s: %v\n", filename, err)
					break
				}
			}
		}
	}

	return fileCount
}

// Checks if any row in the grid contains a streak of at least `streakLength` consecutive robots
func hasHorizontalStreak(grid [][]rune, streakLength int) bool {
	for _, row := range grid {
		count := 0
		for _, cell := range row {
			if cell == '#' {
				count++
				if count > streakLength {
					return true
				}
			} else {
				count = 0
			}
		}
	}
	return false
}

func partTwo(filename string) int {
	getInput(filename)
	simulateAndTrackStates()
	return writeFilteredStatesToFiles()
}
