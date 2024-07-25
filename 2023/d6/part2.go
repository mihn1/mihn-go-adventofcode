package main

import (
	"bufio"
	"strconv"
	"strings"
)

func part2(scanner *bufio.Scanner) int {
	var res int = 1
	var time, distance int
	if scanner.Scan() {
		time = parseLine2(scanner.Text())
	}
	if scanner.Scan() {
		distance = parseLine2(scanner.Text())
	}

	res *= computeWays(time, distance)
	return res
}

func parseLine2(line string) int {
	numStr := strings.Split(line, ":")[1]
	numStr = strings.Replace(numStr, " ", "", -1)
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return num
}
