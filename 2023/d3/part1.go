package main

import (
	"bufio"
	"strconv"
	"unicode"
)

func part1(scanner *bufio.Scanner) int {
	res := 0
	var prev, cur, next string

	for scanner.Scan() {
		line := scanner.Text()

		prev, cur, next = cur, next, line
		res += handleLines(prev, cur, next)
	}

	// calculate for the very last line
	prev, cur, next = cur, next, ""
	res += handleLines(prev, cur, next)

	return res
}

// After reading 3 lines, calculate sum of the middle/current line
func handleLines(prev, cur, next string) int {
	sum := 0

	for i := 0; i < len(cur); {
		num, le := yieldInt(cur, i)
		if le == 0 {
			i++
		} else {
			// check if this num is valid - adjacent to a symbol
			isValid := i > 0 && isSymbol(cur[i-1])
			isValid = isValid || i+le < len(cur) && isSymbol(cur[i+le])
			isValid = isValid || symbolInLine(prev, i-1, i+le) || symbolInLine(next, i-1, i+le)
			if isValid {
				sum += num
			}

			i += le
		}
	}

	return sum
}

func symbolInLine(input string, start, end int) bool {
	for i := max(start, 0); i <= min(end, len(input)-1); i++ {
		if isSymbol(input[i]) {
			return true
		}
	}
	return false
}

func isSymbol(c byte) bool {
	if !unicode.IsDigit(rune(c)) && !unicode.IsSpace(rune(c)) && c != '.' {
		return true
	}
	return false
}

// try to yield an int start from start index, returns the int and the end index
func yieldInt(input string, start int) (int, int) {
	num := 0
	i := start

	for ; i < len(input); i++ {
		digit, err := strconv.Atoi(string(input[i]))
		if err != nil {
			break
		}
		num = num*10 + digit
	}

	return num, i - start
}
