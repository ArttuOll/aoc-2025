package a

import (
	"fmt"
	"strconv"

	"github.com/ArttuOll/aoc-2025/internal/input"
)

type Direction int

const (
	Left  Direction = iota
	Right Direction = 1
)

type Rotation struct {
	direction Direction
	times     int
}

func ParseRotation(rotationString string) (*Rotation, error) {
	result := Rotation{}

	directionPart := string(rotationString[0])
	switch directionPart {
	case "L":
		result.direction = Left
	case "R":
		result.direction = Right
	default:
		return nil, fmt.Errorf("unable to parse direction part of rotation")
	}

	timesPart, err := strconv.Atoi(rotationString[1:])
	if err != nil {
		return nil, fmt.Errorf("unable to parse times part of rotation: %w", err)
	}
	result.times = timesPart

	return &result, nil
}

type Dial struct {
	dialPointer int
}

func (dial *Dial) increment() {
	if dial.dialPointer == 99 {
		dial.dialPointer = 0
	} else {
		dial.dialPointer++
	}
}

func (dial *Dial) decrement() {
	if dial.dialPointer == 0 {
		dial.dialPointer = 99
	} else {
		dial.dialPointer--
	}
}

func (dial *Dial) rotate(rotation Rotation) {
	times := rotation.times % 100
	switch rotation.direction {
	case Left:
		for range times {
			dial.decrement()
		}
	case Right:
		for range times {
			dial.increment()
		}
	}
}

func Solve(inputFilePath string) error {
	input, err := input.Read(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to read the input: %w", err)
	}

	countOfPointingAtZero := 0
	dial := Dial{
		dialPointer: 50,
	}
	for _, rotationString := range input {
		rotation, err := ParseRotation(rotationString)
		if err != nil {
			return fmt.Errorf("unable to parse rotation: %w", err)
		}

		dial.rotate(*rotation)

		if dial.dialPointer == 0 {
			countOfPointingAtZero++
		}
	}

	fmt.Println(countOfPointingAtZero)

	return nil
}
