package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Coordinates struct {
	R, C int
}

type Trajectory struct {
	Position  Coordinates
	Direction int
}

var GUARD_TYPE = []rune{
	'^', // up
	'>', // right
	'v', // down
	'<', // left
}

var GUARD int

const OBSTACLE = '#'

var initialPosition Coordinates
var guardPosition Coordinates
var obstaclePositions = make(map[Coordinates]bool)
var distinctPositions = make(map[Coordinates]bool)
var guardTrajectory = []Trajectory{}
var mapDimensions Coordinates

func getInput(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	r := 0
	c := 0
	for scanner.Scan() {
		c = 0
		for _, position := range scanner.Text() {
			if position == OBSTACLE {
				obstaclePositions[Coordinates{R: r, C: c}] = true
			} else if position == '^' {
				GUARD = 0
				guardPosition = Coordinates{R: r, C: c}
			} else if position == '>' {
				GUARD = 1
				guardPosition = Coordinates{R: r, C: c}
			} else if position == 'v' {
				GUARD = 2
				guardPosition = Coordinates{R: r, C: c}
			} else if position == '<' {
				GUARD = 3
				guardPosition = Coordinates{R: r, C: c}
			}
			c++
		}
		r++
	}
	initialPosition = guardPosition
	mapDimensions = Coordinates{R: r, C: c}
	// NOTE: this is for partOne
	// distinctPositions[guardPosition] = true
	// NOTE: this is for partTwo
	guardTrajectory = append(guardTrajectory, Trajectory{Position: guardPosition, Direction: GUARD})
}

func partOne(fileName string) {
	getInput(fileName)
	mapTrajectory()
	countDistinctPositions()
}

func mapTrajectory() {
	for !outsideMap(guardPosition) {
		tryIncrementPosition()
	}
}

func tryIncrementPosition() {
	newPosition := calculateNextPosition(GUARD, guardPosition)
	if checkCollision(newPosition, obstaclePositions) {
		GUARD = (GUARD + 1) % 4
		newPosition = calculateNextPosition(GUARD, guardPosition)
	}
	guardPosition = newPosition
	guardTrajectory = append(guardTrajectory, Trajectory{Position: guardPosition, Direction: GUARD})
}

func calculateNextPosition(direction int, current Coordinates) Coordinates {
	switch direction {
	case 0: // Up
		return Coordinates{R: current.R - 1, C: current.C}
	case 1: // Right
		return Coordinates{R: current.R, C: current.C + 1}
	case 2: // Down
		return Coordinates{R: current.R + 1, C: current.C}
	case 3: // Left
		return Coordinates{R: current.R, C: current.C - 1}
	}
	log.Fatal("shouldn't have reached here")
	return current
}

func checkCollision(position Coordinates, obstacles map[Coordinates]bool) bool {
	if _, exists := obstacles[position]; exists {
		return true
	}
	return false
}

func outsideMap(position Coordinates) bool {
	return (position.C < 0 || position.R < 0) || (position.C >= mapDimensions.C || position.R >= mapDimensions.R)
}

func countDistinctPositions() {
	fmt.Println(len(distinctPositions) - 1)
}

func partTwo(fileName string) {
	getInput(fileName)
	mapTrajectory()
	findObstaclePositions()
}

func findObstaclePositions() {
	n_obstacles := make(map[Coordinates]bool)
	for _, obstacle := range guardTrajectory {
		newObstacle := calculateObstaclePosition(obstacle)

		if createsLoop(newObstacle) {
			n_obstacles[newObstacle] = true
		}
	}

	fmt.Println(len(n_obstacles))
}

func calculateObstaclePosition(obstacle Trajectory) Coordinates {
	switch obstacle.Direction {
	case 0:
		return Coordinates{R: obstacle.Position.R - 1, C: obstacle.Position.C}
	case 1:
		return Coordinates{R: obstacle.Position.R, C: obstacle.Position.C + 1}
	case 2:
		return Coordinates{R: obstacle.Position.R + 1, C: obstacle.Position.C}
	case 3:
		return Coordinates{R: obstacle.Position.R, C: obstacle.Position.C - 1}
	}
	log.Fatal("shouldn't have happened")
	return Coordinates{}
}

func createsLoop(obstacle Coordinates) bool {
	currentPosition := initialPosition
	visited := make(map[Trajectory]bool)
	obstacles := make(map[Coordinates]bool)
	for k, v := range obstaclePositions {
		obstacles[k] = v
	}
	obstacles[obstacle] = true

	guard := 0

	for !outsideMap(currentPosition) {
		state := Trajectory{Position: currentPosition, Direction: guard}
		if visited[state] {
			// fmt.Println("Loop detected at:", state)
			return true
		}
		visited[state] = true

		switch guard {
		case 0:
			newPosition := Coordinates{R: currentPosition.R - 1, C: currentPosition.C}
			if checkCollision(newPosition, obstacles) {
				guard = (guard + 1) % 4
				currentPosition = Coordinates{R: currentPosition.R, C: currentPosition.C + 1}
			} else {
				currentPosition = newPosition
			}
		case 1:
			newPosition := Coordinates{R: currentPosition.R, C: currentPosition.C + 1}
			if checkCollision(newPosition, obstacles) {
				guard = (guard + 1) % 4
				currentPosition = Coordinates{R: currentPosition.R + 1, C: currentPosition.C}
			} else {
				currentPosition = newPosition
			}
		case 2:
			newPosition := Coordinates{R: currentPosition.R + 1, C: currentPosition.C}
			if checkCollision(newPosition, obstacles) {
				guard = (guard + 1) % 4
				currentPosition = Coordinates{R: currentPosition.R, C: currentPosition.C - 1}
			} else {
				currentPosition = newPosition
			}
		case 3:
			newPosition := Coordinates{R: currentPosition.R, C: currentPosition.C - 1}
			if checkCollision(newPosition, obstacles) {
				guard = (guard + 1) % 4
				currentPosition = Coordinates{R: currentPosition.R - 1, C: currentPosition.C}
			} else {
				currentPosition = newPosition
			}
		default:
			log.Fatal("wtf")
		}
	}
	return false
}

func main() {
	// partOne("input")
	partTwo("input")
}
