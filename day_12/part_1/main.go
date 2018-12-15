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

// startGen added 2 empty Pots to the front and the back of the list on each generation so we can evalute the current generation.
// We also 2 places in to start our work.
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

	//fmt.Println(b)
	//fmt.Println(a)
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
	//gen := startGen(linked)
	count := 0
	for i := 0; i < 20; i++ {
		l2 := list.New()
		startGen(linked, linked)
		startGen(l2, linked)
		buildNextGen(l2, linked)

		// for e := linked.Front(); e != nil; e = e.Next() {
		// 	if e.Value == "#" {
		// 		count++
		// 	}
		// }
		linked = l2
	}

	i := 0
	for e := starting; e != nil; e = e.Prev() {
		if e.Value == "#" {
			count += i
		}
		i--
	}

	i = 0
	for e := starting; e != nil; e = e.Next() {
		if e.Value == "#" {
			count += i
		}
		i++
	}

	fmt.Println(count)
	fmt.Println(time.Since(t))
}

func main() {
	Start()
}
