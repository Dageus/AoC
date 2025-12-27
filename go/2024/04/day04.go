package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func getInput(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	letterMap := [][]rune{}
	padding := []rune("....")

	i := 0
	for scanner.Scan() {
		line := append(padding, []rune(scanner.Text())...)
		line = append(line, padding...)
		letterMap = append(letterMap, line)
		i++
	}
	// Add padding rows at the top and bottom
	paddingRow := []rune(strings.Repeat(".", len(letterMap[0])))
	for i := 0; i < 4; i++ {
		letterMap = append([][]rune{paddingRow}, letterMap...)
		letterMap = append(letterMap, paddingRow)
	}
	return letterMap
}

func partOne(fileName string) int {
	letterMap := getInput(fileName)
	return countXmas(letterMap)
}

func countXmas(letterMap [][]rune) int {
	sum := 0
	directions := []struct{ dx, dy int }{
		{0, 1}, {0, -1}, // Horizontal
		{1, 0}, {-1, 0}, // Vertical
		{1, 1}, {-1, -1}, // Diagonal down-right and up-left
		{1, -1}, {-1, 1}, // Diagonal down-left and up-right
	}

	word := []rune("XMAS")
	rows := len(letterMap)
	cols := len(letterMap[0])

	for r := 4; r < rows-4; r++ {
		for c := 4; c < cols-4; c++ {
			for _, dir := range directions {
				match := true

				for i := 0; i < len(word); i++ {
					nr := r + dir.dx*i // New row index
					nc := c + dir.dy*i // New column index

					// fmt.Println("letterMap[", nr, "][", nc, "] - ", string(letterMap[nr][nc]), "| word[", i, "] - ", string(word[i]))
					if letterMap[nr][nc] != word[i] {
						match = false
						break
					}
				}

				if match {
					sum++
				}
			}
		}
	}

	return sum
}

func partTwo(filename string) int {
	letterMap := getInput(filename)
	return countX_mas(letterMap)
}

func countX_mas(letterMap [][]rune) int {
	sum := 0

	rows := len(letterMap)
	cols := len(letterMap[0])

	isX_mas := func(r int, c int) bool {
		if letterMap[r-1][c-1] == 'M' && letterMap[r+1][c+1] == 'S' &&
			((letterMap[r-1][c+1] == 'M' && letterMap[r+1][c-1] == 'S') ||
				(letterMap[r-1][c+1] == 'S' && letterMap[r+1][c-1] == 'M')) {
			return true
		}
		if letterMap[r-1][c-1] == 'S' && letterMap[r+1][c+1] == 'M' &&
			((letterMap[r-1][c+1] == 'M' && letterMap[r+1][c-1] == 'S') ||
				(letterMap[r-1][c+1] == 'S' && letterMap[r+1][c-1] == 'M')) {
			return true
		}
		return false
	}

	for r := 4; r < rows-4; r++ {
		for c := 4; c < cols-4; c++ {
			if letterMap[r][c] != 'A' {
				continue
			}

			if isX_mas(r, c) {
				sum++
			}
		}
	}

	return sum
}
