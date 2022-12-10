package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type action struct {
	action string
	count  int
}

func main() {
	input, _ := readLines("input.txt")
	actions := getActions(input)
	total := process(actions)

	fmt.Printf("Total: %d\n", total)
}

func process(actions []action) int {
	cycle := 0
	X := 1
	total := 0

	increaseCycle := func(times int) {
		for i := 0; i < times; i++ {
			cycle++
			if (cycle == 20) || ((cycle-20)%40 == 0) && (cycle != 40) {
				fmt.Printf("Cycle %d: X=%d; %d\n", cycle, X, cycle*X)
				total += cycle * X
			}
		}
	}

	for _, action := range actions {
		switch action.action {
		case "noop": //Takes one cycle
			// do nothing
			increaseCycle(1)
		case "addx": //Takes two cycles
			increaseCycle(2)
			X += action.count
		}
	}

	return total
}

func getActions(input []string) []action {
	var actions []action
	for _, line := range input {
		l := strings.Split(line, " ")
		if l[0] != "noop" {
			actions = append(actions, action{l[0], silentStringToInt(l[1])})
		} else {
			actions = append(actions, action{"noop", 0})
		}
	}
	return actions
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func silentStringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
