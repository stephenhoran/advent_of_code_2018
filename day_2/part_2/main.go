package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// This method is pretty terrible, time complexity of O(n2). Using this as a benchmark. If I have enough time I will provide a solution using
// hashing. The solution here was found pretty quick however I only had to iterate over the solution 46 times.

// compareStrings takes our first string and linearly compares each position of a string to the other ID's provided in the solution.
// If we find a match we stop doing further work and present a string containing the sequential results of common letters.
func compareStrings(s string, i []string) {
	ss := strings.Split(s, "") // split our first string. Do this first so we do not keep splitting for no reason.
	for _, id := range i {
		mismatches := 0              // store our mismatches, we are looking for exactly 1 mismatch at the end of the cycle.
		ids := strings.Split(id, "") // split the item recieved from the solution.
		for index, c := range ss {
			if c != ids[index] { // compare the current index value of our string to the index value of the item presented during this iteration.
				mismatches++
			}
		}
		if mismatches == 1 { // This is simple linear comparison and concatination.
			var common []string
			fmt.Printf("Found %d number of mismatches between %s and %s\n", mismatches, s, id)
			for index, c := range ss {
				if c == ids[index] {
					common = append(common, c)
				}
			}
			fmt.Printf("The common character are: %s\n", strings.Join(common, ""))
			os.Exit(0)
		}
	}

}

func main() {
	// Lets first grab our solutions input
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var inv []string // Create a slice of int to store our inventory in.

	// Lets build our slice containing inventory from the file.
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		inv = append(inv, scan.Text())
	}

	for _, i := range inv {
		compareStrings(i, inv)
	}

}
