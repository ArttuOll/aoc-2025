package solution

import (
	"fmt"
	"strconv"

	day1a "github.com/ArttuOll/aoc-2025/internal/day1/a"
	day1b "github.com/ArttuOll/aoc-2025/internal/day1/b"
	day2a "github.com/ArttuOll/aoc-2025/internal/day2/a"
	day2b "github.com/ArttuOll/aoc-2025/internal/day2/b"
)

func Run(args []string) error {
	day, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("given day string couldn't be converted to an int: %v. %w", day, err)
	}

	section := args[1]
	inputFilePath := args[2]

	switch day {
	case 1:
		switch section {
		case "a":
			return day1a.Solve(inputFilePath)
		case "b":
			return day1b.Solve(inputFilePath)
		}
	case 2:
		switch section {
		case "a":
			return day2a.Solve(inputFilePath)
		case "b":
			return day2b.Solve(inputFilePath)
		}
	default:
		return fmt.Errorf("no puzzle found for day %v", day)
	}

	return nil
}
