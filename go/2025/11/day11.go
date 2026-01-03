package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	START_DEVICE       = "you"
	STOP_DEVICE        = "out"
	SERVER_RACK        = "svr"
	MANDATORY_DEVICE_1 = "dac"
	MANDATORY_DEVICE_2 = "fft"
)

var (
	memo map[string]int
)

func getInput(filename string) map[string][]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	deviceMap := make(map[string][]string)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		devices := strings.Split(scanner.Text(), ": ")
		device := devices[0]
		outputs := strings.Split(devices[1], " ")
		deviceMap[device] = outputs
	}

	return deviceMap
}

func countPaths(curr string, end string, adj map[string][]string) int {
	if curr == end {
		return 1
	}

	if count, ok := memo[curr]; ok {
		return count
	}

	total := 0
	for _, neigh := range adj[curr] {
		total += countPaths(neigh, end, adj)
	}
	memo[curr] = total
	return total
}

func runPaths(start string, end string, adj map[string][]string) int {
	memo = make(map[string]int)
	return countPaths(start, end, adj)
}

func partOne(filename string) int {
	deviceMap := getInput(filename)
	memo = make(map[string]int)
	return countPaths(START_DEVICE, STOP_DEVICE, deviceMap)
}

func partTwo(filename string) int {
	deviceMap := getInput(filename)
	paths1 := runPaths(SERVER_RACK, MANDATORY_DEVICE_1, deviceMap) * runPaths(MANDATORY_DEVICE_1, MANDATORY_DEVICE_2, deviceMap) * runPaths(MANDATORY_DEVICE_2, STOP_DEVICE, deviceMap)

	paths2 := runPaths(SERVER_RACK, MANDATORY_DEVICE_2, deviceMap) * runPaths(MANDATORY_DEVICE_2, MANDATORY_DEVICE_1, deviceMap) * runPaths(MANDATORY_DEVICE_1, STOP_DEVICE, deviceMap)

	return paths1 + paths2
}
