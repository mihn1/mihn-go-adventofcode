package main

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const cards2 string = "J23456789TQKA"

func part2(scanner *bufio.Scanner) int {
	var res int = 0
	var bets = make([]bet, 0)

	for scanner.Scan() {
		bet, err := parseLine2(scanner.Text())
		if err != nil {
			fmt.Printf("err parsing line: %v\n", err)
		}
		bets = append(bets, bet)
	}

	fmt.Printf("bets: %v\n", len(bets))

	slices.SortFunc(bets, compareHand2)

	for i, bet := range bets {
		res += bet.bid * (i + 1)
	}

	return res
}

func parseLine2(line string) (bet, error) {
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
	bet.handType = parseHand2(bet.hand)
	return bet, nil
}

func parseHand2(hand string) HandType {
	jokerCnt := 0
	counter := make(map[rune]int)

	for _, c := range hand {
		if c == 'J' {
			jokerCnt += 1
		} else {
			counter[c] += 1
		}
	}

	// Add jokers to the max count of the hand
	if jokerCnt > 0 {
		var maxCard rune = 0
		for c, cnt := range counter {
			if maxCard == 0 || cnt > counter[maxCard] {
				maxCard = c
			}
		}
		counter[maxCard] += jokerCnt
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

func compareHand2(a, b bet) int {
	if a.handType > b.handType {
		return 1
	} else if a.handType < b.handType {
		return -1
	}

	for i := 0; i < len(a.hand); i++ {
		aCard := strings.IndexRune(cards2, rune(a.hand[i]))
		bCard := strings.IndexRune(cards2, rune(b.hand[i]))
		if aCard > bCard {
			return 1
		} else if aCard < bCard {
			return -1
		}
	}

	return 1
}
