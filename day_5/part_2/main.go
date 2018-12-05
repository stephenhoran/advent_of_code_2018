package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
	"unicode"
)

var wg sync.WaitGroup
var mx sync.Mutex
var final []int

func work(s string, b rune) {
	var z []rune
	max := 0
	index := 0
	for _, r := range s {
		if r == b || r == unicode.ToLower(b) {
			continue
		}
		if index == 0 {
			z = append(z, r)
			index++
			max++
		} else if (unicode.IsUpper(r) || unicode.IsUpper(z[index-1])) && (unicode.IsLower(r) || unicode.IsLower(z[index-1])) {
			if unicode.ToLower(r) == unicode.ToLower(z[index-1]) {
				index--
			} else {
				if index < max {
					z[index] = r
					index++
				} else {
					z = append(z, r)
					index++
					max++
				}
			}
		} else {
			if index < max {
				z[index] = r
				index++
			} else {
				z = append(z, r)
				index++
				max++
			}
		}
	}
	mx.Lock()
	final = append(final, index)
	mx.Unlock()
	wg.Done()
}

//Start launches program
func Start() {
	timeoverall := time.Now()
	// Lets first grab our solutions input
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		const ABC = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		s := scan.Text()

		for _, b := range ABC {
			wg.Add(1)
			go func(s string, b rune) {
				work(s, b)
			}(s, b)
		}
		wg.Wait()
		m := 0
		for _, f := range final {
			if m == 0 || f < m {
				m = f
			}
		}
	}

	fmt.Println(time.Since(timeoverall))
}

func main() {
	Start()
}
