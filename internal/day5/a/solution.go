package a

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ArttuOll/aoc-2025/internal/input"
)

type Range struct {
	start int
	end   int
}

func (r *Range) Parse(input string) error {
	parts := strings.Split(input, "-")
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("failed to parse start of range: %s", parts[0])
	}
	r.start = start

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("failed to parse end of range: %s", parts[1])
	}
	r.end = end

	return nil
}

func (r *Range) Contains(id int) bool {
	return id >= r.start && id <= r.end
}

func Solve(inputFilePath string) error {
	input, err := input.Read(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to read the input: %w", err)
	}

	var ranges []Range
	var ids []int
	iteratingRangeStrings := true
	for _, line := range input {
		if line == "" {
			iteratingRangeStrings = false
			continue
		}

		if iteratingRangeStrings {
			idRange := Range{}
			err := idRange.Parse(line)
			if err != nil {
				return err
			}

			ranges = append(ranges, idRange)
		} else {
			id, err := strconv.Atoi(line)
			if err != nil {
				return fmt.Errorf("failed to parse ID string to int: %s", line)
			}

			ids = append(ids, id)
		}
	}

	freshCount := 0
	for _, id := range ids {
		for _, idRange := range ranges {
			if idRange.Contains(id) {
				freshCount++
				break
			}
		}
	}

	fmt.Println(freshCount)

	return nil
}
