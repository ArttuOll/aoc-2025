package a

import (
	"fmt"

	"github.com/ArttuOll/aoc-2025/internal/input"
)

func Solve(inputFilePath string) error {
	input, err := input.Read(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to read the input: %w", err)
	}

	return nil
}
