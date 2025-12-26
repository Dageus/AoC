package main

import (
	"bufio"
	"math"
	"os"
)

func getInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		bank := []int{}
		for _, r := range scanner.Text() {
			bank = append(bank, int(r-'0'))
		}
		banks = append(banks, bank)
	}
}

var banks [][]int

func maxJoltage(line []int) int {
	// maximize 10's digit

	tens := 0
	idx := 0

	for i := 0; i < len(line)-1; i++ {
		if line[i] > tens {
			tens = line[i]
			idx = i
		}
	}

	ones := 0
	for j := idx + 1; j < len(line); j++ {
		if line[j] > ones {
			ones = line[j]
		}
	}
	return tens*10 + ones
}

func maxJoltage12Banks(line []int) int {
	total := 0
	j := 0
	for i := 11; i >= 0; i-- {
		subset := line[j : len(line)-i]
		idx := 0
		mx := 0
		for index, value := range subset {
			if value > mx {
				mx = value
				idx = index
			}
		}

		total += mx * int(math.Pow10(i))
		j += idx + 1
	}
	return total
}

func partOne(filename string) int {
	getInput(filename)

	sum := 0

	for _, line := range banks {
		n := maxJoltage(line)
		println(n)
		sum += n
	}
	return sum
}

func partTwo(filename string) int {
	getInput(filename)

	sum := 0

	for _, line := range banks {
		n := maxJoltage12Banks(line)
		println(n)
		sum += n
	}

	return sum
}

func main() {
	res := partTwo("input")
	println(res)
}
