package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var X = 1
var cycle = 1

var cycles = map[int]int{1: 1}

func main() {
	// open file
	f, err := os.Open("day10-input")

	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if scanner.Text() == "noop" {

			cycle++
			cycles[cycle] = X

		} else {
			slice := strings.Split(scanner.Text(), " ")
			cpuInt, err := strconv.Atoi(slice[1])
			if err != nil {
				panic(err)
			}
			cycle++
			cycles[cycle] = X

			cycle++
			X += cpuInt
			cycles[cycle] = X
		}
	}
	fmt.Println(getStrength(20, cycles[20]) + getStrength(60, cycles[60]) + getStrength(100, cycles[100]) + getStrength(140, cycles[140]) + getStrength(180, cycles[180]) + getStrength(220, cycles[220]))
	fmt.Println(cycles)
	fmt.Println(getStrength(220, cycles[220]))
}

func getStrength(cycle int, cpuval int) int {
	return cycle * cpuval
}
