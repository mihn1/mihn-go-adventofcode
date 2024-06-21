package main

import (
	"bufio"
	"strconv"
	"strings"
)

type seedPair struct {
	start  int
	length int
}

func part2(scanner *bufio.Scanner) int {
	var seedPairs []seedPair
	if scanner.Scan() {
		seedPairs = parseSeedPairs(scanner.Text())
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
	for _, pair := range seedPairs {
		for seed := pair.start; seed < pair.start+pair.length; seed++ {
			var input int = seed
			for _, fn := range pipeline {
				input = fn(input)
			}

			if input < res {
				res = input
			}
		}
	}

	return res
}

func parseSeedPairs(line string) []seedPair {
	pairs := make([]seedPair, 0, 4)
	line = strings.TrimSpace(line[6:])
	numStrs := strings.Split(line, " ")

	for i := range len(numStrs) / 2 {
		seed, err := strconv.Atoi(strings.TrimSpace(numStrs[i*2]))
		if err != nil {
			continue
		}
		length, err := strconv.Atoi(strings.TrimSpace(numStrs[i*2+1]))
		if err != nil {
			continue
		}
		pairs = append(pairs, seedPair{seed, length})

	}
	return pairs
}
