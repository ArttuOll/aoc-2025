package a

import (
	"cmp"
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
	size   int
}

func (jb *JunctionBox) Find() *JunctionBox {
	for jb.parent != jb {
		jb.parent = jb.parent.parent
		jb = jb.parent
	}

	return jb
}

func (jb *JunctionBox) Union(other *JunctionBox) {
	root1 := jb.Find()
	root2 := other.Find()

	// The roots are different, the boxes are not part of the same circuit
	if root1 != root2 {
		if root1.size < root2.size {
			root1.parent = root2
			root2.size += root1.size
		} else {
			root2.parent = root1
			root1.size += root2.size
		}
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
	jb.parent = jb
	jb.size = 1

	return nil
}

func (jb *JunctionBox) DistanceTo(other *JunctionBox) float64 {
	return math.Sqrt(math.Pow(jb.x-other.x, 2) + math.Pow(jb.y-other.y, 2) + math.Pow(jb.z-other.z, 2))
}

func getPairs(junctionBoxes []*JunctionBox) [][2]*JunctionBox {
	var result [][2]*JunctionBox
	n := len(junctionBoxes)

	for i := range n {
		for j := i + 1; j < n; j++ {
			result = append(result, [2]*JunctionBox{junctionBoxes[i], junctionBoxes[j]})
		}
	}

	return result
}

func sortJunctionBoxPairsByDistance(pairs [][2]*JunctionBox) {
	slices.SortFunc(pairs, func(a, b [2]*JunctionBox) int {
		a1 := a[0]
		a2 := a[1]
		b1 := b[0]
		b2 := b[1]

		return cmp.Compare(a1.DistanceTo(a2), b1.DistanceTo(b2))
	})
}

func sortJunctionBoxesBySize(junctionBoxes []*JunctionBox) {
	slices.SortFunc(junctionBoxes, func(a, b *JunctionBox) int {
		return cmp.Compare(a.size, b.size)
	})
}

func Solve(inputFilePath string) error {
	input, err := input.Read(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to read the input: %w", err)
	}

	var junctionBoxes []*JunctionBox
	for _, line := range input {
		junctionBox := &JunctionBox{}
		junctionBox.Parse(line)
		junctionBoxes = append(junctionBoxes, junctionBox)
	}

	junctionBoxPairs := getPairs(junctionBoxes)
	sortJunctionBoxPairsByDistance(junctionBoxPairs)
	for i := range 1000 {
		pair := junctionBoxPairs[i]
		pair[0].Union(pair[1])
	}

	sortJunctionBoxesBySize(junctionBoxes)

	longest := junctionBoxes[len(junctionBoxes)-1]
	secondLongest := junctionBoxes[len(junctionBoxes)-2]
	thirdLongest := junctionBoxes[len(junctionBoxes)-3]

	fmt.Println(longest.size * secondLongest.size * thirdLongest.size)

	return nil
}
