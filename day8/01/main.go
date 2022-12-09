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
	hight   int
	visible bool
}

func main() {
	input, _ := readLines("test.txt")
	grid := make(map[coord]tree)
	maxWidth, maxHeight := gridToMap(input, grid)
	setVisibility(grid, maxWidth, maxHeight)
	fmt.Println(grid)
	visible := countVisibleTrees(grid)
	fmt.Println(visible)
}

func countVisibleTrees(grid map[coord]tree) int {
	visible := 0
	for _, tree := range grid {
		if tree.visible {
			visible++
		}
	}
	return visible
}

func setVisibility(grid map[coord]tree, maxWidth, maxHeight int) {
	for y := 1; y < maxHeight-1; y++ {
		for x := 1; x < maxWidth-1; x++ {
			isVisibleInDirection(grid, maxWidth, maxHeight, x, y, 0, "")
		}
	}
}

func isVisibleInDirection(grid map[coord]tree, maxWidth, maxHeight, x, y, offset int, direction string) {
	if direction == "" {
		direction = "left"
	}
	if offset == 0 {
		offset = 1
	}
	switch direction {
	case "left":
		if grid[coord{x, y}].hight <= grid[coord{x - offset, y}].hight {
			isVisibleInDirection(grid, maxWidth, maxHeight, x, y, 1, "right")
		} else {
			if x-offset > 0 {
				isVisibleInDirection(grid, maxWidth, maxHeight, x, y, offset+1, direction)
			}
		}
	case "right":
		if grid[coord{x, y}].hight <= grid[coord{x + offset, y}].hight {
			isVisibleInDirection(grid, maxWidth, maxHeight, x, y, 1, "up")
		} else {
			if x+offset < maxWidth-1 {
				isVisibleInDirection(grid, maxWidth, maxHeight, x, y, offset+1, direction)
			}
		}
	case "up":
		if grid[coord{x, y}].hight <= grid[coord{x, y - offset}].hight {
			isVisibleInDirection(grid, maxWidth, maxHeight, x, y, 1, "down")
		} else {
			if y-offset > 0 {
				isVisibleInDirection(grid, maxWidth, maxHeight, x, y, offset+1, direction)
			}
		}
	case "down":
		if grid[coord{x, y}].hight <= grid[coord{x, y + offset}].hight {
			//Now all directions are false
			grid[coord{x, y}] = tree{hight: grid[coord{x, y}].hight, visible: false}
			return
		} else {
			if y+offset < maxHeight-1 {
				isVisibleInDirection(grid, maxWidth, maxHeight, x, y, offset+1, direction)
			}
		}
	}
	return
}

// returns map and max x and y
func gridToMap(input []string, grid map[coord]tree) (int, int) {
	maxHeight := 0
	maxWidth := 0
	for y, line := range input {
		maxWidth = len(line)
		maxHeight++
		for x, char := range line {
			grid[coord{x, y}] = tree{hight: silentStringToInt(string(char)), visible: true}
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
