package a

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/ArttuOll/aoc-2025/internal/input"
)

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

func (b *Bank) GetLargestJoltage() int {
	largest := slices.Max(b.batteries)
	indexOfLargest := slices.Index(b.batteries, largest)

	var secondLargest int
	largestJoltageString := ""
	if indexOfLargest == len(b.batteries)-1 {
		secondLargest = slices.Max(b.batteries[:indexOfLargest])
		largestJoltageString = strconv.Itoa(secondLargest) + strconv.Itoa(largest)
	} else {
		secondLargest = slices.Max(b.batteries[indexOfLargest+1:])
		largestJoltageString = strconv.Itoa(largest) + strconv.Itoa(secondLargest)
	}

	largestJoltage, err := strconv.Atoi(largestJoltageString)
	if err != nil {
		log.Fatal("failed to convert largest joltage to integer. You did something wrong: %w", err)
	}
	return largestJoltage
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
