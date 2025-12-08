package b

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

func parseNumber(input []string) (int, error) {
	longestLength := 0
	for _, stringValue := range input {
		len := len(stringValue)
		if len > longestLength {
			longestLength = len
		}
	}

	numberString := ""
	for _, row := range input {
		if row == " " {
			continue
		}

		numberString += string(row[len(row)-1])
	}

	number, err := strconv.Atoi(numberString)
	if err != nil {
		return 0, fmt.Errorf("failed to convert string to int: %s", numberString)
	}

	return number, nil
}

func Solve(inputFilePath string) error {
	input, err := input.Read(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to read the input: %w", err)
	}

	grid := make([][]string, len(input))
	for i, line := range input {
		grid[i] = strings.Split(line, "")
	}

	var problems []Problem
	rowLen := len(grid[0])
	var numbers []int
	operand := ""
	for i := range rowLen {
		var valueStrings []string
		for _, row := range grid {
			valueString := row[i]
			if valueString == "*" || valueString == "+" {
				operand = valueString
				continue
			}

			valueStrings = append(valueStrings, valueString)
		}

		isEmpty := true
		for _, value := range valueStrings {
			if value != " " {
				isEmpty = false
				break
			}
		}

		if isEmpty {
			problems = append(problems, Problem{values: numbers, operand: operand})
			numbers = make([]int, 0)
			continue
		}

		newValue, err := parseNumber(valueStrings)
		if err != nil {
			return err
		}
		numbers = append(numbers, newValue)
	}
	problems = append(problems, Problem{values: numbers, operand: operand})

	result := 0
	for _, problem := range problems {
		result += problem.Result()
	}

	fmt.Println(result)

	return nil
}
