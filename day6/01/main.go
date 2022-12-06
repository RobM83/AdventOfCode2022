package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	line, _ := readLines("input.txt")
	sp := getStartPosition(line)
	fmt.Println(sp)
}

func getStartPosition(line string) int {
	lastFour := []string{"", "", "", ""}
	for i := 0; i < len(line); i++ {
		lastFour = lastFour[1:]
		lastFour = append(lastFour, string(line[i]))

		if (i > 3) && uniqueSlice(lastFour) {
			return i + 1
		}
		fmt.Println(lastFour)
	}

	return 1
}

func uniqueSlice(slice []string) bool {
	seen := make(map[string]bool)
	for _, v := range slice {
		if seen[v] {
			return false
		}
		seen[v] = true
	}
	return true
}

func readLines(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var output string
	for scanner.Scan() {
		output += scanner.Text()
	}
	return output, scanner.Err()
}
