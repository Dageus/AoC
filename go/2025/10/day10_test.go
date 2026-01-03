package main

import (
	_ "embed"
	"testing"

	"github.com/Dageus/advent-of-code/go/utils"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	const sample_output = int64(7)
	const expected = int64(505)

	assert.Equal(t, sample_output, partOne(utils.Sample()))
	assert.Equal(t, expected, partOne(utils.Input()))
}

func TestPartTwo(t *testing.T) {
	const sample_output = int64(33)
	const expected = int64(20002)

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
