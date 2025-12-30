package main

import (
	_ "embed"
	"testing"

	"github.com/Dageus/advent-of-code/go/utils"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	const sample_output = 13
	const expected = 1428

	assert.Equal(t, partOne(utils.Sample()), sample_output)
	assert.Equal(t, partOne(utils.Input()), expected)
}

func TestPartTwo(t *testing.T) {
	const sample_output = 43
	const expected = 8936

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
