package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
	"unicode"
)

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

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		s := scan.Text()
		var z []rune
		max := 0
		index := 0
		for i, r := range s {
			if i == 0 {
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
		fmt.Println(index)
	}
	fmt.Println(time.Since(t))
}

func main() {
	Start()
}
