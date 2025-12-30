package main

import (
	_ "embed"
	"testing"

	"github.com/Dageus/advent-of-code/go/utils"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	const sample_output = 7
	const expected = 505

	assert.Equal(t, partOne(utils.Sample()), sample_output)
	assert.Equal(t, partOne(utils.Input()), expected)
}

func TestPartTwo(t *testing.T) {
	const sample_output = 33
	const expected = 20002

	assert.Equal(t, partTwo(utils.Sample()), sample_output)
	assert.Equal(t, partTwo(utils.Input()), expected)
}

func BenchmarkPartOne(b *testing.B) {
	var inputDay = utils.Input()
	for range b.N {
		partOne(inputDay)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	var inputDay = utils.Input()
	for range b.N {
		partTwo(inputDay)
	}
}
