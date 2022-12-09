package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type coord struct {
	x int
	y int
}

type tree struct {
	hight             int
	visible           bool
	visibleTreesLeft  int
	visibleTreesRight int
	visibleTreesUp    int
	visibleTreesDown  int
}

func main() {
	input, _ := readLines("input.txt")
	grid := make(map[coord]*tree)
	maxWidth, maxHeight := gridToMap(input, grid)
	setVisibility(grid, maxWidth, maxHeight)
	widest := getWidestView(grid)
	fmt.Println(widest)
}

func getWidestView(grid map[coord]*tree) int {
	hightest := 0
	for _, tree := range grid {
		score := tree.visibleTreesLeft * tree.visibleTreesRight * tree.visibleTreesUp * tree.visibleTreesDown
		if score > hightest {
			hightest = score
		}
	}
	return hightest
}

func setVisibility(grid map[coord]*tree, maxWidth, maxHeight int) {
	for y := 0; y < maxHeight; y++ {
		for x := 0; x < maxWidth; x++ {
			isVisibleInDirection(grid, maxWidth, maxHeight, x, y, 0, "")
		}
	}
}

func isVisibleInDirection(grid map[coord]*tree, maxWidth, maxHeight, x, y, offset int, direction string) {
	//fmt.Println(x, y, offset, direction)
	if direction == "" {
		direction = "left"
	}
	if offset == 0 {
		offset = 1
	}
	switch direction {
	case "left":
		if x-offset < 0 {
			isVisibleInDirection(grid, maxWidth, maxHeight, x, y, 1, "right")
		} else {
			grid[coord{x, y}].visibleTreesLeft++
			if grid[coord{x, y}].hight <= grid[coord{x - offset, y}].hight {
				isVisibleInDirection(grid, maxWidth, maxHeight, x, y, 1, "right")
			} else {
				isVisibleInDirection(grid, maxWidth, maxHeight, x, y, offset+1, direction)
			}
		}
	case "right":
		if x+offset > maxWidth-1 {
			isVisibleInDirection(grid, maxWidth, maxHeight, x, y, 1, "up")
		} else {
			grid[coord{x, y}].visibleTreesRight++
			if grid[coord{x, y}].hight <= grid[coord{x + offset, y}].hight {
				isVisibleInDirection(grid, maxWidth, maxHeight, x, y, 1, "up")
			} else {
				isVisibleInDirection(grid, maxWidth, maxHeight, x, y, offset+1, direction)
			}
		}
	case "up":
		if y-offset < 0 {
			isVisibleInDirection(grid, maxWidth, maxHeight, x, y, 1, "down")
		} else {
			grid[coord{x, y}].visibleTreesUp++
			if grid[coord{x, y}].hight <= grid[coord{x, y - offset}].hight {
				isVisibleInDirection(grid, maxWidth, maxHeight, x, y, 1, "down")
			} else {
				isVisibleInDirection(grid, maxWidth, maxHeight, x, y, offset+1, direction)
			}
		}
	case "down":
		if y+offset > maxHeight-1 {
			return
		} else {
			grid[coord{x, y}].visibleTreesDown++
			if grid[coord{x, y}].hight <= grid[coord{x, y + offset}].hight {
				return
			} else {
				isVisibleInDirection(grid, maxWidth, maxHeight, x, y, offset+1, direction)
			}
		}
	}
	return
}

// returns map and max x and y
func gridToMap(input []string, grid map[coord]*tree) (int, int) {
	maxHeight := 0
	maxWidth := 0
	for y, line := range input {
		maxWidth = len(line)
		maxHeight++
		for x, char := range line {
			t := tree{hight: silentStringToInt(string(char)), visible: true}
			grid[coord{x, y}] = &t
		}
	}
	return maxWidth, maxHeight
}

func silentStringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
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
