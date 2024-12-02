package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Report []int
type Reports []Report

func loadInput(file string) (Reports, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	var reports Reports
	for _, line := range lines {
		if line == "" {
			continue
		}
		var report Report
		for _, str := range strings.Fields(line) {
			num, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
			report = append(report, num)
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func isSafe(report Report) bool {
	var diffs []int
	for i := 1; i < len(report); i++ {
		diffs = append(diffs, report[i]-report[i-1])
	}

	for i := 1; i < len(diffs); i++ {
		d := diffs[i]
		if math.Abs(float64(d)) < 1 || math.Abs(float64(d)) > 3 || (d^diffs[0]) < 0 {
			return false
		}
	}
	return true
}

func isSafeSubreport(report Report) bool {
	n := len(report)
	bitset := make([]int, n-1)
	for i := range bitset {
		bitset[i] = 1
	}
	bitset = append(bitset, 0)

	for {
		var subReport Report
		for i := 0; i < n; i++ {
			if bitset[i] == 1 {
				subReport = append(subReport, report[i])
			}
		}
		if isSafe(subReport) {
			return true
		}
		if !prevPermutation(bitset) {
			break
		}
	}
	return false
}

func prevPermutation(arr []int) bool {
	n := len(arr)
	i := n - 2
	for i >= 0 && arr[i] >= arr[i+1] {
		i--
	}
	if i < 0 {
		return false
	}

	j := n - 1
	for arr[j] <= arr[i] {
		j--
	}

	arr[i], arr[j] = arr[j], arr[i]
	sort.Sort(sort.Reverse(sort.IntSlice(arr[i+1:])))
	return true
}

func part1(reports Reports) int {
	count := 0
	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}
	return count
}

func part2(reports Reports) int {
	count := 0
	for _, report := range reports {
		if isSafe(report) || isSafeSubreport(report) {
			count++
		}
	}
	return count
}

func main() {
	actualValues, err := loadInput("input.txt")
	if err != nil {
		fmt.Println("Error loading input:", err)
		return
	}

	fmt.Printf("part1: %d\n", part1(actualValues))
	fmt.Printf("part2: %d\n", part2(actualValues))
}
