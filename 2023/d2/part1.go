package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	totalRed   int = 12
	totalGreen int = 13
	totalBlue  int = 14
)

func part1(reader *bufio.Reader) int {
	res := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading file", err)
				break
			}
			if len(line) == 0 {
				break
			}
		}

		gameId, content := getGameId(line)
		if checkGame(content) {
			res += gameId
		}
	}

	return res
}

func checkGame(content string) bool {
	sets := strings.Split(content, ";")
	for _, set := range sets {
		var red, green, blue int
		balls := strings.Split(set, ",")
		for _, ball := range balls {
			splited := strings.Split(strings.TrimSpace(ball), " ")
			if len(splited) != 2 {
				continue
			}
			cntStr, color := splited[0], splited[1]
			var tmp int
			switch color {
			case "red":
				tmp, _ = strconv.Atoi(cntStr)
				red += tmp
			case "green":
				tmp, _ = strconv.Atoi(cntStr)
				green += tmp
			case "blue":
				tmp, _ = strconv.Atoi(cntStr)
				blue += tmp
			}
		}

		if red > totalRed || green > totalGreen || blue > totalBlue {
			return false
		}
	}

	return true
}

func getGameId(line string) (int, string) {
	splited := strings.Split(line, ":")
	if len(splited) != 2 {
		return 0, ""
	}

	game, content := splited[0], splited[1]

	gameId, err := strconv.Atoi(strings.Split(game, " ")[1])
	if err != nil {
		return 0, ""
	}

	return gameId, content
}
