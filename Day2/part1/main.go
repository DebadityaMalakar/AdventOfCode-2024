package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ReactorReport struct {
	levels []int
}

func parseInput(inputContent string) []ReactorReport {
	var retval []ReactorReport
	scanner := bufio.NewScanner(strings.NewReader(inputContent))
	for scanner.Scan() {
		line := scanner.Text()
		numStrs := strings.Fields(line)
		var levels []int
		for _, numStr := range numStrs {
			num, _ := strconv.Atoi(numStr)
			levels = append(levels, num)
		}
		retval = append(retval, ReactorReport{levels: levels})
	}
	return retval
}

func checkFalling(l, r int) bool {
	return l > r && (l-r) >= 1 && (l-r) <= 3
}

func checkRising(l, r int) bool {
	return r > l && (r-l) >= 1 && (r-l) <= 3
}

func part1(data []ReactorReport) string {
	safeCount := 0
	for _, report := range data {
		lvl := report.levels
		check := checkRising
		if lvl[0] >= lvl[len(lvl)-1] {
			check = checkFalling
		}
		safe := true
		for i := 0; i < len(lvl)-1; i++ {
			if !check(lvl[i], lvl[i+1]) {
				safe = false
				break
			}
		}
		if safe {
			safeCount++
		}
	}
	return strconv.Itoa(safeCount)
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	rawData := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rawData += scanner.Text() + "\n"
	}
	parsedData := parseInput(rawData)
	// Call part1 separately
	resultOne := part1(parsedData)
	fmt.Printf("Result 1: %s\n", resultOne)
}
