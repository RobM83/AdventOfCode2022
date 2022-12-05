package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type crates struct {
	stack map[int][]string
}

type operation struct {
	count int
	from  int
	to    int
}

func NewCrates() *crates {
	c := &crates{}
	c.stack = make(map[int][]string)

	return c
}

func main() {
	c := NewCrates()
	ops, err := readLines("input.txt", c)
	if err != nil {
		panic(err)
	}
	arrange(c, ops)

	output := ""
	for i := 1; i <= len(c.stack); i++ {
		clean := strings.Replace(c.stack[i][len(c.stack[i])-1], "[", "", -1)
		clean = strings.Replace(clean, "]", "", -1)
		output = output + clean
	}
	fmt.Println(output)
}

func arrange(c *crates, ops []*operation) {
	for _, op := range ops {
		c.stack[op.to] = append(c.stack[op.to], c.stack[op.from][len(c.stack[op.from])-op.count:]...)
		c.stack[op.from] = c.stack[op.from][:len(c.stack[op.from])-op.count]
	}
}

func readLines(path string, c *crates) ([]*operation, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	head := true //Read header

	var ops []*operation
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			head = false
			continue
		}

		if head {
			// 4 positions
			//Read per 3 charactes
			i := 0
			col := 1
			for i < len(line) {
				column := line[i : i+3]
				if strings.Contains(column, "[") {
					if _, ok := c.stack[col]; ok {
						c.stack[col] = toBottom(c.stack[col], column)
					} else {
						c.stack[col] = []string{column}
					}
				}
				i += 4
				col++
			}

		} else {
			var op operation
			l := strings.Split(line, " ")
			op.count, _ = strconv.Atoi(l[1])
			op.from, _ = strconv.Atoi(l[3])
			op.to, _ = strconv.Atoi(l[5])
			ops = append(ops, &op)
		}
	}
	return ops, scanner.Err()
}

func pop(stack []string) (string, []string) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}

func toBottom(stack []string, crate string) []string {
	newStack := []string{crate}
	newStack = append(newStack, stack...)
	return newStack
}
