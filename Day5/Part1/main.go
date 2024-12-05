package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read and split input
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content []string
	for scanner.Scan() {
		content = append(content, strings.TrimSpace(scanner.Text()))
	}

	splitIndex := 0
	for i, line := range content {
		if line == "" {
			splitIndex = i
			break
		}
	}

	rules := parseRules(content[:splitIndex])
	updates := parseUpdates(content[splitIndex+1:])

	nPart1 := 0

	// Check if all rules are correct for each update
	for _, values := range updates {
		relevantRules := filterRelevantRules(rules, values)
		if allCorrect(values, relevantRules) {
			midIndex := len(values) / 2
			val := atoi(values[midIndex])
			nPart1 += val
		}
	}

	fmt.Printf("%d\n", nPart1)
}

func parseRules(lines []string) [][2]string {
	var rules [][2]string
	for _, line := range lines {
		parts := strings.Split(line, "|")
		rules = append(rules, [2]string{parts[0], parts[1]})
	}
	return rules
}

func parseUpdates(lines []string) [][]string {
	var updates [][]string
	for _, line := range lines {
		updates = append(updates, strings.Split(line, ","))
	}
	return updates
}

func filterRelevantRules(rules [][2]string, values []string) [][2]string {
	var relevantRules [][2]string
	for _, rule := range rules {
		if contains(values, rule[0]) && contains(values, rule[1]) {
			relevantRules = append(relevantRules, rule)
		}
	}
	return relevantRules
}

func allCorrect(values []string, rules [][2]string) bool {
	for _, rule := range rules {
		if indexOf(values, rule[0]) >= indexOf(values, rule[1]) {
			return false
		}
	}
	return true
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func indexOf(slice []string, item string) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

func atoi(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}
