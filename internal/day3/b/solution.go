package b

import (
	"fmt"
	"math"
	"strconv"

	"github.com/ArttuOll/aoc-2025/internal/input"
)

type CacheKey struct {
	subBank string
	counter int
}

func findLargestJoltage(batteries string, counter int, cache map[CacheKey]int) (int, error) {
	if counter == 0 {
		return 0, nil
	}

	// Shortcut: we need to return all of the remaining digits
	if len(batteries) == counter {
		result, err := strconv.Atoi(batteries)
		if err != nil {
			return 0, fmt.Errorf("converting bank slice %v to integer failed", result)
		}

		return result, nil
	}

	firstDigit, err := strconv.Atoi(string(batteries[0]))
	if err != nil {
		return 0, fmt.Errorf("converting joltage string %v to integer failed", firstDigit)
	}

	substring := batteries[1:]

	cacheKey := CacheKey{substring, counter - 1}
	// Start building the largest number recursively from the first digit
	withCurrentDigit := firstDigit * int(math.Pow(10, float64(counter-1)))
	if cached, exists := cache[cacheKey]; exists {
		withCurrentDigit += cached
	} else {
		value, err := findLargestJoltage(substring, counter-1, cache)
		if err != nil {
			return 0, err
		}

		cache[cacheKey] = value
		withCurrentDigit += value
	}

	// Start building the largest number recursively from the second digit
	cacheKey = CacheKey{substring, counter}
	withoutCurrentDigit := 0
	if cached, exists := cache[cacheKey]; exists {
		withoutCurrentDigit = cached
	} else {
		value, err := findLargestJoltage(substring, counter, cache)
		if err != nil {
			return 0, err
		}

		cache[cacheKey] = value
		withoutCurrentDigit = value
	}

	// Return the larger of them
	if withCurrentDigit > withoutCurrentDigit {
		return withCurrentDigit, nil
	} else {
		return withoutCurrentDigit, nil
	}
}

func GetLargestJoltage(bank string) (int, error) {
	cache := make(map[CacheKey]int)
	return findLargestJoltage(bank, 12, cache)
}

func Solve(inputFilePath string) error {
	input, err := input.Read(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to read the input: %w", err)
	}

	sum := 0
	for _, bankString := range input {
		value, err := GetLargestJoltage(bankString)
		if err != nil {
			return fmt.Errorf("failed to find largest joltage for bank %s", bankString)
		}

		sum += value
	}

	fmt.Println(sum)

	return nil
}
