package a

import (
	"fmt"
	"strings"

	"github.com/ArttuOll/aoc-2025/internal/input"
)

type Grid struct {
	grid [][]string
}

func (g *Grid) Parse(input []string) {
	g.grid = make([][]string, len(input))
	for i, row := range input {
		g.grid[i] = strings.Split(row, "")
	}
}

func (g *Grid) HasTopLeftNeighbor(row int, col int) bool {
	if row == 0 {
		return false
	}

	if col == 0 {
		return false
	}

	return g.grid[row-1][col-1] == "@"
}

func (g *Grid) HasTopNeighbor(row int, col int) bool {
	if row == 0 {
		return false
	}

	return g.grid[row-1][col] == "@"
}

func (g *Grid) HasTopRightNeighbor(row int, col int) bool {
	if row == 0 {
		return false
	}

	if col == len(g.grid[row])-1 {
		return false
	}

	return g.grid[row-1][col+1] == "@"
}

func (g *Grid) HasRightNeighbor(row int, col int) bool {
	if col == len(g.grid[row])-1 {
		return false
	}

	return g.grid[row][col+1] == "@"
}

func (g *Grid) HasBottomRightNeighbor(row int, col int) bool {
	if col == len(g.grid[row])-1 {
		return false
	}

	if row == len(g.grid)-1 {
		return false
	}

	return g.grid[row+1][col+1] == "@"
}

func (g *Grid) HasBottomNeighbor(row int, col int) bool {
	if row == len(g.grid)-1 {
		return false
	}

	return g.grid[row+1][col] == "@"
}

func (g *Grid) HasBottomLeftNeighbor(row int, col int) bool {
	if col == 0 {
		return false
	}

	if row == len(g.grid)-1 {
		return false
	}

	return g.grid[row+1][col-1] == "@"
}

func (g *Grid) HasLeftNeighbor(row int, col int) bool {
	if col == 0 {
		return false
	}

	return g.grid[row][col-1] == "@"
}

func (g *Grid) HasLessThanFourNeighbors(row int, col int) bool {
	neighborsCount := 0

	if g.HasTopNeighbor(row, col) {
		neighborsCount++
	}

	if g.HasTopRightNeighbor(row, col) {
		neighborsCount++
	}

	if g.HasRightNeighbor(row, col) {
		neighborsCount++
	}

	if g.HasBottomRightNeighbor(row, col) {
		neighborsCount++
	}

	if g.HasBottomNeighbor(row, col) {
		neighborsCount++
	}

	if g.HasBottomLeftNeighbor(row, col) {
		neighborsCount++
	}

	if g.HasLeftNeighbor(row, col) {
		neighborsCount++
	}

	if g.HasTopLeftNeighbor(row, col) {
		neighborsCount++
	}

	return neighborsCount < 4
}

func Solve(inputFilePath string) error {
	input, err := input.Read(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to read the input: %w", err)
	}

	grid := Grid{}
	grid.Parse(input)

	result := 0
	for i, row := range grid.grid {
		for j, value := range row {
			if value == "@" && grid.HasLessThanFourNeighbors(i, j) {
				result++
			}
		}
	}

	fmt.Println(result)

	return nil
}
