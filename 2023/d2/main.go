package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const debug = false

func main() {
	var filename string
	if debug {
		filename = "sample2"
	} else {
		if len(os.Args) < 2 {
			fmt.Println("No input file provided")
			return
		}

		filename = os.Args[1]
	}

	fmt.Println("Running file:", filename)
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	fmt.Println("Part 1:", part1(reader))
	// reader.Reset(file)
	file.Seek(0, io.SeekStart)
	fmt.Println("Part 2:", part2(reader))
}
