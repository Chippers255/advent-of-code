package day04

import (
	"fmt"

	"github.com/chippers255/advent-of-code/pkg/solver"
	"github.com/chippers255/advent-of-code/pkg/util/parse"
)

func init() {
	solver.Register(2025, 4, func() solver.Solver {
		return &Solver{}
	})
}

// Solver contains the logic for Advent of Code 2025 Day 4.
type Solver struct{}

func getSurroundingValues(grid [][]rune, x int, y int) []rune {
	values := make([]rune, 0, 8)
	height := len(grid)
	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			if dx == 0 && dy == 0 {
				continue
			}
			nx := x + dx
			ny := y + dy
			if ny < 0 || ny >= height {
				continue
			}
			if nx < 0 || nx >= len(grid[ny]) {
				continue
			}
			if grid[ny][nx] == '@' {
				values = append(values, grid[ny][nx])
			}
		}
	}
	return values
}

func isSafe(grid [][]rune, x int, y int) bool {
	values := getSurroundingValues(grid, x, y)
	return len(values) < 4
}

func removeValues(grid [][]rune, removes [][]int) [][]rune {
	for _, remove := range removes {
		grid[remove[1]][remove[0]] = 'x'
	}
	return grid
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
}

// Part1 solves the first part of the puzzle.
func (Solver) Part1(ctx solver.Context) (string, error) {
	grid, err := parse.RuneGrid(ctx.Input)
	if err != nil {
		return "", err
	}

	answer := 0
	for y, row := range grid {
		for x, cell := range row {
			if cell == '@' {
				if isSafe(grid, x, y) {
					answer++
				}
			}
		}
	}
	return fmt.Sprintf("%d", answer), nil
}

// Part2 solves the second part of the puzzle.
func (Solver) Part2(ctx solver.Context) (string, error) {
	grid, err := parse.RuneGrid(ctx.Input)
	if err != nil {
		return "", err
	}

	answer := 0

	// start a while true loop
	for {
		//printGrid(grid)
		iterationAnswer := 0
		removes := [][]int{}
		for y := range grid {
			for x, value := range grid[y] {
				if value != '@' {
					continue
				}
				if isSafe(grid, x, y) {
					iterationAnswer++
					removes = append(removes, []int{x, y})
				}
			}
		}
		if iterationAnswer == 0 {
			break
		}
		answer += iterationAnswer
		//fmt.Println(iterationAnswer)
		//fmt.Println("--------------------------------")
		grid = removeValues(grid, removes)
	}

	return fmt.Sprintf("%d", answer), nil
}
