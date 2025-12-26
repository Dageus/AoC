package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInputPartOne(fileName string) ([]int, []int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	list1 := []int{}
	list2 := []int{}

	for fileScanner.Scan() {
		items := strings.Split(fileScanner.Text(), "   ")
		item1, err := strconv.Atoi(items[0])
		if err != nil {
			log.Fatal(err)
		}
		item2, err := strconv.Atoi(items[1])
		if err != nil {
			log.Fatal(err)
		}
		list1 = append(list1, item1)
		list2 = append(list2, item2)
	}

	return list1, list2
}

func getInputPartTwo(fileName string) ([]int, map[int]int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	list := []int{}
	table := map[int]int{}

	for fileScanner.Scan() {
		items := strings.Split(fileScanner.Text(), "   ")
		item1, err := strconv.Atoi(items[0])
		if err != nil {
			log.Fatal(err)
		}
		item2, err := strconv.Atoi(items[1])
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, item1)
		if _, ok := table[item2]; ok {
			table[item2]++
		} else {
			table[item2] = 1
		}
	}

	return list, table
}

func partOne() {
	list1, list2 := getInputPartOne("input")
	sort.Ints(list1)
	sort.Ints(list2)
	sum := 0
	for i := range list1 {
		sum += int(math.Abs(float64(list1[i]) - float64(list2[i])))
	}
	fmt.Println(sum)
}

func partTwo() {
	list, table := getInputPartTwo("input")
	fmt.Println(table)
	sum := 0
	for _, val := range list {
		if mul, ok := table[val]; ok {
			sum += val * mul
		}
	}
	fmt.Println(sum)
}

func main() {
	partTwo()
}
