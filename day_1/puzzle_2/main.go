package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
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

// Find does a binary search and returns the index in which this value could be placed in the slice.
// If the value is the size of the array, this means it is larger then all of the numbers in the slice,
// so just return the value.
// Next we check to see if the value of the index returned in the number we are looking for, then we win.
// else, just return the value.
func find(s []int, i int, f int) int {
	//fmt.Printf("Checking is %d in %v\n", i, s)
	n := sort.SearchInts(s, i)
	if n == len(s) {
		return n
	} else if i == s[n] {
		fmt.Printf("Found repeating frequency: %d\n", i)
		fmt.Printf("Found in %d interations!\n", f)
		os.Exit(0)
	} else {
		return n
	}

	return n
}

func main() {
	// Lets first grab our solutions input
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
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
	//freqslice := []int{freq[0] + freq[1]}
	var freqslice []int

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

	freqcount := 0
	num := 0
	rounds := 0
	for count := 0; count <= len(freq); count++ {
		t := time.Now()
		// fmt.Printf("Count: %d and len is %d\n", count, len(freq))
		if count == len(freq) {
			count = 0
			rounds++
			fmt.Println(time.Since(t))
		}
		//fmt.Printf("Adding %d + %d\n", freqslice[freqcount], freq[count])
		num = num + freq[count]
		index := find(freqslice, num, rounds)
		// Since we are using a binary search, we much keep our list in order. Since we have arrived here, we can assume
		// that we have not found our repeating frequency. We now need to take out number and insert it into the proper index
		// to keep our list ordered.
		freqslice = append(freqslice, 0)             // We must extend the size of the slice to insert a new value.
		copy(freqslice[index+1:], freqslice[index:]) // We then shift the slice.
		freqslice[index] = num                       // Adding our new value into the the correct location.
		freqcount++
	}
}
