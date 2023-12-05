package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
  "slices"
)

func openFile(s string) string {
	constents, err := os.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	return string(constents)
}

type Card struct {
  number int
	left  []int
	right []int
}

func parseNumberStringArr(strs []string) []int {
	res := []int{}
	for _, n := range strs {
		if n == "" {
			continue
		}
		num := strings.Trim(n, " ")
		val, err := strconv.Atoi(num)
		if err != nil {
			fmt.Printf("len of n: %v\n", len(num))
			fmt.Printf("Error parsing: %v\n", num)
		}
		res = append(res, val)
	}
	return res
}

func parseInput(s string) (res []Card) {
	s = strings.Trim(s, "\n")
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		splitLine := strings.Split(line, ": ")
		cardNum := strings.Trim(splitLine[0], "Card ")
		nums := strings.Split(splitLine[1], " | ")
		leftNumsStr := strings.Split(strings.Trim(nums[0], " "), " ")
		rightNumsStr := strings.Split(nums[1], " ")

		leftNums := parseNumberStringArr(leftNumsStr)
		rightNums := parseNumberStringArr(rightNumsStr)
    n, err := strconv.Atoi(cardNum)
    if err != nil {
      fmt.Printf("Error parsing: %v\n", cardNum)
    }
    res = append(res, Card{number: n, left: leftNums, right: rightNums})
	}
	return res
}

func part1(s string) {
  cards := parseInput(s)
  res := []int{}
  for _, card := range cards {
    match := false
    v := 0
    for _, num := range card.left {
      if slices.Contains(card.right, num) {
        if match {
          v = 2 * v
        } else {
          v=1
          match = true
        }
      }
    }

    res = append(res, v)
  }

  r := 0
  for _, n := range res {
    r += n
  }
  fmt.Printf("Part 1: %v\n", r)
}

func part2(s string) {
  initialCards := parseInput(s)
  finalCards := map[int]int{}
  // loop through cards
    // Get score for each card
    // add new cards to a map? add to a total
    // 
}

func main() {
	s := openFile("input1.txt")
  part1(s)
  s = openFile("exampleinput1.txt")
  part2(s)
}
