package a

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ArttuOll/aoc-2025/internal/input"
)

type JunctionBox struct {
	x       float64
	y       float64
	z       float64
	circuit *Circuit
}

type Circuit map[JunctionBox]bool

func (jb *JunctionBox) Parse(input string) error {
	values := strings.Split(input, ",")
	x, err := strconv.ParseFloat(values[0], 64)
	if err != nil {
		return fmt.Errorf("failed to convert x coordinate to int: %v", x)
	}
	y, err := strconv.ParseFloat(values[1], 64)
	if err != nil {
		return fmt.Errorf("failed to convert y coordinate to int: %v", y)
	}
	z, err := strconv.ParseFloat(values[2], 64)
	if err != nil {
		return fmt.Errorf("failed to convert z coordinate to int: %v", z)
	}

	jb.x = x
	jb.y = y
	jb.z = z

	return nil
}

func (jb *JunctionBox) DistanceTo(other JunctionBox) float64 {
	return math.Sqrt(math.Pow(jb.x-other.x, 2) + math.Pow(jb.y-other.y, 2) + math.Pow(jb.z-other.z, 2))
}

func findClosestTogether(junctionBoxes []JunctionBox) [2]JunctionBox {
	shortestDistance := junctionBoxes[0].DistanceTo(junctionBoxes[1])
	shortest := [2]JunctionBox{junctionBoxes[0], junctionBoxes[1]}
	for i := 0; i < len(junctionBoxes); i++ {
		first := junctionBoxes[i]
		for j := 0; i < len(junctionBoxes); i++ {
			second := junctionBoxes[j]

			// Both are already a part of some circuit
			if first.circuit != nil && second.circuit != nil {
				continue
			}

			if distance := first.DistanceTo(second); distance < shortestDistance {
				shortestDistance = distance
				shortest = [2]JunctionBox{first, second}
			}
		}
	}

	return shortest
}

func connectJunctionBoxesToCircuit(junctionBoxes [2]JunctionBox, circuits []Circuit) {
	if circuit := junctionBoxes[0].circuit; circuit != nil {
		array := *circuit
		array = append(array, junctionBoxes[1])
	}
}

func Solve(inputFilePath string) error {
	input, err := input.Read(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to read the input: %w", err)
	}

	var junctionBoxes []JunctionBox
	for _, line := range input {
		junctionBox := JunctionBox{}
		junctionBox.Parse(line)
		junctionBoxes = append(junctionBoxes, junctionBox)
	}

	var circuits []Circuit
	for range 1000 {
		closest := findClosestTogether(junctionBoxes)

	}

	fmt.Println(len(junctionBoxes))

	return nil
}
