package main

import (
	"bufio"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) int {
	var res int = 1
	var times, distances []int
	if scanner.Scan() {
		times = parseLine(scanner.Text())
	}
	if scanner.Scan() {
		distances = parseLine(scanner.Text())
	}

	if len(times) != len(distances) {
		panic("times and distances are not the same length")
	}

	for i := 0; i < len(times); i++ {
		res *= computeWays(times[i], distances[i])
	}

	return res
}

func parseLine(line string) []int {
	nums := make([]int, 0)
	numsStr := strings.TrimSpace(strings.Split(line, ":")[1])

	for _, numStr := range strings.Split(numsStr, " ") {
		num, err := strconv.Atoi(strings.TrimSpace(numStr))
		if err != nil {
			continue
		}
		nums = append(nums, num)
	}
	return nums
}

func computeWays(time, distance int) int {
	var ways int

	for i := range time {
		remainingTime := time - i
		traveledDist := i * remainingTime
		if traveledDist > distance {
			ways += 1
		}
	}

	return ways
}