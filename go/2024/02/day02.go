package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const MAX_DIFF = 3

func getInput(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	list := [][]int{}

	for fileScanner.Scan() {
		items := strings.Split(fileScanner.Text(), " ")
		sublist := []int{}
		for _, item := range items {
			item, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal(err)
			}
			sublist = append(sublist, item)
		}
		list = append(list, sublist)
	}

	return list
}

func safeReport(report []int) bool {
	trend := report[0] - report[1]
	for i := 1; i < len(report); i++ {
		diff := report[i-1] - report[i]
		if int(math.Abs(float64(diff))) > MAX_DIFF || int(math.Abs(float64(diff))) < 1 {
			return false
		}
		if (trend > 0 && diff < 0) || (trend < 0 && diff > 0) {
			return false
		}
	}
	return true
}

func checkSafe(report []int) bool {
	trend := report[0] - report[1]

	for i := 1; i < len(report); i++ {
		diff := report[i-1] - report[i]
		if int(math.Abs(float64(diff))) > MAX_DIFF || int(math.Abs(float64(diff))) < 1 {
			return false
		}
		if (trend > 0 && diff < 0) || (trend < 0 && diff > 0) {
			return false
		}
	}
	return true
}

func tolerableReport(report []int) bool {
	if checkSafe(report) {
		return true
	}

	for i := range len(report) {
		modified := make([]int, 0, len(report)-1)
		modified = append(modified, report[:i]...)
		modified = append(modified, report[i+1:]...)
		if checkSafe(modified) {
			return true
		}
	}
	return false
}

func partOne(filename string) int {
	sum := 0
	reports := getInput(filename)
	fmt.Println(len(reports))
	for _, report := range reports {
		if safeReport(report) {
			sum++
		}
	}

	return sum
}

func partTwo(filename string) int {
	sum := 0
	reports := getInput(filename)
	fmt.Println(len(reports))
	for _, report := range reports {
		if tolerableReport(report) {
			sum++
		}
	}

	return sum
}
