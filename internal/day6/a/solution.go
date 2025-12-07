package a

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ArttuOll/aoc-2025/internal/input"
)

type Problem struct {
	values  []int
	operand string
}

func (p *Problem) Result() int {
	result := p.values[0]
	for _, value := range p.values[1:] {
		if p.operand == "*" {
			result *= value
		}

		if p.operand == "+" {
			result += value
		}
	}

	return result
}

func Solve(inputFilePath string) error {
	input, err := input.Read(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to read the input: %w", err)
	}

	grid := make([][]string, len(input))
	for i, line := range input {
		grid[i] = strings.Fields(line)
	}

	var problems []Problem
	rowLen := len(grid[0])
	for i := range rowLen {
		var values []int
		operand := ""
		for _, row := range grid {
			valueString := row[i]
			if valueString == "*" || valueString == "+" {
				operand = valueString
				continue
			}

			value, err := strconv.Atoi(valueString)
			if err != nil {
				return fmt.Errorf("failed to convert string to int: %s", row[i])
			}

			values = append(values, value)

		}

		problems = append(problems, Problem{values, operand})
	}

	result := 0
	for _, problem := range problems {
		result += problem.Result()
	}

	fmt.Println(result)

	return nil
}
