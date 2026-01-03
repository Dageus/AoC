package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	left, right int
}

func partOne(filename string) int {
	var freshItems []Range
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	numbers := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			numbers = true
			continue
		}

		if !numbers {
			edges := strings.Split(line, "-")

			left, _ := strconv.Atoi(edges[0])
			right, _ := strconv.Atoi(edges[1])

			freshItems = append(freshItems, Range{left, right})
		} else {
			num, _ := strconv.Atoi(line)

			// test if num is in range
			for _, i := range freshItems {
				if num >= i.left && num <= i.right {
					sum++
					break
				}
			}
		}
	}
	return sum
}

func partTwo(filename string) int {
	var freshItems []Range
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		edges := strings.Split(line, "-")

		left, _ := strconv.Atoi(edges[0])
		right, _ := strconv.Atoi(edges[1])

		freshItems = append(freshItems, Range{left, right})
	}

	if len(freshItems) == 0 {
		return 0
	}

	sort.Slice(freshItems, func(i, j int) bool { return freshItems[i].left < freshItems[j].left })

	merged := []Range{}
	cur := freshItems[0]

	for i := 1; i < len(freshItems); i++ {
		next := freshItems[i]
		if next.left <= cur.right+1 {
			if next.right > cur.right {
				cur.right = next.right
			}
		} else {
			merged = append(merged, cur)
			cur = next
		}
	}

	merged = append(merged, cur)

	total := 0
	for _, r := range merged {
		total += (r.right - r.left + 1)
	}

	return total
}
