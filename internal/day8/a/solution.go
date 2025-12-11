package a

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/ArttuOll/aoc-2025/internal/input"
)

type JunctionBox struct {
	x float64
	y float64
	z float64
}

type Circuit struct {
	junctionBoxes []JunctionBox
}

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

func findClosestTogether(circuits []Circuit) [2]Circuit {
	shortestDistance := circuits[0].junctionBoxes[0].DistanceTo(circuits[1].junctionBoxes[0])
	var shortest [2]Circuit

	for i := range circuits {
		first := circuits[i]
		for j := range circuits {
			second := circuits[j]

			// We're not connecting circuits with more than one junction box
			if len(first.junctionBoxes) > 1 && len(second.junctionBoxes) > 1 {
				continue
			}

			// The circuits are the same
			if slices.Equal(first.junctionBoxes, second.junctionBoxes) {
				continue
			}

			// ??
			if distance := first.DistanceTo(second); distance < shortestDistance {
				shortestDistance = distance
				shortest = [2]JunctionBox{first, second}
			}
		}
	}

	return shortest
}

func connectJunctionBoxesToCircuitOrReturnNew(junctionBoxes [2]JunctionBox) Circuit {
	first := junctionBoxes[0]
	second := junctionBoxes[1]

	if circuit := first.circuit; circuit != nil {
		circuitMap := *circuit
		circuitMap[second] = true
		return nil
	}

	if circuit := second.circuit; circuit != nil {
		circuitMap := *circuit
		circuitMap[first] = true
		return nil
	}

	circuit := map[JunctionBox]bool{
		first:  true,
		second: true,
	}

	return circuit
}

func Solve(inputFilePath string) error {
	input, err := input.Read(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to read the input: %w", err)
	}

	var circuits []Circuit
	for _, line := range input {
		junctionBox := JunctionBox{}
		junctionBox.Parse(line)
		circuit := Circuit{
			junctionBoxes: map[JunctionBox]bool{junctionBox: true},
		}
	}

	for range 10 {
		closest := findClosestTogether(junctionBoxes)
		newCircuit := connectJunctionBoxesToCircuitOrReturnNew(closest)
		if newCircuit != nil {
			circuits = append(circuits, newCircuit)
		}
	}

	slices.SortFunc(circuits, func(a Circuit, b Circuit) int {
		if len(a) < len(b) {
			return -1
		} else if len(a) > len(b) {
			return 1
		}

		return 0
	})

	longest := circuits[len(circuits)-1]
	secondLongest := circuits[len(circuits)-2]
	thirdLongest := circuits[len(circuits)-3]

	fmt.Println(len(longest) * len(secondLongest) * len(thirdLongest))

	return nil
}
