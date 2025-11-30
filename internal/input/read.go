package input

import (
	"bufio"
	"fmt"
	"os"
)

func Read(inputFilePath string) ([]string, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read the input file. Make sure it's int the same directory as the solution file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
