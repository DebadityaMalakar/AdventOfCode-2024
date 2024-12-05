package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Checks if the update is valid according to the ordering rules
func isValidUpdate(update []string, orderingRules map[string][]string) bool {
	for index, page := range update {
		previousPages := update[:index]
		shouldBeLaterPages := orderingRules[page]
		for _, p := range previousPages {
			if contains(shouldBeLaterPages, p) {
				return false
			}
		}
	}
	return true
}

// Corrects an invalid update based on ordering rules
func correctUpdate(update []string, orderingRules map[string][]string) []string {
	correctedUpdate := append([]string{}, update...)

	for !isValidUpdate(correctedUpdate, orderingRules) {
		for index, page := range correctedUpdate {
			previousPages := correctedUpdate[:index]
			shouldBeLaterPages := orderingRules[page]
			for _, p := range previousPages {
				if contains(shouldBeLaterPages, p) {
					firstPageToSwapIndex := indexOf(correctedUpdate, p)
					correctedUpdate = append(correctedUpdate[:firstPageToSwapIndex], correctedUpdate[firstPageToSwapIndex+1:]...)
					correctedUpdate = append(correctedUpdate[:index], append([]string{p}, correctedUpdate[index:]...)...)
					break
				}
			}
		}
	}
	return correctedUpdate
}

// Main function to solve the problem
func solve(inputFile string) string {
	// Read the file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	// Extract ordering rules
	orderingRuleRegex := regexp.MustCompile(`(\d+)\|(\d+)`)
	matches := orderingRuleRegex.FindAllStringSubmatch(content, -1)
	orderingRules := make(map[string][]string)
	for _, match := range matches {
		orderingRules[match[1]] = append(orderingRules[match[1]], match[2])
	}

	// Extract pages to check
	pageCheckRegex := regexp.MustCompile(`\d+(,\d+)+`)
	pageMatches := pageCheckRegex.FindAllString(content, -1)
	var pagesToCheck [][]string
	for _, match := range pageMatches {
		pagesToCheck = append(pagesToCheck, strings.Split(match, ","))
	}

	// Process all invalid updates sequentially and calculate the sum
	sum := 0
	for _, update := range pagesToCheck {
		if !isValidUpdate(update, orderingRules) {
			correctedUpdate := correctUpdate(update, orderingRules)
			midIndex := len(correctedUpdate) / 2
			value, _ := strconv.Atoi(correctedUpdate[midIndex])
			sum += value
		}
	}

	return strconv.Itoa(sum)
}

// Utility function to check if a slice contains a value
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Utility function to get the index of a value in a slice
func indexOf(slice []string, value string) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func main() {
	// Call the solve function with the input file
	result := solve("input.txt")
	fmt.Printf("%s\n", result)
}
