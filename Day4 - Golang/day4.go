package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day4-input")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	p := 0
	for scanner.Scan() {
		pair1 := strings.Split(strings.Split(scanner.Text(), ",")[0], "-")
		pair2 := strings.Split(strings.Split(scanner.Text(), ",")[1], "-")

		fmt.Println(pair1, pair2)

		pair1map := createMap(pair1[0], pair1[1])
		pair2map := createMap(pair2[0], pair2[1])

		fmt.Println(pair1map, pair2map)

		if detrOverlap(pair1map, pair2map) {
			p++
		} else if detrOverlap(pair2map, pair1map) {
			p++
		}
		print(p)
	}
}

func createMap(a string, b string) map[int]bool {
	c, _ := strconv.Atoi(a)
	d, _ := strconv.Atoi(b)

	m := make(map[int]bool)
	for i := c; i <= d; i++ {
		m[i] = true
	}
	return m
}

func detrOverlap(a map[int]bool, b map[int]bool) bool {
	for i := range a {
		if b[i] {
			continue
		} else {
			return false
		}

	}
	return true
}
