package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// You notice that the device repeats the same frequency change list over and over. To calibrate the device, you need to find the first frequency it reaches twice.

// For example, using the same list of changes above, the device would loop as follows:

// Current frequency  0, change of +1; resulting frequency  1.
// Current frequency  1, change of -2; resulting frequency -1.
// Current frequency -1, change of +3; resulting frequency  2.
// Current frequency  2, change of +1; resulting frequency  3.
// (At this point, the device continues from the start of the list.)
// Current frequency  3, change of +1; resulting frequency  4.
// Current frequency  4, change of -2; resulting frequency  2, which has already been seen.
// In this example, the first frequency reached twice is 2. Note that your device might need to repeat its list of frequency changes many times before a duplicate frequency is found, and that duplicates might be found while in the middle of processing the list.

// Here are other examples:

// +1, -1 first reaches 0 twice.
// +3, +3, +4, -2, -4 first reaches 10 twice.
// -6, +3, +8, +5, -6 first reaches 5 twice.
// +7, +7, -2, -7, -4 first reaches 14 twice.
// What is the first frequency your device reaches twice?

// Find completes a linear search on the provided slice to find the requested int. Not sure if spending the time to sort the list and use
// a binary search would be more efficent.
func find(s *[]int, i int) {
	// fmt.Printf("Checking is %d in %v\n", i, s)
	for _, n := range *s {
		if i == n {
			fmt.Printf("Found repeating frequency: %d\n", i)
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

	var freq []int // Create a slice of int to store our frequencies in.

	// Lets build our slice containing frequencies from the file.
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		i, _ := strconv.Atoi(scan.Text())
		freq = append(freq, i)
	}

	// A slice for the previous frequencies, we use composition to prepopulate the slice with the first 2 pairs added.
	freqslice := []int{freq[0] + freq[1]}

	// The puzzle states that we may need to iterate over our numbers multiple times to find the duplicates, so
	// we state an infinite loop with a counter to track our place. If the current count is longer then the len
	// of the slice, we need to set it back to 0.
	//
	// Next we take the current count as the index and the next iteration and add them together, determine if
	// this frequency is already in the slice using find(). If we don't find this number we will append the number
	// and continue on.
	// Example numbers: [-7 +16 +5 +11]
	// -7 + 16 = 9 -> find(9) -> proceed
	// 9 + 5 = 14 -> find(14) -> proceed
	// 14 + 11 = 25 -> find(25) -> proceed
	// 25 - 7 = 18 -> find(18) -> proceed

	count := 2 // Since we aready addressed the first pair, we need to start on the third number
	freqcount := 0
	for {
		// fmt.Printf("Count: %d and len is %d\n", count, len(freq))
		if count >= len(freq) {
			count = 0
		}
		//fmt.Printf("Adding %d + %d\n", freqslice[freqcount], freq[count])
		i := freqslice[freqcount] + freq[count]
		find(&freqslice, i) // Dereferencing here, The size of this slice gets pretty large, super easy optimization here.
		freqslice = append(freqslice, i)
		count++
		freqcount++
	}
}
