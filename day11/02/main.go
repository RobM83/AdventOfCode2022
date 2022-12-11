package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"math/big"
)

type monkey struct {
	number         int        //MonkeyNumber
	items          []*big.Int //Items (represent with worry level)
	operation      string     //How to get new worry level
	test           *big.Int   //Divisible by
	testTrue       int
	testFalse      int
	inspectedItems int
}

func main() {
	input, _ := readLines("input.txt")
	monkeys := getMonkeys(input)

	rounds := 10000
	for i := 0; i < rounds; i++ {
		playRound(monkeys)
		// if i == 0 || i == 19 || i == 999 || i == 1999 || i == 9999 {
		// 	fmt.Println("round: ", i+1)
		// 	printItems(monkeys)
		// 	printInspectedItems(monkeys)
		// }
		//fmt.Println(i)
	}

	fmt.Println("Monkey business: ", getMonkeyBusiness(monkeys))
}

func getMonkeyBusiness(monkeys map[int]*monkey) uint {
	inspectedItems := []int{}
	for i := 0; i < len(monkeys); i++ {
		inspectedItems = append(inspectedItems, monkeys[i].inspectedItems)
	}
	sort.Ints(inspectedItems)

	monkeyBusiness := uint(inspectedItems[len(inspectedItems)-1]) * uint(inspectedItems[len(inspectedItems)-2])

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

func greatestCommonDenominator(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func leastCommonMultiplier(a, b int64) int64 {
	return a * b / greatestCommonDenominator(a, b)
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
		for i, wl := range monkey.items {
			modVal := new(big.Int)
			//			math := big.NewInt(0)
			zero := big.NewInt(0)

			//Test - When thrown the monkey add the new item to the end of the list
			//if monkey.items[i]%monkey.test == 0 {
			leastCommonMultiple := int64(1)
			for _, m := range monkeys {
				leastCommonMultiple = leastCommonMultiplier(leastCommonMultiple, m.test.Int64())
			}

			wl = wl.Mod(wl, big.NewInt(leastCommonMultiple))
			modVal = modVal.Mod(wl, monkey.test)

			if modVal.Cmp(zero) == 0 {

				//			if math.Mod(monkey.items[i], monkey.test) == big.NewInt(0) {
				dstMonkey := monkeys[monkey.testTrue]
				dstMonkey.items = append(dstMonkey.items, monkey.items[i])
				monkey.inspectedItems++
			} else {
				dstMonkey := monkeys[monkey.testFalse]
				dstMonkey.items = append(dstMonkey.items, monkey.items[i])
				monkey.inspectedItems++
			}
		}
		monkey.items = []*big.Int{}
	}
}

func inspect(operation string, items []*big.Int) []*big.Int {
	//Examples
	// new = old * 19
	// new = old + 6
	// new = old * old
	newItems := []*big.Int{}
	operation = strings.Replace(operation, "new = ", "", -1)
	for _, wl := range items {
		newOps := strings.Replace(operation, "old", wl.String(), -1)
		newOps = strings.TrimLeft(newOps, " ")
		calc := strings.Split(newOps, " ")
		switch calc[1] {
		case "*":
			nrA := silentStringToBigInt(calc[0])
			nrB := silentStringToBigInt(calc[2])
			wl = wl.Mul(nrA, nrB)
		case "+":
			nrA := silentStringToBigInt(calc[0])
			nrB := silentStringToBigInt(calc[2])
			wl = wl.Add(nrA, nrB)
		}
		newItems = append(newItems, wl)
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
				m.items = append(m.items, silentStringToBigInt(item))
			}
			continue
		} else if strings.Contains(line, "Operation:") {
			line = strings.Replace(line, "Operation: ", "", -1)
			m.operation = line
			continue
		} else if strings.Contains(line, "Test:") {
			line = strings.Replace(line, "Test: divisible by ", "", -1)
			line = strings.Replace(line, " ", "", -1)
			m.test = silentStringToBigInt(line)
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

func silentStringToBigInt(str string) *big.Int {
	n := big.Int{}
	n.SetString(str, 10)
	return &n
}
