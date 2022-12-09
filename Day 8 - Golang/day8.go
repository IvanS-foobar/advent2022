package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

func getCoords() map[coord]int {
	// open file
	coords := map[coord]int{}
	f, err := os.Open("day8-input")

	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	y := 0

	for scanner.Scan() {
		slice := strings.Split(scanner.Text(), "")

		for x, i := range slice {
			size, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			coords[coord{x, y}] = size
		}

		y++
	}
	return coords
}

func isVisible(c coord, coords map[coord]int) bool {
	bc := 0
	fmt.Println(c)
	if c.x == 0 || c.y == 0 || c.x == 98 || c.y == 98 {
		fmt.Println("outer tree, returning true")
		return true
	} else {
		//check left to edge
		for i := 0; i < c.x; i++ {
			if coords[coord{i, c.y}] >= coords[c] {
				fmt.Println("can't see it from the left")
				bc++
				break
			}
		}
		//check right to edge
		for i := c.x + 1; i < 99; i++ {
			if coords[coord{i, c.y}] >= coords[c] {
				fmt.Println("can't see it from the right")
				bc++
				break
			}
		}
		//check up to edge
		for i := 0; i < c.y; i++ {
			if coords[coord{c.x, i}] >= coords[c] {
				fmt.Println("can't see it from the top")
				bc++
				break
			}
		}
		//check down to edge
		for i := c.y + 1; i < 99; i++ {
			if coords[coord{c.x, i}] >= coords[c] {
				fmt.Println("can't see it from the bottom")
				bc++
				break
			}
		}
	}
	fmt.Println(bc)
	if bc == 4 {
		fmt.Println("could not see it from any 4 of the directions, return false")
		return false
	} else {
		fmt.Println("we could the see the tree from at least one direction, returning true")
		return true
	}

}

func main() {
	coords := getCoords()
	count := 0
	for i := range coords {
		if isVisible(i, coords) {
			count++
		}
	}
	fmt.Println(len(coords))
	fmt.Println(count)
}
