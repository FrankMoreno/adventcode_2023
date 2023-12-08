package day7

import (
	"sort"
	"strconv"
	"strings"
)

var cardStrength = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"1": 1,
}

var cardStrengthV2 = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"T": 11,
	"9": 10,
	"8": 9,
	"7": 8,
	"6": 7,
	"5": 6,
	"4": 5,
	"3": 4,
	"2": 3,
	"J": 2,
}

type hand struct {
	original string
	cards    map[string]int
	bid      int
	handType int
}

func Part1(inputs []string) int {
	total := 0
	hands := []hand{}
	for _, input := range inputs {
		hands = append(hands, parseHand(input))
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType != hands[j].handType {
			return hands[i].handType < hands[j].handType
		} else {
			for k := range hands[i].original {
				if cardStrength[string(hands[i].original[k])] != cardStrength[string(hands[j].original[k])] {
					return cardStrength[string(hands[i].original[k])] < cardStrength[string(hands[j].original[k])]
				}
			}
			return hands[i].handType < hands[j].handType
		}
	})

	for idx, hand := range hands {
		total += (idx + 1) * hand.bid
	}

	return total
}

func Part2(inputs []string) int {
	total := 0
	hands := []hand{}
	for _, input := range inputs {
		hands = append(hands, parseHand(input))
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType != hands[j].handType {
			return hands[i].handType < hands[j].handType
		} else {
			for k := range hands[i].original {
				if cardStrengthV2[string(hands[i].original[k])] != cardStrengthV2[string(hands[j].original[k])] {
					return cardStrengthV2[string(hands[i].original[k])] < cardStrengthV2[string(hands[j].original[k])]
				}
			}
			return hands[i].handType < hands[j].handType
		}
	})

	for idx, hand := range hands {
		total += (idx + 1) * hand.bid
	}

	return total
}

func parseHand(input string) hand {
	splitInput := strings.Fields(input)
	original := splitInput[0]
	bid, _ := strconv.Atoi(splitInput[1])
	cards := make(map[string]int)
	for _, char := range original {
		cards[string(char)] += 1
	}

	return hand{
		original: original,
		bid:      bid,
		cards:    cards,
		handType: getHandTypeV2(cards),
	}
}

func getHandType(cards map[string]int) int {
	twoPairs := 0
	threePairs := 0

	for _, value := range cards {
		if value == 5 {
			return 7
		}

		if value == 4 {
			return 6
		}

		if value == 3 {
			threePairs += 1
		}

		if value == 2 {
			twoPairs += 1
		}
	}

	if twoPairs > 0 && threePairs > 0 {
		return 5
	}

	if threePairs >= 1 {
		return 4
	}

	if twoPairs > 1 {
		return 3
	}

	if twoPairs == 1 {
		return 2
	}

	return 1
}

func getHandTypeV2(cards map[string]int) int {
	pairs := 0
	threeOfAKind := 0

	for key, value := range cards {
		if key == "J" {
			continue
		}

		if value == 5 || (value+cards["J"] == 5) {
			return 7
		}

		if value == 4 || (value+cards["J"] == 4) {
			return 6
		}

		if value == 3 {
			threeOfAKind += 1
		}

		if value == 2 {
			pairs += 1
		}
	}

	if cards["J"] == 5 {
		return 7
	}

	if cards["J"] == 4 {
		return 6
	}

	if pairs > 0 && threeOfAKind > 0 || (pairs == 1 && cards["J"] == 3) || (threeOfAKind == 1 && cards["J"] == 2) || (pairs == 2 && cards["J"] == 1) {
		return 5
	}

	if threeOfAKind >= 1 || (pairs >= 1 && cards["J"] >= 1) || cards["J"] >= 2 {
		return 4
	}

	if pairs == 2 || (pairs == 1 && cards["J"] >= 1) || cards["J"] >= 2 {
		return 3
	}

	if pairs == 1 || cards["J"] >= 1 {
		return 2
	}

	return 1
}
