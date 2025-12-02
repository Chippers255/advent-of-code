package input

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Loader knows how to locate puzzle inputs.
type Loader struct {
	BasePath string
}

// NewLoader creates a Loader anchored at the provided base path.
func NewLoader(base string) Loader {
	if base == "" {
		base = "inputs"
	}
	return Loader{BasePath: base}
}

// Load returns the bytes for the requested year/day. If sample is true it looks for -sample files.
func (l Loader) Load(year, day int, sample bool) ([]byte, error) {
	path := l.Path(year, day, sample)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read input %s: %w", path, err)
	}
	return data, nil
}

// Path returns the absolute path to the requested input.
func (l Loader) Path(year, day int, sample bool) string {
	suffix := ""
	if sample {
		suffix = "-sample"
	}
	filename := fmt.Sprintf("day%02d%s.txt", day, suffix)
	return filepath.Join(l.BasePath, fmt.Sprintf("%d", year), filename)
}

// ParseIntPairs interprets the provided input as space-separated integer pairs.
func ParseIntPairs(data []byte) ([]int, []int, error) {
	lines := bytes.Split(bytes.TrimSpace(data), []byte("\n"))
	left := make([]int, 0, len(lines))
	right := make([]int, 0, len(lines))

	for idx, rawLine := range lines {
		line := strings.TrimSpace(string(rawLine))
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 2 {
			return nil, nil, fmt.Errorf("parse line %d: want at least two columns, got %d", idx+1, len(fields))
		}

		a, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, nil, fmt.Errorf("parse line %d left value: %w", idx+1, err)
		}
		b, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, nil, fmt.Errorf("parse line %d right value: %w", idx+1, err)
		}

		left = append(left, a)
		right = append(right, b)
	}

	return left, right, nil
}
