package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func readInput() {
	var sum int
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Close the file when the function returns

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit := getFirstDigit(line)
		lastDigit := getLastDigit(line)

		if firstDigit != -1 && lastDigit != -1 {
			combined := fmt.Sprintf("%d%d", firstDigit, lastDigit)
			num, _ := strconv.Atoi(combined)
			sum += num
		} else if lastDigit == -1 {
			combined := fmt.Sprintf("%d%d", firstDigit, firstDigit)
			num, _ := strconv.Atoi(combined)
			sum += num
		}
	}

	fmt.Println("Sum:", sum)

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func getFirstDigit(s string) int {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return int(char - '0')
		}
	}
	return -1 // No digit found
}

func getLastDigit(s string) int {
	var lastDigit rune
	for _, char := range s {
		if unicode.IsDigit(char) {
			lastDigit = char
		}
	}
	if lastDigit != 0 {
		return int(lastDigit - '0')
	}
	return -1 // No digit found
}

func main() {
	readInput()
}
