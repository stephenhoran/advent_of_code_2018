package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// findDuplicates accepts a string in which you wish to check for duplicates. It returns 2 ints, one containing
// an 0 or 1 if 2 duplicates are found. The second int returns a 0 or 1 if 3 duplicates are found. This is not
// the most efficent as it will check every character even the chracters we have alread found duplicates for.
// We do exit in the event we have found a two and a three as no additional work needs to be done.
func findDuplicates(s string) (int, int) {
	split := strings.Split(s, "") // Split the character
	two := 0
	three := 0
	for _, c := range split { // range over the split
		count := strings.Count(s, c) // Count the number of times that character appears in this string.
		if count == 2 {
			two = 1
		} else if count == 3 {
			three = 1
		} else if (two == 1) && (three == 1) {
			return two, three
		}
	}
	return two, three
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

	//Lets store our results in a map for now.
	totals := map[string]int{
		"two":   0,
		"three": 0,
	}

	for _, i := range inv {
		two, three := findDuplicates(i)
		totals["two"] = totals["two"] + two
		totals["three"] = totals["three"] + three
	}

	fmt.Printf("Number of twos found: %d\n", totals["two"])
	fmt.Printf("Number of threes found: %d\n", totals["three"])
	fmt.Printf("Your hash is: %d\n", totals["two"]*totals["three"])

}
