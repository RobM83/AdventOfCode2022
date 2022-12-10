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
	process(actions)
}

func process(actions []action) {
	cycle := 0
	X := 1

	spritePos := [3]int{0, 1, 2}
	line := [40]string{}
	linePos := 0

	setSpritePosition := func(pos int) {
		spritePos[0] = pos - 1
		spritePos[1] = pos
		spritePos[2] = pos + 1
	}

	drawLine := func(line []string) {
		for _, l := range line {
			fmt.Printf("%s", l)
		}
		fmt.Printf("\n")
	}

	increaseCycle := func(times int) {
		for i := 0; i < times; i++ {
			if linePos == spritePos[0] || linePos == spritePos[1] || linePos == spritePos[2] {
				line[linePos] = "#"
			} else {
				line[linePos] = "."
			}
			cycle++
			linePos++
			if cycle%40 == 0 {
				drawLine(line[:])
				line = [40]string{}
				linePos = 0
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
			setSpritePosition(X)
		}
	}
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
