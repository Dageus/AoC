package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	Track = '.'
	Start = 'S'
	Wall  = '#'
	End   = 'E'
	Vidx  = 0
	Tidx  = 1
)

// Globals -----------------------

// Directions   U, D, L, R
var dr = []int{-1, 1, 0, 0}
var dc = []int{0, 0, -1, 1}
var start, end Pos

// -------------------------------

type Pos struct {
	row, col int
}

func shortestPath(raceTrack [][]byte) (map[Pos]Pos, [][][2]int) {
	var q []Pos
	var p Pos

	// [0] -> 0 if not visited, 1 if visited
	// [1] -> time/cost from start
	visitedCost := make([][][2]int, len(raceTrack))

	m := len(raceTrack)
	prev := map[Pos]Pos{start: start}
	for i := range len(raceTrack) {
		visitedCost[i] = make([][2]int, m)
	}

	q = append(q, start)
	for len(q) > 0 {
		p, q = q[0], q[1:]
		if p == end {
			return prev, visitedCost
		}
		visitedCost[p.row][p.col][Vidx] = 1
		for i := range 4 {
			next := Pos{p.row + dr[i], p.col + dc[i]}
			if raceTrack[next.row][next.col] == Wall || visitedCost[next.row][next.col][Vidx] == 1 {
				continue
			}
			q = append(q, next)
			visitedCost[next.row][next.col][Tidx] = visitedCost[p.row][p.col][Tidx] + 1
			prev[next] = p
		}
	}

	return nil, nil
}

func solve(fileName string, diff int) int {
	raceTrack := [][]byte{}

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := []byte{}
		for _, char := range scanner.Text() {
			line = append(line, byte(char))
		}
		raceTrack = append(raceTrack, line)
	}
	n, m := len(raceTrack), len(raceTrack[0])

	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			cell := raceTrack[i][j]
			switch cell {
			case Start:
				start = Pos{i, j}
			case End:
				end = Pos{i, j}
			}
		}
	}

	path, visited := shortestPath(raceTrack)

	cheats := map[struct{ p1, p2 Pos }]int{}
	tSaved := map[int]int{}

	var res int
	for p1 := range path {
		for p2 := range path {
			if p1 == p2 {
				continue
			}
			dx := int(math.Abs(float64(p1.row) - float64(p2.row)))
			dy := int(math.Abs(float64(p1.col) - float64(p2.col)))
			t1 := visited[p1.row][p1.col][Tidx]
			t2 := visited[p2.row][p2.col][Tidx]
			dt := int(math.Abs(float64(t1)-float64(t2))) - dx - dy
			if dx+dy <= 20 && dt >= diff {
				pPair := struct{ p1, p2 Pos }{p1, p2}
				if _, ok := cheats[struct{ p1, p2 Pos }{p2, p1}]; ok {
					pPair = struct{ p1, p2 Pos }{p2, p1}
				}
				if t, ok := cheats[pPair]; !ok {
					cheats[pPair] = dt
				} else {
					cheats[pPair] = min(t, dt)
				}
			}
		}
	}

	for _, v := range cheats {
		tSaved[v] += 1
	}

	for _, v := range tSaved {
		res += v
	}

	return res
}

func main() {
	tests := []struct {
		fileName string
		diff     int
		want     int
	}{
		{"input2", 50, 285},
		{"input", 100, 985737},
	}

	for _, test := range tests {
		got := solve(test.fileName, test.diff)
		if got != test.want {
			fmt.Printf("Failed Test %s\n\tGot %d, Want %d\n", test.fileName, got, test.want)
			continue
		}
		fmt.Printf("%s: %d\n", test.fileName, got)
	}
}
