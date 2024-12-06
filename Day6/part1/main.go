package main

import (
	"fmt"
	"os"
	"strings"
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

func partOne(rawInput string) int {
	grid, startPosition, startDirection := loadGrid(rawInput)
	visitedPositions := getVisitedPositions(grid, startPosition, startDirection)
	return len(visitedPositions)
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	rawInput := string(data)
	result := partOne(rawInput)
	fmt.Printf("Part One: %d\n", result)
}
