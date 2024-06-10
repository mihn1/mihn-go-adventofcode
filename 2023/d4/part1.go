package main

import (
	"bufio"
	"example/adventofcode/shared"
	"strings"
)

func part1(scanner *bufio.Scanner) int {
	res := 0

	for scanner.Scan() {
		line := scanner.Text()
		points := handleLine(line)
		res += points
	}

	return res
}

func handleLine(line string) int {
	points := 0
	splited := strings.Split(line, ":")
	if len(splited) != 2 {
		return points
	}

	cardContent := splited[1]
	splited = strings.Split(cardContent, "|")
	if len(splited) != 2 {
		return points
	}

	winningNums := strings.Fields(strings.TrimSpace(splited[0]))
	myNums := strings.Fields(strings.TrimSpace(splited[1]))
	factor := -1

	for _, num := range myNums {
		if shared.Contains(winningNums, num) {
			factor++
		}
	}

	if factor == -1 {
		return 0
	}
	return 1 << factor
}
