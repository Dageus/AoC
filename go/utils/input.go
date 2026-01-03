package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Input returns the input filename based on the calling file's name.
// e.g., if called from "day01_test.go", it looks for "inputs/day01.input"
// or it downloads the file
func Input() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Could not recover caller information")
	}

	base := filepath.Base(filename)

	name := strings.TrimSuffix(base, ".go")
	name = strings.TrimSuffix(name, "_test")

	dir := filepath.Dir(filename)
	dir = filepath.Dir(dir)

	year := filepath.Base(dir)

	dir = filepath.Dir(dir)
	dir = filepath.Dir(dir)

	inputPath := filepath.Join(dir, "inputs", year, name+".input")

	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		fmt.Println("Test input doesn't exist, downloading from ")
		if b, err := DownloadInput(); err == nil {
			if err := os.WriteFile(filename, b, 0644); err != nil {
				panic(err)
			}
			return inputPath
		}
	}
	return inputPath
}

// Sample returns the sample filename based on the calling file's name.
// e.g., if called from "day01_test.go", it looks for "inputs/day01.sample"
func Sample() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Could not recover caller information")
	}

	base := filepath.Base(filename)

	name := strings.TrimSuffix(base, ".go")
	name = strings.TrimSuffix(name, "_test")

	dir := filepath.Dir(filename)
	dir = filepath.Dir(dir)

	year := filepath.Base(dir)

	dir = filepath.Dir(dir)
	dir = filepath.Dir(dir)

	samplePath := filepath.Join(dir, "inputs", year, name+".sample")

	if _, err := os.Stat(samplePath); os.IsNotExist(err) {
		panic(fmt.Errorf("Sample input doesn't exist, please create and populate %v with the sample.", samplePath))
	}
	return samplePath
}
