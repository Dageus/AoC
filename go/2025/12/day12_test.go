package main

import (
	_ "embed"
	"testing"

	"github.com/Dageus/advent-of-code/go/utils"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	const expected = 550

	assert.Equal(t, expected, partOne(utils.Input()))
}

func BenchmarkPartOne(b *testing.B) {
	var inputDay = utils.Input()
	for b.Loop() {
		partOne(inputDay)
	}
}
