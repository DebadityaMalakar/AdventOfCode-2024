package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafe(r []int) bool {
	if r[0] > r[1] {
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
	}
	for i := 0; i < len(r)-1; i++ {
		if r[i+1]-r[i] < 1 || r[i+1]-r[i] > 3 {
			return false
		}
	}
	return true
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safe := 0

	for scanner.Scan() {
		line := scanner.Text()
		report := make([]int, 0)
		for _, numStr := range strings.Fields(line) {
			num, _ := strconv.Atoi(numStr)
			report = append(report, num)
		}

		if isSafe(report) {
			safe++
		} else {
			for level := range report {
				temp := append(append([]int{}, report[:level]...), report[level+1:]...)
				if isSafe(temp) {
					safe++
					break
				}
			}
		}
	}

	fmt.Println(safe)
}