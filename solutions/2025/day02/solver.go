package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chippers255/advent-of-code/pkg/solver"
	"github.com/chippers255/advent-of-code/pkg/util/parse"
)

func init() {
	solver.Register(2025, 2, func() solver.Solver {
		return &Solver{}
	})
}

// Solver contains the logic for Advent of Code 2025 Day 2.
type Solver struct{}

// look for duplicate digits in the value
func checkBad1(value int) bool {
	strValue := strconv.Itoa(value)
	length := len(strValue)

	if length%2 != 0 {
		return false
	}

	half := length / 2
	firstHalf := strValue[:half]
	secondHalf := strValue[half:]
	return firstHalf == secondHalf
}

// Part1 solves the first part of the puzzle.
func (Solver) Part1(ctx solver.Context) (string, error) {
	ranges, err := parse.RangeList(ctx.Input)
	if err != nil {
		return "", err
	}

	answer := 0
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			if checkBad1(i) {
				answer += i
			}
		}
	}

	return fmt.Sprintf("%d", answer), nil
}

// look for repeating digits in the value
func checkBad2(value int) bool {
	strValue := strconv.Itoa(value)
	length := len(strValue)
	half := length / 2

	for i := 1; i <= half; i++ {
		if length%i != 0 {
			continue
		}
		digits := strValue[:i]
		count := strings.Count(strValue, digits)
		if count == length/i {
			return true
		}
	}
	return false
}

// Part2 solves the second part of the puzzle.
func (Solver) Part2(ctx solver.Context) (string, error) {
	ranges, err := parse.RangeList(ctx.Input)
	if err != nil {
		return "", err
	}

	answer := 0
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			if checkBad2(i) {
				answer += i
			}
		}
	}

	return fmt.Sprintf("%d", answer), nil
}
