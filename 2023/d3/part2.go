package main

import (
	"bufio"
)

var prevContainer, curContainer, nextContainer map[int][]int

func part2(scanner *bufio.Scanner) int {
	res := 0
	prevContainer, curContainer, nextContainer = make(map[int][]int), make(map[int][]int), make(map[int][]int)
	var prev, cur, next string

	for scanner.Scan() {
		line := scanner.Text()

		prev, cur, next = cur, next, line
		handleLines2(prev, cur, next)
		for _, lst := range prevContainer {
			if len(lst) == 2 {
				res += lst[0] * lst[1]
			}
		}
		prevContainer, curContainer, nextContainer = curContainer, nextContainer, make(map[int][]int)
	}

	// calculate for the last 2 lines
	prev, cur, next = cur, next, ""
	handleLines2(prev, cur, next)
	for _, lst := range prevContainer {
		if len(lst) == 2 {
			res += lst[0] * lst[1]
		}
	}
	for _, lst := range curContainer {
		if len(lst) == 2 {
			res += lst[0] * lst[1]
		}
	}

	return res
}

// After reading 3 lines, calculate sum of the middle/current line
func handleLines2(prev, cur, next string) {
	for i := 0; i < len(cur); {
		num, le := yieldInt(cur, i)
		if le == 0 { // no number found
			i++
		} else {
			if i > 0 && cur[i-1] == '*' {
				curContainer[i-1] = append(curContainer[i-1], num)
			}
			if i+le < len(cur) && cur[i+le] == '*' {
				curContainer[i+le] = append(curContainer[i+le], num)
			}
			for _, prevIdx := range asterisksInLine(prev, i-1, i+le) {
				prevContainer[prevIdx] = append(prevContainer[prevIdx], num)
			}
			for _, nextIdx := range asterisksInLine(next, i-1, i+le) {
				nextContainer[nextIdx] = append(nextContainer[nextIdx], num)
			}

			i += le
		}
	}
}

func asterisksInLine(input string, start, end int) []int {
	lst := make([]int, 0)
	for i := max(start, 0); i <= min(end, len(input)-1); i++ {
		if input[i] == '*' {
			lst = append(lst, i)
		}
	}
	return lst
}
