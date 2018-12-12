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

var grid [][]int

var out = make(chan []int)

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

func getMaxPower(size int) {
	fmt.Println(size)
	maxPower := 0
	maxX := 0
	maxY := 0
	var ret []int
	for x := 0; x < len(grid)-size+1; x++ {
		for y := 0; y < len(grid)-size+1; y++ {
			power := 0
			for x2 := 0; x2 < size; x2++ {
				for y2 := 0; y2 < size; y2++ {
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
	ret = append(ret, maxPower, maxX, maxY, size)
	out <- ret
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

	grid = make([][]int, gridSize)

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
	size := 0
	for i := 0; i < gridSize; i++ {

		go func(i int) {
			getMaxPower(i)
		}(i)

	}

	count := 0
	for i := range out {
		count++
		if count == gridSize {
			close(out)
		}
		if i[0] > maxPower {
			maxPower = i[0]
			maxX = i[1]
			maxY = i[2]
			size = i[3]
		}
	}

	fmt.Println(maxPower, maxX, maxY, size)

	fmt.Println(time.Since(t))
}

func main() {
	Start()
}
