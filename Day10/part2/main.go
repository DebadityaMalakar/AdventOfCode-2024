package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Step struct {
	X, Y, Height int
}

// GetNeighbours retrieves the neighbors of a step from the map.
func (s Step) GetNeighbours(grid map[[2]int]Step) []Step {
	neighbors := []Step{}
	directions := [][2]int{
		{-1, 0}, {1, 0}, {0, 1}, {0, -1},
	}

	for _, dir := range directions {
		neighbor, exists := grid[[2]int{s.X + dir[0], s.Y + dir[1]}]
		if exists {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

// CountTrails counts the number of trails starting from a given step.
func CountTrails(end Step, grid map[[2]int]Step) int {
	count := 0
	toVisit := []Step{end}

	for len(toVisit) > 0 {
		// Dequeue the first element
		currentTrail := toVisit[0]
		toVisit = toVisit[1:]

		if currentTrail.Height == 0 {
			count++
		}

		// Add valid neighbors to the queue
		for _, neighbor := range currentTrail.GetNeighbours(grid) {
			if currentTrail.Height-neighbor.Height == 1 {
				toVisit = append(toVisit, neighbor)
			}
		}
	}

	return count
}

// Solve processes the input and computes the total number of trails.
func Solve(input string) string {
	lines := strings.Split(input, "\n")
	var steps []Step

	// Parse input into steps
	for i, line := range lines {
		for j, point := range line {
			if point >= '0' && point <= '9' {
				steps = append(steps, Step{X: i, Y: j, Height: int(point - '0')})
			}
		}
	}

	// Create a map for grid lookup
	grid := make(map[[2]int]Step)
	for _, step := range steps {
		grid[[2]int{step.X, step.Y}] = step
	}

	// Find all steps with height 9 and count trails from them
	totalTrails := 0
	for _, step := range steps {
		if step.Height == 9 {
			totalTrails += CountTrails(step, grid)
		}
	}

	return fmt.Sprintf("%d", totalTrails)
}

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	// Read the file contents
	var input strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input.WriteString(scanner.Text())
		input.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	// Solve the problem with the file input
	result := Solve(strings.TrimSpace(input.String()))
	fmt.Println(result)
}
