package solver

import (
	"fmt"
	"sort"
	"sync"
)

var (
	registry = map[string]Factory{}
	mu       sync.RWMutex
)

// Metadata describes a registered solver.
type Metadata struct {
	Year int
	Day  int
}

// Register adds a solver factory for the given year/day.
func Register(year, day int, factory Factory) {
	key := makeKey(year, day)
	mu.Lock()
	defer mu.Unlock()
	if _, exists := registry[key]; exists {
		panic(fmt.Sprintf("solver already registered for %s", key))
	}
	registry[key] = factory
}

// Get looks up and constructs a solver for the given year/day.
func Get(year, day int) (Solver, error) {
	key := makeKey(year, day)
	mu.RLock()
	factory, ok := registry[key]
	mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("no solver registered for %s", key)
	}
	return factory(), nil
}

// List returns all registered solvers sorted by year/day.
func List() []Metadata {
	mu.RLock()
	defer mu.RUnlock()
	items := make([]Metadata, 0, len(registry))
	for key := range registry {
		year, day := parseKey(key)
		items = append(items, Metadata{Year: year, Day: day})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].Year == items[j].Year {
			return items[i].Day < items[j].Day
		}
		return items[i].Year < items[j].Year
	})
	return items
}

func makeKey(year, day int) string {
	return fmt.Sprintf("%04d-%02d", year, day)
}

func parseKey(key string) (int, int) {
	var year, day int
	fmt.Sscanf(key, "%d-%d", &year, &day)
	return year, day
}
