package a

import (
	"fmt"
	"math"

	"github.com/ArttuOll/aoc-2025/internal/input"
)

type JunctionBox struct {
	x float64
	y float64
	z float64
}

func (jb *JunctionBox) DistanceTo(other *JunctionBox) float64 {
	return math.Sqrt(math.Pow(jb.x-other.x, 2) + math.Pow(jb.y-other.y, 2) + math.Pow(jb.z-other.z, 2))
}

func Solve(inputFilePath string) error {
	input, err := input.Read(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to read the input: %w", err)
	}

	fmt.Println()

	return nil
}
