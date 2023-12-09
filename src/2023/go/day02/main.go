package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	var part1Sum int
	var part2Sum int
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
		fewestCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		subsets := strings.Split(strings.Split(line, ":")[1], ";")
		for _, subset := range subsets {
			for _, cube := range strings.Split(subset, ",") {
				cube := cube[1:]
				val := strings.Split(cube, " ")
				count, _ := strconv.Atoi(val[0])
				color := val[1]
				if count > maxCubes[color] {
					possible = false
				}
				if count > fewestCubes[color] {
					fewestCubes[color] = count
				}
			}
		}

		if possible {
			part1Sum += gameNumber
		}

		part2Sum += fewestCubes["red"] * fewestCubes["green"] * fewestCubes["blue"]
	}

	fmt.Println("Part 1 answer:", part1Sum)
	fmt.Println("Part 2 answer:", part2Sum)
}

func main() {
	part1()
}
