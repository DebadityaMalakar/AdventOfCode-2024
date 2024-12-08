package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func inBounds(pos [2]int, w, h int) bool {
	i, j := pos[0], pos[1]
	return i >= 0 && i < h && j >= 0 && j < w
}

func main() {
	// Read input data from file
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	var inputData [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		inputData = append(inputData, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	// Get grid dimensions
	width := len(inputData[0])
	height := len(inputData)

	// Antennas map where key=frequency and value=list of positions
	antennas := make(map[string][][2]int)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			freq := inputData[i][j]
			pos := [2]int{i, j}
			if freq != "." {
				antennas[freq] = append(antennas[freq], pos)
			}
		}
	}

	// Build list of unique antinode positions
	uniqueAntinodes := make(map[[2]int]bool)
	for _, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				a := positions[i]
				b := positions[j]
				slope := [2]int{b[0] - a[0], b[1] - a[1]}

				// Calculate antinodes
				antinodeA := [2]int{a[0] - slope[0], a[1] - slope[1]}
				antinodeB := [2]int{b[0] + slope[0], b[1] + slope[1]}

				if inBounds(antinodeA, width, height) {
					uniqueAntinodes[antinodeA] = true
				}
				if inBounds(antinodeB, width, height) {
					uniqueAntinodes[antinodeB] = true
				}
			}
		}
	}

	// Print the count of unique antinodes
	fmt.Println(len(uniqueAntinodes))
}
