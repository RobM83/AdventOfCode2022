package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	elves, _ := readLines("input.txt")
	fmt.Println(highestAmountOfFood(elves))
}

type elf struct {
	food int
}

func highestAmountOfFood(elves []elf) int {
	highest := 0
	for _, elf := range elves {
		if elf.food > highest {
			highest = elf.food
		}
	}
	return highest
}

func readLines(path string) ([]elf, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var elves []elf
	e := elf{}
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			elves = append(elves, e)
			e = elf{}
		} else {
			i, err := strconv.Atoi(input)
			if err != nil {
				return nil, err
			}
			e.food += i
		}
	}
	return elves, scanner.Err()
}
