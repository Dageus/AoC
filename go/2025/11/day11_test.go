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

const part2 = `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`

func TestPartOne(t *testing.T) {
	const sample_output = 5
	const expected = 472

	// write the given test case to the sample
	os.WriteFile(utils.Sample(), []byte(part1), 0644)

	assert.Equal(t, sample_output, partOne(utils.Sample()))
	assert.Equal(t, expected, partOne(utils.Input()))
}

func TestPartTwo(t *testing.T) {
	const sample_output = 2
	const expected = 526811953334940

	// write the given test case to the sample
	os.WriteFile(utils.Sample(), []byte(part2), 0644)

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
