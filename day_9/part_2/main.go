package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var marble int
var currentPlayer int

// Map for scoring the players scores.
type playersScores map[int]int

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

	players := 0
	lastMarble := 0
	// Lets read in our solutions file
	for scan.Scan() {
		s := strings.Split(scan.Text(), " ")
		players, _ = strconv.Atoi(s[0])
		lastMarble, _ = strconv.Atoi(s[6])
	}

	// Create a map of scores as well as our board game. Also keep track of the current marble and player
	scores := make(playersScores)
	marble = 1

	// Build our list of players before starting.
	for p := 1; p <= players; p++ {
		scores[p] = 0
	}

	circle := &ring.Ring{Value: 0}

	for m := marble; m <= lastMarble*100; m++ {
		if m%23 == 0 {
			circle = circle.Move(-8)
			player := m % players
			r := circle.Unlink(1)
			scores[player] += m + r.Value.(int)
			circle = circle.Next()
		} else {
			circle = circle.Next()

			r := ring.New(1)
			r.Value = m

			circle.Link(r)
			circle = circle.Next()
		}
	}

	winner := []int{0, 0}
	for k, v := range scores {
		if v > winner[1] {
			winner[0] = k
			winner[1] = v
		}
	}

	fmt.Println(winner)
	fmt.Println(time.Since(t))
}

func main() {
	Start()
}
