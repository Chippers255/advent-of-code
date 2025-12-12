package parse

import (
	"bytes"
	"fmt"
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
// Returns an error if the input does not contain any rows.
func RuneGrid(data []byte) ([][]rune, error) {
	trimmed := bytes.TrimRight(data, "\n")
	if len(trimmed) == 0 {
		return nil, fmt.Errorf("no rune rows found in input")
	}

	lines := bytes.Split(trimmed, []byte("\n"))
	grid := make([][]rune, 0, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			return nil, fmt.Errorf("empty line found in input")
		}
		grid = append(grid, []rune(string(line)))
	}
	return grid, nil
}

// DigitGrid converts each non-empty line of digits into a row of ints.
func DigitGrid(data []byte) ([][]int, error) {
	rawLines := bytes.Split(bytes.TrimSpace(data), []byte("\n"))
	grid := make([][]int, 0, len(rawLines))

	for _, raw := range rawLines {
		line := bytes.TrimSpace(raw)
		if len(line) == 0 {
			continue
		}

		row := make([]int, len(line))
		for i, b := range line {
			if b < '0' || b > '9' {
				return nil, fmt.Errorf("invalid digit %q in input", b)
			}
			row[i] = int(b - '0')
		}
		grid = append(grid, row)
	}

	if len(grid) == 0 {
		return nil, fmt.Errorf("no digit rows found in input")
	}

	return grid, nil
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

// read a list of ranges and return as a 2d slice of ints
func RangeList(data []byte) ([][]int, error) {
	line := strings.TrimSpace(string(data))
	if line == "" {
		return nil, fmt.Errorf("empty input")
	}

	ranges := strings.Split(line, ",")
	if len(ranges) == 0 {
		return nil, fmt.Errorf("no ranges found")
	}
	result := make([][]int, 0, len(ranges))

	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if r == "" {
			return nil, fmt.Errorf("empty range")
		}

		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid range format: %s (expected start-end)", r)
		}

		start, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid start value in range %s: %w", r, err)
		}

		end, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, fmt.Errorf("invalid end value in range %s: %w", r, err)
		}

		result = append(result, []int{start, end})
	}

	return result, nil
}

// RangeLinesAndInts parses input split by a blank line into two sections.
// The first section contains dash-separated ranges (one per line) returned as [][]int.
// The second section contains integers (one per line) returned as []int.
func RangeLinesAndInts(data []byte) ([][]int, []int, error) {
	normalized := strings.ReplaceAll(string(data), "\r\n", "\n")
	parts := strings.Split(strings.TrimSpace(normalized), "\n\n")
	if len(parts) != 2 {
		return nil, nil, fmt.Errorf("expected two sections separated by a blank line, got %d", len(parts))
	}

	parseRange := func(line string) ([]int, error) {
		fields := strings.Split(line, "-")
		if len(fields) != 2 {
			return nil, fmt.Errorf("invalid range %q (expected start-end)", line)
		}
		start, err := strconv.Atoi(strings.TrimSpace(fields[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid start value in %q: %w", line, err)
		}
		end, err := strconv.Atoi(strings.TrimSpace(fields[1]))
		if err != nil {
			return nil, fmt.Errorf("invalid end value in %q: %w", line, err)
		}
		return []int{start, end}, nil
	}

	rangeLines := strings.Split(strings.TrimSpace(parts[0]), "\n")
	ranges := make([][]int, 0, len(rangeLines))
	for _, raw := range rangeLines {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		r, err := parseRange(line)
		if err != nil {
			return nil, nil, err
		}
		ranges = append(ranges, r)
	}

	targetLines := strings.Split(strings.TrimSpace(parts[1]), "\n")
	targets := make([]int, 0, len(targetLines))
	for _, raw := range targetLines {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		value, err := strconv.Atoi(line)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid value %q: %w", line, err)
		}
		targets = append(targets, value)
	}

	return ranges, targets, nil
}
