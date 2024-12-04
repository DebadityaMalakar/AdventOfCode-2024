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

var directions = [8][2]int{
	{-1, 0},  // up
	{1, 0},   // down
	{0, -1},  // left
	{0, 1},   // right
	{-1, -1}, // up-left
	{-1, 1},  // up-right
	{1, -1},  // down-left
	{1, 1},   // down-right
}

// Check if coordinates are within bounds
func checkBounds(rows, cols, x, y int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

// Search for a specific word in all directions
func searchWord(grid [][]rune, word string) int {
	rows := len(grid)
	cols := len(grid[0])
	matches := 0

	for x, row := range grid {
		for y := range row {
			for _, d := range directions {
				match := true
				for i := 0; i < len(word); i++ {
					newX := x + i*d[0]
					newY := y + i*d[1]
					if !checkBounds(rows, cols, newX, newY) || grid[newX][newY] != rune(word[i]) {
						match = false
						break
					}
				}
				if match {
					matches++
				}
			}
		}
	}

	return matches
}

// Part 1: Search for occurrences of "XMAS"
func part1() {
	grid := readFile()
	matches := searchWord(grid, "XMAS")
	fmt.Println("Part 1 answer:", matches)
}

func main() {
	part1()
}
