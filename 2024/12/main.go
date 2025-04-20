package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	Plant rune
	R, C  int
}
type Point struct {
	X, Y int
}

const (
	AREA      = 0
	PERIMETER = 1
	SIDES     = 2
)

var region_map [][]rune
var seen = make(map[Point]bool, dimension*dimension)
var dimension int
var total_cost int

func getInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		// line := []rune{'.'}
		// line = append(line, []rune(scanner.Text())...)
		// line = append(line, '.')
		// region_map = append(region_map, line)
		region_map = append(region_map, []rune(scanner.Text()))
	}
	dimension = len(region_map[0])
	// padding := []rune{}
	// for range dimension {
	// 	padding = append(padding, '.')
	// }
	// region_map = append([][]rune{padding}, region_map...)
	// region_map = append(region_map, padding)
	printMap()
}

func printMap() {
	for _, line := range region_map {
		fmt.Println(string(line))
	}
}

func calculateTotalCost() {
	for r, line := range region_map {
		for c, plant := range string(line) {
			point := Coordinate{plant, r, c}
			if plant == '.' || seen[Point{point.R, point.C}] {
				continue
			}
			floodFill(point)
		}
	}
}

func calculateDiscountCost() {
	fmt.Println("dimension:", dimension)
	for r := range dimension {
		for c := range dimension {
			point := Point{r, c}
			if region := measureRegion(point); region != nil {
				fmt.Println("A region of", string(region_map[r][c]), "plants with price", region[AREA], "*", region[SIDES], "=", region[AREA]*region[SIDES])
				total_cost += region[0] * region[1]
			}
		}
	}
}

func measureRegion(point Point) []int {
	if seen[point] {
		return nil
	}
	region := make([]int, 3)
	measureRegionRecursive(point, region_map[point.Y][point.X], region)
	return region
}

// Recursive helper function to measure the region
func measureRegionRecursive(point Point, plant rune, region []int) {
	if seen[point] {
		return
	}
	region[AREA]++
	seen[point] = true

	// Up
	if point.Y > 0 && region_map[point.Y-1][point.X] == plant {
		measureRegionRecursive(Point{point.X, point.Y - 1}, plant, region)
	} else {
		region[PERIMETER]++
		sideContinuation := point.X > 0 && region_map[point.Y][point.X-1] == plant && (point.Y == 0 || region_map[point.Y-1][point.X-1] != plant)
		if !sideContinuation {
			region[SIDES]++
		}
	}

	// Right
	if point.X < dimension-1 && region_map[point.Y][point.X+1] == plant {
		measureRegionRecursive(Point{point.X + 1, point.Y}, plant, region)
	} else {
		region[PERIMETER]++
		sideContinuation := point.Y > 0 && region_map[point.Y-1][point.X] == plant && (point.X == len(region_map[0])-1 || region_map[point.Y-1][point.X+1] != plant)
		if !sideContinuation {
			region[SIDES]++
		}
	}

	// Down
	if point.Y < len(region_map)-1 && region_map[point.Y+1][point.X] == plant {
		measureRegionRecursive(Point{point.X, point.Y + 1}, plant, region)
	} else {
		region[PERIMETER]++
		sideContinuation := point.X < len(region_map[0])-1 && region_map[point.Y][point.X+1] == plant && (point.Y == len(region_map)-1 || region_map[point.Y+1][point.X+1] != plant)
		if !sideContinuation {
			region[SIDES]++
		}
	}

	// Left
	if point.X > 0 && region_map[point.Y][point.X-1] == plant {
		measureRegionRecursive(Point{point.X - 1, point.Y}, plant, region)
	} else {
		region[PERIMETER]++
		sideContinuation := point.Y < len(region_map)-1 && region_map[point.Y+1][point.X] == plant && (point.X == 0 || region_map[point.Y+1][point.X-1] != plant)
		if !sideContinuation {
			region[SIDES]++
		}
	}
}

func floodFill(point Coordinate) {
	directions := []Coordinate{
		{0, 1, 0},
		{0, -1, 0},
		{0, 0, 1},
		{0, 0, -1},
	}

	stack := []Coordinate{point}
	area := 0
	corners := 4
	seen[Point{point.R, point.C}] = true

	for len(stack) > 0 {
		plant := stack[0]
		stack = stack[1:]

		area++

		for _, direction := range directions {
			new_plant := Coordinate{0, plant.R + direction.R, plant.C + direction.C}
			new_plant.Plant = region_map[new_plant.R][new_plant.C]
			if new_plant.Plant == '.' || new_plant.Plant != plant.Plant {
				// TODO:
			} else if !seen[Point{new_plant.C, new_plant.R}] {
				seen[Point{new_plant.C, new_plant.R}] = true
				stack = append(stack, new_plant)
			}
		}
	}

	fmt.Println("region of plant", point.Plant, "has", corners, "sides")
	total_cost += area * corners
}

func newCorner(new_plant Coordinate, polygon []Coordinate) bool {
	for _, plant := range polygon {
		if new_plant.C != plant.C && new_plant.R != plant.R {
			return true
		}
	}
	return false
}

func partOne(filename string) {
	getInput(filename)
	calculateTotalCost()
	fmt.Println("total_cost:", total_cost)
}

func partTwo(filename string) {
	getInput(filename)
	calculateDiscountCost()
	fmt.Println("discount_cost:", total_cost)
}

func main() {
	// partOne("input2")
	partTwo("input3")
}
