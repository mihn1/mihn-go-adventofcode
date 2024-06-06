package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func part2(reader *bufio.Reader) int {
	res := 0
	for i := 1; true; i++ {
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

		_, content := getGameId(line)
		gameRes := calulateGame(content)
		res += gameRes
	}

	return res
}

func calulateGame(content string) int {
	var red, green, blue int
	sets := strings.Split(content, ";")

	for _, set := range sets {
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
				red = max(red, tmp)
			case "green":
				tmp, _ = strconv.Atoi(cntStr)
				green = max(green, tmp)
			case "blue":
				tmp, _ = strconv.Atoi(cntStr)
				blue = max(blue, tmp)
			}
		}
	}

	return red * green * blue
}
