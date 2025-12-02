package day02

import (
	"fmt"

	"github.com/chippers255/advent-of-code/pkg/solver"
	"github.com/chippers255/advent-of-code/pkg/util/parse"
)

func init() {
	solver.Register(2024, 2, func() solver.Solver {
		return &Solver{}
	})
}

// Solver contains the logic for Advent of Code 2024 Day 2.
type Solver struct{}

// Part1 solves the first part of the puzzle.
func (Solver) Part1(ctx solver.Context) (string, error) {
	reports, err := parse.IntMatrix(ctx.Input)
	if err != nil {
		return "", err
	}

	safe := 0
	for _, report := range reports {
		if checkReport(report) {
			safe++
		}
	}

	return fmt.Sprintf("%d", safe), nil
}

// Part2 solves the second part of the puzzle.
func (Solver) Part2(ctx solver.Context) (string, error) {
	reports, err := parse.IntMatrix(ctx.Input)
	if err != nil {
		return "", err
	}

	safe := 0
	for _, report := range reports {
		if checkReport(report) {
			safe++
			continue
		}

		for i := range report {
			if checkReport(removeLevel(report, i)) {
				safe++
				break
			}
		}
	}

	return fmt.Sprintf("%d", safe), nil
}

type state int

const (
	unknown state = iota
	increasing
	decreasing
)

type reindeerGuard struct {
	current state
	valid   bool
}

func newReindeerGuard() *reindeerGuard {
	return &reindeerGuard{current: unknown, valid: true}
}

func (rg *reindeerGuard) Next(prev, curr int) {
	diff := curr - prev
	if diff == 0 || diff > 3 || diff < -3 {
		rg.valid = false
		return
	}

	if rg.current == unknown {
		if diff > 0 {
			rg.current = increasing
		} else {
			rg.current = decreasing
		}
		return
	}

	if rg.current == increasing && diff < 0 {
		rg.valid = false
	} else if rg.current == decreasing && diff > 0 {
		rg.valid = false
	}
}

func checkReport(report []int) bool {
	if len(report) <= 1 {
		return true
	}

	guard := newReindeerGuard()
	for i := 1; i < len(report); i++ {
		guard.Next(report[i-1], report[i])
		if !guard.valid {
			return false
		}
	}
	return true
}

func removeLevel(report []int, index int) []int {
	out := make([]int, 0, len(report)-1)
	out = append(out, report[:index]...)
	out = append(out, report[index+1:]...)
	return out
}
