package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	HISTORIAN_COMPUTER = 't'
)

var nodes = make(map[string]*Node)
var processedPairs = make(map[[3]string]bool)

type Node struct {
	Name       string
	Neighbours []string
}

func getInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		parseConnection(scanner.Text())
	}
}

func parseConnection(conn string) {
	connections := strings.Split(conn, "-")
	conn1 := connections[0]
	conn2 := connections[1]
	if _, exists := nodes[conn1]; !exists {
		node := Node{Name: conn1, Neighbours: []string{conn2}}
		nodes[conn1] = &node
	} else {
		node := nodes[conn1]
		node.Neighbours = append(node.Neighbours, conn2)
	}

	if _, exists := nodes[conn2]; !exists {
		node := Node{Name: conn2, Neighbours: []string{conn1}}
		nodes[conn2] = &node
	} else {
		node := nodes[conn2]
		node.Neighbours = append(node.Neighbours, conn1)
	}
}

func partOne(filename string) {
	getInput(filename)
	connected_computers := 0
	for _, node := range nodes {
		if node.Name[0] == HISTORIAN_COMPUTER && len(node.Neighbours) >= 2 { // search for 3 inter connected computers
			connected_computers += searchForConnection(node.Name, node.Neighbours)
		}
	}
	fmt.Println("len:", connected_computers)
}

func searchForConnection(computer string, neighbours []string) int {
	intersections := 0
	for i := 0; i < len(neighbours)-1; i++ {
		iNeighbours := nodes[neighbours[i]].Neighbours
		inter := intersection(neighbours, iNeighbours)

		if len(inter) != 0 {

			for _, node := range inter {
				pairKey := []string{computer, neighbours[i], node}
				sort.Strings(pairKey)
				key := [3]string{pairKey[0], pairKey[1], pairKey[2]}
				fmt.Println(key)

				if processedPairs[key] {
					continue
				}

				fmt.Println("for", computer, "and", neighbours[i], "the intersection is:", inter)
				processedPairs[key] = true

				intersections++
			}

		}

	}
	return intersections
}

func intersection(nums1 []string, nums2 []string) []string {
	var result []string

	set := make(map[string]struct{}) // There is no set in Golang, use map to implement. struct{} is an empty structure to save memory.
	for _, v := range nums2 {
		set[v] = struct{}{}
	}

	for _, v := range nums1 {
		if _, ok := set[v]; ok {
			result = append(result, v)
		}
	}
	return result
}

func createPairKey(node1, node2 string) string {
	if node1 < node2 {
		return node1 + "-" + node2
	}
	return node2 + "-" + node1
}

func partTwo(filename string) {
	getInput(filename)

	largest_connection := []string{}
	for _, node := range nodes {
		connected_computers := findLargestClique(node.Name)
		if len(connected_computers) > len(largest_connection) {
			largest_connection = connected_computers
			fmt.Println("new largest:", largest_connection)
		}
	}
	sort.Strings(largest_connection)
	fmt.Println("password:", strings.Join(largest_connection, ","))
}

func findLargestClique(start string) []string {
	// Start with the node itself
	current_clique := []string{start}
	neighbors := nodes[start].Neighbours

	// Check if adding each neighbor maintains full connectivity
	for _, neighbor := range neighbors {
		if isFullyConnected(current_clique, neighbor) {
			current_clique = append(current_clique, neighbor)
		}
	}

	return current_clique
}

// Checks if adding a new node to the clique maintains full connectivity
func isFullyConnected(clique []string, newNode string) bool {
	for _, node := range clique {
		if !isNeighbor(node, newNode) {
			return false
		}
	}
	return true
}

// Helper to check if two nodes are neighbors
func isNeighbor(node1, node2 string) bool {
	for _, neighbor := range nodes[node1].Neighbours {
		if neighbor == node2 {
			return true
		}
	}
	return false
}

func main() {
	// partOne("input")
	partTwo("input")
}
