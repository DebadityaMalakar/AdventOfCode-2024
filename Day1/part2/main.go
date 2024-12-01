package main

import (
	"bufio"
	"fmt"
	"os"
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

	var l1 []int
	l2 := make(map[int]int)

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
		l2[rightNum]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	sum := 0
	for _, n1 := range l1 {
		sum += n1 * l2[n1]
	}

	fmt.Println(sum)
}
