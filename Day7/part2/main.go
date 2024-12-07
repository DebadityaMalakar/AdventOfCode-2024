package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Define the function type for operators
type operator func(int, int) int

// Parse the input file and return equations as slices of integers
func parseInput(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var equations [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return r < '0' || r > '9'
		})
		var eq []int
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				panic(err)
			}
			eq = append(eq, num)
		}
		equations = append(equations, eq)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return equations
}

// Generate the Cartesian product of operators repeated n times
func product(operators []operator, n int) [][]operator {
	if n == 0 {
		return [][]operator{}
	}
	if n == 1 {
		var result [][]operator
		for _, op := range operators {
			result = append(result, []operator{op})
		}
		return result
	}

	smaller := product(operators, n-1)
	var result [][]operator
	for _, op := range operators {
		for _, combo := range smaller {
			newCombo := append([]operator{op}, combo...)
			result = append(result, newCombo)
		}
	}
	return result
}

// Check if an equation satisfies the test value using given operators
func checkEq(eq []int, operators []operator) bool {
	testVal := eq[0]
	nums := eq[1:]

	for _, ops := range product(operators, len(nums)-1) {
		acc := nums[0]
		for i, op := range ops {
			acc = op(acc, nums[i+1])
		}
		if acc == testVal {
			return true
		}
	}
	return false
}

// Custom operator for concatenation
func concat(a, b int) int {
	multiplier := 1
	for b/multiplier > 0 {
		multiplier *= 10
	}
	return a*multiplier + b
}

// Solve part two of the problem
func partTwo(eqs [][]int) int {
	add := func(a, b int) int { return a + b }
	mul := func(a, b int) int { return a * b }
	operators := []operator{add, mul, concat}

	sum := 0
	for _, eq := range eqs {
		if checkEq(eq, operators) {
			sum += eq[0]
		}
	}
	return sum
}

func main() {
	eqs := parseInput("input.txt")
	fmt.Printf("Part 2: %d\n", partTwo(eqs))
}
