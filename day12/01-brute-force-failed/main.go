package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	m, s, e, maxHeight, maxWidth := readMap(input)
	paths := []int{}
	visited := []coord{}
	getPaths(m, s, e, maxHeight, maxWidth, 0, visited, &paths)

	sort.Ints(paths)

	fmt.Println(paths[0])
}

func getPaths(m map[coord]*rock, start coord, end coord, maxHeight, MaxWidth, hopCount int, visited []coord, paths *[]int) {

	//fmt.Printf("Start: %v\t Hopcount: %d\t Visited: %v\n", start, hopCount, visited)

	if start.x == end.x && start.y == end.y {
		*paths = append(*paths, len(visited))
		//fmt.Printf("Hops: %v \n", visited)
		hopCount = 0
		return
	} else {
		hopCount++
		visited = append(visited, start)
	}

	pd := checkDirections(start, m, maxHeight, MaxWidth)

	if !pd[UP] && !pd[DOWN] && !pd[LEFT] && !pd[RIGHT] {
		return //No more moves
	}

	isVisited := func(c coord) bool {
		for _, v := range visited {
			if v.x == c.x && v.y == c.y {
				return true
			}
		}
		return false
	}

	if pd[UP] {
		newCoord := coord{start.x, start.y - 1}
		if !isVisited(newCoord) {
			getPaths(m, newCoord, end, maxHeight, MaxWidth, hopCount, visited, paths)
		}
	}
	if pd[DOWN] {
		newCoord := coord{start.x, start.y + 1}
		if !isVisited(newCoord) {
			getPaths(m, newCoord, end, maxHeight, MaxWidth, hopCount, visited, paths)
		}
	}
	if pd[LEFT] {
		newCoord := coord{start.x - 1, start.y}
		if !isVisited(newCoord) {
			getPaths(m, newCoord, end, maxHeight, MaxWidth, hopCount, visited, paths)
		}
	}
	if pd[RIGHT] {
		newCoord := coord{start.x + 1, start.y}
		if !isVisited(newCoord) {
			getPaths(m, newCoord, end, maxHeight, MaxWidth, hopCount, visited, paths)
		}
	}
}

func checkDirections(c coord, m map[coord]*rock, maxHeight, maxWidth int) []bool {

	doAble := func(from, to int) bool {
		if from >= to {
			return true
		} else if to-from == 1 {
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
