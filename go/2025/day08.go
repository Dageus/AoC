package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ref for DSU impl:
// https://cp-algorithms.com/data_structures/disjoint_set_union.html

type DSU struct {
	parent []int
	size   []int
}

func newDSU(n int) *DSU {
	dsu := DSU{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := range n {
		dsu.parent[i] = i
		dsu.size[i] = 1
	}

	return &dsu
}

func (dsu *DSU) find(v int) int {
	if v == dsu.parent[v] {
		return v
	}
	dsu.parent[v] = dsu.find(dsu.parent[v])
	return dsu.parent[v]
}

func (dsu *DSU) union(v, u int) bool {
	a := dsu.find(v)
	b := dsu.find(u)
	if a != b {
		if dsu.size[a] < dsu.size[b] {
			dsu.parent[a] = b
			dsu.size[b] += dsu.size[a]
			return true
		} else {
			dsu.parent[b] = a
			dsu.size[a] += dsu.size[b]
			return true
		}
	}
	return false
}

// --------------------

const LIMIT = 1000

type Position struct {
	x, y, z int
}

type Edge struct {
	u, v             int
	squared_distance int
}

func getInput(filename string) []Position {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var list []Position

	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		list = append(list, Position{x, y, z})
	}
	return list
}

func calculateSquaredDistance(a, b Position) int {
	x := a.x - b.x
	y := a.y - b.y
	z := a.z - b.z
	return x*x + y*y + z*z
}

func partOne(filename string) int {
	list := getInput(filename)

	var edges []Edge

	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			distance := calculateSquaredDistance(list[i], list[j])
			edges = append(edges, Edge{u: i, v: j, squared_distance: distance})
		}
	}

	sort.Slice(edges, func(i, j int) bool { return edges[i].squared_distance < edges[j].squared_distance })

	dsu := newDSU(len(list))
	var limit int

	limit = min(len(edges), LIMIT)

	for i := range limit {
		edge := edges[i]
		dsu.union(edge.u, edge.v)
	}

	circuitSizes := make(map[int]int)
	for i := range len(list) {
		// find root of circuits and add them to the map
		root := dsu.find(i)
		circuitSizes[root] = dsu.size[root]
	}

	sizes := []int{}
	for _, s := range circuitSizes {
		sizes = append(sizes, s)
	}

	sort.Ints(sizes)

	n := len(sizes)

	return sizes[n-1] * sizes[n-2] * sizes[n-3]
}

func partTwo(filename string) int {
	list := getInput(filename)

	var edges []Edge

	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			distance := calculateSquaredDistance(list[i], list[j])
			edges = append(edges, Edge{u: i, v: j, squared_distance: distance})
		}
	}

	sort.Slice(edges, func(i, j int) bool { return edges[i].squared_distance < edges[j].squared_distance })

	dsu := newDSU(len(list))

	merges := len(list) - 1
	count := 0

	for _, e := range edges {
		if dsu.union(e.u, e.v) {
			count++
			if count == merges {
				return list[e.u].x * list[e.v].x
			}
		}
	}

	return 0
}

func main() {
	res := partTwo("input")
	println(res)
}
