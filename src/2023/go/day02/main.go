package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	var sum int
	gameNumber := 0
	maxCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		gameNumber += 1
		line := scanner.Text()
		possible := true

		subsets := strings.Split(strings.Split(line, ":")[1], ";")
		for _, subset := range subsets {
			for _, cube := range strings.Split(subset, ",") {
				cube := cube[1:]
				val := strings.Split(cube, " ")
				count, _ := strconv.Atoi(val[0])
				color := val[1]
				if count > maxCubes[color] {
					possible = false
					break
				}
			}
		}

		if possible {
			sum += gameNumber
		}
	}

	fmt.Println("Part 1 answer:", sum)
}

func main() {
	part1()
}
