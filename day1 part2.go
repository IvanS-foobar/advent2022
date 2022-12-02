package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	// open file
	f, err := os.Open("day1-input")

	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	var m []int

	calorieCount := 0

	for scanner.Scan() {

		if scanner.Text() != "" {
			interInt, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}

			calorieCount = interInt + calorieCount
			fmt.Println(calorieCount)
			continue
		}
		m = append(m, calorieCount)
		calorieCount = 0
		fmt.Printf("line: %s\n", scanner.Text())
	}
	m = append(m, calorieCount)

	sort.Sort(sort.Reverse(sort.IntSlice(m)))
	fmt.Println(m)
	fmt.Println(m[0] + m[1] + m[2])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
