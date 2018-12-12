package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

const gridSize int = 300

var serialNum int

func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}

func getPowerLevel(x int, y int) int {
	rackid := x + 10
	powerlevel := rackid * y
	powerlevel += serialNum
	powerlevel *= rackid
	powerlevel = digit(powerlevel, 3)
	return powerlevel - 5
}

//Start launches program
func Start() {
	t := time.Now()

	// Lets first grab our solutions input
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	// Lets read in our solutions file
	for scan.Scan() {
		s := scan.Text()
		serialNum, _ = strconv.Atoi(s)
	}

	grid := make([][]int, gridSize)

	for i := 0; i < gridSize; i++ {
		grid[i] = make([]int, gridSize)
	}

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid); y++ {
			grid[x][y] = getPowerLevel(x, y)
		}
	}

	maxPower := 0
	maxX := 0
	maxY := 0

	for x := 0; x < len(grid)-2; x++ {
		for y := 0; y < len(grid)-2; y++ {
			power := 0
			for x2 := 0; x2 < 3; x2++ {
				for y2 := 0; y2 < 3; y2++ {
					power += grid[x+x2][y+y2]
				}
			}
			if power > maxPower {
				maxPower = power
				maxX = x
				maxY = y
			}
		}
	}

	fmt.Println(maxPower, maxX, maxY)

	fmt.Println(time.Since(t))
}

func main() {
	Start()
}
