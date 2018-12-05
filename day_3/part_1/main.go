package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Start begin the program
func Start() int {
	// Lets first grab our solutions input
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	fab := make(map[int][]int)

	scan := bufio.NewScanner(file)
	top := 0
	left := 0
	for scan.Scan() {
		s := strings.Split(scan.Text(), " ")
		id, _ := strconv.Atoi(s[0][1:])
		place := strings.Split(s[2], ",")
		dem := strings.Split(s[3], "x")
		l, _ := strconv.Atoi(place[0])
		t, _ := strconv.Atoi(strings.TrimRight(place[1], ":"))
		w, _ := strconv.Atoi(dem[0])
		h, _ := strconv.Atoi(dem[1])
		width := t + w
		height := l + h
		if width > top {
			top = width
		}
		if height > left {
			left = height
		}

		fab[id] = []int{l, t, w, h}
	}

	fabric := make([][]uint8, top)
	for i := range fabric {
		fabric[i] = make([]uint8, left)
	}

	for _, v := range fab {
		et := v[0] + v[2]
		bl := v[1] + v[3]
		for i := v[0]; i < et; i++ {
			for j := v[1]; j < bl; j++ {
				fabric[i][j]++
			}
		}
	}

	count := 0
	for _, j := range fabric {
		for _, i := range j {
			if i > 1 {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println(Start())
}
