package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	SENSOR    = "S"
	BEACON    = "B"
	DISTRESS  = "D"
	NONBEACON = "#"
)

type coord struct {
	x int
	y int
}

type sensor struct {
	coord    coord
	beacon   coord
	distance int
}

func main() {
	input, _ := readLines("input.txt")

	sensors := getSensor(input)
	coords := getCoords(sensors)

	signals := getBeaconSignalsOnRow(coords, 2000000)

	printMap(coords)
	fmt.Println(signals)
}

func getBeaconSignalsOnRow(coords map[coord]string, row int) int {
	signals := 0

	for coord, value := range coords {
		if coord.y == row && (value == NONBEACON || value == SENSOR) {
			signals++
		}
	}
	return signals
}

func printMap(coords map[coord]string) {
	coordsX := make([]int, 0)
	coordsY := make([]int, 0)

	for coord := range coords {
		coordsX = append(coordsX, coord.x)
		coordsY = append(coordsY, coord.y)
	}

	sort.Ints(coordsX)
	sort.Ints(coordsY)

	for y := coordsY[0]; y <= coordsY[len(coordsY)-1]; y++ {
		for x := coordsX[0]; x <= coordsX[len(coordsX)-1]; x++ {
			if _, ok := coords[coord{x, y}]; ok {
				fmt.Print(coords[coord{x, y}])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func getCoords(sensors map[coord]*sensor) map[coord]string {
	coords := make(map[coord]string)
	for _, sensor := range sensors {
		coords[sensor.coord] = SENSOR
		coords[sensor.beacon] = BEACON

		//Get NonBeacon
		for i := 0; i <= sensor.distance; i++ {
			for d := 0; d <= diff(sensor.distance, i); d++ {
				if _, ok := coords[coord{sensor.coord.x + i, sensor.coord.y - d}]; !ok {
					coords[coord{sensor.coord.x + i, sensor.coord.y - d}] = NONBEACON
				}
				if _, ok := coords[coord{sensor.coord.x + i, sensor.coord.y + d}]; !ok {
					coords[coord{sensor.coord.x + i, sensor.coord.y + d}] = NONBEACON
				}
				if _, ok := coords[coord{sensor.coord.x - i, sensor.coord.y - d}]; !ok {
					coords[coord{sensor.coord.x - i, sensor.coord.y - d}] = NONBEACON
				}
				if _, ok := coords[coord{sensor.coord.x - i, sensor.coord.y + d}]; !ok {
					coords[coord{sensor.coord.x - i, sensor.coord.y + d}] = NONBEACON
				}
			}
		}
	}
	return coords
}

func getSensor(input []string) map[coord]*sensor {
	sensors := make(map[coord]*sensor)
	for _, line := range input {
		var sx, sy, bx, by int
		//Sensor at x=20, y=1: closest beacon is at x=15, y=3
		//Sensor at x=2, y=18: closest beacon is at x=-2, y=15
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		//Calc distance
		distance := diff(sx, bx) + diff(sy, by)
		sensors[coord{sx, sy}] = &sensor{coord{sx, sy}, coord{bx, by}, distance}
	}
	return sensors
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

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
