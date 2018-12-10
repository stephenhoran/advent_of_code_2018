package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

var work []string
var current string
var workers []*worker // A slice of workers

var queue [][]string   // Queue from Part 1.
var donequeue []string // Everything that has completed working in part 2.
var pt2q []string      // Our queue for part 2.

var cost = make(map[string]int)

// Struct for our worker.
type worker struct {
	currentNode string
	timeLeft    int
	available   bool
}

// startWork takes the current worker node and begins the clock.
func (w *worker) startWork(s string) {
	w.currentNode = s
	w.available = false
	w.timeLeft = cost[s]
}

// step decrements the time, makes the worker available and adds its node to the done queue.
func (w *worker) step() {
	w.timeLeft--
	if w.timeLeft < 0 {
		w.timeLeft = 0
	}
	if w.timeLeft == 0 && w.available == false {
		donequeue = append(donequeue, w.currentNode)
		w.available = true
	}

}

// Determine if any worker has started working on the current node.
func hasStarted(s string) bool {
	for _, w := range workers {
		if s == w.currentNode {
			return true
		}
	}
	return false
}

// To build a dynamic list of work available for part 2. We need to check to see if all of a nodes dependencies are now
// in the donequeue.
func areDepsComplete(s string, n map[string][]string) bool {
	count := 0 // Count needs to be the length of the slice of dependencies.
	for _, i := range n[s] {
		for _, q := range donequeue {
			if s == q { // If this node is already in the donequeue, fail fast.
				return false
			}
			if q == i {
				count++
			}
		}
	}
	if len(n[s]) == 0 { // if len is zero, these have already been addressed in the first iteration.
		return false
	}
	if len(n[s]) == count { // All dependencies are in the donequeue.
		return true
	}

	return false
}

//Start launches program
func Start() {
	t := time.Now()

	nodes := make(map[string][]string)
	//var order []string
	// Lets first grab our solutions input
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		s := strings.Split(scan.Text(), " ")
		nodes[s[7]] = append(nodes[s[7]], s[1])
		if _, ok := nodes[s[1]]; !ok {
			nodes[s[1]] = make([]string, 0)
		}
	}

	// Small refactors to part 1.
	var complete []string
	for len(complete) != len(nodes) {
		for k, v := range nodes {
			for _, i := range v {
				if i == string(current) {
					//nodes[k] = append(nodes[k][:count], nodes[k][count+1:]...)
				}
			}
			if len(nodes[k]) == 0 {
				work = append(work, k)
			}
		}
		sort.Strings(work)
		current = work[0]
		queue = append(queue, work)
		complete = append(complete, work[0])
		work = nil

	}

	// Here is our starting point for part 2. All on the items with no dependencies
	pt2q := queue[0]
	fmt.Println(pt2q)
	fmt.Println(nodes)

	// Lets build the cost of each value now.
	count := 1
	for i := 65; i < 65+26; i++ {
		cost[string(i)] = count + 60
		count++
	}

	// Create a slice of workers
	for i := 0; i < 5; i++ {
		workers = append(workers, &worker{
			available: true,
		})
	}

	step := 0 // Track our steps
	// fmt.Printf("Sec\tWorker1\tWorker2\tWorker3\tWorker4\tWorker5\tComplete\n")
	// fmt.Printf("Sec\tWorker1\tWorker2\tComplete\n")
	for len(donequeue) < len(nodes) { // While the donequeue does not contain all nodes.
		for _, i := range pt2q { // range over queue
			for _, w := range workers { // range over workers
				if !hasStarted(i) { // If this node is not started on any other node.
					if w.available { // If this node is available for work
						w.startWork(i) // Start node working on this worker.
					}
				}
			}
		}
		//fmt.Printf("%d\t%s,%d\t%s,%d\t%s,%d\t%s,%d\t%s,%d\t%s\n", step, workers[0].currentNode, workers[0].timeLeft, workers[1].currentNode, workers[1].timeLeft, workers[2].currentNode, workers[2].timeLeft, workers[3].currentNode, workers[3].timeLeft, workers[4].currentNode, workers[4].timeLeft, donequeue)
		//fmt.Printf("%d\t%s,%d\t%s,%d\t%s\n", step, workers[0].currentNode, workers[0].timeLeft, workers[1].currentNode, workers[1].timeLeft, donequeue)
		for _, w := range workers {
			w.step() // Increment the steps on all the workers.
		}

		// Rebuild the working queue with anything whose dependencies are resolved.
		pt2q = make([]string, 0)
		if len(donequeue) != 0 {
			for k := range nodes {
				if areDepsComplete(k, nodes) {
					pt2q = append(pt2q, k)
				}
			}
		}
		sort.Strings(pt2q)
		step++
	}
	fmt.Println(donequeue, step)
	fmt.Println(time.Since(t))

}

func main() {
	Start()
}
