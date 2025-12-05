package day03

import (
	"fmt"

	"github.com/chippers255/advent-of-code/pkg/solver"
	"github.com/chippers255/advent-of-code/pkg/util/parse"
)

func init() {
	solver.Register(2025, 3, func() solver.Solver {
		return &Solver{}
	})
}

// Solver contains the logic for Advent of Code 2025 Day 3.
type Solver struct{}

type cell struct {
	power    int
	position int
}

func examineBattery(battery []int) (int, error) {
	bestPower := cell{power: battery[0], position: 0}
	secondBestPower := cell{power: -1, position: 0}

	for i := 1; i < len(battery); i++ {
		if battery[i] > bestPower.power && i < len(battery)-1 {
			secondBestPower = cell{power: -1, position: 0}
			bestPower = cell{power: battery[i], position: i}
		} else if battery[i] > secondBestPower.power {
			secondBestPower = cell{power: battery[i], position: i}
		} else if secondBestPower.power == -1 {
			secondBestPower = cell{power: battery[i], position: i}
		}
	}

	if bestPower.position < secondBestPower.position {
		return bestPower.power*10 + secondBestPower.power, nil
	} else {
		return secondBestPower.power*10 + bestPower.power, nil
	}
}

// Part1 solves the first part of the puzzle.
func (Solver) Part1(ctx solver.Context) (string, error) {
	batteryGrid, err := parse.DigitGrid(ctx.Input)
	if err != nil {
		return "", err
	}

	answer := 0
	for _, battery := range batteryGrid {
		power, err := examineBattery(battery)
		if err != nil {
			return "", err
		}
		fmt.Println(power)
		answer += power
	}
	return fmt.Sprintf("%d", answer), nil
}

func buildBestCells(battery []int) []cell {
	bestCells := make([]cell, 12)
	if len(battery) == 0 {
		for i := range bestCells {
			bestCells[i] = cell{power: -1, position: 0}
		}
		return bestCells
	}

	limit := len(bestCells)
	if len(battery) < limit {
		limit = len(battery)
	}

	for i := 0; i < limit; i++ {
		if i == 0 {
			bestCells[i] = cell{power: battery[i], position: i}
		} else {
			bestCells[i] = cell{power: -1, position: 0}
		}
	}
	for i := limit; i < len(bestCells); i++ {
		bestCells[i] = cell{power: -1, position: 0}
	}

	return bestCells
}

func resetBestCells(bestCells []cell, position int) []cell {
	for i := 0; i < len(bestCells); i++ {
		if i > position {
			bestCells[i] = cell{power: -1, position: 0}
		}
	}
	return bestCells
}

func examineBigBattery(battery []int) (int, error) {
	bestCells := buildBestCells(battery)

	for i := 1; i < len(battery); i++ {
		for j := 0; j < len(bestCells); j++ {
			if battery[i] > bestCells[j].power && i < len(battery)-(11-j) {
				bestCells[j] = cell{power: battery[i], position: i}
				resetBestCells(bestCells, j)
				break
			}
		}
	}

	power := 0
	for i := 0; i < len(bestCells); i++ {
		digit := bestCells[i].power
		if digit < 0 {
			digit = 0
		}
		power = power*10 + digit
	}
	return power, nil
}

// Part2 solves the second part of the puzzle.
func (Solver) Part2(ctx solver.Context) (string, error) {
	batteryGrid, err := parse.DigitGrid(ctx.Input)
	if err != nil {
		return "", err
	}

	answer := 0
	for _, battery := range batteryGrid {
		power, err := examineBigBattery(battery)
		if err != nil {
			return "", err
		}
		fmt.Println(power)
		answer += power
	}
	return fmt.Sprintf("%d", answer), nil
}
