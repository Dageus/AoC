package main

import (
	"bufio"
	"math"
	"os"
)

func getInput(filename string) [][]int {
	var banks [][]int

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
	return banks
}

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
	banks := getInput(filename)

	sum := 0

	for _, line := range banks {
		n := maxJoltage(line)
		sum += n
	}
	return sum
}

func partTwo(filename string) int {
	banks := getInput(filename)

	sum := 0

	for _, line := range banks {
		n := maxJoltage12Banks(line)
		sum += n
	}

	return sum
}
