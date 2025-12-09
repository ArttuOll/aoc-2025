package a

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

func countConnectedSplitters(grid [][]string, row int, col int, visitedSplitters map[CacheKey]bool) int {
	cacheKey := CacheKey{row, col}
	if _, exists := visitedSplitters[cacheKey]; exists {
		return 0
	}

	// We're past the last row
	if row >= len(grid) {
		return 0
	}

	cell := grid[row][col]

	if cell == "^" {
		// Found a splitter, continue downward on its both sides. Mark the splitter as visited.
		visitedSplitters[CacheKey{row, col}] = true
		return 1 + countConnectedSplitters(grid, row+1, col-1, visitedSplitters) + countConnectedSplitters(grid, row+1, col+1, visitedSplitters)
	}

	return countConnectedSplitters(grid, row+1, col, visitedSplitters)
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

	count := countConnectedSplitters(grid, 0, startingPoint, make(map[CacheKey]bool))

	fmt.Println(count)

	return nil
}
