package main

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const cards string = "23456789TJQKA"

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPairs
	ThreeOfaKind
	FullHouse
	FourOfaKind
	FiveOfaKind
)

type bet struct {
	hand     string
	bid      int
	handType HandType
}

func part1(scanner *bufio.Scanner) int {
	var res int = 0
	var bets = make([]bet, 0)

	for scanner.Scan() {
		bet, err := parseLine(scanner.Text())
		if err != nil {
			fmt.Printf("err parsing line: %v\n", err)
		}
		bets = append(bets, bet)
	}

	fmt.Printf("bets: %v\n", len(bets))

	slices.SortFunc(bets, compareHand)

	// fmt.Printf("sorted bets: %v\n", bets)

	for i, bet := range bets {
		res += bet.bid * (i + 1)
	}

	return res
}

func parseLine(line string) (bet, error) {
	bet := bet{}
	split := strings.Split(strings.TrimSpace(line), " ")
	if len(split) != 2 {
		return bet, fmt.Errorf("invalid line")
	}

	bid, err := strconv.Atoi(split[1])
	if err != nil {
		fmt.Println("err parsing bid: ", err)
		return bet, err
	}

	bet.hand = split[0]
	bet.bid = bid
	bet.handType = parseHand(bet.hand)
	return bet, nil
}

func parseHand(hand string) HandType {
	counter := make(map[rune]int)
	for _, c := range hand {
		counter[c] += 1
	}

	var hType HandType
	for _, cnt := range counter {
		switch cnt {
		case 5:
			return FiveOfaKind
		case 4:
			return FourOfaKind
		case 3:
			if hType == OnePair {
				return FullHouse
			}
			hType = ThreeOfaKind
		case 2:
			if hType == OnePair {
				return TwoPairs
			}
			if hType == ThreeOfaKind {
				return FullHouse
			}
			hType = OnePair
		}
	}

	return hType
}

func compareHand(a, b bet) int {
	if a.handType > b.handType {
		return 1
	} else if a.handType < b.handType {
		return -1
	}

	for i := 0; i < len(a.hand); i++ {
		aCard := strings.IndexRune(cards, rune(a.hand[i]))
		bCard := strings.IndexRune(cards, rune(b.hand[i]))
		if aCard > bCard {
			return 1
		} else if aCard < bCard {
			return -1
		}
	}

	return 1
}
