package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	lines := readInput()
	partOne(lines)
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

func partOne(lines []string) {
	lineMatches := make([][]string, 0)
	re := regexp.MustCompile("\\d")
	for _, line := range lines {
		found := re.FindAllString(line, -1)
		lineMatches = append(lineMatches, found)
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

		log.Printf("%d + %d = %d", total, comboInt, total+comboInt)
		total = total + comboInt
	}

	log.Printf("part 1 %d", total)
}
