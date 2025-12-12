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

	parent *JunctionBox
}

func newJunctionBox(x, y, z float64, parent *JunctionBox) JunctionBox {
	jb := JunctionBox{}
	jb.x = x
	jb.y = y
	jb.z = z
	jb.parent = &jb
	return jb
}

func (jb *JunctionBox) find(other JunctionBox) JunctionBox {
	// We have reached the root
	if other.parent == &other {
		return other
	}

	return jb.find(*other.parent)
}

func (jb *JunctionBox) union(other JunctionBox) {
	root1 := jb.find(*jb)
	root2 := other.find(other)

	// The roots are different, the boxes are not part of the same circuit
	if root1 != root2 {
		root2.parent = &root1
	}
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

func (c *Circuit) Equals(other Circuit) bool {
	if len(c.junctionBoxes) != len(other.junctionBoxes) {
		return false
	}

	for junctionBox := range c.junctionBoxes {
		if _, ok := other.junctionBoxes[junctionBox]; !ok {
			return false
		}
	}

	for junctionBox := range other.junctionBoxes {
		if _, ok := c.junctionBoxes[junctionBox]; !ok {
			return false
		}
	}

	return true
}

func (c *Circuit) Merge(other Circuit) {
	for junctionBox := range other.junctionBoxes {
		c.junctionBoxes[junctionBox] = true
	}
}

func findClosestTogether(circuits []Circuit) [2]Circuit {
	shortestDistance := 100000.0
	var shortest [2]Circuit

	for i := range circuits {
		first := circuits[i]
		for j := range circuits {
			second := circuits[j]

			// The circuits are the same
			if first.Equals(second) {
				continue
			}

			for fjb := range first.junctionBoxes {
				for sjb := range second.junctionBoxes {
					if distance := fjb.DistanceTo(sjb); distance < shortestDistance {
						shortestDistance = distance
						shortest = [2]Circuit{first, second}
					}
				}
			}
		}
	}

	return shortest
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

		circuits = append(circuits, circuit)
	}

	for range 10 {
		closest := findClosestTogether(circuits)
		first := closest[0]
		second := closest[1]
		first.Merge(second)

		slices.Delete(circuits, slices.Index(circuits, first))
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
