package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type item struct {
	dir  bool
	name string
	size int
	//items []*item
}

func main() {
	input, _ := readLines("input.txt")
	folders := processLines(input)
	folderSizes := make(map[string]int)
	getFolderSize(folders, "/", &folderSizes)
	total := getSumFoldersBelowSize(&folderSizes, 100000)
	fmt.Println(total)
}

func processLines(lines []string) *map[string][]*item {
	folders := make(map[string][]*item)
	path := []string{}
	currentPath := ""
	for _, line := range lines {
		//First char $ is command
		if line[0] == '$' {
			cmd := strings.Split(line[1:], " ")

			switch cmd[1] {
			case "cd": //Change directory
				if cmd[2] != ".." { //Go to cmd[2]
					path = append(path, cmd[2])
					currentPath = strings.Join(path, "/")
					//Items
				} else {
					path = path[:len(path)-1]
					currentPath = strings.Join(path, "/")
				}
			case "ls": //List directory
				continue
			}
		} else { //Read item(s)
			s := 0     //Size
			d := false //is Directory
			itemRaw := strings.Split(line, " ")
			if itemRaw[0] == "dir" { //Directory
				d = true
			} else {
				s, _ = strconv.Atoi(itemRaw[0])
			}
			itm := &item{dir: d, name: itemRaw[1], size: s}
			folders[currentPath] = append(folders[currentPath], itm)
		}
	}
	return &folders
}

func getFolderSize(folders *map[string][]*item, folder string, folderSizes *map[string]int) int {
	items := (*folders)[folder]
	size := 0
	for _, item := range items {
		if item.dir == false {
			size += item.size
		} else {
			folder := folder + "/" + item.name
			size += getFolderSize(folders, folder, folderSizes)
		}
	}
	(*folderSizes)[folder] = size
	fmt.Println(folder, size)
	return size
}

func getSumFoldersBelowSize(folderSizes *map[string]int, maxSize int) int {
	total := 0
	for _, size := range *folderSizes {
		if size <= maxSize {
			total += size
		}
	}
	return total
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
