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

func parseNode(license []int) (int, []int, int) {
	children := license[:1][0]
	metadata := license[1:2][0]
	licenseData := license[2:]
	count := 0
	var values []int

	// fmt.Printf("Children: %d\tMetadata: %d\tLicense: %v\n", children, metadata, licenseData[metadata:])
	for c := 0; c < children; c++ {
		value := 0
		result := 0
		result, licenseData, value = parseNode(licenseData)
		count += result
		values = append(values, value)
	}

	for _, m := range licenseData[:metadata] {
		count += m
	}

	value := 0
	if children == 0 {
		for _, m := range licenseData[:metadata] {
			value += m
		}
		//fmt.Println(licenseData[:metadata], values)
		return count, licenseData[metadata:], value
	}
	// If metadata is great then zero and in length of metadata. I previous tried to throw away
	for _, m := range licenseData[:metadata] {
		if m <= len(values) && m > 0 {
			value += values[m-1] // Add the previous values
		}
	}

	return count, licenseData[metadata:], value
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

	_, _, values := parseNode(license)
	fmt.Println(values)

	fmt.Println(time.Since(t))
}

func main() {
	Start()
}
