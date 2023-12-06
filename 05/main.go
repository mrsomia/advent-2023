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

type seedMap struct {
	name      string
	srcElem   string
	dstElem   string
	numRanges []numRange
}

type numRange struct {
	srcStart int
	dstStart int
	rangeLen int
}

func parseInput(s string) ([]int, map[string]seedMap) {
	s = strings.Trim(s, "\n")
	sections := strings.Split(s, "\n\n")
	seedsStr := strings.Trim(sections[0], "seeds: ")
	seedStrings := strings.Split(seedsStr, " ")

	seeds := []int{}
	maps := map[string]seedMap{}

	// Gets slice of seed numbers
	for _, seedStr := range seedStrings {
		n, err := strconv.Atoi(seedStr)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			seeds = append(seeds, n)
		}
	}

	sections = sections[1:]
	for _, section := range sections {
		lines := strings.Split(section, "\n")
		name := strings.Trim(lines[0], " map:")
		elements := strings.Split(name, "-")
		srcElement := elements[0]
		dstElement := elements[2]
		seedMap := seedMap{name: name, srcElem: srcElement, dstElem: dstElement}

		for _, seedRange := range lines[1:] {
			values := strings.Split(seedRange, " ")
			dstStart, err := strconv.Atoi(values[0])
			if err != nil {
				log.Fatalf("Unable to parse: %v\n", values[0])
				continue
			}
			srcStart, err := strconv.Atoi(values[1])
			if err != nil {
				log.Fatalf("Unable to parse: %v\n", values[1])
				continue
			}
			rangeLen, err := strconv.Atoi(values[2])
			if err != nil {
				log.Fatalf("Unable to parse: %v\n", values[2])
				continue
			}
			r := numRange{srcStart: srcStart, dstStart: dstStart, rangeLen: rangeLen}
			seedMap.numRanges = append(seedMap.numRanges, r)

		}
		maps[dstElement] = seedMap

	}
	return seeds, maps
}

func convertToElem(seed int, seedMap seedMap) int {
	for _, numRange := range seedMap.numRanges {
		l := numRange.rangeLen
		if seed >= numRange.srcStart && seed < (numRange.srcStart+l) {
			diff := seed - numRange.srcStart
			return numRange.dstStart + diff
		}
	}
	return seed
}

func part1(s string) {
	seeds, seedMaps := parseInput(s)

	var min int = 999999999999999999

	for _, seed := range seeds {
		val := convertToElem(seed, seedMaps["soil"])
		val = convertToElem(val, seedMaps["fertilizer"])
		val = convertToElem(val, seedMaps["water"])
		val = convertToElem(val, seedMaps["light"])
		val = convertToElem(val, seedMaps["temperature"])
		val = convertToElem(val, seedMaps["humidity"])
		val = convertToElem(val, seedMaps["location"])
		if val < min {
			min = val
		}
	}
	fmt.Println(min)
}

func main() {
	s := openFile("05/input.txt")
	part1(s)
}
