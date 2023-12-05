package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var exampleInput1 = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
  Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
  Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
  Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
  Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func openFile(fn string) string {
	content, err := os.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

type Round struct {
	blue  int
	green int
	red   int
}

const (
	totalRed   = 12
	totalGreen = 13
	totalBlue  = 14
)

type Game struct {
	id     int
	rounds []Round
}

func parseGameLine(line string) Game {
	splitLine := strings.Split(line, ":")
	game := Game{}

	idRe := regexp.MustCompile(`\d+`)
	gameIdString := idRe.FindString(splitLine[0])
	if gameIdString == "" {
		log.Fatal(fmt.Sprintf("Unable to parse game id from: %v\n", line))
	}
	gameId, err := strconv.Atoi(gameIdString)
	if err != nil {
		log.Fatal(err)
	}
	game.id = gameId

	rounds := strings.Split(splitLine[1], ";")
	for _, round := range rounds {
		// Parse round
		round = strings.Trim(round, " ")
		r := Round{}
		cubes := strings.Split(round, ", ")
		for _, cubeStr := range cubes {
			// Parse each cube pair to get correct item
			cubeTuple := strings.Split(cubeStr, " ")
			value, err := strconv.Atoi(cubeTuple[0])
			if err != nil {
				log.Fatal(err)
			}

			switch cubeTuple[1] {
			case "blue":
				r.blue = value
			case "red":
				r.red = value
			case "green":
				r.green = value
			default:
				fmt.Printf("Unexpected Color: %v, for id: %v", cubeTuple[1], gameId)
			}
		}

		game.rounds = append(game.rounds, r)
	}
	return game
}

func parseInput(str string) []Game {
	// split by line
  str = strings.Trim(str, "\n")
	gameLines := strings.Split(str, "\n")
	games := []Game{}
	for _, gameLine := range gameLines {
		// Get game ID
		game := parseGameLine(gameLine)
		// fmt.Println(game)
		games = append(games, game)
	}
	return games
}

func validateGamePossible(g *Game) bool {
	// check total of rounds is < that total cubes (constants)
	for _, round := range g.rounds {
    if round.blue > totalBlue {
      return false
    }
    if round.red > totalRed {
      return false
    }
    if round.green > totalGreen {
      return false
    }
	}

	return true
}

func solve1(s string) {
  validGames := []int{}
  games := parseInput(s)
  for _, game := range games {
    if validateGamePossible(&game) {
      validGames = append(validGames, game.id)
    }
  }
	// sum IDs

  r := 0
  for _, n := range validGames {
    r += n
  }
  fmt.Println(r)
}

func minimunNumOfCubes (g *Game) []int {
  maxRed := 0
  maxBlue := 0
  maxGreen := 0
  for _, round := range g.rounds {
    if round.red > maxRed {
      maxRed = round.red
    }
    if round.blue > maxBlue {
      maxBlue = round.blue
    }
    if round.green > maxGreen {
      maxGreen = round.green
    }
  }
  return []int{maxRed, maxGreen, maxBlue}
}

func solve2(s string) {
  res := []int{}
  games := parseInput(s)
  for _, game := range games {
    // get the valid number of cubes
    minCubes := minimunNumOfCubes(&game)
    // fmt.Println(minCubes)
    // multiply cubes together
    p := 1
    for _, minCube := range minCubes {
      p *= minCube
    }
    res = append(res, p)
    // append to results
  }
  r := 0
  for _, n := range res {
    r += n
  }

  fmt.Println(r)
}

func main() {
	str := exampleInput1
	str = openFile("02/input1.txt")
	solve1(str)
  solve2(str)
}
