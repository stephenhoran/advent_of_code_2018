package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type point struct {
	X    float64
	Y    float64
	Area int
}

type grid struct {
	MAXX int
	MAXY int
	MINX int
	MINY int
}

func manhattenDistance(x int, y int, b point) float64 {
	return math.Abs(float64(x)-b.X) + math.Abs(float64(y)-b.Y)
}

func isWinner(i int, s []int) bool {
	for _, c := range s {
		if i == c {
			return false
		}
	}
	return true
}

//Start launches program
func Start() {
	t := time.Now()

	var sliceOfPoints []point
	var disqualified []int
	// Lets first grab our solutions input
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	currentGrid := grid{
		MINX: 0,
		MINY: 0,
		MAXX: 0,
		MAXY: 0,
	}

	for scan.Scan() {
		s := strings.Split(scan.Text(), ",")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(strings.TrimSpace(s[1]))
		sliceOfPoints = append(sliceOfPoints, point{
			X: float64(x),
			Y: float64(y),
		})

		if x > currentGrid.MAXX {
			currentGrid.MAXX = x
		}
		if y > currentGrid.MAXY {
			currentGrid.MAXY = y
		}
		if x < currentGrid.MINX || currentGrid.MINX == 0 {
			currentGrid.MINX = x
		}
		if y < currentGrid.MINY || currentGrid.MINY == 0 {
			currentGrid.MINY = y
		}

	}

	for x := currentGrid.MINY; x <= currentGrid.MAXX; x++ {
		for y := currentGrid.MINY; y <= currentGrid.MAXY; y++ {
			var index int
			min := float64(0)
			dup := false
			for i, n := range sliceOfPoints {
				m := manhattenDistance(x, y, n)
				if i == 0 {
					min = m
					continue
				}
				if m == 0 {
					index = i
					break
				} else if m < min && dup == true {
					min = m
					dup = false
					index = i
				} else if m < min {
					min = m
					index = i
				} else if m == min {
					dup = true
				} else {
					continue
				}
			}

			if dup == true {
				continue
			} else if x == 0 || y == 0 || x == currentGrid.MAXX || y == currentGrid.MAXY {
				// fmt.Println(index, x, y)
				sliceOfPoints[index].Area++
				disqualified = append(disqualified, index)
			} else {
				// fmt.Println(index, x, y)
				sliceOfPoints[index].Area++
			}
		}
	}

	largestArea := 0
	winner := 0
	for i, p := range sliceOfPoints {
		if p.Area > largestArea {
			w := isWinner(i, disqualified)
			if w == true {
				winner = i
				largestArea = p.Area
			}
		}
	}
	fmt.Println(winner, sliceOfPoints[winner])
	fmt.Println(time.Since(t))
}

func main() {
	Start()
}
