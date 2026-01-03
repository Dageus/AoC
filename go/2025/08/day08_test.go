package main

import (
	_ "embed"
	"testing"

	"github.com/Dageus/advent-of-code/go/utils"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	const sample_output = 40
	const expected = 63920

	assert.Equal(t, sample_output, partOne(utils.Sample(), 10))
	assert.Equal(t, expected, partOne(utils.Input(), 1000))
}

func TestPartTwo(t *testing.T) {
	const sample_output = 25272
	const expected = 1026594680

	assert.Equal(t, sample_output, partTwo(utils.Sample()))
	assert.Equal(t, expected, partTwo(utils.Input()))
}

func BenchmarkPartOne(b *testing.B) {
	var inputDay = utils.Input()
	for b.Loop() {
		partOne(inputDay, 1000)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	var inputDay = utils.Input()
	for b.Loop() {
		partTwo(inputDay)
	}
}
