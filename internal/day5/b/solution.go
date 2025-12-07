package b

import (
	"fmt"
	"slices"
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

func (r *Range) Merge(other Range) {
	if r.start > other.start {
		r.start = other.start
	}

	if r.end < other.end {
		r.end = other.end
	}
}

func (r *Range) Contains(other Range) bool {
	if r.start <= other.start && r.end >= other.end {
		return true
	}

	return false
}

func (r *Range) Overlaps(other Range) bool {
	if r.end >= other.start && other.end > r.end {
		return true
	}

	if other.end >= r.start && r.end > other.end {
		return true
	}

	return false
}

func (r *Range) isEmpty() bool {
	return r.start == 0 && r.end == 0
}

func mergeRanges(range1 Range, range2 Range) [2]Range {
	if range1.Contains(range2) {
		return [2]Range{range1, {}}
	} else if range2.Contains(range1) {
		return [2]Range{{}, range2}
	}

	if range1.Overlaps(range2) {
		range1.Merge(range2)
		return [2]Range{range1, {}}
	}

	return [2]Range{range1, range2}
}

func Solve(inputFilePath string) error {
	input, err := input.Read(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to read the input: %w", err)
	}

	var ranges []Range
	for _, line := range input {
		if line == "" {
			break
		}

		idRange := Range{}
		err := idRange.Parse(line)
		if err != nil {
			return err
		}
		ranges = append(ranges, idRange)
	}

	slices.SortFunc(ranges, func(range1 Range, range2 Range) int {
		if range1.start < range2.start {
			return -1
		} else if range1.start > range2.start {
			return 1
		} else {
			return 0
		}
	})

	mergedRanges := ranges
	for {
		for i := 1; i < len(mergedRanges); i++ {
			merged := mergeRanges(mergedRanges[i-1], mergedRanges[i])
			mergedRanges[i-1] = merged[0]
			mergedRanges[i] = merged[1]
		}

		var filtered []Range
		for _, idRange := range mergedRanges {
			if idRange.isEmpty() {
				continue
			}

			filtered = append(filtered, idRange)
		}

		if len(filtered) < len(mergedRanges) {
			mergedRanges = filtered
			continue
		}

		break
	}

	count := 0
	for _, idRange := range mergedRanges {
		count += (idRange.end - idRange.start) + 1
	}

	fmt.Println(count)

	return nil
}
