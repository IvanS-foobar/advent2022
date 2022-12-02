package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var m1 = map[string]int{
	"A X": 4,
	"B X": 1,
	"C X": 7,
	"A Y": 8,
	"B Y": 5,
	"C Y": 2,
	"A Z": 3,
	"B Z": 9,
	"C Z": 6,
}

var m2 = map[string]int{
	"A X": 3,
	"B X": 1,
	"C X": 2,
	"A Y": 4,
	"B Y": 5,
	"C Y": 6,
	"A Z": 8,
	"B Z": 9,
	"C Z": 7,
}

func main() {
	f, err := os.Open("day2-input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	p := 0
	for scanner.Scan() {
		fmt.Printf("line: %s\n", scanner.Text())
		p = p + m2[scanner.Text()]
	}

	fmt.Println(p)
}
