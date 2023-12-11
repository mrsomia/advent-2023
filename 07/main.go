package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func openFile(s string) string {
	constents, err := os.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	return string(constents)
}

type Hand struct {
	cards string
	bid   int
}

func parse(s string) []Hand {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	hands := make([]Hand, 0)
	for _, line := range lines {
		items := strings.Fields(line)
		bid, _ := strconv.Atoi(items[1])
		hand := Hand{cards: items[0], bid: bid}
		hands = append(hands, hand)
	}
	return hands
}

func getHandScore(s string) int {
	m := make(map[rune]int)
	for _, char := range s {
		m[char]++
	}

	for _, v := range m {
		switch {
		case v == 5:
			// five of a kind
			return 1
		case v == 4:
			// four of a kind
			return 2
		case v == 3 && len(m) == 2:
			// full house
			return 3
		case v == 3 && len(m) == 3:
			// 3 of a kind
			return 4
		case v == 2 && len(m) == 3:
			// Two pair
			return 5
		case v == 2 && len(m) == 4:
			return 6
		case v == 1 && len(m) == 5:
			return 7
		}
	}
	return 8
}

var charMap = map[byte]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

func getHandScoreTie(left string, right string) bool {
	for i := 0; i < 5; i++ {
		charL := charMap[left[i]]
		charR := charMap[right[i]]
		if charL == charR {
			continue
		} else {
			return charL < charR
		}
	}
	return false
}


var charMap2 = map[byte]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

func getHandScoreTie2(left string, right string) bool {
	for i := 0; i < 5; i++ {
		charL := charMap2[left[i]]
		charR := charMap2[right[i]]
		if charL == charR {
			continue
		} else {
			return charL < charR
		}
	}
	return false
}

func part1(s string) {
	hands := parse(s)
	sort.Slice(hands, func(i, j int) bool {
		scoreI := getHandScore(hands[i].cards)
		scoreJ := getHandScore(hands[j].cards)
		if scoreI != scoreJ {
			return scoreI > scoreJ
		}
		return getHandScoreTie(hands[i].cards, hands[j].cards)
	})

	total := 0
	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}

	fmt.Println(total)
}

type cardCount struct {
	card  rune
	count int
}

func getHandScore2(s1 string) int {
	m := make(map[rune]int)
	jokers := 0
	for _, char := range s1 {
		m[char]++
		if char == 'J' {
			jokers++
		}
	}
  if jokers == 5 {
    return 1
  }

  max := cardCount{
    card: 'X',
    count: 0,
  }

  if jokers > 0 {
    for k,v := range m {
      if v > max.count && k != 'J' {
        max.count = v
        max.card = k
      }
    }

    m[max.card] += jokers
    delete(m, 'J')
  }

	for _, v := range m {
		switch {
		case v == 5:
			// five of a kind
			return 1
		case v == 4:
			// four of a kind
			return 2
		case v == 3 && len(m) == 2:
			// full house
			return 3
		case v == 3 && len(m) == 3:
			// 3 of a kind
			return 4
		case v == 2 && len(m) == 3:
			// Two pair
			return 5
		case v == 2 && len(m) == 4:
			return 6
		case v == 1 && len(m) == 5:
			return 7
		}
	}
	return 8
}

func part2(s string) {
	hands := parse(s)
	sort.Slice(hands, func(i, j int) bool {
		scoreI := getHandScore2(hands[i].cards)
		scoreJ := getHandScore2(hands[j].cards)
		if scoreI != scoreJ {
			return scoreI > scoreJ
		}
		return getHandScoreTie2(hands[i].cards, hands[j].cards)
	})

	total := 0
	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}

	fmt.Println(total)
}

func main() {
	s := openFile("07/input.txt")
	part1(s)
  part2(s)
}
