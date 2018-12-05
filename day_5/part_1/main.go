// COMMENTS ARE IN PART 1. This solution was one nested loop, which I feel is self explanitory if you read part
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
)

//Start launches program
func Start() {

	// Lets first grab our solutions input
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		t := time.Now()
		s := scan.Text()
		for i := 0; i < len(s)-1; i++ {
			var sb strings.Builder
			if unicode.IsUpper(rune(s[i])) {
				l := unicode.ToLower(rune(s[i]))
				// if (l == rune(s[i-1])) && (l == rune(s[i+1])) {
				// 	sb.WriteByte(s[i-1])
				// 	sb.WriteByte(s[i])
				// 	sb.WriteByte(s[i+1])
				// 	s = strings.Replace(s, sb.String(), "", 1)
				// 	fmt.Println(s)
				// 	i = i - 3
				if l == rune(s[i-1]) {
					sb.WriteByte(s[i-1])
					sb.WriteByte(s[i])
					s = strings.Replace(s, sb.String(), "", 1)
					i = i - 3
				} else if l == rune(s[i+1]) {
					sb.WriteByte(s[i])
					sb.WriteByte(s[i+1])
					s = strings.Replace(s, sb.String(), "", 1)
					i = i - 2
				} else {
					continue
				}
			}

		}
		fmt.Println(len(s))
		fmt.Println(time.Since(t))
	}

}

func main() {
	Start()
}
