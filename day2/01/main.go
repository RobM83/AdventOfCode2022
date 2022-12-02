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
		"X": 1, //Rock
		"Y": 2, //Paper
		"Z": 3, //Scissors
	}

	games, _ := readLines("input.txt")

	totalScore := 0
	for _, game := range games {

		myScore := pointMap[game.Hand2]

		if pointMap[game.Hand1] == pointMap[game.Hand2] {
			myScore += 3
			totalScore += myScore
			continue
		}
		//Scissors vs Rock
		if pointMap[game.Hand1] == 3 && pointMap[game.Hand2] == 1 {
			myScore += 6
			totalScore += myScore
			continue
		}
		//Rock vs Scissors
		if pointMap[game.Hand1] == 1 && pointMap[game.Hand2] == 3 {
			myScore += 0
			totalScore += myScore
			continue
		}
		//Other cases
		if pointMap[game.Hand1] < pointMap[game.Hand2] { //Win
			myScore += 6
			totalScore += myScore
			continue
		} else {
			myScore += 0
			totalScore += myScore
			continue
		}

	}
	fmt.Println(totalScore)
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
