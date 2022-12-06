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
	buffer := []string{}
	for i := 0; i < len(line); i++ {
		buffer = append(buffer, string(line[i]))
		if !uniqueSlice(buffer) {
			buffer = buffer[1:]
		} else if len(buffer) == 14 {
			return i + 1
		}
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
