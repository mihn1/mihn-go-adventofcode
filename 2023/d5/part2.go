package main

import (
	"bufio"
	"example/adventofcode/shared"
	"strings"
)

func part2(scanner *bufio.Scanner) int {
	res := 0

	counter := make(map[int]int)

	for i := 0; scanner.Scan(); i++ {
		if _, exists := counter[i]; !exists {
			counter[i] = 0
		}
		counter[i] += 1 // original card
		res += counter[i]

		line := scanner.Text()
		winningCards := handleLine2(line)
		for nxt := 1; nxt <= winningCards; nxt++ {
			if _, exists := counter[i+nxt]; !exists {
				counter[i+nxt] = 0
			}
			counter[i+nxt] += counter[i]
		}
	}

	return res
}

func handleLine2(line string) int {
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
	cnt := 0

	for _, num := range myNums {
		if shared.Contains(winningNums, num) {
			cnt++
		}
	}

	return cnt
}
