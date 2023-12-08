package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func part1() {
	var sum int

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Close the file when the function returns

	scanner := bufio.NewScanner(file)

	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		firstDigit := getFirstDigit(line)
		lastDigit := getLastDigit(line)
		num, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstDigit, lastDigit))
		sum += num
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Part 1 answer:", sum)
}

func part2() {
	letterDigits := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	var sum int

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Close the file when the function returns

	scanner := bufio.NewScanner(file)

	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		var digits []int

		for i, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, int(char-'0'))
			}
			for j, val := range letterDigits {
				if strings.HasPrefix(line[i:], val) {
					digits = append(digits, j+1)
				}
			}
		}

		num, _ := strconv.Atoi(fmt.Sprintf("%d%d", digits[0], digits[len(digits)-1]))
		sum += num
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Part 2 answer:", sum)
}

func getFirstDigit(s string) int {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return int(char - '0')
		}
	}
	return 0 // No digit found
}

func getLastDigit(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			return int(s[i] - '0')
		}
	}
	return 0 // No digit found
}

func main() {
	part1()
	part2()
}
