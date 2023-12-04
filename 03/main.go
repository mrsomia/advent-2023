package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
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

func parseLines(s string) [][]string {
	s = strings.Trim(s, "\n")
	lines := strings.Split(s, "\n")
	chars := [][]string{}
	for _, line := range lines {
		row := strings.Split(line, "")
		chars = append(chars, row)
	}
	return chars
}

func isCharANumber(s string) bool {
	re := regexp.MustCompile(`\d`)
	return re.MatchString(s)
}

func getNeighbours(chars [][]string, x, y int) []string {
	dirs := [][]int{{1, 0}, {1, 1}, {0, 1}, {0, -1}, {-1, -1}, {-1, 0}, {1, -1}, {-1, 1}}
	res := []string{}
	for _, dir := range dirs {
		dx := dir[0]
		dy := dir[1]
		newY := y + dy
		newX := x + dx
		if newX >= len(chars[y]) || newX < 0 {
			continue
		}
		if newY >= len(chars) || newY < 0 {
			continue
		}
		res = append(res, chars[newY][newX])
	}
	return res
}

func isDot(s string) bool {
	return s == "."
}

func part1(s string) {
	chars := parseLines(s)
	// fmt.Println(chars)

	validNums := []int{}
	currentNum := ""
	isValid := false

	for y, row := range chars {
		for x, char := range row {
			isNumber := isCharANumber(char)
			if !isNumber && isDot(char) && isValid {
				//add to validnums
				n, err := strconv.Atoi(currentNum)
				if err != nil {
					fmt.Println(err)
				} else {
					validNums = append(validNums, n)
				}
				currentNum = ""
				isValid = false
			}

			if isNumber {
				currentNum = currentNum + char
				// Check neighbours for symbol
				neighbours := getNeighbours(chars, x, y)
				for _, neighbour := range neighbours {
					if !isCharANumber(neighbour) && !isDot(neighbour) {
						isValid = true
					}
				}

			} else if isValid {
				// Add to validnums
				n, err := strconv.Atoi(currentNum)
				if err != nil {
					fmt.Println(err)
				} else {
					validNums = append(validNums, n)
				}
				currentNum = ""
				isValid = false
			}

			if isDot(char) && !isValid {
				currentNum = ""
			}

		}
		if isValid {
			n, err := strconv.Atoi(currentNum)
			if err != nil {
				fmt.Println(err)
			} else {
				validNums = append(validNums, n)
			}
			currentNum = ""
			isValid = false
		}
	}
	fmt.Println(validNums)

	r := 0
	for _, n := range validNums {
		r += n
	}

	fmt.Println(r)
}

func main() {
	str := openFile("input1.txt")
	part1(str)
}
