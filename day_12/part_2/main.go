package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var starting *list.Element

var rules [][]string

func startGen(l *list.List, o *list.List) *list.Element {
	e := o.Front()
	startingNode := e
	if e.Value == "#" {
		l.PushFront(".")
	}
	e = o.Back()
	if e.Value == "#" {
		l.PushBack(".")
	}

	return startingNode
}

func compareSlice(s1 []string) bool {
	for _, i := range rules {
		found := true
		for n := range i {
			if s1[n] != i[n] {
				found = false
			}
		}
		if found == true {
			return true
		}
	}
	return false
}

func getPots(e *list.Element) []string {
	pots := make([]string, 5)
	pots[2] = e.Value.(string)
	l := e.Prev()
	if l == nil {
		pots[1] = "."
		pots[0] = "."
	} else {
		pots[1] = l.Value.(string)
		l = l.Prev()
		if l == nil {
			pots[0] = "."
		} else {
			pots[0] = l.Value.(string)
		}
	}
	e = e.Next()
	if e == nil {
		pots[3] = "."
		pots[4] = "."
	} else {
		pots[3] = e.Value.(string)
		e = e.Next()
		if e == nil {
			pots[4] = "."
		} else {
			pots[4] = e.Value.(string)
		}
	}

	return pots
}

func buildNextGen(n *list.List, o *list.List) {
	old := o.Front()
	if n.Len() == 0 {
		n.PushFront(".")
	}
	start := n.Front()
	for oldele := old; oldele != nil; oldele = oldele.Next() {
		if start == nil {
			n.PushFront(".")
		}
		if oldele == starting {
			starting = start
		}
		current := getPots(oldele)
		if compareSlice(current) {
			start.Value = "#"
		} else {
			start.Value = "."
		}
		n.PushBack(".")
		start = start.Next()
	}
}

func endGameOfLife(c int, d int, s int) {
	total := (50000000000 - (s + 1)) * d
	fmt.Printf("Total: %d\n", c+total)
	os.Exit(0)
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

	var init string
	scan := bufio.NewScanner(file)
	// Lets read in our solutions file
	for scan.Scan() {
		s := scan.Text()
		if strings.Contains(s, "initial") {
			init = strings.TrimSpace(strings.Split(s, ":")[1])
		}

		if strings.Contains(s, "=>") {
			rule := strings.TrimSpace(strings.Split(s, " ")[0])
			r := []string{}
			for _, i := range rule {
				r = append(r, string(i))
			}
			rules = append(rules, r)
		}

	}

	linked := list.New()
	for _, i := range init {
		linked.PushBack(string(i))
	}

	starting = startGen(linked, linked)
	var gameOfLife, previousGen, gameOfLifeCount int
	for i := 0; i < 1000; i++ {
		count := 0
		l2 := list.New()
		startGen(linked, linked)
		startGen(l2, linked)
		buildNextGen(l2, linked)

		linked = l2

		f := 0
		for e := starting; e != nil; e = e.Prev() {
			if e.Value == "#" {
				count += f
			}
			f--
		}

		f = 0
		for e := starting; e != nil; e = e.Next() {
			if e.Value == "#" {
				count += f
			}
			f++
		}

		dif := count - previousGen
		if dif == gameOfLife && gameOfLifeCount == 5 {
			endGameOfLife(count, dif, i)
		} else if dif == gameOfLife {
			gameOfLifeCount++
		} else {
			gameOfLifeCount = 0
			gameOfLife = dif
		}
		previousGen = count

	}

	fmt.Println(time.Since(t))
}

func main() {
	Start()
}
