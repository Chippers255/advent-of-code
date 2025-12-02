package parse

import (
	"bytes"
	"strconv"
	"strings"
)

// Lines splits the input into non-empty trimmed lines.
func Lines(data []byte) []string {
	raw := bytes.Split(bytes.TrimSpace(data), []byte("\n"))
	lines := make([]string, 0, len(raw))
	for _, line := range raw {
		trimmed := strings.TrimSpace(string(line))
		if trimmed == "" {
			continue
		}
		lines = append(lines, trimmed)
	}
	return lines
}

// Ints converts each newline-delimited value into an int.
func Ints(data []byte) ([]int, error) {
	lines := Lines(data)
	nums := make([]int, 0, len(lines))
	for _, line := range lines {
		v, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		nums = append(nums, v)
	}
	return nums, nil
}

// RuneGrid converts the input into a 2D slice of runes.
func RuneGrid(data []byte) [][]rune {
	lines := bytes.Split(bytes.TrimRight(data, "\n"), []byte("\n"))
	grid := make([][]rune, 0, len(lines))
	for _, line := range lines {
		grid = append(grid, []rune(string(line)))
	}
	return grid
}

// IntMatrix treats each line as a whitespace-delimited list of integers.
func IntMatrix(data []byte) ([][]int, error) {
	lines := bytes.Split(bytes.TrimSpace(data), []byte("\n"))
	matrix := make([][]int, 0, len(lines))

	for _, raw := range lines {
		line := strings.TrimSpace(string(raw))
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		row := make([]int, len(fields))
		for i, field := range fields {
			val, err := strconv.Atoi(field)
			if err != nil {
				return nil, err
			}
			row[i] = val
		}
		matrix = append(matrix, row)
	}

	return matrix, nil
}
