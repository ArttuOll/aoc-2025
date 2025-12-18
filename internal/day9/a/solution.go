package a

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ArttuOll/aoc-2025/internal/input"
)

type Tile struct {
	row int
	col int
}

func parseTile(input string) (Tile, error) {
	parts := strings.Split(input, ",")
	col, err := strconv.Atoi(parts[0])
	if err != nil {
		return Tile{}, fmt.Errorf("failed to parse column of tile %v: %v", parts[0], err)
	}

	row, err := strconv.Atoi(parts[1])
	if err != nil {
		return Tile{}, fmt.Errorf("failed to parse row of tile %v: %v", parts[1], err)
	}

	return Tile{
		row,
		col,
	}, nil

}

func getPairs(tiles []Tile) [][2]Tile {
	var result [][2]Tile
	n := len(tiles)

	for i := range n {
		for j := i + 1; j < n; j++ {
			result = append(result, [2]Tile{tiles[i], tiles[j]})
		}
	}

	return result
}

func calculateArea(pair [2]Tile) float64 {
	height := math.Abs(float64(pair[0].row)-float64(pair[1].row)) + 1
	width := math.Abs(float64(pair[0].col)-float64(pair[1].col)) + 1

	return height * width
}

func Solve(inputFilePath string) error {
	input, err := input.Read(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to read the input: %w", err)
	}

	var tiles []Tile
	for _, line := range input {
		tile, err := parseTile(line)
		if err != nil {
			return err
		}

		tiles = append(tiles, tile)
	}

	pairs := getPairs(tiles)

	largestArea := 0.0
	for _, pair := range pairs {
		area := calculateArea(pair)
		if area > largestArea {
			largestArea = area
		}
	}

	fmt.Printf("%f\n", largestArea)

	return nil
}
