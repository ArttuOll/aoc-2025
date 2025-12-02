package b

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ArttuOll/aoc-2025/internal/input"
)

func allSame(idStrings []string) bool {
	for _, idString := range idStrings {
		if idString != idStrings[0] {
			return false
		}
	}

	return true
}

func validateId(id int) bool {
	idString := strconv.Itoa(id)

	for partLength := 1; partLength < len(idString); partLength++ {
		if len(idString)%partLength != 0 {
			// The ID can't be divided into equal length parts of this size
			continue
		}

		// Divide the ID into parts of equal length
		parts := make([]string, 0)
		for i := 0; i+partLength <= len(idString); i = i + partLength {
			parts = append(parts, idString[i:i+partLength])
		}

		if allSame(parts) {
			return false
		}
	}

	return true
}

type IdRange struct {
	start int
	end   int
}

func (idRange *IdRange) Parse(rangeString string) error {
	parts := strings.Split(rangeString, "-")
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("failed to parse start part of range: %w", err)
	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("failed to parse end part of range: %w", err)
	}

	idRange.start = start
	idRange.end = end

	return nil
}

func (idRange *IdRange) InvalidIds() []int {
	invalidIds := make([]int, 0)
	for i := idRange.start; i <= idRange.end; i++ {
		if !validateId(i) {
			invalidIds = append(invalidIds, i)
		}
	}

	return invalidIds
}

func Solve(inputFilePath string) error {
	input, err := input.Read(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to read the input: %w", err)
	}

	idRangeStrings := strings.Split(input[0], ",")

	invalidIds := make([]int, 0)
	for _, idRangeString := range idRangeStrings {
		idRange := IdRange{}
		idRange.Parse(idRangeString)
		invalidIds = append(invalidIds, idRange.InvalidIds()...)
	}

	sum := 0
	for _, invalidId := range invalidIds {
		sum += invalidId
	}

	fmt.Println(sum)

	return nil
}
