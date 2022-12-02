package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	m := make(map[int]int)

	elf := 1
	calorieCount := 0

	for scanner.Scan() {

		if scanner.Text() == "" {
			m[elf] = calorieCount
			elf++
			calorieCount = 0

		} else {
			interInt, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			calorieCount = interInt + calorieCount

		}

		fmt.Printf("line: %s\n", scanner.Text())

	}

	max := 0
	elfnum := 0
	for key, value := range m {
		if value > max {
			max = value
			elfnum = key
		}
	}
	fmt.Println(elfnum, max)
	fmt.Println(m)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
