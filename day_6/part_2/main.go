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

	"github.com/pkg/profile"
)

type point struct {
	X float64
	Y float64
}

type grid struct {
	MAXX int
	MAXY int
	MINX int
	MINY int
}

func manhattenDistance(x float64, y float64, b point) float64 {
	return math.Abs(x-b.X) + math.Abs(y-b.Y)
}

//Start launches program
func Start() {
	t := time.Now()

	var sliceOfPoints []point
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

	area := 0
	for x := currentGrid.MINX; x <= currentGrid.MAXX; x++ {
		for y := currentGrid.MINY; y <= currentGrid.MAXY; y++ {
			c := 0
			for _, n := range sliceOfPoints {
				m := manhattenDistance(float64(x), float64(y), n)
				c += int(m)
			}
			if c < 10000 {
				area++
			}
		}
	}
	fmt.Println(area)
	fmt.Println(time.Since(t))
}

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	for i := 0; i < 10; i++ {
		Start()
	}

}
