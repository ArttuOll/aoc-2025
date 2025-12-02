package a

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ArttuOll/aoc-2025/internal/input"
)

func validateID(id int) bool {
	idString := strconv.Itoa(id)
	return idString[0:len(idString)/2] != idString[len(idString)/2:]
}

type IDRange struct {
	start int
	end   int
}

func (idRange *IDRange) Parse(rangeString string) error {
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

func (idRange *IDRange) InvalidIDs() []int {
	var invalidIds []int
	for i := idRange.start; i <= idRange.end; i++ {
		if !validateID(i) {
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

	var invalidIDs []int
	for _, idRangeString := range idRangeStrings {
		idRange := IDRange{}

		if err := idRange.Parse(idRangeString); err != nil {
			return fmt.Errorf("unable to parse ID range: %s", idRangeString)
		}

		invalidIDs = append(invalidIDs, idRange.InvalidIDs()...)
	}

	sum := 0
	for _, invalidID := range invalidIDs {
		sum += invalidID
	}

	fmt.Println(sum)

	return nil
}
