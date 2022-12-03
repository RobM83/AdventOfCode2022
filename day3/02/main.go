package main

import (
	"bufio"
	"fmt"
	"os"
)

type elfGroup struct {
	rugsacks       []*rugsack
	common         string
	commonPriority int
}

type rugsack struct {
	content string
	items   map[string]int
}

func main() {
	elfGroups, _ := readLines("input.txt")

	sumOfPriorities := 0
	for _, eg := range elfGroups {
		checkRugsacks(eg.rugsacks)
		getBadge(eg)
		// fmt.Println(eg.common)
		// fmt.Println(eg.commonPriority)
		sumOfPriorities += eg.commonPriority
	}

	fmt.Printf("Sum of priorities: %d\n", sumOfPriorities)
}

func checkRugsacks(rugsacks []*rugsack) {
	for _, r := range rugsacks {
		r.items = make(map[string]int)

		for i := 0; i < len(r.content); i++ {
			r.items[r.content[i:i+1]] = getPriority(r.content[i : i+1])
		}
	}
}

func getBadge(eg *elfGroup) {
	rs1 := eg.rugsacks[0]
	rs2 := eg.rugsacks[1]
	rs3 := eg.rugsacks[2]

	for k := range rs1.items {
		if val, ok := rs2.items[k]; ok {
			if _, ok2 := rs3.items[k]; ok2 {
				eg.common = k
				eg.commonPriority = val
				break
			}
		}
	}
}

func getPriority(char string) int {
	r := []rune(char)
	p := int(r[0])
	if p > 96 {
		return p - 96 // a =  97 - 1 =
	} else {
		return p - 38 // A = 65 - 27 =
	}
}

func readLines(path string) ([]*elfGroup, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var elfGroups []*elfGroup
	groupSize := 3
	lineNumber := 0

	eg := &elfGroup{}
	for scanner.Scan() {
		line := scanner.Text()
		eg.rugsacks = append(eg.rugsacks, &rugsack{content: line})
		lineNumber++
		if lineNumber%groupSize == 0 { //new group
			elfGroups = append(elfGroups, eg)
			eg = &elfGroup{}
		}
	}
	return elfGroups, scanner.Err()
}
