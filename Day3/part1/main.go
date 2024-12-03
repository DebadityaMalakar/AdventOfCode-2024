package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func main() {
	// Read the entire file content
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Compile the regular expression to match "mul(a,b)"
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// Find all matches in the content
	matches := re.FindAllStringSubmatch(string(content), -1)

	// Calculate the sum of the products
	sum := 0
	for _, match := range matches {
		a, _ := strconv.Atoi(match[1]) // Convert the first number to int
		b, _ := strconv.Atoi(match[2]) // Convert the second number to int
		sum += a * b                   // Add the product to the sum
	}

	// Print the result
	fmt.Println(sum)
}
