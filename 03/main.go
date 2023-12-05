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

type NeighbourWithCoords struct {
  x int
  y int
  value string
}

func getNeighboursWithCoords(chars [][]string, x, y int) []NeighbourWithCoords {
	dirs := [][]int{{1, 0}, {1, 1}, {0, 1}, {0, -1}, {-1, -1}, {-1, 0}, {1, -1}, {-1, 1}}
	res := []NeighbourWithCoords{}
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
    res = append(res, NeighbourWithCoords{value: chars[newY][newX], x: newX, y: newY })
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

func isStar (s string) bool {
  return "*" == s
}

func part2(s string) {
	chars := parseLines(s)
	// fmt.Println(chars)

	stars := map[string][]int{}
	currentNum := ""
  nextToStarAt := ""

	for y, row := range chars {
		for x, char := range row {
			isNumber := isCharANumber(char)
			if !isNumber && isDot(char) && nextToStarAt != "" {
				//add to validnums
				n, err := strconv.Atoi(currentNum)
				if err != nil {
					fmt.Println(err)
				} else {
          v, ok := stars[nextToStarAt]
          if ok {
            stars[nextToStarAt] = append(v, n)
          } else {
            stars[nextToStarAt] = []int{n}
          }
				}
				currentNum = ""
        nextToStarAt = ""
			}

			if isDot(char) && nextToStarAt != ""{
				currentNum = ""
			}

			if isNumber {
				currentNum = currentNum + char
				// Check neighbours for symbol
				neighbours := getNeighboursWithCoords(chars, x, y)
				for _, neighbour := range neighbours {
					if isStar(neighbour.value) {
						nextToStarAt = fmt.Sprintf("%v,%v", neighbour.x, neighbour.y)
					}
				}

			} else if nextToStarAt != "" {
				// Add to validnums
				n, err := strconv.Atoi(currentNum)
				if err != nil {
					fmt.Println(err)
				} else {
          v, ok := stars[nextToStarAt]
          if ok {
            stars[nextToStarAt] = append(v, n)
          } else {
            stars[nextToStarAt] = []int{n}
          }
				}
				currentNum = ""
        nextToStarAt = ""
			}

			if isDot(char) && nextToStarAt == "" && !isNumber {
				currentNum = ""
			}

      // fmt.Printf("\nchar: %v\nnextToStarAt: %v\ncurrentNum: %v\nstars: %v\n", char, nextToStarAt, currentNum, stars)
		}
		if nextToStarAt != "" {
				n, err := strconv.Atoi(currentNum)
				if err != nil {
					fmt.Println(err)
				} else {
          v, ok := stars[nextToStarAt]
          if ok {
            stars[nextToStarAt] = append(v, n)
          } else {
            stars[nextToStarAt] = []int{n}
          }
				}
		}
    currentNum = ""
    nextToStarAt = ""

	}
	// fmt.Println(stars)

  c := 0

  for _, v := range stars {
    if len(v) == 2 {
      c+= (v[0] * v[1])
    }
  }
  fmt.Println(c)
}


func main() {
	str := openFile("03/input1.txt")
	part1(str)
  part2(str)
}
