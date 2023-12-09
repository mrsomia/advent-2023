package main

import (
	"fmt"
	"log"
	"math"
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

var seedMaps map[string]seedMap

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
	seedMaps = maps
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

type seedRange struct {
	minSeed  int
	len      int
	mapToCat string
	mapToNum int
}

func part2(s string) {
	seeds, seedMaps := parseInput(s)

	seedRanges := make([]seedRange, 0)
	for i := 0; i < len(seeds)-1; i += 2 {
		seed := seeds[i]
		rangeLen := seeds[i+1]
		seedrange := seedRange{
			minSeed:  seed,
			len:      rangeLen,
			mapToCat: "seed",
			mapToNum: seed,
		}
		seedRanges = append(seedRanges, seedrange)
	}

	categoryOrder := []string{
		"soil",
		"fertilizer",
		"water",
		"light",
		"temperature",
		"humidity",
		"location",
	}

	currentMin := math.MaxInt

	for _, category := range categoryOrder {
		catMap := seedMaps[category]
		seedRanges = processSeedRanges(seedRanges, catMap)
	}

	for _, sr := range seedRanges {
		if sr.mapToNum < currentMin {
			currentMin = sr.mapToNum
			fmt.Printf("New Current Min: %v\n", currentMin)
		}
	}
}

func processSeedRanges(seedRanges []seedRange, catMap seedMap) (mapped []seedRange) {
	for _, sr := range seedRanges {
		// loop over catMap numranges
		wasMapped := false
		for _, mapRange := range catMap.numRanges {
			mapRangeLeft := mapRange.srcStart
			mapRangeRight := mapRange.srcStart + mapRange.rangeLen
			srLeft := sr.mapToNum
			srRight := sr.mapToNum + sr.len
			switch {
			case srLeft >= mapRangeLeft && srRight <= mapRangeRight:
				// Whole of the seed range is covered
				sr.mapToCat = catMap.dstElem
				mapped = append(mapped, sr)
				wasMapped = true
			case srLeft < mapRangeLeft && srRight <= mapRangeRight:
				// Right of left that is covered
				newRange := seedRange{
					minSeed:  sr.minSeed + mapRangeLeft - srLeft,
					mapToCat: catMap.dstElem,
					mapToNum: mapRange.dstStart,
					len:      srRight - mapRangeLeft,
				}
				mapped = append(mapped, newRange)
				// start of Left is not covered
				sr.len = mapRangeLeft - srLeft
				seedRanges = append(seedRanges, sr) // Add to back to see if covered by another region
				wasMapped = true

			case srLeft > mapRangeLeft && srRight > mapRangeRight:
				// right side that's not covered
				newRange := seedRange{
					minSeed:  sr.minSeed + mapRangeRight - srLeft,
					mapToCat: catMap.srcElem,
					mapToNum: sr.mapToNum + mapRangeRight - srLeft,
					len:      mapRangeRight - srRight,
				}
				seedRanges = append(seedRanges, newRange)

				sr.mapToCat = catMap.dstElem
				sr.len = mapRangeRight - srLeft
				sr.mapToNum = mapRange.dstStart + srLeft - mapRangeLeft
				mapped = append(mapped, sr)
				wasMapped = true

			case srLeft < mapRangeLeft && srRight > mapRangeRight:
				// map the middle and create 2 new ranges for either side
				leftRange := seedRange{
					minSeed:  sr.minSeed,
					mapToCat: sr.mapToCat,
					mapToNum: sr.mapToNum,
					len:      mapRangeLeft - srLeft,
				}
				seedRanges = append(seedRanges, leftRange)
				rightRange := seedRange{
					minSeed:  sr.minSeed + mapRangeRight - srLeft,
					mapToCat: sr.mapToCat,
					mapToNum: sr.mapToNum + mapRangeRight - srLeft,
					len:      srRight - mapRangeRight,
				}

				seedRanges = append(seedRanges, rightRange)

				sr.minSeed = sr.minSeed + mapRangeLeft - srLeft
				sr.mapToCat = catMap.dstElem
				sr.mapToNum = mapRange.dstStart
				sr.len = mapRange.rangeLen

				mapped = append(mapped, sr)
				wasMapped = true
			default:
			}
		}
		// if range is not covered map to same numbers in new cat
		if !wasMapped {
			newRange := seedRange{
				minSeed:  sr.minSeed,
				len:      sr.len,
				mapToCat: catMap.dstElem,
				mapToNum: sr.mapToNum,
			}
			mapped = append(mapped, newRange)

		}
	}
	return mapped
}

func main() {
	s := openFile("05/exampleinput.txt")
	part1(s)
	part2(s)
}
