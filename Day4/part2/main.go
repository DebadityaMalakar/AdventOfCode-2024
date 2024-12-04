package main

import (
	"bufio"
	"fmt"
	"os"
)

// Read input from file and convert it to a 2D rune slice
func readFile() [][]rune {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		input = append(input, []rune(text))
	}

	return input
}

// Check if the surrounding characters form a valid "XMAS" pattern
func checkXmasPattern(topLeft, topRight, bottomLeft, bottomRight rune) bool {
	return (topLeft == 'M' && topRight == 'M' && bottomLeft == 'S' && bottomRight == 'S') ||
		(topLeft == 'S' && topRight == 'S' && bottomLeft == 'M' && bottomRight == 'M') ||
		(topLeft == 'M' && topRight == 'S' && bottomLeft == 'M' && bottomRight == 'S') ||
		(topLeft == 'S' && topRight == 'M' && bottomLeft == 'S' && bottomRight == 'M')
}

// Check if coordinates are within bounds
func checkBounds(rows, cols, x, y int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

// Search for custom "XMAS" patterns
func searchXmas(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	matches := 0

	for x, row := range grid {
		for y := range row {
			if grid[x][y] != 'A' {
				continue
			}

			topLeftX, topLeftY := x-1, y-1
			topRightX, topRightY := x-1, y+1
			bottomLeftX, bottomLeftY := x+1, y-1
			bottomRightX, bottomRightY := x+1, y+1

			if !checkBounds(rows, cols, topLeftX, topLeftY) ||
				!checkBounds(rows, cols, topRightX, topRightY) ||
				!checkBounds(rows, cols, bottomLeftX, bottomLeftY) ||
				!checkBounds(rows, cols, bottomRightX, bottomRightY) {
				continue
			}

			topLeft := grid[topLeftX][topLeftY]
			topRight := grid[topRightX][topRightY]
			bottomLeft := grid[bottomLeftX][bottomLeftY]
			bottomRight := grid[bottomRightX][bottomRightY]

			if checkXmasPattern(topLeft, topRight, bottomLeft, bottomRight) {
				matches++
			}
		}
	}

	return matches
}

// Part 2: Search for custom "XMAS" patterns
func part2() {
	grid := readFile()
	matches := searchXmas(grid)
	fmt.Println("Part 2 answer:", matches)
}

func main() {
	part2()
}
