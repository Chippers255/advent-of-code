package day01

import (
	"fmt"
	"strconv"

	"github.com/chippers255/advent-of-code/pkg/solver"
	"github.com/chippers255/advent-of-code/pkg/util/mathx"
	"github.com/chippers255/advent-of-code/pkg/util/parse"
)

func init() {
	solver.Register(2025, 1, func() solver.Solver {
		return &Solver{}
	})
}

// Solver contains the logic for Advent of Code 2025 Day 1.
type Solver struct{}

// function to move the dial
func moveDial(position int, direction string, amount int) int {
	if direction == "L" {
		return (position - amount) % 100
	} else {
		return (position + amount) % 100
	}
}

// function to parse the instruction (eg: "L10" -> ("L", 10))
func parseInstruction(instruction string) (string, int, error) {
	direction := string(instruction[0])
	amount, err := strconv.Atoi(instruction[1:])
	if err != nil {
		return "", 0, err
	}
	return direction, amount, nil
}

// Part1 solves the first part of the puzzle.
func (Solver) Part1(ctx solver.Context) (string, error) {
	currentPosition := 50
	zeroCount := 0
	lines := parse.Lines(ctx.Input)

	for _, line := range lines {
		direction, amount, err := parseInstruction(line)
		if err != nil {
			return "", err
		}
		currentPosition = moveDial(currentPosition, direction, amount)
		if currentPosition == 0 {
			zeroCount++
		}
	}
	return fmt.Sprintf("%d", zeroCount), nil
}

// count how many times the dial passes through 0
func zeroDetector(start int, end int) int {
	diff := end - start
	if diff == 0 {
		return 0
	}
	if diff > 0 {
		return mathx.FloorDiv(end, 100) - mathx.FloorDiv(start, 100)
	}
	return mathx.FloorDiv(start-1, 100) - mathx.FloorDiv(end-1, 100)
}

// move the dial, count the positions, then normalize the position
func moveDial2(position int, direction string, amount int) (int, int) {
	newPosition := 0
	if direction == "L" {
		newPosition = position - amount
	} else {
		newPosition = position + amount
	}
	normalizedPosition := newPosition % 100
	return zeroDetector(position, newPosition), normalizedPosition
}

// Part2 solves the second part of the puzzle.
func (Solver) Part2(ctx solver.Context) (string, error) {
	lines := parse.Lines(ctx.Input)
	position := 50
	zeroCount := 0
	tempZeroCount := 0
	for _, line := range lines {
		direction, amount, err := parseInstruction(line)
		if err != nil {
			return "", err
		}
		tempZeroCount, position = moveDial2(position, direction, amount)
		zeroCount += tempZeroCount
	}
	return fmt.Sprintf("%d", zeroCount), nil
}
