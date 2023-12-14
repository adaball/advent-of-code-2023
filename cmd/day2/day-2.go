package main

import (
	"github.com/adamball/advent-of-code-2023/util"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Game 1: 4 green, 7 blue; 2 blue, 4 red; 5 blue, 2 green, 2 red; 1 green, 3 red, 9 blue; 3 green, 9 blue; 7 green, 2 blue, 2 red
// Determine which games would have been possible if the bag had been loaded with only `12` red cubes, `13` green cubes, and `14` blue cubes. What is the sum of the IDs of those games?

func gamePossible(diceHandfuls []map[string]int64) bool {
	for _, diceHandful := range diceHandfuls {
		if diceHandful["red"] > 12 || diceHandful["green"] > 13 || diceHandful["blue"] > 14 {
			return false
		}
	}

	return true
}

func main() {
	lines := util.ReadInput(2)
	partOne(lines)
}

func partOne(lines []string) {

	impossibleGames := make([]int64, 0)
	for _, line := range lines {
		diceRegex := regexp.MustCompile("(\\d+) (red|green|blue)")
		gameIdRegex := regexp.MustCompile("^Game (\\d+)")

		lineTokens := strings.Split(line, ":")

		gameIdSubMatches := gameIdRegex.FindAllSubmatch([]byte(lineTokens[0]), -1)
		gameId, err := strconv.ParseInt(string(gameIdSubMatches[0][1]), 10, 64)

		if err != nil {
			log.Fatalf("Unable to parse int from %s: %v", gameIdSubMatches[0][1], err)
		}

		diceSetTokens := strings.Split(lineTokens[1], ";")
		diceHandfuls := make([]map[string]int64, 0)
		for _, diceSetToken := range diceSetTokens {
			diceSubMatches := diceRegex.FindAllStringSubmatch(diceSetToken, -1)
			handful := make(map[string]int64)

			for _, diceSubMatch := range diceSubMatches {
				qty, err := strconv.ParseInt(diceSubMatch[1], 10, 64)

				if err != nil {
					log.Fatalf("unable to parse int from %s: %v", diceSubMatch[1], err)
				}

				handful[diceSubMatch[2]] = qty
			}

			diceHandfuls = append(diceHandfuls, handful)
		}

		if !gamePossible(diceHandfuls) {
			log.Printf("\ngame id: %d\ndiceHandfuls: %v", gameId, diceHandfuls)
			impossibleGames = append(impossibleGames, gameId)
		}
	}

	total := int64(0)
	for _, gameId := range impossibleGames {
		total += gameId
	}

	log.Print(impossibleGames)

	log.Printf("Part 1: %d", total)

}
