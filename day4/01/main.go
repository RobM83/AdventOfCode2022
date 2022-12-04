package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	input   string
	s1start int
	s1end   int
	s2start int
	s2end   int
}

func main() {
	pairs, _ := readLines("input.txt")
	splitInput(pairs)
	overlap := nrOfOverlap(pairs)
	fmt.Println(overlap)
}

func nrOfOverlap(pairs []*pair) int {
	overlap := 0
	checkOverlap := func(s1start, s1end, s2start, s2end int) bool {
		if s1start >= s2start && s1end <= s2end {
			return true
		}
		return false
	}

	for _, pair := range pairs {
		if checkOverlap(pair.s1start, pair.s1end, pair.s2start, pair.s2end) ||
			checkOverlap(pair.s2start, pair.s2end, pair.s1start, pair.s1end) {
			overlap++
		}
	}
	return overlap
}

func splitInput(pairs []*pair) {
	getStartAndEnd := func(input string) (int, int) {
		start, _ := strconv.Atoi(strings.Split(input, "-")[0])
		end, _ := strconv.Atoi(strings.Split(input, "-")[1])
		return start, end
	}

	for _, pair := range pairs {
		// split input
		input := strings.Split(pair.input, ",")
		pair.s1start, pair.s1end = getStartAndEnd(input[0])
		pair.s2start, pair.s2end = getStartAndEnd(input[1])
	}
}

func readLines(path string) ([]*pair, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var pairs []*pair
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		pair := &pair{
			input: line,
		}
		pairs = append(pairs, pair)
	}
	return pairs, scanner.Err()
}
