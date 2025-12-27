package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Space struct {
	Size, Start int
}

var fileSize = make(map[rune]int)
var stringFileSize = make(map[string]Space)
var id int

// NOTE: maps size of empty space to start index
var emptySpace []Space

func reformatDiskString(disk []rune) []string {
	newDisk := []string{}
	disk_index := 0

	for idx, r := range disk {
		val := runeToInt(r)
		char := strconv.Itoa(id)
		if idx%2 != 0 {
			char = "."
			emptySpace = append(emptySpace, Space{Start: disk_index, Size: val})
		}
		if char != "." {
			stringFileSize[char] = Space{Start: disk_index, Size: val}
			fmt.Println(stringFileSize[char])
		}
		for range val {
			newDisk = append(newDisk, char)
			disk_index++
		}
		id += idx % 2
	}

	// fmt.Println(fileSize)
	fmt.Println("EMPTYSPACE", emptySpace)

	return newDisk
}

func reformatDisk(disk []rune) []rune {
	newDisk := []rune{}
	id := 0
	disk_index := 0

	for idx, r := range disk {
		val := runeToInt(r)
		char := intToRune(id)
		if idx%2 != 0 {
			char = '.'
			emptySpace = append(emptySpace, Space{Start: disk_index, Size: val})
		}
		for range val {
			newDisk = append(newDisk, char)
			disk_index++
		}
		id += idx % 2
		fileSize[char] = val
	}

	// fmt.Println(fileSize)

	return newDisk
}

func reorganizeSpace(disk []rune) int {
	// NOTE: two pointer?
	left := 0
	right := len(disk) - 1
	for left < right {
		if disk[left] != '.' {
			left++
			continue
		}
		if disk[right] == '.' {
			right--
			continue
		}
		disk[left], disk[right] = disk[right], disk[left]
		left++
		right--
	}
	return computeChecksum(disk)
}

func reorganizeSpaceFullFiles(disk []string) int {

	fmt.Println("FILES:", stringFileSize)
	fmt.Println("N_IDS:", id)
	for i := id; i >= 0; i-- {
		fmt.Println("->trying id:", i)
		index := strconv.Itoa(i)
		file := stringFileSize[index]
		// Check for the leftmost free space large enough to fit the file
		for j, freespace := range emptySpace {
			if freespace.Start >= file.Start {
				// fmt.Println("not valid (", freespace.Start, ">", file.Start, ")")
				break
			}

			if freespace.Size < file.Size {
				continue
			}

			// Move the file to the free space
			// fmt.Printf("Moving file %d (size %d) to free space at %d\n", i, file.Size, freespace.Start)
			for k := 0; k < file.Size; k++ {
				disk[freespace.Start+k] = disk[file.Start+k]
				disk[file.Start+k] = "."
			}

			// Update free space
			emptySpace[j].Start += file.Size
			emptySpace[j].Size -= file.Size

			// If the free space is now used up, remove it
			if emptySpace[j].Size == 0 {
				emptySpace = append(emptySpace[:j], emptySpace[j+1:]...)
			}

			// Update file's new position
			file.Start = freespace.Start
			stringFileSize[index] = file

			break
		}
	}

	return computeChecksumString(disk)
}

func findIndexForSize(size int) int {
	for i, freespace := range emptySpace {
		if freespace.Size >= size {
			return i
		}
	}
	fmt.Println("NO FREE SPACE FOUND FOR SIZE =", size)
	panic("")
}

func getFreeSpaceForIndex(index int) int {
	for _, freespace := range emptySpace {
		if freespace.Start == index {
			return freespace.Size
		}
	}
	return -1
}

func findNextFile(disk []rune, left int) int {
	for i := left; i < len(disk); i++ {
		if disk[i] != '.' {
			return i - left
		}
	}
	return -1
}

func computeChecksum(disk []rune) int {
	checksum := 0

	for idx, val := range disk {
		if val != '.' {
			checksum += idx * runeToInt(val)
		}
	}
	return checksum
}

func computeChecksumString(disk []string) int {
	checksum := 0

	for idx, val := range disk {
		if val != "." {
			val, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			checksum += idx * val
		}
	}
	return checksum
}

func runeToInt(val rune) int {
	return int(val - '0')
}

func intToRune(val int) rune {
	return rune('0' + val)
}

func partOne(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		newDisk := reformatDisk([]rune(scanner.Text()))

		return reorganizeSpace(newDisk)
	}
	return 0
}

func partTwo(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		newDisk := reformatDiskString([]rune(scanner.Text()))

		return reorganizeSpaceFullFiles(newDisk)
	}
	return 0
}
