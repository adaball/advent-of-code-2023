package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := readInput()
	partOne(lines)
	partTwo(lines)
}

func readInput() []string {
	fileSystem := os.DirFS(".")

	file, err := fileSystem.Open("input/day-1")
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatalf("error closing input file: %+v", err)
		}
		log.Print("input file closed")
	}()

	if err != nil {
		log.Fatalf("error opening input file: %+v", err)
	}

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// this is a much more involved way of solving it, but I somehow wasn't getting any success using regex alone 🤷‍
//
// nb: integers 48-57 inclusive are decimal 0-9 in utf8
func findFirstAndLastDigit(line string) (string, string) {
	wordToDigit := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	digitWords := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine")

	// move 2 pointers towards the center of line (one starting at 0, one starting at len(line)-1)
	// when a single digit or a spelled out number is found for each, return the found strings
	finalA, finalB := "", ""
	for ptrA, ptrB := 0, len(line)-1; ptrB >= 0 && ptrA < len(line); ptrA, ptrB = ptrA+1, ptrB-1 {

		if finalA == "" && ptrA < len(line) {
			charA := line[ptrA]
			strA := line[0 : ptrA+1]
			aMatches := digitWords.FindAllString(strA, -1)

			if charA >= 48 && charA <= 57 {
				finalA = fmt.Sprintf("%c", charA)
			} else if aMatches != nil {
				finalA = aMatches[0]
			}
		}

		if finalB == "" && ptrB >= 0 {
			charB := line[ptrB]
			strB := line[ptrB:]
			bMatches := digitWords.FindAllString(strB, -1)

			if charB >= 48 && charB <= 57 {
				finalB = fmt.Sprintf("%c", charB)
			} else if bMatches != nil {
				finalB = bMatches[len(bMatches)-1]
			}
		}
	}

	v, ok := wordToDigit[finalA]
	if ok {
		finalA = strings.Clone(v)
	}

	v, ok = wordToDigit[finalB]
	if ok {
		finalB = strings.Clone(v)
	}

	return finalA, finalB
}

func partTwo(lines []string) {
	total := int64(0)
	for _, line := range lines {
		digitA, digitB := findFirstAndLastDigit(line)
		combined := fmt.Sprintf("%s%s", digitA, digitB)
		converted, err := strconv.ParseInt(combined, 10, 64)

		if err != nil {
			log.Fatalf("error converting %s: %v", combined, err)
		}

		total += converted
	}

	log.Printf("Part 2: %d", total)
}

func partOne(lines []string) {
	lineMatches := make([][]string, 0)
	re := regexp.MustCompile("\\d")
	for _, line := range lines {
		found := re.FindAllString(line, -1)
		if found != nil {
			lineMatches = append(lineMatches, found)
		}
	}

	total := int64(0)
	for _, matches := range lineMatches {
		a := matches[0]
		b := matches[len(matches)-1]
		combo := fmt.Sprintf("%s%s", a, b)
		comboInt, err := strconv.ParseInt(combo, 10, 64)

		if err != nil {
			log.Fatalf("error converting %s to int: %v", combo, err)
		}

		total = total + comboInt
	}

	log.Printf("Part 1: %d", total)
}
