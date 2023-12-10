package main

import (
	"fmt"
	"log"
	"os"
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

func parse(s string) ([]int, []int) {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	timesLine := strings.TrimSpace(lines[0])
	distancesLine := strings.TrimSpace(lines[1])
	timesStr := strings.Fields(timesLine)
	distancesStr := strings.Fields(distancesLine)
	times := make([]int, 0)
	distances := make([]int, 0)

	for _, time := range timesStr[1:] {
		t, _ := strconv.Atoi(time)
		times = append(times, t)
	}
	for _, distance := range distancesStr[1:] {
		d, _ := strconv.Atoi(distance)
		distances = append(distances, d)
	}
	return times, distances
}

func parse2(s string) (int, int) {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	timesLine := strings.TrimSpace(lines[0])
	distancesLine := strings.TrimSpace(lines[1])
	timesFields := strings.Fields(timesLine)
	distancesFields := strings.Fields(distancesLine)
  timeStr := strings.Join(timesFields[1:], "")
  distanceStr := strings.Join(distancesFields[1:], "")

  t, _ := strconv.Atoi(timeStr)
  d, _ := strconv.Atoi(distanceStr)
	return t, d
}

func getAllNaiveSolutions(d int, t int) []int {
	ms := make([]int, 0)
	for i := 1; i < t; i++ {
		if d < i*(t-i) {
			ms = append(ms, i)
		}
	}
  return ms
}

func part1(s string) {
	times, distances := parse(s)
  n := make([]int, 0)
  for i, t := range times {
    d := distances[i]
    solves := getAllNaiveSolutions(d, t)
    n = append(n, len(solves))
  }
  r := 1
  for _, num := range n {
    r *= num
  }

  fmt.Println(r)
}

func part2(s string) {
  t, d := parse2(s)
  min := 1

	for i := 1; i < t; i++ {
		if d < i*(t-i) {
			min = i
      break
		}
	}

  max := t-1

	for i := t; i > 0; i-- {
		if d < i*(t-i) {
			max = i + 1
      break
		}
	}
  fmt.Println(max - min)
}

func main() {
	s := openFile("06/input.txt")
	part1(s)
  part2(s)
}
