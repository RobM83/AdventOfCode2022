package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	HEAD = 0
	TAIL = 9
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
	knots := []coord{}
	for i := 0; i < 10; i++ {
		knots = append(knots, coord{0, 0})
	} //knots[0] = H / knot[9] = T

	tailPositions := doMovement(actions, knots)

	fmt.Println(len(tailPositions))
}

func doMovement(actions []action, knots []coord) map[coord]bool {
	tailPositions := make(map[coord]bool)
	tailPositions[knots[TAIL]] = true

	for _, action := range actions {
		//Move head (add count)
		switch action.direction {
		case "R":
			for a := 0; a < action.count; a++ {
				knots[HEAD].x++
				for i := 1; i < len(knots); i++ {
					knots[i] = moveTail(knots[i-1], knots[i])
					tailPositions[knots[TAIL]] = true
				}
			}
		case "L":
			for a := 0; a < action.count; a++ {
				knots[HEAD].x--
				for i := 1; i < len(knots); i++ {
					knots[i] = moveTail(knots[i-1], knots[i])
					tailPositions[knots[TAIL]] = true
				}
			}
		case "U":
			for a := 0; a < action.count; a++ {
				knots[HEAD].y++
				for i := 1; i < len(knots); i++ {
					knots[i] = moveTail(knots[i-1], knots[i])
					tailPositions[knots[TAIL]] = true
				}
			}
		case "D":
			for a := 0; a < action.count; a++ {
				knots[HEAD].y--
				for i := 1; i < len(knots); i++ {
					knots[i] = moveTail(knots[i-1], knots[i])
					tailPositions[knots[TAIL]] = true
				}
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
			if diff(head.x, tail.x) == 1 {
				tail.x = head.x
			} else if head.x > tail.x {
				tail.x++
			} else {
				tail.x--
			}
		} else {
			tail.y--
			if diff(head.x, tail.x) == 1 {
				tail.x = head.x
			} else if head.x > tail.x {
				tail.x++
			} else {
				tail.x--
			}
		}
	} else {
		if head.x > tail.x {
			tail.x++
			if diff(head.y, tail.y) == 1 {
				tail.y = head.y
			} else if head.y > tail.y {
				tail.y++
			} else {
				tail.y--
			}
		} else {
			tail.x--
			if diff(head.y, tail.y) == 1 {
				tail.y = head.y
			} else if head.y > tail.y {
				tail.y++
			} else {
				tail.y--
			}
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
