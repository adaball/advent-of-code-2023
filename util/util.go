package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadInput(day int) []string {
	fileSystem := os.DirFS(".")

	file, err := fileSystem.Open(fmt.Sprintf("input/day-%d", day))
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
