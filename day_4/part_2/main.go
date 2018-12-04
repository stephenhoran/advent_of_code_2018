package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var id string
var foundid string
var start time.Time
var end time.Time
var index int
var min int8

type sleeplog struct {
	minutes   int
	timeslice []int8
}

type event struct {
	ts    time.Time
	event []string
}

type sortedevents []event

func (e sortedevents) Len() int {
	return len(e)
}

func (e sortedevents) Less(i, j int) bool {
	return e[i].ts.Before(e[j].ts)
}

func (e sortedevents) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func find(s []string, x string) bool {
	for _, n := range s {
		if x == n {
			return true
		}
	}
	return false
}

func main() {
	t := time.Now()
	var events []event
	// Lets first grab our solutions input
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		var ts strings.Builder
		s := strings.Split(scan.Text(), " ")
		ts.WriteString(s[0][1:])
		ts.WriteString(" ")
		ts.WriteString(s[1][:5])
		t, _ := time.Parse("2006-01-02 15:04", ts.String())
		events = append(events, event{
			ts:    t,
			event: s[2:],
		})
	}

	sort.Sort(sortedevents(events))

	guard := make(map[string]*sleeplog)

	for _, i := range events {
		if find(i.event, "Guard") {
			id = i.event[1][1:]
			_, ok := guard[id]
			if ok {
				continue
			} else {
				guard[id] = &sleeplog{
					timeslice: make([]int8, 60),
					minutes:   0,
				}
			}
		}
		if find(i.event, "falls") {
			start = i.ts
		}
		if find(i.event, "wakes") {
			end = i.ts
			e := end.Sub(start)
			s := int64(time.Duration(start.Minute()))
			d := int64(e / time.Minute)
			for j := s; j < s+d; j++ {
				guard[id].timeslice[j]++
				guard[id].minutes++
			}
		}
	}

	var long int8
	for k, v := range guard {
		for i, j := range v.timeslice {
			if j > long {
				long = j
				index = i
				foundid = k
			}
		}
	}

	z, _ := strconv.Atoi(foundid)

	fmt.Println(foundid, index, z*index)

	fmt.Println(time.Since(t))

}
