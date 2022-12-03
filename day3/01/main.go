package main

import (
	"bufio"
	"fmt"
	"os"
)

type rugsack struct {
	content           string
	itemsCompartment1 map[string]int
	itemsCompartment2 map[string]int
	common            string
	commonPriority    int
}

func main() {
	rugsacks, _ := readLines("input.txt")
	checkRugsacks(rugsacks)

	sumOfPriorities := 0
	for _, r := range rugsacks {
		// fmt.Println(r.content)
		// fmt.Println(r.itemsCompartment1)
		// fmt.Println(r.itemsCompartment2)
		// fmt.Println(r.common)
		// fmt.Println(r.commonPriority)
		sumOfPriorities += r.commonPriority
	}

	fmt.Printf("Sum of priorities: %d\n", sumOfPriorities)
}

func checkRugsacks(rugsacks []*rugsack) {
	for _, r := range rugsacks {
		r.itemsCompartment1 = make(map[string]int)
		r.itemsCompartment2 = make(map[string]int)

		for i := 0; i < len(r.content)/2; i++ {
			r.itemsCompartment1[r.content[i:i+1]] = getPriority(r.content[i : i+1])
			pos := i + (len(r.content) / 2)
			r.itemsCompartment2[r.content[pos:pos+1]] = getPriority(r.content[pos : pos+1])
		}

		for k := range r.itemsCompartment1 {
			if val, ok := r.itemsCompartment2[k]; ok {
				r.common = k
				r.commonPriority = val
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

func readLines(path string) ([]*rugsack, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rugsacks []*rugsack
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rugsacks = append(rugsacks, &rugsack{content: line})
	}
	return rugsacks, scanner.Err()
}
