package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Present struct {
	Shape [][]bool
}

type Region struct {
	X, Y   int
	Layout []int
}

func getInput(filename string) ([]Present, []Region) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	presents := []Present{}
	regions := []Region{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	presentCounter := 0

	for scanner.Scan() {
		if presentCounter <= 5 {
			shape := [][]bool{}
			for scanner.Scan() {
				text := scanner.Text()
				fmt.Println(text)
				if text == "" {
					break
				}

				shapeRow := []bool{}
				for _, v := range text {
					switch v {
					case '#':
						shapeRow = append(shapeRow, true)
					case '.':
						shapeRow = append(shapeRow, false)
					}
				}

				shape = append(shape, shapeRow)
			}
			presents = append(presents, Present{Shape: shape})
			presentCounter++
		} else {
			list := strings.Split(scanner.Text(), ": ")
			dimensionsStr, layoutStr := list[0], list[1]
			dimensions := strings.Split(dimensionsStr, "x")
			x, _ := strconv.Atoi(dimensions[0])
			y, _ := strconv.Atoi(dimensions[1])

			layout := []int{}

			for v := range strings.SplitSeq(layoutStr, " ") {
				num, _ := strconv.Atoi(v)
				layout = append(layout, num)
			}
			regions = append(regions, Region{
				X:      x,
				Y:      y,
				Layout: layout,
			})
		}

	}

	return presents, regions
}

func sumList(l []int) int {
	s := 0
	for _, i := range l {
		s += i
	}
	return s
}

func partOne(filename string) int {
	_, regions := getInput(filename)
	sum := 0

	for _, region := range regions {
		if region.X/3*region.Y/3 >= sumList(region.Layout) {
			sum++
		}
	}

	return sum
}

func partTwo(filename string) int {
	_, regions := getInput(filename)
	sum := 0

	for _, region := range regions {
		if region.X/3*region.Y/3 >= sumList(region.Layout) {
			sum++
		}
	}

	return sum
}

func main() {
	res := partOne("input")
	println(res)
}
