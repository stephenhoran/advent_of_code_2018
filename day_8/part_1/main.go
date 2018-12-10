package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseNode(license []int) (int, []int) {
	children := license[:1][0]
	metadata := license[1:2][0]
	licenseData := license[2:]
	count := 0

	// fmt.Printf("Children: %d\tMetadata: %d\tLicense: %v\n", children, metadata, licenseData[metadata:])
	for c := 0; c < children; c++ {
		result := 0
		result, licenseData = parseNode(licenseData)
		count += result
	}

	for _, m := range licenseData[:metadata] {
		count += m
	}

	return count, licenseData[metadata:]
}

//Start launches program
func Start() {
	t := time.Now()

	var license []int

	// Lets first grab our solutions input
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	// Lets read in our solutions file and store it in a slice of int's (license)
	for scan.Scan() {
		b := scan.Text()
		for _, s := range strings.Split(b, " ") {
			i, _ := strconv.Atoi(s)
			license = append(license, i)
		}
	}

	count, _ := parseNode(license)
	fmt.Println(count)

	fmt.Println(time.Since(t))
}

func main() {
	Start()
}
