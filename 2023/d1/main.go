package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var wordToDigit = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No input file provided")
		return
	}

	filename := os.Args[1]
	fmt.Println("Running file:", filename)

	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var sumAtoi, sumDirect int
	for i := 1; true; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error reading file", err)
			}
			break
		}

		sumAtoi += handleLineAtoi(line)
		sumDirect += handleLineDirect(line)
	}

	fmt.Printf("Sum Atoi: %d\nSum Direct: %d\n", sumAtoi, sumDirect)
}

func handleLineAtoi(line string) int {
	var first, second int
	l := 0

	for l = range len(line) {
		parsed, err := strconv.Atoi(string(line[l]))
		if err != nil {
			digit := getDigitFromWord(line, l, false)
			if digit == -1 {
				continue
			} else {
				first = digit
				break
			}
		} else {
			first = parsed
			break
		}
	}

	for r := len(line) - 1; r >= l; r-- {
		parsed, err := strconv.Atoi(string(line[r]))
		if err != nil {
			// Try to parse word to digit
			digit := getDigitFromWord(line, r, true)
			if digit == -1 {
				continue
			} else {
				second = digit
				break
			}
		} else {
			second = parsed
			break
		}
	}

	return first*10 + second
}

func handleLineDirect(line string) int {
	var first, second int

	l := 0
	for l = 0; l < len(line); l++ {
		if line[l] >= '0' && line[l] <= '9' {
			first = int(line[l] - '0')
			break
		} else {
			digit := getDigitFromWord(line, l, false)
			if digit == -1 {
				continue
			} else {
				first = digit
				break
			}
		}
	}

	for r := len(line) - 1; r >= l; r-- {
		if line[r] >= '0' && line[r] <= '9' {
			second = int(line[r] - '0')
			break
		} else {
			digit := getDigitFromWord(line, r, true)
			if digit == -1 {
				continue
			} else {
				second = digit
				break
			}
		}
	}

	return first*10 + second
}

func getDigitFromWord(line string, i int, isEnd bool) int {
	for word, digit := range wordToDigit {
		var start, end int
		if isEnd {
			start = i
			end = i + len(word)
		} else {
			start = i - len(word) + 1
			end = i + 1
		}
		if start < 0 || end > len(line) {
			continue
		}
		if line[start:end] == word {
			return digit
		}
	}
	return -1
}
