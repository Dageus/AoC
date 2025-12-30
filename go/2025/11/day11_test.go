package main

import (
	_ "embed"
	"os"
	"testing"

	"github.com/Dageus/advent-of-code/go/utils"
	"github.com/stretchr/testify/assert"
)

const part1 = `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`

const part2 = `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`

func TestPartOne(t *testing.T) {
	const sample_output = 5
	const expected = 472

	// write the given test case to the sample
	os.WriteFile(utils.Sample(), []byte(part1), 0644)

	assert.Equal(t, partOne(utils.Sample()), sample_output)
	assert.Equal(t, partOne(utils.Input()), expected)
}

func TestPartTwo(t *testing.T) {
	const sample_output = 2
	const expected = 526811953334940

	// write the given test case to the sample
	os.WriteFile(utils.Sample(), []byte(part2), 0644)

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
