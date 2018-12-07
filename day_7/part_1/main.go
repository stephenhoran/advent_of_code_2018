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

//Start launches program
func Start() {
	t := time.Now()

	nodes := make(map[string][]string)
	var complete []string
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

	for len(nodes) > 0 {
		for k, v := range nodes {
			for count, i := range v {
				if i == string(current) {
					nodes[k] = append(nodes[k][:count], nodes[k][count+1:]...)
				}
			}
			if len(nodes[k]) == 0 {
				work = append(work, k)
			}
		}
		sort.Strings(work)
		current = work[0]
		delete(nodes, current)
		complete = append(complete, work[0])
		work = nil
	}
	fmt.Println(strings.Join(complete, ""))
	fmt.Println(time.Since(t))
}

func main() {
	Start()
}
