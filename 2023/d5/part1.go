package main

import (
	"bufio"
	"slices"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) int {
	var seeds []int
	if scanner.Scan() {
		seeds = parseSeeds(scanner.Text())
	}
	scanner.Scan() // empty line
	seedToSoil := parseMap(scanner)
	soilToFertilizer := parseMap(scanner)
	fertilizerToWater := parseMap(scanner)
	waterToLight := parseMap(scanner)
	lightToTemperature := parseMap(scanner)
	temperatureToHumidity := parseMap(scanner)
	humidityToLocation := parseMap(scanner)

	pipeline := make([]func(input int) int, 0, 10)
	pipeline = append(pipeline,
		findOutputFromMapFunc(seedToSoil),
		findOutputFromMapFunc(soilToFertilizer),
		findOutputFromMapFunc(fertilizerToWater),
		findOutputFromMapFunc(waterToLight),
		findOutputFromMapFunc(lightToTemperature),
		findOutputFromMapFunc(temperatureToHumidity),
		findOutputFromMapFunc(humidityToLocation),
	)

	res := int(1e9)
	for _, seed := range seeds {
		var input int = seed
		for _, fn := range pipeline {
			input = fn(input)
		}

		if input < res {
			res = input
		}
	}

	return res
}

func findOutputFromMapFunc(mapping [][]int) func(input int) int {
	return func(input int) int {
		return findOutputFromMap(mapping, input)
	}
}

func findOutputFromMap(mapping [][]int, input int) int {
	l, r := 0, len(mapping)
	for l < r {
		m := (l + r) / 2
		if mapping[m][0] > input {
			r = m
		} else {
			l = m + 1
		}
	}
	upperBound := l
	if upperBound == 0 {
		return input
	}

	if mapping[upperBound-1][0]+mapping[upperBound-1][2]-1 < input {
		return input
	}

	return input - mapping[upperBound-1][0] + mapping[upperBound-1][1]
}

func parseSeeds(line string) []int {
	seeds := make([]int, 0, 4)
	line = line[6:]
	for _, seed := range strings.Split(line, " ") {
		seedNum, err := strconv.Atoi(strings.TrimSpace(seed))
		if err != nil {
			continue
		}
		seeds = append(seeds, seedNum)
	}
	return seeds
}

func parseMap(scanner *bufio.Scanner) [][]int {
	container := make([][]int, 0)
	scanner.Scan() // header
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}

		nums := strings.Split(line, " ")
		if len(nums) != 3 {
			break
		}

		dest, err := strconv.Atoi(nums[0])
		if err != nil {
			break
		}

		src, err := strconv.Atoi(nums[1])
		if err != nil {
			break
		}

		count, err := strconv.Atoi(nums[2])
		if err != nil {
			break
		}
		container = append(container, []int{src, dest, count})
	}

	slices.SortFunc(container, func(a, b []int) int {
		if a[0] < b[0] {
			return -1
		} else if a[0] > b[0] {
			return 1
		}
		return 0
	})

	return container
}
