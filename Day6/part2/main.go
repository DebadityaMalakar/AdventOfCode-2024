package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type CellType int

const (
	SolidObject CellType = iota
	EmptySpace
)

type Direction struct {
	rowDelta int
	colDelta int
}

var (
	DirectionUp    = Direction{-1, 0}
	DirectionDown  = Direction{1, 0}
	DirectionLeft  = Direction{0, -1}
	DirectionRight = Direction{0, 1}
)

func loadGrid(rawInput string) ([][]CellType, [2]int, Direction) {
	grid := [][]CellType{}
	var position [2]int
	var direction Direction

	for rowIndex, row := range strings.Split(rawInput, "\n") {
		gridRow := []CellType{}
		for colIndex, char := range row {
			switch char {
			case '#':
				gridRow = append(gridRow, SolidObject)
			case '^':
				gridRow = append(gridRow, EmptySpace)
				position = [2]int{rowIndex, colIndex}
				direction = DirectionUp
			case 'v':
				gridRow = append(gridRow, EmptySpace)
				position = [2]int{rowIndex, colIndex}
				direction = DirectionDown
			case '<':
				gridRow = append(gridRow, EmptySpace)
				position = [2]int{rowIndex, colIndex}
				direction = DirectionLeft
			case '>':
				gridRow = append(gridRow, EmptySpace)
				position = [2]int{rowIndex, colIndex}
				direction = DirectionRight
			default:
				gridRow = append(gridRow, EmptySpace)
			}
		}
		grid = append(grid, gridRow)
	}
	return grid, position, direction
}

func copyGrid(grid [][]CellType) [][]CellType {
	newGrid := make([][]CellType, len(grid))
	for i := range grid {
		newGrid[i] = make([]CellType, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	return newGrid
}

func runStep(grid [][]CellType, position [2]int, direction Direction) ([2]int, Direction, bool) {
	row, col := position[0], position[1]
	newRow := row + direction.rowDelta
	newCol := col + direction.colDelta

	if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[0]) {
		return [2]int{}, Direction{}, true
	}

	if grid[newRow][newCol] == SolidObject {
		switch direction {
		case DirectionUp:
			return runStep(grid, position, DirectionRight)
		case DirectionDown:
			return runStep(grid, position, DirectionLeft)
		case DirectionLeft:
			return runStep(grid, position, DirectionUp)
		case DirectionRight:
			return runStep(grid, position, DirectionDown)
		}
	}

	return [2]int{newRow, newCol}, direction, false
}

func getVisitedPositions(grid [][]CellType, startPosition [2]int, startDirection Direction) map[[2]int]bool {
	position := startPosition
	direction := startDirection
	visitedPositions := make(map[[2]int]bool)
	finished := false

	for !finished {
		visitedPositions[position] = true
		position, direction, finished = runStep(grid, position, direction)
	}

	return visitedPositions
}

func isGuardInALoop(grid [][]CellType, startPosition [2]int, startDirection Direction) bool {
	// Floyd's Cycle-Finding Algorithm (Tortoise and Hare)
	position1, direction1 := startPosition, startDirection
	position2, direction2, exited := runStep(grid, startPosition, startDirection)

	if exited {
		return false
	}

	for position1 != position2 || direction1 != direction2 {
		position1, direction1, _ = runStep(grid, position1, direction1)
		position2, direction2, exited = runStep(grid, position2, direction2)
		if exited {
			return false
		}
		position2, direction2, exited = runStep(grid, position2, direction2)
		if exited {
			return false
		}
	}
	return true
}

func checkIfSettingPositionSolidMakesALoop(grid [][]CellType, position [2]int, startDirection Direction, startPosition [2]int) bool {
	// Create a copy of the grid to avoid modifying the original
	gridCopy := copyGrid(grid)

	// Set the position to solid object
	gridCopy[position[0]][position[1]] = SolidObject

	// Check if this creates a loop
	result := isGuardInALoop(gridCopy, startPosition, startDirection)

	return result
}

func partTwo(rawInput string) int {
	grid, startPosition, startDirection := loadGrid(rawInput)
	visitedPositions := getVisitedPositions(grid, startPosition, startDirection)

	loopCount := 0
	for position := range visitedPositions {
		// Skip the start position
		if position == startPosition {
			continue
		}

		if checkIfSettingPositionSolidMakesALoop(grid, position, startDirection, startPosition) {
			loopCount++
		}
	}

	return loopCount
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	rawInput := string(data)

	startTime := time.Now()
	result := partTwo(rawInput)
	duration := time.Since(startTime)

	fmt.Printf("Part Two: %d (Duration: %v)\n", result, duration)
}
