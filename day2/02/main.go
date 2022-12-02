package main

import (
	"bufio"
	"fmt"
	"os"
)

type game struct {
	Hand1 string
	Hand2 string
}

func main() {
	pointMap := map[string]int{
		"A": 1, //Rock
		"B": 2, //Paper
		"C": 3, //Scissors
		"X": 1, //Lose
		"Y": 2, //Win
		"Z": 3, //Draw
	}

	games, _ := readLines("input.txt")

	totalScore := 0
	for _, game := range games {

		switch game.Hand2 {
		case "X": //Lose
			totalScore += pointMap[getCounterHandLose(game.Hand1)]
		case "Y": //Draw
			totalScore += pointMap[game.Hand1] + 3
		case "Z": //win
			totalScore += pointMap[getCounterHandWin(game.Hand1)] + 6
		}
	}
	fmt.Println(totalScore)
}

func getCounterHandWin(hand string) string {
	switch hand {
	case "A":
		return "B"
	case "B":
		return "C"
	case "C":
		return "A"
	}
	return ""
}

func getCounterHandLose(hand string) string {
	switch hand {
	case "A":
		return "C"
	case "B":
		return "A"
	case "C":
		return "B"
	}
	return ""
}

func readLines(path string) ([]game, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var games []game
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		games = append(games, game{line[0:1], line[2:3]})
	}
	return games, scanner.Err()
}
