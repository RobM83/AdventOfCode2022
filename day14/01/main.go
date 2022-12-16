package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	ROCK   = "#"
	AIR    = "."
	SOURCE = "+"
	SAND   = "o"
)

type coord struct {
	x int
	y int
}

func main() {
	input, _ := readLines("input.txt")

	cave, abyss := createCave(input)

	dropSand(cave, coord{500, 0}, abyss)

	count := 0
	for c, s := range cave {
		if s == SAND {
			count++
		}
		fmt.Printf("Coord: [%v], Symbol: [%s]\n", c, s)
	}

	printCave(cave)

	fmt.Println("Count: ", count)
}

// For debugging purposes
func printCave(cave map[coord]string) {
	xCoords := make([]int, 0)
	yCoords := make([]int, 0)

	for c := range cave {
		xCoords = append(xCoords, c.x)
		yCoords = append(yCoords, c.y)
	}

	sort.Ints(xCoords)
	sort.Ints(yCoords)

	for y := yCoords[0]; y <= yCoords[len(yCoords)-1]; y++ {
		for x := xCoords[0]; x <= xCoords[len(xCoords)-1]; x++ {
			if _, ok := cave[coord{x, y}]; ok {
				fmt.Printf("%s", cave[coord{x, y}])
			} else {
				fmt.Printf("%s", AIR)
			}
		}
		fmt.Println()
	}
	fmt.Println("Max-y	", yCoords[len(yCoords)-1])
}

// Returns abyss (reached) and drop count
func dropSand(cave map[coord]string, start coord, abyss int) bool {
	end := false
	sand := start

	//	for falling := true; falling; falling = !(checkNextCoord(cave, sand, abyss) == coord{-1, -1}) {
	for {
		sand, end = checkNextCoord(cave, sand, abyss)
		if end || (sand == coord{-1, -1}) {
			break
		}
	}

	if !end {
		//cave[sand] = SAND
		end = dropSand(cave, start, abyss)
	}

	return end
}

// returns (-1, -1) if no next coord and abyss
func checkNextCoord(cave map[coord]string, start coord, maxDepth int) (coord, bool) {
	abyss := false
	//Direction down
	if _, ok := cave[coord{start.x, start.y + 1}]; !ok {
		if start.y+1 >= maxDepth {
			abyss = true
		}

		return coord{start.x, start.y + 1}, abyss
	}
	//Direction down-left
	if _, ok := cave[coord{start.x - 1, start.y + 1}]; !ok {
		if start.y+1 >= maxDepth {
			abyss = true
		}
		return coord{start.x - 1, start.y + 1}, abyss
	}
	//Direction down-right
	if _, ok := cave[coord{start.x + 1, start.y + 1}]; !ok {
		if start.y+1 >= maxDepth {
			abyss = true
		}
		return coord{start.x + 1, start.y + 1}, abyss
	}
	cave[start] = SAND
	return coord{-1, -1}, abyss
}

// Retruns cave and abyss
func createCave(input []string) (map[coord]string, int) {
	cave := make(map[coord]string)
	cave[coord{500, 0}] = SOURCE

	for _, line := range input {
		path := strings.Split(line, "->")
		from := coord{-1, -1}
		for _, p := range path {
			p = strings.TrimSpace(p)
			if from.x == -1 && from.y == -1 {
				from = coord{silentStringToInt(strings.Split(p, ",")[0]), silentStringToInt(strings.Split(p, ",")[1])}
				cave[from] = ROCK
			} else {
				to := coord{silentStringToInt(strings.Split(p, ",")[0]), silentStringToInt(strings.Split(p, ",")[1])}
				if from.x == to.x {
					if from.y > to.y {
						for y := from.y; y >= to.y; y-- { //ABS
							cave[coord{from.x, y}] = ROCK
						}
					} else {
						for y := from.y; y <= to.y; y++ { //ABS
							cave[coord{from.x, y}] = ROCK
						}
					}
				} else {
					if from.x > to.x {
						for x := from.x; x >= to.x; x-- { //ABS
							cave[coord{x, from.y}] = ROCK
						}
					} else {
						for x := from.x; x <= to.x; x++ { //ABS
							cave[coord{x, from.y}] = ROCK
						}

					}
				}
				from = coord{silentStringToInt(strings.Split(p, ",")[0]), silentStringToInt(strings.Split(p, ",")[1])}
			}
		}
	}

	var yCoords []int
	for c := range cave {
		yCoords = append(yCoords, c.y)
	}

	sort.Ints(yCoords)

	return cave, yCoords[len(yCoords)-1]
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

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
