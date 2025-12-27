package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var order = make(map[int][]int)
var updates [][]int

func getInput(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// scan rules
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		nums := strings.Split(scanner.Text(), "|")
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal(err)
		}
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatal(err)
		}
		if _, ok := order[num1]; ok {
			order[num1] = append(order[num1], num2)
		} else {
			order[num1] = []int{num2}
		}
	}

	fmt.Println(order)

	// scan updates
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), ",")
		line := []int{}
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			line = append(line, n)
		}
		updates = append(updates, line)
	}
}

func verifyUpdates() {
	validUpdates := [][]int{}
	for _, update := range updates {
		if !checkIntegrity(update) {
			validUpdates = append(validUpdates, update)
		}
	}
	updates = validUpdates
}

func checkIntegrity(update []int) bool {
	indexMap := make(map[int]int)
	for i, num := range update {
		indexMap[num] = i
	}

	for _, page := range update {
		if dependencies, exists := order[page]; exists {
			for _, dependent := range dependencies {
				depIndex, depExists := indexMap[dependent]
				if depExists {
					if indexMap[page] >= depIndex {
						return false
					}
				}
			}
		}
	}

	return true
}

func fixUpdate(update []int) []int {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	nodes := make(map[int]bool)

	for _, page := range update {
		nodes[page] = true
		if dependencies, exists := order[page]; exists {
			for _, dependent := range dependencies {
				if contains(update, dependent) {
					graph[page] = append(graph[page], dependent)
					inDegree[dependent]++
				}
			}
		}
	}

	sorted := []int{}
	queue := []int{}

	for node := range nodes {
		if inDegree[node] == 0 {
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		sorted = append(sorted, current)

		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// If the sorted array doesn't contain all nodes, there's a cycle (invalid rules)
	if len(sorted) != len(update) {
		log.Fatal("Cyclic dependency detected in rules!")
	}

	return sorted
}

// Helper function to check if a value exists in a slice
func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func verifyIncorrectUpdates() {
	invalidUpdates := [][]int{}
	for _, update := range updates {
		if !checkIntegrity(update) {
			fmt.Println("before:", update)
			newUpdate := fixUpdate(update)
			invalidUpdates = append(invalidUpdates, newUpdate)
			fmt.Println("after:", newUpdate)
		}
	}
	updates = invalidUpdates
}

func sumMiddleNumber() int {
	sum := 0

	for _, update := range updates {
		sum += update[len(update)/2]
	}
	return sum
}

func partOne(filename string) int {
	getInput(filename)
	verifyUpdates()
	return sumMiddleNumber()
}

func partTwo(fileName string) int {
	getInput(fileName)
	verifyIncorrectUpdates()
	return sumMiddleNumber()

}
