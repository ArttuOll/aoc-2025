package b

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ArttuOll/aoc-2025/internal/input"
)

type TreeNode struct {
	id       string
	parents  []*TreeNode
	children []*TreeNode
}

func buildTree(grid [][]string, row, col int, nodes map[string]*TreeNode, parent *TreeNode) *TreeNode {
	// We're past the last row
	if row >= len(grid) || col < 0 || col >= len(grid[row]) {
		return nil
	}

	cell := grid[row][col]

	if cell != "^" && cell != "S" {
		return buildTree(grid, row+1, col, nodes, parent)
	}

	id := fmt.Sprintf("%v%v", row, col)

	node, exists := nodes[id]
	if !exists {
		node = &TreeNode{
			id:       id,
			children: []*TreeNode{},
			parents:  []*TreeNode{},
		}
		nodes[id] = node
	}

	if parent != nil {
		node.parents = append(node.parents, parent)
		parent.children = append(parent.children, node)
	}

	buildTree(grid, row+1, col-1, nodes, node)
	buildTree(grid, row+1, col+1, nodes, node)

	return node
}

func (t *TreeNode) getPaths(path []*TreeNode) [][]*TreeNode {
	newPath := make([]*TreeNode, len(path))
	copy(newPath, path)

	// We're at a leaf node
	if len(t.children) == 0 {
		return [][]*TreeNode{newPath}
	}

	var paths [][]*TreeNode
	for _, node := range t.children {
		childPaths := node.getPaths(newPath)
		paths = append(paths, childPaths...)
	}

	return paths
}

func Solve(inputFilePath string) error {
	input, err := input.Read(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to read the input: %w", err)
	}

	grid := make([][]string, len(input))
	for row, line := range input {
		grid[row] = strings.Split(line, "")
	}

	startingCol := slices.Index(grid[0], "S")

	root := TreeNode{
		id:       fmt.Sprintf("0%v", startingCol),
		children: make([]*TreeNode, 0),
		parents:  make([]*TreeNode, 0),
	}

	buildTree(grid, 0, startingCol, make(map[string]*TreeNode), &root)

	paths := root.getPaths(nil)

	fmt.Println(len(paths))

	return nil
}
