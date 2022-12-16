package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type input struct {
	left  string
	right string
}

type pair struct {
	left  []any
	right []any
}

func main() {
	input, _ := readLines("input.txt")
	pairs := disectPairs(input)

	packets := []any{}

	for _, p := range pairs {
		packets = append(packets, p.left)
		packets = append(packets, p.right)
	}

	//Add [[2]] and [[6]]
	packets = append(packets, []any{[]any{2.0}})
	packets = append(packets, []any{[]any{6.0}})

	//Sort packets
	sort.Slice(packets, func(i, j int) bool { return compare(packets[i], packets[j]) < 0 })

	//Find index of packet [[2]] and [[6]]
	score := 1 //Can't multiply with 0
	for i, p := range packets {
		fmt.Println(p)
		if fmt.Sprint(p) == "[[2]]" {
			fmt.Printf("Index of [[2]]: %d", i+1)
			score = score * (i + 1)
		}
		if fmt.Sprint(p) == "[[6]]" {
			fmt.Printf("Index of [[6]]: %d", i+1)
			score = score * (i + 1)
		}
	}
	fmt.Println(score)
}

func disectPairs(input []*input) []pair {
	var pairs []pair
	for _, p := range input {
		left := []any{}
		right := []any{}
		json.Unmarshal([]byte(p.left), &left)
		json.Unmarshal([]byte(p.right), &right)
		pairs = append(pairs, pair{left, right})
	}
	return pairs
}

func compare(left, right any) int {
	l, lok := left.([]any)
	r, rok := right.([]any)

	switch {
	case !lok && !rok:
		return int(left.(float64) - right.(float64))
	case !lok:
		l = []any{left}
	case !rok:
		r = []any{right}
	}

	for i := 0; i < len(l) && i < len(r); i++ {
		if c := compare(l[i], r[i]); c != 0 {
			return c
		}
	}
	return len(l) - len(r)
}

func readLines(path string) ([]*input, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var pairs []*input
	for scanner.Scan() {
		left := scanner.Text()
		scanner.Scan()
		right := scanner.Text()
		scanner.Scan() //Empty
		pairs = append(pairs, &input{left, right})
	}
	return pairs, scanner.Err()
}

func silentStringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
