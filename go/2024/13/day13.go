package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	A_TOKENS = 3
	B_TOKENS = 1
	OFFSET   = 10000000000000
)

type Button struct {
	X, Y int
}

type Prize struct {
	X, Y int
}

type ClawMachine struct {
	A    Button
	B    Button
	Goal Prize
}

var clawMachines []ClawMachine
var aGoal, bGoal int

func getInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		buttonA := parseButton(scanner.Text())
		if !scanner.Scan() {
			panic("couldn't scan")
		}

		buttonB := parseButton(scanner.Text())
		if !scanner.Scan() {
			panic("couldn't scan")
		}

		prize := parsePrize(scanner.Text())

		clawMachines = append(clawMachines, ClawMachine{buttonA, buttonB, prize})
	}
	fmt.Println(clawMachines)
}

func parsePrize(text string) Prize {
	els := strings.Split(text, ":")
	buttons := strings.Trim(els[1], " ")
	buttonValues := strings.Split(buttons, ", ")
	xPrize := strings.Split(buttonValues[0], "=")
	xPrizeValue, err := strconv.Atoi(xPrize[1])
	if err != nil {
		panic(err)
	}
	yPrize := strings.Split(buttonValues[1], "=")
	yPrizeValue, err := strconv.Atoi(yPrize[1])
	if err != nil {
		panic(err)
	}

	return Prize{xPrizeValue + OFFSET, yPrizeValue + OFFSET}
}

func parseButton(text string) Button {
	els := strings.Split(text, ":")
	buttons := strings.Trim(els[1], " ")
	buttonValues := strings.Split(buttons, ", ")
	xButton := strings.Split(buttonValues[0], "+")
	xButtonValue, err := strconv.Atoi(xButton[1])
	if err != nil {
		panic(err)
	}
	yButton := strings.Split(buttonValues[1], "+")
	yButtonValue, err := strconv.Atoi(yButton[1])
	if err != nil {
		panic(err)
	}

	return Button{xButtonValue, yButtonValue}
}

func getOptimalResult(clawMachine ClawMachine) *Prize {
	topDivision := clawMachine.Goal.Y*clawMachine.A.X - clawMachine.Goal.X*clawMachine.A.Y
	bottomDivision := clawMachine.B.Y*clawMachine.A.X - clawMachine.B.X*clawMachine.A.Y
	if topDivision%bottomDivision != 0 {
		return nil
	}
	B := topDivision / bottomDivision

	topDivision = (clawMachine.Goal.X - B*clawMachine.B.X)
	bottomDivision = clawMachine.A.X
	if topDivision%bottomDivision != 0 {
		return nil
	}
	A := topDivision / bottomDivision

	return &Prize{X: A, Y: B}
}

func partOne(filename string) {
	getInput(filename)
	for _, clawMachine := range clawMachines {
		if result := getOptimalResult(clawMachine); result != nil {
			aGoal += result.X
			bGoal += result.Y
		}
	}
	fmt.Println(aGoal*A_TOKENS + bGoal*B_TOKENS)
}

func main() {
	partOne("input")
}
