package solution

import (
	"fmt"
	"strconv"

	day1a "github.com/ArttuOll/aoc-2025/internal/day1/a"
	day1b "github.com/ArttuOll/aoc-2025/internal/day1/b"
	day2a "github.com/ArttuOll/aoc-2025/internal/day2/a"
	day2b "github.com/ArttuOll/aoc-2025/internal/day2/b"
	day3a "github.com/ArttuOll/aoc-2025/internal/day3/a"
	day3b "github.com/ArttuOll/aoc-2025/internal/day3/b"
	day4a "github.com/ArttuOll/aoc-2025/internal/day4/a"
	day4b "github.com/ArttuOll/aoc-2025/internal/day4/b"
	day5a "github.com/ArttuOll/aoc-2025/internal/day5/a"
	day5b "github.com/ArttuOll/aoc-2025/internal/day5/b"
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
	case 3:
		switch section {
		case "a":
			return day3a.Solve(inputFilePath)
		case "b":
			return day3b.Solve(inputFilePath)
		}
	case 4:
		switch section {
		case "a":
			return day4a.Solve(inputFilePath)
		case "b":
			return day4b.Solve(inputFilePath)
		}
	case 5:
		switch section {
		case "a":
			return day5a.Solve(inputFilePath)
		case "b":
			return day5b.Solve(inputFilePath)
		}
	default:
		return fmt.Errorf("no puzzle found for day %v", day)
	}

	return nil
}
