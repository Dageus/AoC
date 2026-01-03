package main

import (
	_ "embed"
	"testing"

	"github.com/Dageus/advent-of-code/go/utils"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	const sample_output = 357
	const expected = 17144

	assert.Equal(t, sample_output, partOne(utils.Sample()))
	assert.Equal(t, expected, partOne(utils.Input()))
}

func TestPartTwo(t *testing.T) {
	const sample_output = 3121910778619
	const expected = 170371185255900

	assert.Equal(t, sample_output, partTwo(utils.Sample()))
	assert.Equal(t, expected, partTwo(utils.Input()))
}

func BenchmarkPartOne(b *testing.B) {
	var inputDay = utils.Input()
	for b.Loop() {
		partOne(inputDay)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	var inputDay = utils.Input()
	for b.Loop() {
		partTwo(inputDay)
	}
}
