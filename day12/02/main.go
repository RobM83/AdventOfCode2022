package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	UP    = 0
	DOWN  = 1
	LEFT  = 2
	RIGHT = 3
)

type coord struct {
	x int
	y int
}

type rock struct {
	elevation string //Original input
	height    int    //Heigh in decimal (a = 1, b = 2, etc)
}

func main() {
	input, _ := readLines("input.txt")
	m, _, e, maxHeight, maxWidth := readMap(input)

	visited := make(map[coord]int)
	unvisited := make(map[coord]int)

	for y := 0; y <= maxHeight; y++ {
		for x := 0; x <= maxWidth; x++ {
			coord := coord{x, y}
			unvisited[coord] = m[coord].height
		}
	}

	//Start with end
	getCost(m, e, maxHeight, maxWidth, visited, unvisited, 0)

	fmt.Println(getNearestA(m, visited))

}

func getNearestA(m map[coord]*rock, visited map[coord]int) int {
	steps := 9999
	for k, v := range m {
		if v.elevation == "a" || v.elevation == "S" {
			fmt.Printf("Found A at %v, with value %d\n", k, visited[k])
			if visited[k] < steps && visited[k] > 0 {
				steps = visited[k]
			}
		}
	}
	return steps
}

func getCost(m map[coord]*rock, c coord, maxHeight, maxWidth int, visited, unvisited map[coord]int, cost int) {

	//Remove from unvisited
	delete(unvisited, c)

	if _, ok := visited[c]; ok {
		if cost < visited[c] {
			visited[c] = cost
		}
	} else {
		visited[c] = cost
	}

	pd := checkDirections(c, m, maxHeight, maxWidth)

	if !pd[UP] && !pd[DOWN] && !pd[LEFT] && !pd[RIGHT] {
		return //No more moves (from this point)
	}

	if pd[UP] {
		newCoord := coord{c.x, c.y - 1}
		if _, ok := unvisited[newCoord]; ok {
			getCost(m, newCoord, maxHeight, maxWidth, visited, unvisited, cost+1)
		} else if visited[newCoord] > cost+1 {
			getCost(m, newCoord, maxHeight, maxWidth, visited, unvisited, cost+1)
		}
	}
	if pd[DOWN] {
		newCoord := coord{c.x, c.y + 1}
		if _, ok := unvisited[newCoord]; ok {
			getCost(m, newCoord, maxHeight, maxWidth, visited, unvisited, cost+1)
		} else if visited[newCoord] > cost+1 {
			getCost(m, newCoord, maxHeight, maxWidth, visited, unvisited, cost+1)
		}
	}
	if pd[LEFT] {
		newCoord := coord{c.x - 1, c.y}
		if _, ok := unvisited[newCoord]; ok {
			getCost(m, newCoord, maxHeight, maxWidth, visited, unvisited, cost+1)
		} else if visited[newCoord] > cost+1 {
			getCost(m, newCoord, maxHeight, maxWidth, visited, unvisited, cost+1)
		}
	}
	if pd[RIGHT] {
		newCoord := coord{c.x + 1, c.y}
		if _, ok := unvisited[newCoord]; ok {
			getCost(m, newCoord, maxHeight, maxWidth, visited, unvisited, cost+1)
		} else if visited[newCoord] > cost+1 {
			getCost(m, newCoord, maxHeight, maxWidth, visited, unvisited, cost+1)
		}
	}
}

func checkDirections(c coord, m map[coord]*rock, maxHeight, maxWidth int) []bool {

	//Reverse rules going from End to Start
	doAble := func(from, to int) bool {
		if from <= to {
			return true
		} else if from-to == 1 {
			return true
		} else {
			return false
		}
	}

	possibleDirections := []bool{false, false, false, false}
	if c.y-1 >= 0 {
		if doAble(m[c].height, m[coord{c.x, c.y - 1}].height) {
			possibleDirections[UP] = true
		}
	}
	if c.y+1 <= maxHeight {
		if doAble(m[c].height, m[coord{c.x, c.y + 1}].height) {
			possibleDirections[DOWN] = true
		}
	}
	if c.x-1 >= 0 {
		if doAble(m[c].height, m[coord{c.x - 1, c.y}].height) {
			possibleDirections[LEFT] = true
		}
	}
	if c.x+1 <= maxWidth {
		if doAble(m[c].height, m[coord{c.x + 1, c.y}].height) {
			possibleDirections[RIGHT] = true
		}
	}

	return possibleDirections
}

// returns map, start, end, max height, max width
func readMap(input []string) (map[coord]*rock, coord, coord, int, int) {
	m := map[coord]*rock{}
	start := coord{}
	end := coord{}
	mh := 0
	mw := 0
	for y, line := range input {
		mh = y
		for x, char := range line {
			mw = x
			coord := coord{x, y}
			rock := &rock{string(char), elevationToHeight(string(char))}
			m[coord] = rock
			if string(char) == "S" {
				start = coord
			}
			if string(char) == "E" {
				end = coord
			}
		}
	}
	return m, start, end, mh, mw
}

func elevationToHeight(elevation string) int {
	transmap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26, "S": 1, "E": 26}
	return transmap[elevation]
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
