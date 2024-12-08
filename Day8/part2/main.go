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

	// Antenna dictionary where key = frequency, value = list of positions
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

				// Collect antinodes
				antinodes := []([2]int){a, b}

				// Antinodes moving away from `a`
				for k := 1; ; k++ {
					antinode := [2]int{a[0] - k*slope[0], a[1] - k*slope[1]}
					if !inBounds(antinode, width, height) {
						break
					}
					antinodes = append(antinodes, antinode)
				}

				// Antinodes moving away from `b`
				for k := 1; ; k++ {
					antinode := [2]int{b[0] + k*slope[0], b[1] + k*slope[1]}
					if !inBounds(antinode, width, height) {
						break
					}
					antinodes = append(antinodes, antinode)
				}

				// Add antinodes to the unique set
				for _, antinode := range antinodes {
					uniqueAntinodes[antinode] = true
				}
			}
		}
	}

	// Print the count of unique antinodes
	fmt.Println(len(uniqueAntinodes))
}
