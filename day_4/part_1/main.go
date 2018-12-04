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

var id string       // This id contains the last seen id of our guard, this allows us to manage the transaction.
var foundid string  // Foundid contains the ID of the guard we will be reporting or the solution.
var start time.Time // Start time of the current transaction.
var end time.Time   // End time of the current transaction.
var index int       // The index of the hour in which is currently the highest.
var max int8        // Currently the count of the most slept hour.

// Sleeplog is a struct containing the current count of the guards sleep time and a slice size of 60 which contains the minutes
// slept by a guard.
type sleeplog struct {
	minutes   int
	timeslice []int8
}

// Event stores the timestamp and a slice of strings containing details of the event.
type event struct {
	ts    time.Time
	event []string
}

// Sortedevents is a slice of event that will fulfill the sort interface requirements to all us to sort the slice of events.
type sortedevents []event

// Len fulfills the sort interface. It is the len of elements in the collection.
func (e sortedevents) Len() int {
	return len(e)
}

// Less fulfills the sort interface. Here we say if i came before j return a bool using the timestamp in the struct.
func (e sortedevents) Less(i, j int) bool {
	return e[i].ts.Before(e[j].ts)
}

// Swap fulfills the sort interface. Here if less returns true, this would be the arrangement they are stored.
func (e sortedevents) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// Find simple does a linear lookup of a string in a slice. Returns true or false.
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
		// We will be using string builder to efficently parse the text.
		var ts strings.Builder
		s := strings.Split(scan.Text(), " ")
		ts.WriteString(s[0][1:]) // we split up the timestamp
		ts.WriteString(" ")
		ts.WriteString(s[1][:5])                            // and now the event
		t, _ := time.Parse("2006-01-02 15:04", ts.String()) // parsing the timestamp
		events = append(events, event{                      // append to our slice of events an event struct literally containing our parsed information
			ts:    t,
			event: s[2:],
		})
	}

	sort.Sort(sortedevents(events)) // Now we sort those events cronologically.

	guard := make(map[string]*sleeplog) // Create a map to hold our guards and their sleep information.

	// Time to iterate over our events.
	for _, i := range events {
		// If we find a guard we need to store that guard in a variable. This will let us know when the transaction for this
		// guard is complete.
		if find(i.event, "Guard") {
			id = i.event[1][1:]
			// We check to see if an entry for this guard exists in this map.
			_, ok := guard[id]
			if ok {
				continue
			} else {
				// If not lets compose one.
				guard[id] = &sleeplog{
					timeslice: make([]int8, 60),
					minutes:   0,
				}
			}
		}
		// If we find "falls", lets assign this starting timestamp.
		if find(i.event, "falls") {
			start = i.ts
		}
		// If we find wakes, we know this iteration of sleeping is done. Grab the timestamp and determine how long that guard has slept for.
		// Then convert the start minute to an int and the duration to an int.
		if find(i.event, "wakes") {
			end = i.ts
			e := end.Sub(start)
			s := int64(time.Duration(start.Minute()))
			d := int64(e / time.Minute)
			// This loop basically starts at the minute the guard starting sleeping and iterates for each minute the guard slept. It
			// then increments each minute by 1 along with the count of minutes slept.
			for j := s; j < s+d; j++ {
				guard[id].timeslice[j]++
				guard[id].minutes++
			}
		}
	}

	long := 0
	// Here we loop over off our guard and determine who slept the move amount of minutes.
	for k, v := range guard {
		if v.minutes > long {
			long = v.minutes
			foundid = k
		}
	}

	// Here we look for the minute is which was slept the most. Since we only increment by one we are just looking for the largest number.
	for i, v := range guard[foundid].timeslice {
		if v > max {
			max = v
			index = i
		}
	}

	z, _ := strconv.Atoi(foundid)

	fmt.Println(foundid, index, z*index)

	fmt.Println(time.Since(t))

}
