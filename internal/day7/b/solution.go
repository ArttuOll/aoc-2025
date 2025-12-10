package b

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ArttuOll/aoc-2025/internal/input"
)

type CacheKey struct {
	row int
	col int
}

func countPossiblePaths(grid [][]string, row int, col int, cache map[CacheKey]int) int {
	// We're past the last row
	if row >= len(grid) {
		return 1
	}

	cacheKey := CacheKey{row, col}
	if paths, exists := cache[cacheKey]; exists {
		return paths
	}

	count := 0

	cell := grid[row][col]
	if cell == "^" {
		// Found a splitter, count possible paths on its both sides
		count = countPossiblePaths(grid, row+1, col-1, cache) + countPossiblePaths(grid, row+1, col+1, cache)
	} else {
		count = countPossiblePaths(grid, row+1, col, cache)
	}

	cache[CacheKey{row, col}] = count
	return count
}

func Solve(inputFilePath string) error {
	input, err := input.Read(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to read the input: %w", err)
	}

	grid := make([][]string, len(input))
	for row, line := range input {
		grid[row] = strings.Split(line, "")
	}

	startingPoint := slices.Index(grid[0], "S")

	count := countPossiblePaths(grid, 0, startingPoint, make(map[CacheKey]int))

	fmt.Println(count)

	return nil
}
