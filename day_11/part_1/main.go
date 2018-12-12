package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
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

	// Lets read in our solutions file
	for scan.Scan() {

	}

	fmt.Println(time.Since(t))
}

func main() {
	Start()
}
