package solution

import (
	"fmt"
	"strconv"

	day1a "github.com/ArttuOll/aoc-2025/internal/day1/a"
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

		}
	default:
		return fmt.Errorf("no puzzle found for day %v", day)
	}

	return nil
}
