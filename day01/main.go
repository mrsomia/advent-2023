package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var exampleInput = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

var exampleInput2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func openFile(fn string) string {
	content, err := os.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func solve1(str string) {
  lines := strings.Split(str, "\n")
  // fmt.Printf("%v\n", lines)
  res := []int{}

  re := regexp.MustCompile(`\d`)
  for _, line := range lines {
    matches := re.FindAllString(line, -1)
    if matches == nil {
      continue
    }
    num, err:= strconv.Atoi(matches[0] + matches[len(matches)- 1])
    if err != nil {
      log.Fatal(err)
    }
    res = append(res, num)
  }

  // fmt.Println(res)
  answer := 0
  for _, n := range res {
    answer += n
  }
  fmt.Println(answer)
}

var m = map[string]string{
  "one" : "1",
  "two" : "2",
  "three" : "3",
  "four" : "4",
  "five" : "5",
  "six" : "6",
  "seven" : "7",
  "eight" : "8",
  "nine" : "9",
}

func replace(s string) string {
  current := s
  current = strings.Replace(current, "one", "on1e", -1)
  current = strings.Replace(current, "two", "tw2o", -1)
  current = strings.Replace(current, "three", "thr3ee", -1)
  current = strings.Replace(current, "four", "fo4ur", -1)
  current = strings.Replace(current, "five", "fi5ve", -1)
  current = strings.Replace(current, "six", "si6x", -1)
  current = strings.Replace(current, "seven", "sev7en", -1)
  current = strings.Replace(current, "eight", "eig8ht", -1)
  current = strings.Replace(current, "nine", "ni9ne", -1)
  return current
}

func solve2(str string) {
  lines := strings.Split(str, "\n")
  // fmt.Printf("%v\n", lines)
  res := []int{}

  re := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine|\d`)
  digitRe := regexp.MustCompile(`^\d$`)
  for _, line := range lines {
    matches := re.FindAllString(line, -1)
    if matches == nil {
      continue
    }
    first := matches[0]
    if !digitRe.MatchString(first) {
      first = m[first]
    }
    second := matches[len(matches)-1]
    if !digitRe.MatchString(second) {
      second = m[second]
    }

    num, err:= strconv.Atoi(first + second)
    if err != nil {
      log.Fatal(err)
    }
    res = append(res, num)
  }
  // fmt.Println(res)
  answer := 0
  for _, n := range res {
    answer += n
  }
  fmt.Println(answer)
}

func main() {
  str := openFile("input1.txt")
  solve1(str)
  str = openFile("input2.txt")
  // str = exampleInput2
  solve2(replace(str))
}
