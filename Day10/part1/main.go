package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"
)

type Step struct {
	X, Y, Height int
}

// PriorityQueue implements a min-heap for Step based on the distance map.
type PriorityQueue struct {
	steps       []Step
	distanceMap map[Step]int
}

func (pq *PriorityQueue) Len() int {
	return len(pq.steps)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.distanceMap[pq.steps[i]] < pq.distanceMap[pq.steps[j]]
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.steps[i], pq.steps[j] = pq.steps[j], pq.steps[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	pq.steps = append(pq.steps, x.(Step))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.steps
	n := len(old)
	step := old[n-1]
	pq.steps = old[:n-1]
	return step
}

func canReach(start Step, end Step, grid map[[2]int]Step) bool {
	distanceMap := make(map[Step]int)
	for _, step := range grid {
		distanceMap[step] = math.MaxInt / 2
	}
	distanceMap[start] = 0

	pq := &PriorityQueue{distanceMap: distanceMap}
	heap.Init(pq)
	heap.Push(pq, start)

	for pq.Len() > 0 {
		currentStep := heap.Pop(pq).(Step)
		if currentStep == end {
			return true
		}

		currentDistance := distanceMap[currentStep]
		neighbors := []Step{}

		// Gather neighbors
		directions := [][2]int{
			{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		}
		for _, dir := range directions {
			neighbor, exists := grid[[2]int{currentStep.X + dir[0], currentStep.Y + dir[1]}]
			if exists {
				neighbors = append(neighbors, neighbor)
			}
		}

		// Update distances
		for _, neighbor := range neighbors {
			cost := math.MaxInt / 2
			if neighbor.Height-currentStep.Height == 1 {
				cost = 1
			}
			newDistance := currentDistance + cost
			if distanceMap[neighbor] > newDistance {
				distanceMap[neighbor] = newDistance
				heap.Push(pq, neighbor)
			}
		}
	}

	return false
}

func solve(input string) string {
	lines := strings.Split(input, "\n")
	var steps []Step
	for i, line := range lines {
		for j, point := range line {
			if point >= '0' && point <= '9' {
				steps = append(steps, Step{i, j, int(point - '0')})
			}
		}
	}

	grid := make(map[[2]int]Step)
	for _, step := range steps {
		grid[[2]int{step.X, step.Y}] = step
	}

	var startingPoints, endingPoints []Step
	for _, step := range steps {
		if step.Height == 0 {
			startingPoints = append(startingPoints, step)
		} else if step.Height == 9 {
			endingPoints = append(endingPoints, step)
		}
	}

	count := 0
	for _, start := range startingPoints {
		for _, end := range endingPoints {
			if canReach(start, end, grid) {
				count++
			}
		}
	}
	return fmt.Sprintf("%d", count)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

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

	fmt.Println(solve(input.String()))
}
