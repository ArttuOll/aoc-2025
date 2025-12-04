package b

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/ArttuOll/aoc-2025/internal/input"
)

func lastIndex(s []int, target int) int {
	result := 0
	for i, e := range s {
		if e == target {
			result = i
		}
	}

	return result
}

type Bank struct {
	batteries []int
}

func (b *Bank) Parse(inputString string) error {
	stringJoltages := strings.Split(inputString, "")

	var batteries []int
	for _, joltage := range stringJoltages {
		joltageInt, err := strconv.Atoi(joltage)
		if err != nil {
			return fmt.Errorf("unable to convert joltage to int: %s", joltage)
		}

		batteries = append(batteries, joltageInt)
	}

	b.batteries = batteries

	return nil
}

func pruneJoltage(digits []int) int {
	var filtered []string
	for _, digit := range digits {
		if digit == 0 {
			continue
		}

		filtered = append(filtered, strconv.Itoa(digit))
	}

	joltage, err := strconv.Atoi(strings.Join(filtered, ""))
	if err != nil {
		log.Fatal("combining joltages to an integer failed. You did something wrong.")
	}

	return joltage
}

func findLargestJoltage(batteries []int, digits []int, window []int, counter int) {
	if counter == 3 {
		return
	}

	max := slices.Max(window)
	indexOfMax := lastIndex(batteries, max)

	digits[indexOfMax] = max
	batteries[indexOfMax] = 0

	if len(batteries[indexOfMax:]) == 0 || slices.Max(batteries[indexOfMax:]) == 0 {
		findLargestJoltage(batteries, digits, batteries[:indexOfMax], counter+1)
	} else {
		findLargestJoltage(batteries, digits, batteries[indexOfMax:], counter+1)
	}
}

func (b *Bank) GetLargestJoltage() int {
	batteries := make([]int, len(b.batteries))
	copy(batteries, b.batteries)

	digits := make([]int, len(b.batteries))

	findLargestJoltage(batteries, digits, batteries, 0)

	return pruneJoltage(digits)
}

func Solve(inputFilePath string) error {
	input, err := input.Read(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to read the input: %w", err)
	}

	sum := 0
	for _, bankString := range input {
		bank := Bank{}
		bank.Parse(bankString)
		sum += bank.GetLargestJoltage()
	}

	fmt.Println(sum)

	return nil
}
