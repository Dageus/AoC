package main

import (
	"bufio"
	"os"
	"strconv"
)

const (
	LEFT  = 'L'
	RIGHT = 'R'

	START_DIAL = 50
)

type Rotation struct {
	direction rune
	distance  int
}

func getInput(filename string) []Rotation {
	var movements []Rotation
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		direction := rune(scanner.Text()[0])
		distance, _ := strconv.Atoi(scanner.Text()[1:])
		movements = append(movements, Rotation{direction, distance})
	}

	return movements
}

func partOne(filename string) int {
	movements := getInput(filename)

	dial := START_DIAL

	zeros := 0

	for _, movement := range movements {
		dial, _ = processRotation(dial, movement)
		if dial == 0 {
			zeros++
		}
	}
	return zeros
}

func processRotation(curr_dial int, rotation Rotation) (int, int) {
	switch rotation.direction {
	case LEFT:
		normalized := rotation.distance % 100
		full_rotations := rotation.distance / 100
		new_dial := curr_dial - normalized
		if new_dial == 0 {
			return new_dial, full_rotations + 1
		} else if new_dial < 0 {
			if curr_dial == 0 {
				return new_dial + 100, full_rotations
			}
			return new_dial + 100, full_rotations + 1
		} else {
			return new_dial, full_rotations
		}
	case RIGHT:
		normalized := rotation.distance % 100
		full_rotations := rotation.distance / 100
		new_dial := curr_dial + normalized
		if new_dial == 100 {
			return 0, full_rotations + 1
		} else if new_dial > 99 {
			if curr_dial == 100 {
				return new_dial + 100, full_rotations
			}
			return new_dial % 100, full_rotations + 1
		} else {
			return new_dial, full_rotations
		}
	default:
		panic("Something went wrong")
	}
}

func partTwo(filename string) int {
	movements := getInput(filename)

	dial := START_DIAL

	zeros := 0

	for _, movement := range movements {
		new_dial, n := processRotation(dial, movement)
		dial = new_dial
		zeros += n
	}
	return zeros
}
