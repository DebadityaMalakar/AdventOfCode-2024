package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := "input.txt"
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var l1, l2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		spaceIdx := strings.IndexAny(line, " \t")
		if spaceIdx == -1 {
			fmt.Println("Invalid input format in line:", line)
			continue
		}

		leftNum, err1 := strconv.Atoi(strings.TrimSpace(line[:spaceIdx]))
		rightNum, err2 := strconv.Atoi(strings.TrimSpace(line[spaceIdx+1:]))
		if err1 != nil || err2 != nil {
			fmt.Printf("Error parsing numbers in line: %s\n", line)
			continue
		}

		l1 = append(l1, leftNum)
		l2 = append(l2, rightNum)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	sort.Ints(l1)
	sort.Ints(l2)

	sum := 0
	for i := 0; i < len(l1); i++ {
		sum += abs(l1[i] - l2[i])
	}

	fmt.Println(sum)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
