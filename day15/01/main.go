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

	sensors, coords := getSensorsAndBeacons(input)
	//coords := getCoords(sensors)

	//signals := getBeaconSignalsOnRow(sensors, coords, 10)
	signals := getBeaconSignalsOnRow(sensors, coords, 2000000)

	//printMap(coords)
	fmt.Println(signals)
}

func getBeaconSignalsOnRow(sensors map[coord]*sensor, coords map[coord]string, row int) int {
	signals := 0

	xCoords := make([]int, 0)
	for c, s := range sensors {
		xCoords = append(xCoords, c.x+s.distance)
		xCoords = append(xCoords, c.x-s.distance)
	}
	sort.Ints(xCoords)

	y := row
	for x := xCoords[0]; x < xCoords[len(xCoords)-1]; x++ {
		//Loop throug X and check if it is covered by a sensor
		for c, s := range sensors {
			//Check if sensor covers row
			//fmt.Printf("Distance: %d, x: %d, y: %d - checkX: %d, checkY: %d\n", s.distance, c.x, c.y, x, y)
			if s.distance >= diff(c.x, x)+diff(c.y, y) {
				if _, ok := coords[coord{x, y}]; !ok {
					signals++ //Not a signal or beacon
					break     //Next coord
				}
			}
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

// Get map with sensor as value and map with string as value
func getSensorsAndBeacons(input []string) (map[coord]*sensor, map[coord]string) {
	sensors := make(map[coord]*sensor)
	coords := make(map[coord]string)

	for _, line := range input {
		var sx, sy, bx, by int
		//Sensor at x=20, y=1: closest beacon is at x=15, y=3
		//Sensor at x=2, y=18: closest beacon is at x=-2, y=15
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		//Calc distance
		distance := diff(sx, bx) + diff(sy, by)
		sensors[coord{sx, sy}] = &sensor{coord{sx, sy}, coord{bx, by}, distance}
		coords[coord{sx, sy}] = SENSOR
		coords[coord{bx, by}] = BEACON

	}
	return sensors, coords
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
