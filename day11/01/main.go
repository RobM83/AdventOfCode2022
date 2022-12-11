package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	number         int    //MonkeyNumber
	items          []int  //Items (represent with worry level)
	operation      string //How to get new worry level
	test           int    //Divisible by
	testTrue       int
	testFalse      int
	inspectedItems int
}

func main() {
	input, _ := readLines("input.txt")
	monkeys := getMonkeys(input)

	rounds := 20
	for i := 0; i < rounds; i++ {
		playRound(monkeys)
		fmt.Println("round: ", i)
		printItems(monkeys)
	}

	printInspectedItems(monkeys)

	fmt.Println("Monkey business: ", getMonkeyBusiness(monkeys))
}

func getMonkeyBusiness(monkeys map[int]*monkey) int {
	inspectedItems := []int{}
	for i := 0; i < len(monkeys); i++ {
		inspectedItems = append(inspectedItems, monkeys[i].inspectedItems)
	}
	sort.Ints(inspectedItems)

	monkeyBusiness := inspectedItems[len(inspectedItems)-1] * inspectedItems[len(inspectedItems)-2]

	return monkeyBusiness
}

func printInspectedItems(monkeys map[int]*monkey) {
	for i := 0; i < len(monkeys); i++ {
		fmt.Printf("Monkey %d inspected %d items\n", monkeys[i].number, monkeys[i].inspectedItems)
	}
}

func printItems(monkeys map[int]*monkey) {
	for i := 0; i < len(monkeys); i++ {
		fmt.Printf("Monkey %d has items %d \n", monkeys[i].number, monkeys[i].items)
	}
}

func playRound(monkeys map[int]*monkey) {
	//Each round every monkey plays
	for m := 0; m < len(monkeys); m++ {
		monkey := monkeys[m]
		if len(monkey.items) == 0 {
			continue //Next monkey
		} else {
			//Inspect
			monkey.items = inspect(monkey.operation, monkey.items)
		}
		//Divide by 3 and round down
		for i, item := range monkey.items {
			monkey.items[i] = int(math.Floor(float64(item) / 3))
			//Test - When thrown the monkey add the new item to the end of the list
			if monkey.items[i]%monkey.test == 0 {
				dstMonkey := monkeys[monkey.testTrue]
				dstMonkey.items = append(dstMonkey.items, monkey.items[i])
				monkey.inspectedItems++
			} else {
				dstMonkey := monkeys[monkey.testFalse]
				dstMonkey.items = append(dstMonkey.items, monkey.items[i])
				monkey.inspectedItems++
			}
		}
		monkey.items = []int{}
	}
}

func inspect(operation string, items []int) []int {
	//Examples
	// new = old * 19
	// new = old + 6
	// new = old * old
	newItems := []int{}
	operation = strings.Replace(operation, "new = ", "", -1)
	for _, wl := range items {
		newOps := strings.Replace(operation, "old", strconv.Itoa(wl), -1)
		newOps = strings.TrimLeft(newOps, " ")
		calc := strings.Split(newOps, " ")
		worryLevel := 0
		switch calc[1] {
		case "*":
			worryLevel = silentStringToInt(calc[0]) * silentStringToInt(calc[2])
		case "+":
			worryLevel = silentStringToInt(calc[0]) + silentStringToInt(calc[2])
		}
		newItems = append(newItems, worryLevel)
	}
	return newItems
}

func getMonkeys(input []string) map[int]*monkey {
	monkeys := make(map[int]*monkey)
	var m *monkey
	for _, line := range input {

		if strings.Contains(line, "Monkey") {
			//New Monkey
			monkey := monkey{}
			m = &monkey
			line = strings.Replace(line, "Monkey ", "", -1)
			line = strings.Replace(line, ":", "", -1)
			m.number = silentStringToInt(line)
			continue
		} else if strings.Contains(line, "Starting items:") {
			line = strings.Replace(line, "Starting items: ", "", -1)
			line = strings.Replace(line, " ", "", -1)
			items := strings.Split(line, ",")
			for _, item := range items {
				m.items = append(m.items, silentStringToInt(item))
			}
			continue
		} else if strings.Contains(line, "Operation:") {
			line = strings.Replace(line, "Operation: ", "", -1)
			m.operation = line
			continue
		} else if strings.Contains(line, "Test:") {
			line = strings.Replace(line, "Test: divisible by ", "", -1)
			line = strings.Replace(line, " ", "", -1)
			m.test = silentStringToInt(line)
			continue
		} else if strings.Contains(line, "If true:") {
			line = strings.Replace(line, "If true: throw to monkey ", "", -1)
			line = strings.Replace(line, " ", "", -1)
			m.testTrue = silentStringToInt(line)
			continue
		} else if strings.Contains(line, "If false:") {
			line = strings.Replace(line, "If false: throw to monkey ", "", -1)
			line = strings.Replace(line, " ", "", -1)
			m.testFalse = silentStringToInt(line)
			monkeys[m.number] = m
			continue
		} else {
			continue
		}
	}
	return monkeys
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
