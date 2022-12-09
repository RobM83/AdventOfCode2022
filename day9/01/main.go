package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type action struct {
	direction string
	count     int
}

type coord struct {
	x int
	y int
}

func main() {
	input, _ := readLines("input.txt")
	actions := getActions(input)
	head := coord{0, 0}
	tail := coord{0, 0}
	tailPositions := doMovement(actions, head, tail)

	fmt.Println(len(tailPositions))
}

func doMovement(actions []action, head coord, tail coord) map[coord]bool {
	headPositions := make(map[coord]bool)
	tailPositions := make(map[coord]bool)

	headPositions[head] = true
	tailPositions[tail] = true

	for _, action := range actions {
		//Move head (add count)
		switch action.direction {
		case "R":
			for i := 0; i < action.count; i++ {
				head.x++
				//Move T
				tail = moveTail(head, tail)
				headPositions[head] = true
				tailPositions[tail] = true
			}
		case "L":
			for i := 0; i < action.count; i++ {
				head.x--
				tail = moveTail(head, tail)
				headPositions[head] = true
				tailPositions[tail] = true
			}
		case "U":
			for i := 0; i < action.count; i++ {
				head.y++
				tail = moveTail(head, tail)
				headPositions[head] = true
				tailPositions[tail] = true
			}
		case "D":
			for i := 0; i < action.count; i++ {
				head.y--
				tail = moveTail(head, tail)
				headPositions[head] = true
				tailPositions[tail] = true
			}
		default:
			panic("Unknown direction")
		}
	}

	return tailPositions
}

func moveTail(head coord, tail coord) coord {
	//Move tail if needed
	if diff(head.x, tail.x)+diff(head.y, tail.y) > 2 {
		//Move diagonal
		tail = moveDiagonal(head, tail)
	} else if diff(head.x, tail.x) == 2 {
		//move X
		if head.x > tail.x {
			tail.x++
		} else {
			tail.x--
		}
	} else if diff(head.y, tail.y) == 2 {
		//move Y
		if head.y > tail.y {
			tail.y++
		} else {
			tail.y--
		}
	}

	return tail
}

func moveDiagonal(head coord, tail coord) coord {
	if diff(head.y, tail.y) == 2 {
		if head.y > tail.y {
			tail.y++
			tail.x = head.x
		} else {
			tail.y--
			tail.x = head.x
		}
	} else {
		if head.x > tail.x {
			tail.x++
			tail.y = head.y
		} else {
			tail.x--
			tail.y = head.y
		}
	}
	return tail
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func getActions(input []string) []action {
	var actions []action
	for _, line := range input {
		l := strings.Split(line, " ")
		actions = append(actions, action{l[0], silentStringToInt(l[1])})
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
