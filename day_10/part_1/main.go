package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type point struct {
	posX  int
	posY  int
	voloX int
	voloY int
}

//Start launches program
func Start() {
	t := time.Now()

	pos := [][]int{[]int{1, -1}, []int{1, 0}, []int{1, 1}, []int{0, -1}, []int{0, 1}, []int{-1, -1}, []int{-1, 0}, []int{-1, 1}}
	// Lets first grab our solutions input
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	points := make([]*point, 0)
	// Lets read in our solutions file
	for scan.Scan() {
		s := strings.Split(scan.Text(), "=")
		p := (strings.Split(s[1], "<"))
		p = strings.Split(p[1], ">")
		p = strings.Split(p[0], ",")
		v := (strings.Split(s[2], "<"))
		v = strings.Split(v[1], ">")
		v = strings.Split(v[0], ",")
		px, _ := strconv.Atoi(strings.TrimSpace(strings.TrimRight(p[0], ",")))
		py, _ := strconv.Atoi(strings.TrimSpace(p[1]))
		vx, _ := strconv.Atoi(strings.TrimSpace(v[0]))
		vy, _ := strconv.Atoi(strings.TrimSpace(v[1]))

		points = append(points, &point{
			posX:  px,
			posY:  py,
			voloX: vx,
			voloY: vy,
		})

	}

	found := false
	for found == false {
		for _, p := range points {
			p.posX = p.posX + p.voloX
			p.posY = p.posY + p.voloY
		}

		for _, r := range points {
			for _, p := range pos {
				for _, r2 := range points {
					if r2.posX == r.posX+p[0] && r2.posY == r.posY+p[1] {
						goto Next
					}
				}
			}
			goto End
		Next:
		}
		found = true
	End:
	}

	maxX := 0
	maxY := 0
	for _, j := range points {
		if j.posX > maxX {
			maxX = j.posX
		}
		if j.posY > maxY {
			maxY = j.posY
		}
	}

	var array [][]int
	for i := 0; i < maxY+1; i++ {
		array = append(array, make([]int, maxX+1))
	}

	for _, r := range points {
		array[r.posY][r.posX] = 1
	}

	for _, a := range array {
		fmt.Println(a)
	}

	fmt.Println(time.Since(t))
}

func main() {
	Start()
}
